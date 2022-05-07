// package spawner
package spawner

import (
	"github.com/niklucky/docker/internal/spawner/api"
	"github.com/niklucky/docker/internal/spawner/docker"
)

// Spawner - struct contains common info for spawning containers
type Spawner struct {
	options *Options
	api     *api.Server
	docker  docker.Controller
}

type Options struct {
	API    *api.Config
	Docker *docker.Config
}

func New(options *Options) *Spawner {
	s := &Spawner{
		options: options,
	}
	s.docker = docker.New(options.Docker)
	s.api = api.New(options.API, s.docker)
	return s
}

func (s *Spawner) Start() error {
	return s.api.Start()
}
