package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"sync"
)

var (
	// startingPort is the port at which nodes will be generated at, incrementing by 1 each time
	startingPort = 7000
	// nodeCount is the number of nodes to be spawned
	nodeCount = 5
)

func main() {
	flag.IntVar(&nodeCount, "nodeCount", nodeCount, "the number of nodes to be spawned")
	flag.IntVar(&startingPort, "startingPort", startingPort, "the port at which to start spawning the nodes at, increments by 1")
	flag.Parse()

	buildCmd := exec.Command("go", "build")
	buildCmd.Dir = "../server"
	if err := buildCmd.Run(); err != nil {
		fmt.Printf("failed to build server binary: %s", err)
		return
	}

	var wg sync.WaitGroup

	for i := startingPort; i < startingPort+nodeCount; i++ {

		var args []string
		args = append(args, fmt.Sprintf("-livePort=%d", i))

		for j := startingPort; j < startingPort+nodeCount; j++ {
			if i == j {
				continue
			}
			args = append(args, fmt.Sprintf("-portList=%d", j))
		}

		cmd := exec.Command("./server", args...)

		cmd.Dir = "../server"
		wg.Add(1)
		go executer(&wg, cmd)
	}
	wg.Wait()

}

func executer(wg *sync.WaitGroup, cmd *exec.Cmd) {
	defer wg.Done()
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		fmt.Printf("failed to execute binary with given flags: %s\n", err)
		return
	}
}
