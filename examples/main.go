package main

import (
	"fmt"

	"github.com/niklucky/docker/spawner"
)

func main() {
	s := spawner.New(&spawner.SpawnerOptions{
		Image:      "postgres:13",
		AutoRemove: true,
	})
	s.CreateContainer(spawner.CreateContainerOptions{
		Start: true,
		Name:  "test",
	})
	s.ListContainer()
	fmt.Println("Hello")
}
