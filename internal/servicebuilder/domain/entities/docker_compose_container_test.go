package entities

import (
	"fmt"
	"log/slog"
	"os"
	"testing"

	"github.com/deniskarpenko/composebuilder/internal/servicebuilder/domain/valueobjects"
)

func createImage(t *testing.T, imageName string, tag string) valueobjects.Image {
	image, err := valueobjects.NewImage(imageName, tag)

	if err != nil {
		t.Fatalf("could not create image: %v", err)
	}

	return image
}

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

func TestContainerBuilderWithImage(t *testing.T) {
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	containerName := "nginx-container"

	builder := NewContainerBuilder(containerName, logger)

	imageName := "nginx"

	tag := "latest"

	yamlData := fmt.Sprintf("%s:%s", imageName, tag)

	nginxImage := createImage(t, imageName, tag)

	builder.WithImage(&nginxImage)

	container := builder.Build()

	if container.image.ImageName() != nginxImage.ImageName() {
		t.Fatal("Expected container.image.ImageName to be ", nginxImage.ImageName())
	}

	if container.image.Tag() != nginxImage.Tag() {
		t.Fatal("Expected container.image.Tag to be ", nginxImage.Tag())
	}

	if string(container.image.ToYaml()) != string([]byte(yamlData)) {
		t.Fatal("Expected container.image.ToYaml to be ", yamlData)
	}

}

func TestContainerBuilderWithBuild(t *testing.T) {
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	containerName := "nginx-container"
	builder := NewContainerBuilder(containerName, logger)

	context := "."

	dockerfile := "php.Dockerfile"

	build := valueobjects.NewBuild(context, dockerfile)

	builder.WithBuild(&build)

	container := builder.Build()

	if container.build.Context() != build.Context() {
		t.Fatal("Expected container.build.Dockerfile to be ", build.Context())
	}

	if container.build.Dockerfile() != build.Dockerfile() {
		t.Fatal("Expected container.build.Dockerfile to be ", build.Dockerfile())
	}

}

func TestContainerBuilderCompleteFlow(t *testing.T) {
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	containerName := "complete-flow"

	image := createImage(t, "mysql", "latest")

	builder := NewContainerBuilder(containerName, logger)

	builder.WithImage(&image)
}
