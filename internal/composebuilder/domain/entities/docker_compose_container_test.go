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

}
