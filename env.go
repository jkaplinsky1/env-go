package env_go

import (
	"errors"
	"log"
	"os"
	"reflect"
)

// Key-Value mapping of os.Environ
type EnvVars map[string]string

type tag struct {
	Key      string
	Default  string
	Required bool
}

func GetEnv(env interface{}) []string {
	environ := os.Environ()
	log.Print(environ)
	return environ
}

// Get Environ

// Convert to EnvVars (k/v map)

// Unmarshal to env struct{}

func unmarshal(envVars EnvVars, env interface{}) error {
	v := reflect.ValueOf(env)
	if v.Kind() != reflect.Ptr || v.IsNil() {
		return errors.New("invalid value")
	}

	// take value of each field
	t := v.Type()
	for i := 0; i < t.NumField(); i++ {
		// if struct loop through struct to allow nested structs
		field := v.Field(i)
		if field.Kind() == reflect.Struct {
			// struct recurse
		}
		// else take value based on tag, marshal to passed env interface{}

		tField := t.Field()
		envTag := tField.Tag.Get("env")
	}

	return nil
}
