package services

import (
	"fmt"
	"testing"

	"github.com/deniskarpenko/composebuilder/internal/container_builder/application/dto"
)

func TestContainerApplicationService_CreateApplication(t *testing.T) {
	network := "test"

	mySqlContainer := createMysqlContainer(network, []string{})
	phpContainer := createPhpContainer(network, []string{mySqlContainer.Name()})
	nginxContainer := createNginxContainer(network, []string{phpContainer.Name()})

	containers := []dto.Container{
		mySqlContainer,
		phpContainer,
		nginxContainer,
	}

	project := dto.NewProject("test-project", &containers)

	fmt.Println("project:", project)
}

func createMysqlContainer(network string, dependsOn []string) dto.Container {
	image := dto.NewImage("mysql", "latest")

	return dto.NewContainer(
		"mysql-db",
		&image,
		nil,
		[]dto.Port{dto.NewPort(3306, 3306)},
		nil,
		[]dto.Env{
			dto.NewEnv("MYSQL_ROOT_PASSWORD", "root"),
			dto.NewEnv("MYSQL_DATABASE", "db"),
			dto.NewEnv("MYSQL_USER", "user"),
			dto.NewEnv("MYSQL_PASSWORD", "pass"),
		},
		nil,
		[]string{network},
		dependsOn,
	)
}

func createPhpContainer(network string, dependsOn []string) dto.Container {
	phpBuild := dto.NewBuild("php.Dockerfile", "./dockerfiles")

	volumes := []dto.Volume{
		dto.NewVolume("./src", "/var/www/html"),
		dto.NewVolume("./php/php.ini", "/usr/local/etc/php/php.ini"),
	}

	return dto.NewContainer(
		"php",
		nil,
		&phpBuild,
		nil,
		volumes,
		nil,
		[]string{"envs/php.env"},
		[]string{network},
		dependsOn,
	)
}

func createNginxContainer(network string, dependsOn []string) dto.Container {
	image := dto.NewImage("nginx", "latest")
	return dto.NewContainer(
		"nginx",
		&image,
		nil,
		[]dto.Port{
			dto.NewPort(80, 80),
			dto.NewPort(443, 443),
		},
		[]dto.Volume{
			dto.NewVolume("./src", "/var/www/html"),
			dto.NewVolume("./nginx/nginx.conf", "/etc/nginx/nginx.conf"),
			dto.NewVolume("./nginx/sites/", "/etc/nginx/conf.d/"),
			dto.NewVolume("./nginx/ssl/", "/etc/nginx/ssl/"),
		},
		nil,
		nil,
		[]string{network},
		dependsOn,
	)
}
