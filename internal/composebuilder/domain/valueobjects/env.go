package valueobjects

type Env struct {
	variable string
	value    string
}

func NewEnv(variable string, value string) Env {
	return Env{
		variable: variable,
		value:    value,
	}
}
