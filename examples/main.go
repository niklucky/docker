package main

import (
	"github.com/niklucky/docker/internal/spawner/docker"
	"github.com/niklucky/docker/spawner"
)

func main() {
	s := spawner.New(&spawner.Config{
		Docker: &docker.Config{
			Image:      "test",
			AutoRemove: false,
		},
	})
	// Starting
	s.Start()
}
