package json

import (
	"os"
	"testing"
)

func TestConfigSource_LoadConfig_ValidJSONFile(t *testing.T) {
	jsonData := `{"key1": "value1", "key2": 42, "key3": true}`
	tempFile, err := os.CreateTemp("", "valid_json_*.json")
	if err != nil {
		t.Fatalf("Error creating temporary file: %v", err)
	}
	defer os.Remove(tempFile.Name())
	defer tempFile.Close()

	if _, err := tempFile.WriteString(jsonData); err != nil {
		t.Fatalf("Error writing to temporary file: %v", err)
	}

	configSource := &ConfigSource{}

	config, err := configSource.LoadConfig(tempFile.Name())

	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if len(config.Fields) == 0 {
		t.Error("Expected non-empty Fields map, but got empty")
	}

	expectedFields := map[string]interface{}{
		"key1": "value1",
		"key2": float64(42),
		"key3": true,
	}

	for key, expectedValue := range expectedFields {
		if actualValue, ok := config.Fields[key]; ok {
			if actualValue != expectedValue {
				t.Errorf("Expected %s=%v, but got %s=%v", key, expectedValue, key, actualValue)
			}
		} else {
			t.Errorf("Expected key %s in Fields, but not found", key)
		}
	}
}

func TestConfigSource_LoadConfig_InvalidJSONFile(t *testing.T) {
	invalidJSONData := `{"key1": "value1", "key2": 42, "key3": true,}`
	tempFile, err := os.CreateTemp("", "invalid_json_*.json")
	if err != nil {
		t.Fatalf("Error creating temporary file: %v", err)
	}
	defer os.Remove(tempFile.Name())
	defer tempFile.Close()

	if _, err := tempFile.WriteString(invalidJSONData); err != nil {
		t.Fatalf("Error writing to temporary file: %v", err)
	}

	configSource := &ConfigSource{}

	config, err := configSource.LoadConfig(tempFile.Name())

	if err == nil {
		t.Error("Expected error, but got nil")
	}

	if config != nil {
		t.Error("Expected empty Fields map, but got non-empty")
	}
}

func TestConfigSource_LoadConfig_NonexistentFile(t *testing.T) {
	configSource := &ConfigSource{}

	config, err := configSource.LoadConfig("nonexistent_file.json")

	if err == nil {
		t.Error("Expected error, but got nil")
	}

	if config != nil {
		t.Error("Expected empty Fields map, but got non-empty")
	}
}

func TestConfigSource_LoadConfig_Directory(t *testing.T) {
	tempDir, err := os.MkdirTemp("", "temp_dir_*")
	if err != nil {
		t.Fatalf("Error creating temporary directory: %v", err)
	}
	defer os.Remove(tempDir)

	configSource := &ConfigSource{}

	config, err := configSource.LoadConfig(tempDir)

	if err == nil {
		t.Error("Expected error, but got nil")
	}

	if config != nil {
		t.Error("Expected empty Fields map, but got non-empty")
	}
}
