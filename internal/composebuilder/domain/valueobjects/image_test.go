package valueobjects

import "testing"

func TestNewImage(t *testing.T) {
	tests := []struct {
		name      string
		imageName string
		tag       string
		expected  Image
		expectErr bool
		errorMsg  string
	}{
		{
			name:      "creates image with valid name and tag",
			imageName: "nginx",
			tag:       "1.21.0",
			expected:  Image{imageName: "nginx", tag: "1.21.0"},
			expectErr: false,
		},
		{
			name:      "creates image with valid name and tag",
			imageName: "nginx",
			tag:       "latest",
			expected:  Image{imageName: "nginx", tag: "latest"},
			expectErr: false,
		},
		{
			name:      "creates image with empty tag",
			imageName: "nginx",
			tag:       "",
			expected:  Image{},
			expectErr: true,
			errorMsg:  "image or can't be empty",
		},
		{
			name:      "creates with empty image",
			imageName: "",
			tag:       "latest",
			expected:  Image{},
			expectErr: true,
			errorMsg:  "image or can't be empty",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result, err := NewImage(test.imageName, test.tag)

			if test.expectErr {
				if err == nil {
					t.Error("NewImage() expected error but got none")
					return
				}

				if err.Error() != test.errorMsg {
					t.Errorf("NewImage() expected error message %s but got %s", test.errorMsg, err.Error())
					return
				}
			}

			if result != test.expected {
				t.Errorf("NewImage() = %v, want %v", result, test.expected)
			}
		})
	}
}
