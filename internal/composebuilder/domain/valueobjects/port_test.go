package valueobjects

import "testing"

func TestNewPorts(t *testing.T) {
	tests := []struct {
		name          string
		host          int
		containerPort *int
		wantPorts     Ports
	}{
		{
			name:          "usual ports",
			host:          80,
			containerPort: func() *int { i := 8000; return &i }(),
			wantPorts: Ports{
				host:          80,
				containerPort: func() *int { i := 8000; return &i }(),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewPorts(tt.host, tt.containerPort)
			if got.host != tt.host || got.containerPort != tt.containerPort {
				t.Errorf("NewPorts() got = %v, want = %v", got, tt.wantPorts)
			}
		})
	}
}
