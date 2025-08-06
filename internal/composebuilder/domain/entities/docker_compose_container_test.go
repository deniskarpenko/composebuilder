package entities

import (
	"log/slog"
	"os"
	"testing"
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
