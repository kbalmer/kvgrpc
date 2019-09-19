package node


import (
	"context"
	"fmt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	pb "kvgrpc/kv"
)

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
			"%d and the timestamp of the last action was %d.", s.livePort, request.Timestamp, s.lastAction)
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
				return nil, fmt.Errorf("failed to post/put to store with key %s and value %s: %s\n", request.Key, request.Value, err)
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
			"%d and the timestamp of the last action was %d.", s.livePort, request.Timestamp, s.lastAction)
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
				return nil, fmt.Errorf("failed to delete value associated to key %s from store: %s\n", request.Key, err)
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