package valueobjects

import "testing"

func TestNewImage(t *testing.T) {
	tests := []struct {
		name      string
		imageName string
		tag       string
		expected  Image
	}{
		{
			name:      "creates image with valid name and tag",
			imageName: "nginx",
			tag:       "1.21.0",
			expected:  Image{imageName: "nginx", tag: "1.21.0"},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := NewImage(test.imageName, test.tag)
			if result != test.expected {
				t.Errorf("NewImage() = %v, want %v", result, test.expected)
			}
		})
	}
}
