package dto

type Port struct {
	hostPort      string
	containerPort uint
}

type Volume struct {
	hostPath      string
	containerPath string
}

type Build struct {
	dockerfilePath string
	contextPath    string
}

type Env struct {
	variable string
	value    string
}

type Container struct {
	name      string
	image     *string
	tag       *string
	build     *Build
	ports     []Port
	volumes   []Volume
	envs      []Env
	envFiles  []string
	dependsOn []string
}

type Project struct {
	projectName string
	containers  []*Container
}

func NewProject(projectName string, containers *Container) *Project {
	return &Project{
		projectName: projectName,
		containers:  []*Container{containers},
	}
}
