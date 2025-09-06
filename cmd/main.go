package main

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/deniskarpenko/composebuilder/internal/container_builder/application/services"
)

func main() {
	fmt.Println("Starting up compose builder ...")

	logger := services.NewSlogAdapter(slog.New(slog.NewTextHandler(os.Stdout, nil)))

	service := services.NewContainerApplicationService(logger) // Remove the & here

	fmt.Print(service)

}
