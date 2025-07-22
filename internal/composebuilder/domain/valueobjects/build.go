package valueobjects

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

func (b Build) ToYamlData() (interface{}, error) {
	return struct {
		Context    string `yaml:"context"`
		Dockerfile string `yaml:"dockerfile"`
	}{
		Context:    b.context,
		Dockerfile: b.dockerfile,
	}, nil
}
