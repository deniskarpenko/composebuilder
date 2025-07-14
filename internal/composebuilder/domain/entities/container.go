package entities

import "github.com/deniskarpenko/composebuilder/internal/composebuilder/domain/valueobjects"

type Container struct {
	containerName string
	image         *valueobjects.Image
	build         *valueobjects.Build
	ports         *[]valueobjects.Ports
	volumes       *[]valueobjects.Volume
	envs          *[]valueobjects.Env
	envFiles      *[]string
	networks      *[]string
	dependsOn     *[]string
}
