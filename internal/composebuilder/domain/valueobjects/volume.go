package valueobjects

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
