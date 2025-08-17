package valueobjects

import (
	"testing"
)

func TestNewVolume(t *testing.T) {
	tests := []struct {
		name   string
		source string
		volume string
		want   Volume
	}{
		{
			name:   "Usual",
			source: "./app",
			volume: "/var/run/www/",
			want: Volume{
				source: "./app",
				target: "/var/run/www/",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewVolume(tt.source, tt.volume); got != tt.want {
				t.Errorf("NewVolume() = %v, want %v", got, tt.want)
			}
		})
	}
}
