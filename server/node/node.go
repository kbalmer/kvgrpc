package node

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	pb "kvgrpc/kv"
	"log"
	"net"
	"sync"
	"time"
)

// the server used to implement kv.KVServer
type node struct {
	store      map[string][]byte
	mu         sync.RWMutex
	network    []nodeConnections
	lastAction int64
	livePort string
	portList []string
}

// nodeConnections stores
type nodeConnections struct {
	conn     *grpc.ClientConn
	kvClient pb.KVClient
}

// NewNode returns a blank node with the map for KV storage and the read-write mutex for safely accessing data
func NewNode(livePort string, portList []string) *node {
	return &node{
			store: make(map[string][]byte),
			mu: sync.RWMutex{},
			lastAction: time.Now().UnixNano(),
			livePort: livePort,
			portList: portList,
			network: make([]nodeConnections, len(portList)),
	}
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

// FetchConnections communicates from the node spawned by this file to other nodes in the portList, it first confirms
// a connection can be made. Once successful it then populates the network with grpc.ClientConns and KVClients.
func (s *node) FetchConnections() error {
	for i, port := range s.portList {
		// dial the port to set create the client from this node (on 'livePort') to another node (on 'port')
		newConn, err := grpc.Dial("localhost:"+port, grpc.WithInsecure())
		if err != nil {
			return fmt.Errorf("failed to dial node %s from %s: %s", port, s.livePort, err)
		}

		// attempt to hit the Health endpoint of the KVClient on the node living on 'port', a total of 20 seconds with 100
		// attempts is deemed sufficient in order for all ports and nodes to be listening and serving
		for j := 0; j < 100; j++ {
			if _, err = pb.NewKVClient(newConn).Health(context.Background(), &pb.Empty{}); err != nil {
				time.Sleep(time.Millisecond * 200)
				continue
			}
			break
		}

		if err != nil {
			return fmt.Errorf("failed to hit health endpoint of port %s from port%s: %s", port, s.livePort, err)
		}

		// having succeeded, the network line is populated with the grpc.ClientConn and the KVClient
		s.network[i].conn = newConn
		s.network[i].kvClient = pb.NewKVClient(newConn)
	}

	return nil
}

// Health returns nothing and is used to confirm nodes can communicate with each other within the network
func (s *node) Health(_ context.Context, _ *pb.Empty) (*pb.Empty, error) {
	return &pb.Empty{}, nil
}
