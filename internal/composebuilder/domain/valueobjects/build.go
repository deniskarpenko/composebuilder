package valueobjects

import "gopkg.in/yaml.v3"

type Build struct {
	context    string
	dockerfile string
}

func NewBuild(context, dockerfile string) Build {
	return Build{context: context, dockerfile: dockerfile}
}

func (b Build) Context() string {
	return b.context
}

func (b Build) Dockerfile() string {
	return b.dockerfile
}

func (b Build) ToYaml() ([]byte, error) {
	data := struct {
		Context    string `yaml:"context"`
		Dockerfile string `yaml:"dockerfile"`
	}{
		Context:    b.context,
		Dockerfile: b.dockerfile,
	}

	return yaml.Marshal(data)
}
