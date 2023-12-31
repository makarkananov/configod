package yaml

import (
	"fmt"
	"github.com/makarkananov/configod/core/configloader"
	"os"
)

type ConfigSource struct {
}

func (y *ConfigSource) LoadConfig(filePath string) (*configloader.Config, error) {
	fileInfo, err := os.Stat(filePath)
	if err != nil {
		return nil, fmt.Errorf("error checking file existence: %w", err)
	}

	if fileInfo.IsDir() {
		return nil, fmt.Errorf("provided path is a directory, not a file")
	}

	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("error reading file: %w", err)
	}

	yamlData, err := ParseYAML(data)
	if err != nil {
		return nil, fmt.Errorf("error parsing YAML config: %w", err)
	}

	return &configloader.Config{Fields: yamlData}, nil
}
