package interfaces

type YamlSerializable interface {
	ToYaml() ([]byte, error)
}

type YamlDataSerializable interface {
	ToYamlData() (interface{}, error)
}
