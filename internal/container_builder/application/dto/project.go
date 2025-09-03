package dto

type Project struct {
	projectName string
	containers  *[]Container
}

func NewProject(projectName string, containers *[]Container) *Project {
	return &Project{
		projectName: projectName,
		containers:  containers,
	}
}
