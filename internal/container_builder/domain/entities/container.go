package entities

import (
	"log/slog"

	"github.com/deniskarpenko/composebuilder/internal/container_builder/domain/valueobjects"
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
