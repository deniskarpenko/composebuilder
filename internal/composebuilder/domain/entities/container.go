package entities

import "github.com/deniskarpenko/composebuilder/internal/composebuilder/domain/valueobjects"

type Container struct {
	containerName string
	image         *valueobjects.Image
	build         *valueobjects.Build
	ports         *[]valueobjects.Ports
}
