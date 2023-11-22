package env

import (
	"configod/core/configloader"
)

type ConfigSource struct{}

func (e *ConfigSource) LoadConfig(_ string) (*configloader.Config, error) {
	envData := ParseEnv()

	return &configloader.Config{Fields: envData}, nil
}
