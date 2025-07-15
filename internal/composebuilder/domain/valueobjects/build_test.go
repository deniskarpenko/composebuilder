package valueobjects

import (
	"testing"
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
