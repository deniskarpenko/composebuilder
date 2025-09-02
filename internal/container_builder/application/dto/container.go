package dto

type Port struct {
	hostPort      uint
	containerPort uint
}

func NewPort(hostPort uint, containerPort uint) Port {
	return Port{hostPort: hostPort, containerPort: containerPort}
}

type Volume struct {
	hostPath      string
	containerPath string
}

func NewVolume(hostPath string, containerPath string) Volume {
	return Volume{hostPath: hostPath, containerPath: containerPath}
}

type Image struct {
	imageName string
	tag       string
}

func NewImage(imageName string, tag string) Image {
	return Image{imageName: imageName, tag: tag}
}

type Build struct {
	dockerfilePath string
	contextPath    string
}

func NewBuild(dockerfilePath string, contextPath string) Build {
	return Build{dockerfilePath: dockerfilePath, contextPath: contextPath}
}

type Env struct {
	variable string
	value    string
}

func NewEnv(variable string, value string) Env {
	return Env{variable: variable, value: value}
}

type Container struct {
	name      string
	image     *Image
	build     *Build
	ports     []Port
	volumes   []Volume
	envs      []Env
	envFiles  []string
	networks  []string
	dependsOn []string
}

func NewContainer(
	name string,
	image *Image,
	build *Build,
	ports []Port,
	volumes []Volume,
	envs []Env,
	envFiles []string,
	networks []string,
	dependsOn []string,
) Container {
	return Container{
		name:      name,
		image:     image,
		build:     build,
		ports:     ports,
		volumes:   volumes,
		envs:      envs,
		envFiles:  envFiles,
		dependsOn: dependsOn,
	}
}

func (c *Container) Name() string {
	return c.name
}
