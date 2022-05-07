// package spawner
package spawner

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/network"
	"github.com/docker/docker/client"
	"github.com/docker/go-connections/nat"
	gonanoid "github.com/matoous/go-nanoid/v2"
	"github.com/niklucky/docker/internal/spawner/api"
)

// Spawner - struct contains common info for spawning containers
type spawner struct {
	options *SpawnerOptions
	cli     *client.Client
	api     *api.Server
}

type CreateContainerOptions struct {
	Name       string
	Config     *container.Config
	HostConfig *container.HostConfig
	Start      bool
}
type Spawner interface {
	Start() error
	CreateContainer(options CreateContainerOptions) (container.ContainerCreateCreatedBody, error)
	ListContainer() error
}

type SpawnerOptions struct {
	ContainerName string
	Image         string
	AutoRemove    bool
	API           *api.Config
}

func New(options *SpawnerOptions) Spawner {
	s := &spawner{
		options: options,
	}
	s.api = api.New(options.API)
	return s
}

func (s *spawner) Start() error {
	return s.api.Start()
}

func (s *spawner) init() {
	var err error
	s.cli, err = client.NewEnvClient()
	if err != nil {
		log.Fatalln("Unable to create docker client")
	}
}
func (s *spawner) ListContainer() error {
	s.init()

	containers, err := s.cli.ContainerList(context.Background(), types.ContainerListOptions{})
	if err != nil {
		panic(err)
	}

	if len(containers) > 0 {
		for _, container := range containers {
			fmt.Printf("Container ID: %s", container.ID)
		}
	} else {
		fmt.Println("There are no containers running")
	}
	return nil
}

func (s *spawner) CreateContainer(options CreateContainerOptions) (container.ContainerCreateCreatedBody, error) {
	s.init()
	cont, err := s.cli.ContainerCreate(
		context.Background(),
		s.parseConfig(options.Config),
		s.parseHostConfig(options.HostConfig),
		&network.NetworkingConfig{},
		nil,
		s.getContainerName(options.Name),
	)
	if err != nil {
		panic(err)
	}
	if options.Start {
		s.ContainerStart(cont.ID)
	}
	return cont, nil
}

func (s *spawner) ContainerStart(containerID string) {
	s.init()
	s.cli.ContainerStart(context.Background(), containerID, types.ContainerStartOptions{})
	fmt.Printf("Container %s is started", containerID)
}

func (s *spawner) parseConfig(config *container.Config) *container.Config {
	if config == nil {
		config = &container.Config{}
	}
	if config.Image == "" {
		if s.options.Image == "" {
			log.Fatalln("You need to specify Image in Spawner options or in ContainerCreate")
		}
		config.Image = s.options.Image
	}
	return config
}

func (s *spawner) parseHostConfig(hostConfig *container.HostConfig) *container.HostConfig {
	if hostConfig == nil {
		hostConfig = &container.HostConfig{}
	}
	hostBinding := nat.PortBinding{
		HostIP:   "0.0.0.0",
		HostPort: "8000",
	}
	containerPort, err := nat.NewPort("tcp", "80")
	if err != nil {
		panic("Unable to get the port")
	}

	portBinding := nat.PortMap{containerPort: []nat.PortBinding{hostBinding}}
	hostConfig.PortBindings = portBinding
	if s.options.AutoRemove && !hostConfig.AutoRemove {
		hostConfig.AutoRemove = s.options.AutoRemove
	}
	return hostConfig
	// return &container.HostConfig{}
}

func (s *spawner) getContainerName(name string) string {
	a := []string{"spawned_"}
	if s.options.ContainerName != "" {
		a = append(a, s.options.ContainerName)
	}
	if name != "" {
		a = append(a, name)
	}
	a = append(a, s.generateID())
	return strings.Join(a, "_")
}

func (s *spawner) generateID() string {
	id, _ := gonanoid.New()
	return id
}
