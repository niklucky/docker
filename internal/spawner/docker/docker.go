package docker

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
)

type docker struct {
	config *Config
	cli    *client.Client
}

type CreateContainerOptions struct {
	Name       string
	Config     *container.Config
	HostConfig *container.HostConfig
	Start      bool
}

type Controller interface {
	CreateContainer(options CreateContainerOptions) (container.ContainerCreateCreatedBody, error)
	ListContainers() error
}

func New(config *Config) Controller {
	return &docker{
		config: config,
	}
}

func (d *docker) init() {
	var err error
	d.cli, err = client.NewEnvClient()
	if err != nil {
		log.Fatalln("Unable to create docker client")
	}
}
func (d *docker) ListContainers() error {
	d.init()

	containers, err := d.cli.ContainerList(context.Background(), types.ContainerListOptions{})
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

func (d *docker) CreateContainer(options CreateContainerOptions) (container.ContainerCreateCreatedBody, error) {
	d.init()
	cont, err := d.cli.ContainerCreate(
		context.Background(),
		d.parseConfig(options.Config),
		d.parseHostConfig(options.HostConfig),
		&network.NetworkingConfig{},
		nil,
		d.getContainerName(options.Name),
	)
	if err != nil {
		panic(err)
	}
	if options.Start {
		d.ContainerStart(cont.ID)
	}
	return cont, nil
}

func (d *docker) ContainerStart(containerID string) {
	d.init()
	d.cli.ContainerStart(context.Background(), containerID, types.ContainerStartOptions{})
	fmt.Printf("Container %s is started", containerID)
}

func (d *docker) parseConfig(config *container.Config) *container.Config {
	if config == nil {
		config = &container.Config{}
	}
	if config.Image == "" {
		if d.config.Image == "" {
			log.Fatalln("You need to specify Image in Spawner config or in ContainerCreate")
		}
		config.Image = d.config.Image
	}
	return config
}

func (d *docker) parseHostConfig(hostConfig *container.HostConfig) *container.HostConfig {
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
	if d.config.AutoRemove && !hostConfig.AutoRemove {
		hostConfig.AutoRemove = d.config.AutoRemove
	}
	return hostConfig
	// return &container.HostConfig{}
}

func (d *docker) getContainerName(name string) string {
	a := []string{"spawned_"}
	if d.config.ContainerName != "" {
		a = append(a, d.config.ContainerName)
	}
	if name != "" {
		a = append(a, name)
	}
	a = append(a, d.generateID())
	return strings.Join(a, "_")
}

func (d *docker) generateID() string {
	id, _ := gonanoid.New()
	return id
}
