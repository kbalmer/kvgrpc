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
	"strings"
	"time"
)

var (
	nodeConnectionPort = "7000"
)

type client struct {
	conn     *grpc.ClientConn
	kvClient pb.KVClient
}

func main() {
	flag.StringVar(&nodeConnectionPort, "nodeConnectionPort", "7000", "the port of the node which the client will connect to")
	flag.Parse()
	// take a port value to point at a specific node... check it can connect etc
	c, err := NewClient()
	if err != nil {
		log.Fatalf("failed to create client: %s", err)
	}
	defer c.conn.Close()

	reader := bufio.NewReader(os.Stdin)

	fmt.Printf("==========================================================================================\n"+
		"	This is a CRUD system. You are connect to the system on port: [%s].\n		Use: Get, Post, Put or Delete to access the store.\n"+
		"	Follow the appropriate command with space separated keys and/or values.\n"+
		"==========================================================================================\n", nodeConnectionPort)

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
			if command != "LIST" {
				fmt.Printf("Invalid request; please use Get, Post, Put or Delete...\n")
				continue
			}
			c.lister()
		case 2:
			if command == "GET" {
				c.getter(items[1])
			}
			if command == "DELETE" {
				c.deleter(items[1])
			}
		case 3:
			if command == "POST" || command == "PUT" {
				c.poster(items[1], items[2])
			}
		default:
			fmt.Printf("request invalid, please refer to CRUD commands")
			continue
		}
	}
}

func NewClient() (*client, error) {
	newConn, err := grpc.Dial("localhost:"+nodeConnectionPort, grpc.WithInsecure())
	if err != nil {
		fmt.Printf("did not connect: %v", err)
	}

	newClient := pb.NewKVClient(newConn)

	return &client{conn: newConn, kvClient: newClient}, nil
}

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

func (client *client) poster(key string, value string) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	_, err := client.kvClient.Post(ctx, &pb.PostRequest{Key: key, Value: []byte(value), External: true, Timestamp: time.Now().UnixNano()})
	if err != nil {
		fmt.Printf("failed to post/put to store with key %s and value %s: %s", key, value, err)
		return
	}

	fmt.Printf("successfully post/put key '%s' with value '%s' to store\n", key, value)
}
