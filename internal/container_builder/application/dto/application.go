package dto

import (
	"github.com/deniskarpenko/composebuilder/internal/container_builder/application/interfaces"
)

type Application struct {
	logger   interfaces.Logger
	projects []Project
}
