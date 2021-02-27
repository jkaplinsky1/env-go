package env_go

import (
	"log"
	"os"
)

// Key-Value mapping of os.Environ
var EnvVars map[string]string

type tag struct {
	Keys     []string
	Default  string
	Required bool
}

func GetEnv(v interface{}) {
	environ := os.Environ()
	log.Print(environ)
}