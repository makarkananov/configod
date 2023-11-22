package env

import (
	"os"
	"strings"
)

func ParseEnv() map[string]interface{} {
	envData := make(map[string]interface{})

	for _, envVar := range os.Environ() {
		pair := strings.SplitN(envVar, "=", 2)
		if len(pair) == 2 {
			envData[pair[0]] = pair[1]
		}
	}

	return envData
}
