package services

import (
	"github.com/deniskarpenko/composebuilder/internal/container_builder/application/dto"
	"github.com/deniskarpenko/composebuilder/internal/container_builder/application/interfaces"
)

type ContainerApplicationService struct {
	logger interfaces.Logger
}

func NewContainerApplicationService(log interfaces.Logger) *ContainerApplicationService {
	return &ContainerApplicationService{
		logger: log,
	}
}

func (s *ContainerApplicationService) CreateApplication(app dto.Application) {

}
