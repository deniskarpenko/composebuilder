package valueobjects

import "testing"

func TestNewEnv(t *testing.T) {
	tests := []struct {
		name     string
		variable string
		value    string
		want     Env
		wantErr  bool
		errorMsg string
	}{
		{
			name:     "valid variable and value",
			variable: "DEBUG",
			value:    "true",
			want:     Env{variable: "DEBUG", value: "true"},
			wantErr:  false,
		},
		{
			name:     "invalid variable",
			variable: "",
			value:    "true",
			want:     Env{},
			wantErr:  true,
			errorMsg: "variable and value can't be empty",
		},
		{
			name:     "invalid value",
			variable: "DEBUG",
			value:    "",
			want:     Env{},
			wantErr:  true,
			errorMsg: "variable and value can't be empty",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewEnv(tt.variable, tt.value)

			if (err != nil) != tt.wantErr {
				t.Errorf("NewEnv() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if tt.wantErr {
				if err == nil {
					t.Error("NewEnv() expected error, but got nil")
					return
				}

				if err.Error() != tt.errorMsg {
					t.Errorf("NewEnv() expected error message %s, but got %s", tt.errorMsg, err.Error())
					return
				}
			}

			if tt.want != got {
				t.Errorf("NewEnv() got = %v, want %v", got, tt.want)
			}

		})
	}
}

func TestEnv_ToYaml(t *testing.T) {
	tests := []struct {
		name     string
		env      Env
		expected string
	}{
		{
			name:     "DEBUG: true",
			env:      Env{variable: "DEBUG", value: "true"},
			expected: "DEBUG: true",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.env.ToYaml(); got != tt.expected {
				t.Errorf("ToYaml() = %v, want %v", got, tt.expected)
			}
		})
	}
}
