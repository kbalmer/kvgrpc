package main

import (
	"context"
	"flag"
	"fmt"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	pb "kvgrpc/kv"
	"log"
	"net"
	"sync"
	"time"
)

var (
	// postList is the list of ports to be consumed by the node
	portList portCollection
	//livePort is the port in which the node will send and receive messages on
	livePort string
)

// the server used to implement kv.KVServer
type node struct {
	store      map[string][]byte
	mu         sync.RWMutex
	network    []nodeConnections
	lastAction int64
}

// portCollection is used to store multiple ports to be consumed as a flag for the node connections
type portCollection []string

// nodeConnections stores
type nodeConnections struct {
	conn     *grpc.ClientConn
	kvClient pb.KVClient
}

// NewNode returns a blank node with the map for KV storage and the read-write mutex for safely accessing data
func NewNode() *node {
	return &node{store: make(map[string][]byte), mu: sync.RWMutex{}, lastAction: time.Now().UnixNano()}
}

func (i *portCollection) String() string {
	return "this is just used to satisfy the interface for a port 'Value'"
}

func (i *portCollection) Set(value string) error {
	*i = append(*i, value)
	return nil
}

// main launches the tcp gRPC server
func main() {
	flag.StringVar(&livePort, "livePort", "7000", "the port which the node lives on")
	flag.Var(&portList, "portList", "list of ports to be consumed by the node")
	flag.Parse()

	node := NewNode()

	var eg errgroup.Group
	eg.Go(func() error {
		if err := node.Start(livePort); err != nil {
			return fmt.Errorf("failed to start node on port %s: %s", livePort, err)
		}
		return nil
	})

	eg.Go(func() error {
		if err := node.FetchConnections(); err != nil {
			return fmt.Errorf("failed to fetch connections to other nodes in the system: %s", err)
		}
		return nil
	})

	if err := eg.Wait(); err != nil {
		fmt.Println(err)
		return
	}
}

// FetchConnections communicates from the node spawned by this file to other nodes in the portList, it first confirms
// a connection can be made. Once successful it then populates the network with grpc.ClientConns and KVClients.
func (s *node) FetchConnections() error {
	s.network = make([]nodeConnections, len(portList))

	for i, port := range portList {
		// dial the port to set create the client from this node (on 'livePort') to another node (on 'port')
		newConn, err := grpc.Dial("localhost:"+port, grpc.WithInsecure())
		if err != nil {
			return fmt.Errorf("failed to dial node %s from %s: %s", port, livePort, err)
		}

		// attempt 3 times to hit the Health endpoint of the KVClient on the node living on 'port', a 500ms sleep is allowed
		// between attempts, this should be sufficient to allow all nodes to start up and all connections in the network
		// to resolve
		for j := 0; j < 100; j++ {
			if _, err = pb.NewKVClient(newConn).Health(context.Background(), &pb.Empty{}); err != nil {
				time.Sleep(time.Millisecond * 200)
				continue
			}
			break
		}

		if err != nil {
			return fmt.Errorf("failed to hit health endpoint of port %s from port%s: %s", port, livePort, err)
		}

		// having succeeded, the network line is populated with the grpc.ClientConn and the KVClient
		s.network[i].conn = newConn
		s.network[i].kvClient = pb.NewKVClient(newConn)
	}

	return nil
}

// Start launches the node on its port ready to listen and serve for external calls and internal calls from other nodes
// within the network.
func (s *node) Start(livePort string) error {
	newServer := grpc.NewServer()
	pb.RegisterKVServer(newServer, s)
	lis, err := net.Listen("tcp", ":"+livePort)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	fmt.Printf("launching Node on port: %s\n", livePort)
	if err := newServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

	return nil
}

// Get will return a value given a key or an error if unsuccessful
func (s *node) Get(ctx context.Context, key *pb.Request) (*pb.ValueReturn, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	if _, ok := s.store[key.Key]; !ok {
		errMsg := fmt.Sprintf("the requested key '%s' is not present in the store\n", key.Key)
		return nil, status.Error(codes.NotFound, errMsg)
	}

	returnedVal := s.store[key.Key]
	return &pb.ValueReturn{Value: returnedVal}, nil
}

// Post will update the map with a key and value, it will return nothing but an error upon an unsuccessful request,
// this also acts as a traditional CRUD 'PUT' request
func (s *node) Post(ctx context.Context, request *pb.PostRequest) (*pb.Empty, error) {
	// this first check is the case in which this is an internal request (i.e. we're receiving a Post from another node in the system)
	// then we need to verify (using timestamps) if we've already received this request before
	if !request.External && request.Timestamp < s.lastAction {
		return nil, fmt.Errorf("This is a conflicting Post request to node on port %s.\nThe timestamp of this request is "+
			"%d and the timestamp of the last action was %d.", livePort, request.Timestamp, s.lastAction)
	}

	// add the key/value pair to the store
	s.mu.Lock()
	s.store[request.Key] = request.Value
	s.mu.Unlock()

	// If the Post request originated from an external client, it needs to be sent on to the rest of the nodes in the network.
	// Note: it is sent on with 'External = false' as to avoid relaying messages around the network infinitely
	if request.External {
		for _, node := range s.network {
			_, err := node.kvClient.Post(ctx, &pb.PostRequest{Key: request.Key, Value: []byte(request.Value), External: false, Timestamp: request.Timestamp})
			if err != nil {
				return nil, fmt.Errorf("failed to post/put to store with key %s and value %s: %s", request.Key, request.Value, err)
			}
		}
	}

	return &pb.Empty{}, nil
}

// Delete removes a key and value from the map, it will return nothing but an error upon an unsuccessful request
func (s *node) Delete(ctx context.Context, request *pb.Request) (*pb.Empty, error) {
	// this first check is the case in which this is an internal request (i.e. we're receiving a Delete from another node in the system)
	// then we need to verify (using timestamps) if we've already received this request before
	if !request.External && request.Timestamp < s.lastAction {
		return nil, fmt.Errorf("This is a conflicting Delete request to node on port %s.\nThe timestamp of this request is "+
			"%d and the timestamp of the last action was %d.", livePort, request.Timestamp, s.lastAction)
	}

	// delete the key/value pair from the store
	s.mu.RLock()
	delete(s.store, request.Key)
	s.mu.RUnlock()

	// If the Delete request originated from an external client, it needs to be sent on to the rest of the nodes in the network.
	// Note: it is sent on with 'External = false' as to avoid relaying messages around the network infinitely
	if request.External {
		for _, node := range s.network {
			_, err := node.kvClient.Delete(ctx, &pb.Request{Key: request.Key, External: false, Timestamp: request.Timestamp})
			if err != nil {
				return nil, fmt.Errorf("failed to delete value associated to key %s from store: %s", request.Key, err)
			}
		}
	}

	return &pb.Empty{}, nil
}

// List will return all keys and values in the map, it will return an error upon an unsuccessful request
func (s *node) List(_ *pb.Empty, stream pb.KV_ListServer) error {
	s.mu.RLock()
	defer s.mu.RUnlock()
	for _, v := range s.store {
		if err := stream.Send(&pb.ValueReturn{Value: v}); err != nil {
			return err
		}
	}

	return nil
}

// Health returns nothing and is used to confirm nodes can communicate with each other within the network
func (s *node) Health(_ context.Context, _ *pb.Empty) (*pb.Empty, error) {
	return &pb.Empty{}, nil
}
