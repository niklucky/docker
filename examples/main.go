package main

import (
	"github.com/niklucky/docker/internal/spawner/docker"
	"github.com/niklucky/docker/spawner"
)

func main() {
	s := spawner.New(&spawner.Config{
		Docker: &docker.Config{
			Image:      "nginx",
			AutoRemove: true,
			// API: &api.Config{
			// 	Address: ":8081",
			// },
		},
	})
	// Starting
	s.Start()
}
