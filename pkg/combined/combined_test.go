package combined

import (
	"configod/core/configloader"
	"errors"
	"testing"
)

type MockConfigSource struct {
	ConfigData map[string]interface{}
	LoadErr    error
}

func (m *MockConfigSource) LoadConfig(filePath string) (*configloader.Config, error) {
	if m.LoadErr != nil {
		return nil, m.LoadErr
	}

	return &configloader.Config{Fields: m.ConfigData}, nil
}

func TestConfigSource_LoadConfig_Success(t *testing.T) {
	source1 := &MockConfigSource{
		ConfigData: map[string]interface{}{"key1": "value1", "key2": 42},
	}
	source2 := &MockConfigSource{
		ConfigData: map[string]interface{}{"key2": 43, "key3": true},
	}
	combinedSource := NewCombinedConfigSource(source1, source2)

	config, err := combinedSource.LoadConfig("fakeFilePath")

	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	expectedConfig := map[string]interface{}{"key1": "value1", "key2": 43, "key3": true}
	if !compareMaps(config.Fields, expectedConfig) {
		t.Errorf("Expected config.Fields to be %v, but got %v", expectedConfig, config.Fields)
	}
}

func TestConfigSource_LoadConfig_ErrorInSource(t *testing.T) {
	source1 := &MockConfigSource{
		ConfigData: map[string]interface{}{"key1": "value1", "key2": 42},
	}
	source2 := &MockConfigSource{
		LoadErr: errors.New("error in source2"),
	}
	combinedSource := NewCombinedConfigSource(source1, source2)

	config, err := combinedSource.LoadConfig("fakeFilePath")

	if err == nil {
		t.Error("Expected error, but got nil")
	}

	if config != nil {
		t.Errorf("Expected nil config, but got %+v", config)
	}
}

func compareMaps(map1, map2 map[string]interface{}) bool {
	if len(map1) != len(map2) {
		return false
	}

	for key, value1 := range map1 {
		if value2, ok := map2[key]; !ok || value1 != value2 {
			return false
		}
	}

	return true
}
