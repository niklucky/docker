package main

import (
	"github.com/niklucky/docker/pkg/spawner/docker"
	"github.com/niklucky/docker/spawner"
)

func main() {
	s := spawner.New(&spawner.Config{
		Docker: &docker.Config{
			Image:      "mossport_server",
			AutoRemove: false,
		},
	})
	// Starting
	s.Start()
}
