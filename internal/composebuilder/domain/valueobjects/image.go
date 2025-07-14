package valueobjects

import (
	"errors"
	"fmt"
)

type Image struct {
	imageName string
	tag       string
}

func NewImage(imageName string, tag string) (Image, error) {
	if imageName == "" || tag == "" {
		return Image{}, errors.New("image or can't be empty")
	}

	return Image{
		imageName: imageName,
		tag:       tag,
	}, nil
}

func (i Image) Tag() string {
	return i.tag
}

func (i Image) ImageName() string {
	return i.imageName
}

func (i Image) toYaml() string {
	return fmt.Sprintf("%s:%s", i.imageName, i.tag)
}
