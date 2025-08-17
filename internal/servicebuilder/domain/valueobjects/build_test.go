package valueobjects

import (
	"reflect"
	"testing"

	"gopkg.in/yaml.v3"
)

func TestNewBuild(t *testing.T) {
	tests := []struct {
		name       string
		context    string
		dockerfile string
		want       Build
	}{
		{
			name:       "NewBuild",
			context:    "",
			dockerfile: "./Dockerfile",
			want:       Build{context: "", dockerfile: "./Dockerfile"},
		}, {
			name:       "NewBuild",
			context:    "./my-app  ",
			dockerfile: "",
			want:       Build{context: "./my-app  ", dockerfile: ""},
		},
		{
			name:       "NewBuild",
			context:    "./my-app  ",
			dockerfile: "php.Dockerfile",
			want:       Build{context: "./my-app  ", dockerfile: "php.Dockerfile"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewBuild(tt.context, tt.dockerfile); got != tt.want {
				t.Errorf("NewBuild() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBuild_ToYaml(t *testing.T) {
	tests := []struct {
		name       string
		build      Build
		expected   interface{}
	}{
		{
			name:  "valid build with context and dockerfile",
			build: NewBuild("./app", "Dockerfile"),
			: map[string]string{
				"context":    "./app",
				"dockerfile": "Dockerfile",
			},
		},
		{
			name:  "valid build with context and dockerfile",
			build: NewBuild("./app", "Dockerfile"),
			wantFields: map[string]string{
				"context":    "./app",
				"dockerfile": "Dockerfile",
			},
		},
		{
			name:  "build with empty context",
			build: NewBuild("", "Dockerfile.prod"),
			wantFields: map[string]string{
				"context":    "",
				"dockerfile": "Dockerfile.prod",
			},
		},
		{
			name:  "build with empty Dockerfile",
			build: NewBuild("/app/Dockerfiles", ""),
			wantFields: map[string]string{
				"context":    "/app/Dockerfiles",
				"dockerfile": "",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.build.ToYamlData()

			if err != nil {
				t.Errorf("ToYaml() error = %v", err)
				return
			}

			var result map[string]string

			err = yaml.Unmarshal(got, &result)

			if err != nil {
				t.Errorf("ToYaml() error = %v", err)
				return
			}

			if !reflect.DeepEqual(result, tt.wantFields) {
				t.Errorf("ToYaml() got = %v, want %v", result, tt.wantFields)

			}

		})
	}
}
