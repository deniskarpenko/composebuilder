package valueobjects

import (
	"fmt"
)

type Ports struct {
	host          int
	containerPort *int
}

func NewPorts(host int, containerPort *int) Ports {
	return Ports{host, containerPort}
}

func (p Ports) Host() int {
	return p.host
}

func (p Ports) ContainerPort() *int {
	return p.containerPort
}

func (p Ports) HasContainerPort() bool {
	return p.containerPort != nil
}

func (p Ports) ToYaml() string {
	if p.HasContainerPort() {
		return fmt.Sprintf("%d", p.host)
	}

	return fmt.Sprintf("%d:%d", p.host, p.containerPort)
}
