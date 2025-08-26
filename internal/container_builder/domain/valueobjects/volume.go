package valueobjects

import "fmt"

type Volume struct {
	source string
	target string
}

func NewVolume(source string, target string) Volume {
	return Volume{
		source: source,
		target: target,
	}
}

func (v Volume) Source() string {
	return v.source
}

func (v Volume) Target() string {
	return v.target
}

func (v Volume) ToYaml() ([]byte, error) {
	return []byte(fmt.Sprintf("%s: %s", v.source, v.target)), nil
}
