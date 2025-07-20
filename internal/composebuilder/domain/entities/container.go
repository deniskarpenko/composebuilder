package entities

import (
	"log/slog"

	"github.com/deniskarpenko/composebuilder/internal/composebuilder/domain/valueobjects"
)

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
	logger        *slog.Logger
}

func NewContainer(
	containerName string,
	image *valueobjects.Image,
	build *valueobjects.Build,
	ports *[]valueobjects.Ports,
	volumes *[]valueobjects.Volume,
	envs *[]valueobjects.Env,
	envFiles *[]string,
	networks *[]string,
	dependsOn *[]string,
	logger *slog.Logger,
) Container {
	return Container{
		containerName: containerName,
		image:         image,
		build:         build,
		ports:         ports,
		volumes:       volumes,
		envs:          envs,
		envFiles:      envFiles,
		networks:      networks,
		dependsOn:     dependsOn,
		logger:        logger,
	}
}

func (c Container) ToYaml() ([]byte, error) {
	service := make(map[string]interface{})

	if c.image != nil {
		imageYaml, err := c.image.ToYaml()
	}

}

func (c Container) AddImageToService(service map[string]interface{}) error {
	if c.image == nil {
		return nil
	}

	imageYaml, err := c.image.ToYaml()

	if err != nil {
		return err
	}

	return nil
}
