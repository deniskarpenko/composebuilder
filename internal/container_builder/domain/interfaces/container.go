package interfaces

import (
	"github.com/deniskarpenko/composebuilder/internal/container_builder/domain/entities"
	"github.com/deniskarpenko/composebuilder/internal/container_builder/domain/valueobjects"
)

type ContainerBuilder interface {
	WithImage(image valueobjects.Image) ContainerBuilder
	WithBuild(build valueobjects.Build) ContainerBuilder
	WithPorts(ports ...valueobjects.Ports) ContainerBuilder
	WithVolumes(volumes ...valueobjects.Volume) ContainerBuilder
	WithEnvs(envs ...valueobjects.Env) ContainerBuilder
	WithEnvFiles(envFiles ...string) ContainerBuilder
	WithNetworks(networks ...string)
	WithDependencies(dependencies ...string) ContainerBuilder
	Build() entities.Container
}
