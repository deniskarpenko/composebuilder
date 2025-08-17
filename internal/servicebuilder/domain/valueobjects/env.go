package valueobjects

import (
	"errors"
	"fmt"
)

type Env struct {
	variable string
	value    string
}

func NewEnv(variable string, value string) (Env, error) {
	if variable == "" || value == "" {
		return Env{}, errors.New("variable and value can't be empty")
	}

	return Env{
		variable: variable,
		value:    value,
	}, nil
}

func (env Env) Variable() string {
	return env.variable
}

func (env Env) Value() string {
	return env.value
}

func (env Env) ToYaml() ([]byte, error) {
	yamlData := fmt.Sprintf("%s: %s", env.variable, env.value)
	return []byte(yamlData), nil
}

func (env Env) ToYamlData() (interface{}, error) {
	return map[string]interface{}{
		env.variable: env.value,
	}, nil
}
