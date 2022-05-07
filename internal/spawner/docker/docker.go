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
	Start      bool
	Image      string
	Name       string
	AutoRemove bool `json:"autoRemove"`
	Env        []string
	Ports      []string
}

type Controller interface {
	ContainerList() ([]types.Container, error)
	ContainerCreate(options *CreateContainerOptions) (container.ContainerCreateCreatedBody, error)
	ContainerStart(containerID string) error
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
func (d *docker) ContainerList() ([]types.Container, error) {
	d.init()

	result := make([]types.Container, 0)
	containers, err := d.cli.ContainerList(context.Background(), types.ContainerListOptions{})
	for i := 0; i < len(containers); i++ {
		var cnt = containers[i]
		if cnt.Image == d.config.Image {
			result = append(result, cnt)
		}
	}
	return result, err
}

func (d *docker) ContainerCreate(options *CreateContainerOptions) (container.ContainerCreateCreatedBody, error) {
	d.init()
	fmt.Printf("%v", options)
	cont, err := d.cli.ContainerCreate(
		context.Background(),
		d.getContainerConfig(options),
		d.getContainerHostConfig(options),
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

func (d *docker) ContainerStart(containerID string) error {
	d.init()
	return d.cli.ContainerStart(context.Background(), containerID, types.ContainerStartOptions{})
}

func (d *docker) getContainerConfig(options *CreateContainerOptions) *container.Config {
	config := &container.Config{
		Image: options.Image,
		Env:   options.Env,
	}
	if config.Image == "" {
		if d.config.Image == "" {
			log.Fatalln("You need to specify Image in Spawner config or in ContainerCreate")
		}
		config.Image = d.config.Image
	}
	return config
}

func (d *docker) getContainerHostConfig(options *CreateContainerOptions) *container.HostConfig {
	hostConfig := &container.HostConfig{
		AutoRemove: options.AutoRemove,
	}
	if len(options.Ports) > 0 {
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
	}
	if d.config.AutoRemove && !hostConfig.AutoRemove {
		hostConfig.AutoRemove = d.config.AutoRemove
	}
	return hostConfig
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
