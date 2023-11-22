package json

import (
	"configod/core/configloader"
	"fmt"
	"os"
)

type ConfigSource struct{}

func (j *ConfigSource) LoadConfig(filePath string) (*configloader.Config, error) {
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

	jsonData, err := ParseJSON(data)
	if err != nil {
		return nil, fmt.Errorf("error parsing JSON config: %w", err)
	}

	return &configloader.Config{Fields: jsonData}, nil
}
