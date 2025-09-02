package services

import (
	"testing"

	"github.com/deniskarpenko/composebuilder/internal/container_builder/application/dto"
)

func TestContainerApplicationService_CreateApplication(t *testing.T) {
}

func createPhpContainer() dto.Container {
	phpBuild := dto.NewBuild("php.Dockerfile", "./dockerfiles")

	return dto.NewContainer(
		"php",
		nil,
		&phpBuild,
		nil,
		nil,
		nil,
		nil,
		nil,
		nil,
	)
}
