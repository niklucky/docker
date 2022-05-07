package main

import (
	"github.com/niklucky/docker/internal/spawner/docker"
	"github.com/niklucky/docker/spawner"
)

func main() {
	s := spawner.New(&spawner.Options{
		Docker: &docker.Config{
			Image:      "postgres:13",
			AutoRemove: true,
			// API: &api.Config{
			// 	Address: ":8081",
			// },
		},
	})
	// Starting
	s.Start()
}
