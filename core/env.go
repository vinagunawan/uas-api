package core

import (
	"os"
)

func GetEnv(key string) string {
	return os.Getenv(key)
}
