package main

import (
	"context"
	"fmt"
	"github.com/ITRI-ICL-Peregrine/x-tracer/pkg"
	"github.com/docker/docker/client"
	"log"
	"os"
	"strings"
	"time"
)


func main() {

	log.Println("Start api...")

	containerId := os.Getenv("containerId")

	if containerId == "" {
		containerId = "ec9515bb14a2"
	}

	serverIp := os.Getenv("masterIp")
	if containerId == "" {
		containerId = "ec9515bb14a2"
	}

	probeName := os.Getenv("tools")


	Probe := strings.Split(probeName, ",")
	cli, err := client.NewClientWithOpts(client.WithHost("unix:///var/run/docker.sock"), client.WithAPIVersionNegotiation())
	if err != nil {
		panic(err)
	}
	topResult, err := cli.ContainerTop(context.Background(), containerId, []string{"o", "pid"})
	if err != nil {
		panic(err)
	}
	fmt.Println(topResult.Processes)

	log.Printf("Start new client")

	testClient := pkg.New("6666", serverIp)
	testClient.StartClient(Probe, topResult.Processes)

	for {
		fmt.Println("x-agent - Sleeping")
		time.Sleep(10 * time.Second)
	}

}