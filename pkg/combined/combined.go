package combined

import "github.com/makarkananov/configod/core/configloader"

type ConfigSource struct {
	Sources []configloader.ConfigSource
}

func NewCombinedConfigSource(sources ...configloader.ConfigSource) *ConfigSource {
	return &ConfigSource{
		Sources: sources,
	}
}

func (c *ConfigSource) LoadConfig(filePath string) (*configloader.Config, error) {
	combinedData := make(map[string]interface{})

	for _, source := range c.Sources {
		sourceData, err := source.LoadConfig(filePath)
		if err != nil {
			return nil, err
		}

		combinedData = mergeMaps(combinedData, sourceData.Fields)
	}

	return &configloader.Config{Fields: combinedData}, nil
}

func mergeMaps(map1, map2 map[string]interface{}) map[string]interface{} {
	result := make(map[string]interface{})

	for key, value := range map1 {
		result[key] = value
	}

	for key, value := range map2 {
		result[key] = value
	}

	return result
}
