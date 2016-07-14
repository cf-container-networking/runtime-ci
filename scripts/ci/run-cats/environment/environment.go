package environment

import (
	"fmt"
	"os"
	"strconv"
)

type environment struct{}

func New() *environment {
	return &environment{}
}

func (e *environment) GetBoolean(varName string) (bool, error) {
	switch os.Getenv(varName) {
	case "true":
		return true, nil
	case "false", "":
		return false, nil
	default:
		return false, fmt.Errorf("Invalid environment variable: '%s' must be a boolean 'true' or 'false'", varName)
	}
}

func (e *environment) GetBackend() (string, error) {
	value := os.Getenv("BACKEND")
	switch value {
	case "dea", "diego", "":
		return value, nil
	default:
		return "", fmt.Errorf("Invalid environment variable: 'BACKEND' was '%s', but must be 'diego', 'dea', or empty", value)
	}
}

func (e *environment) GetString(varName string) string {
	return os.Getenv(varName)
}

func (e *environment) GetInteger(varName string) (int, error) {
	value := os.Getenv(varName)
	if value == "" {
		return 0, nil
	}

	intValue, err := strconv.Atoi(value)
	if err != nil || intValue <= 0 {
		return 0, fmt.Errorf("Invalid environment variable: '%s' must be an integer greater than 0", varName)
	}

	return intValue, nil
}