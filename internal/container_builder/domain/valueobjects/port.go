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

func (p Ports) ToYaml() ([]byte, error) {
	if !p.HasContainerPort() {
		return []byte(fmt.Sprintf("%d", p.host)), nil
	}

	return []byte(fmt.Sprintf("%d:%d", p.host, p.containerPort)), nil
}
