package interfaces

type YamlSerializable interface {
	ToYaml() ([]byte, error)
}
