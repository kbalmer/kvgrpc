package main

import (
	"bufio"
	"context"
	"flag"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
	"io"
	pb "kvgrpc/kv"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

var (
	firstPort int
	numOfPorts int
)

// client is the type which holds all properties about the connection between this client and the node it's connected to
type client struct {
	conn     *grpc.ClientConn
	port int
	kvClient pb.KVClient
	latency time.Duration
}

func main() {
	flag.IntVar(&firstPort, "nodeConnectionPort", 7000, "the first possible port the client could connect to")
	flag.IntVar(&numOfPorts, "numOfPorts", 10, "the number of nodes in the network which can be connected to")
	flag.Parse()

	var clients []client

	// before creating a client, all possible nodes will be health checked in the network to find the one with the lowest latency
	for i := firstPort; i < firstPort+numOfPorts; i++ {
		c, err := NewClient(i)
		if err != nil {
			log.Fatalf("failed to create client: %s", err)
		}

		start := time.Now()
		if _, err := c.kvClient.Health(context.Background(), &pb.Empty{}); err != nil {
			return
		}

		c.latency = time.Since(start)
		clients = append(clients, *c)
	}

	sort.Slice(clients, func(i, j int) bool {
		return clients[i].latency < clients[j].latency
	})

	bestClient := clients[0]

	fmt.Printf("==========================================================================================\n"+
		"	This is a CRUD system. You are connected to the system on port: [%d]\n		Use: Get, Post, Put or Delete to access the store.\n"+
		"	Follow the appropriate command with space separated keys and/or values.\n"+
		"==========================================================================================\n", bestClient.port)

	bestClient.inputParser()
}

// NewClient defines a new grpc.ClientConn, KVClient and the port it's connected via
func NewClient(port int) (*client, error) {
	newConn, err := grpc.Dial("localhost:"+strconv.Itoa(port), grpc.WithInsecure())
	if err != nil {
		fmt.Printf("did not connect: %v", err)
	}

	newClient := pb.NewKVClient(newConn)

	return &client{conn: newConn, kvClient: newClient, port: port}, nil
}

// inputParser takes the os.Stdin to deduce what CRUD command to call on the key-value store
func (client *client) inputParser() {
	reader := bufio.NewReader(os.Stdin)
	for {
		newLine, err := reader.ReadString('\n')
		if err != nil {
			fmt.Printf("unable to read from stdin: %s\n", err)
			return
		}

		items := strings.Split(newLine, " ")

		for i := range items {
			items[i] = strings.TrimSpace(items[i])
		}

		command := strings.ToUpper(items[0])
		switch len(items) {
		case 1:
			if command == "EXIT" {
				os.Exit(0)
			}
			if command != "LIST" {
				fmt.Printf("Invalid request; please use Get, Post, Put or Delete...\n")
				continue
			}
			client.lister()
		case 2:
			if command == "GET" {
				client.getter(items[1])
			}
			if command == "DELETE" {
				client.deleter(items[1])
			}
		case 3:
			if command == "POST" || command == "PUT" {
				client.poster(items[1], items[2])
			}
		default:
			fmt.Printf("request invalid, please refer to CRUD commands")
			continue
		}
	}
}

// lister is the client method for retrieving the entire key-value store
func (client *client) lister() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
	defer cancel()

	stream, err := client.kvClient.List(ctx, &pb.Empty{})
	if err != nil {
		log.Printf("failed to initiate stream: %s", err)
		return
	}

	for {
		item, err := stream.Recv()
		if err == io.EOF {
			break
		}

		if err != nil {
			log.Printf("failed to stream values from store: %s", err)
			return
		}

		fmt.Printf("%+v\n", string(item.Value))
	}
}

// getter is the client method for retrieving a value from the key-value store given a specific key
func (client *client) getter(key string) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
	defer cancel()

	val, err := client.kvClient.Get(ctx, &pb.Request{Key: key})
	if err != nil {
		st, ok := status.FromError(err)
		if !ok {
			fmt.Print(err)
			return
		}
		fmt.Printf("%s[failure code: %d]\n", st.Message(), st.Code())
		return
	}

	fmt.Printf("%s\n", string(val.Value))
}

// deleter is the client method for removing a value from the key-value store given a specific key
func (client *client) deleter(key string) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
	defer cancel()

	_, err := client.kvClient.Delete(ctx, &pb.Request{Key: key, External: true, Timestamp: time.Now().UnixNano()})
	if err != nil {
		fmt.Printf("failed to delete value associated to key %s from store: %s", key, err)
		return
	}

	fmt.Printf("successfully deleted '%s' from the store\n", key)
}

// poster is the client method for adding a value to the key-value store with a corresponding key
func (client *client) poster(key string, value string) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
	defer cancel()

	_, err := client.kvClient.Post(ctx, &pb.PostRequest{Key: key, Value: []byte(value), External: true, Timestamp: time.Now().UnixNano()})
	if err != nil {
		fmt.Printf("failed to post/put to store with key %s and value %s: %s", key, value, err)
		return
	}

	fmt.Printf("successfully post/put key '%s' with value '%s' to store\n", key, value)
}
