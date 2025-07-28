package entities

import (
	"log/slog"

	"github.com/deniskarpenko/composebuilder/internal/composebuilder/domain/valueobjects"
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
