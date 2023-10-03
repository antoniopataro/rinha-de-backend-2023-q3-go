package envs

import (
	"log"
	"os"
)

func GetEnvOrDie(env string) (val string) {
	val = os.Getenv(env)

	if val == "" {
		log.Fatalf("missing environment variable %s", env)
	}

	return val
}
