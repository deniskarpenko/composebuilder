package entities

import (
	"log/slog"
	"os"
	"testing"

	"github.com/deniskarpenko/composebuilder/internal/servicebuilder/domain/valueobjects"
)

func TestNewContainerBuilder(t *testing.T) {
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	containerName := "test container"

	builder := NewContainerBuilder(containerName, logger)

	if builder == nil {
		t.Fatal("Expected builder to be non-nil")
	}

	if builder.container.containerName != containerName {
		t.Fatal("Expected container.containerName to be ", containerName)
	}

	if builder.container.Logger() != logger {
		t.Fatal("Expected container.logger to be ", logger)
	}

	if builder.container.ports == nil || len(builder.container.ports) != 0 {
		t.Fatal("Expected container.ports to be non-nil and len 0")
	}

	if builder.container.volumes == nil || len(builder.container.volumes) != 0 {
		t.Fatal("Expected container.volumes to be non-nil and len 0")
	}

	if builder.container.envs == nil || len(builder.container.envs) != 0 {
		t.Fatal("Expected container.envs to be non-nil and len 0")
	}

	if builder.container.envFiles == nil || len(builder.container.envFiles) != 0 {
		t.Fatal("Expected container.envFiles to be non-nil and len 0")
	}

}

func TestContainerBuilderCompleteFlow(t *testing.T) {
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	containerName := "Complete-Container"

	nginxImage, err := valueobjects.NewImage("nginx", "latest")

	if err != nil {
		t.Fatal("Failed to create image object")
	}
	build := valueobjects.NewBuild("./dockerfiles/", "nginx.Dockerfile")
	ports := valueobjects.NewPorts(80, func() *int { i := 80; return &i }())
	volumes := valueobjects.NewVolume("./app/", "/var/www/html/")

	envs, err := valueobjects.NewEnv("LOG_LEVEL", "DEBUG")

	if err != nil {
		t.Fatal("Failed to create env object")
	}

	envFiles := []string{
		"/envs/test.env",
	}

	networks := []string{
		"test_network",
	}

	dependsOn := []string{
		"mysql",
	}

	builder := NewContainerBuilder(containerName, logger)

	builder = builder.WithImage(&nginxImage).WithBuild(&build).WithPorts(ports).WithVolumes(volumes).WithEnvs(envs)
	builder = builder.WithEnvFiles(envFiles...).WithNetworks(networks...).WithDependencies(dependsOn...)

}
