package main

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/deniskarpenko/composebuilder/internal/container_builder/domain/entities"
)

func main() {
	fmt.Println("Starting up compose builder ...")

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	cb := entities.NewContainerBuilder("PHP", logger)
}
