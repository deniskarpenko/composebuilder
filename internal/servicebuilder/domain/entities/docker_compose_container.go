package entities

import (
	"log/slog"

	"github.com/deniskarpenko/composebuilder/internal/servicebuilder/domain/valueobjects"
)

type Container struct {
	containerName string
	image         *valueobjects.Image
	build         *valueobjects.Build
	ports         []valueobjects.Ports
	volumes       []valueobjects.Volume
	envs          []valueobjects.Env
	envFiles      []string
	networks      []string
	dependsOn     []string
	logger        *slog.Logger
}

func (c *Container) Logger() *slog.Logger {
	return c.logger
}

type ContainerBuilder struct {
	container Container
}

func NewContainerBuilder(containerName string, logger *slog.Logger) *ContainerBuilder {
	return &ContainerBuilder{
		container: Container{
			containerName: containerName,
			logger:        logger,
			ports:         make([]valueobjects.Ports, 0),
			volumes:       make([]valueobjects.Volume, 0),
			envs:          make([]valueobjects.Env, 0),
			envFiles:      make([]string, 0),
			networks:      make([]string, 0),
			dependsOn:     make([]string, 0),
		},
	}
}

func (cb *ContainerBuilder) WithImage(image *valueobjects.Image) *ContainerBuilder {
	cb.container.image = image
	return cb
}

func (cb *ContainerBuilder) WithBuild(build *valueobjects.Build) *ContainerBuilder {
	cb.container.build = build
	return cb
}

func (cb *ContainerBuilder) WithPorts(ports ...valueobjects.Ports) *ContainerBuilder {
	cb.container.ports = ports
	return cb
}

func (cb *ContainerBuilder) WithVolumes(volumes ...valueobjects.Volume) *ContainerBuilder {
	cb.container.volumes = volumes
	return cb
}

func (cb *ContainerBuilder) WithEnvs(envs ...valueobjects.Env) *ContainerBuilder {
	cb.container.envs = envs
	return cb
}

func (cb *ContainerBuilder) WithEnvFiles(envFiles ...string) *ContainerBuilder {
	cb.container.envFiles = envFiles
	return cb
}

func (cb *ContainerBuilder) WithNetworks(networks ...string) *ContainerBuilder {
	cb.container.networks = networks
	return cb
}

func (cb *ContainerBuilder) WithDependencies(dependencies ...string) *ContainerBuilder {
	cb.container.dependsOn = dependencies
	return cb
}

func (cb *ContainerBuilder) Build() Container {
	return cb.container
}
