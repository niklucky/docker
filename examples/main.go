package main

import (
	"fmt"

	"github.com/niklucky/docker/spawner"
)

func main() {
	s := spawner.New(&spawner.SpawnerOptions{
		Image:      "postgres:13",
		AutoRemove: true,
		// API: &api.Config{
		// 	Address: ":8081",
		// },
	})
	// Starting
	fmt.Println("Hello")
	s.Start()

	// s.CreateContainer(spawner.CreateContainerOptions{
	// 	Start: true,
	// 	Name:  "test",
	// })
	// s.ListContainer()
}
