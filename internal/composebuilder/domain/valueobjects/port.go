package valueobjects

type Ports struct {
	host          int
	containerPort int
}

func NewPorts(host, containerPort int) Ports {
	return Ports{host, containerPort}
}

func (p Ports) Host() int {
	return p.host
}

func (p Ports) ContainerPort() int {
	return p.containerPort
}
