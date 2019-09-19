package main

import (
	"flag"
	"fmt"
	"golang.org/x/sync/errgroup"
	"kvgrpc/server/node"
)

var (
	// postList is the list of ports to be consumed by the node

)

// portCollection is used to store multiple ports to be consumed as a flag for the node connections
type portCollection []string


func (i *portCollection) String() string {
	return "this is just used to satisfy the interface for a port 'Value'"
}

func (i *portCollection) Set(value string) error {
	*i = append(*i, value)
	return nil
}

// main launches the tcp gRPC server
func main() {
	var portList portCollection
	var livePort string

	flag.StringVar(&livePort, "livePort", "7000", "the port which the node lives on")
	flag.Var(&portList, "portList", "list of ports to be consumed by the node")
	flag.Parse()

	blankNode := node.NewNode(livePort, portList)

	var eg errgroup.Group
	eg.Go(func() error {
		if err := blankNode.Start(livePort); err != nil {
			return fmt.Errorf("failed to start node on port %s: %s", livePort, err)
		}
		return nil
	})

	eg.Go(func() error {
		if err := blankNode.FetchConnections(); err != nil {
			return fmt.Errorf("failed to fetch connections to other nodes in the system: %s", err)
		}
		return nil
	})

	if err := eg.Wait(); err != nil {
		fmt.Println(err)
		return
	}
}