package docker

type Config struct {
	ContainerName string
	Image         string
	AutoRemove    bool
}
