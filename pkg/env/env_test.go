package env

import (
	"os"
	"testing"
)

func TestConfigSource_LoadConfig_LoadingEnvironmentConfig(t *testing.T) {
	configSource := &ConfigSource{}

	if err := os.Setenv("TEST_KEY1", "value1"); err != nil {
		t.Fatalf("Error setting TEST_KEY1: %v", err)
	}
	defer func() {
		if err := os.Unsetenv("TEST_KEY1"); err != nil {
			t.Fatalf("Error unsetting TEST_KEY1: %v", err)
		}
	}()

	if err := os.Setenv("TEST_KEY2", "value2"); err != nil {
		t.Fatalf("Error setting TEST_KEY2: %v", err)
	}
	defer func() {
		if err := os.Unsetenv("TEST_KEY2"); err != nil {
			t.Fatalf("Error unsetting TEST_KEY2: %v", err)
		}
	}()

	expectedEnvVars := map[string]string{
		"TEST_KEY1": "value1",
		"TEST_KEY2": "value2",
	}

	config, err := configSource.LoadConfig("")

	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if len(config.Fields) == 0 {
		t.Error("Expected non-empty Fields map, but got empty")
	}

	for key, value := range expectedEnvVars {
		if val, ok := config.Fields[key]; ok {
			if valStr, ok := val.(string); ok {
				if valStr != value {
					t.Errorf("Expected %s=%s, but got %s=%s", key, value, key, valStr)
				}
			} else {
				t.Errorf("Expected string value for %s, but got %v", key, val)
			}
		} else {
			t.Errorf("Expected key %s in Fields, but not found", key)
		}
	}
}

func TestConfigSource_LoadConfig_EmptyEnvironmentConfig(t *testing.T) {
	configSource := &ConfigSource{}

	_, err := configSource.LoadConfig("")

	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
}

func TestConfigSource_LoadConfig_MultipleEnvironmentVariables(t *testing.T) {
	configSource := &ConfigSource{}

	if err := os.Setenv("TEST_KEY1", "value1"); err != nil {
		t.Fatalf("Error setting TEST_KEY1: %v", err)
	}
	defer func() {
		if err := os.Unsetenv("TEST_KEY1"); err != nil {
			t.Fatalf("Error unsetting TEST_KEY1: %v", err)
		}
	}()

	if err := os.Setenv("TEST_KEY2", "value2"); err != nil {
		t.Fatalf("Error setting TEST_KEY2: %v", err)
	}
	defer func() {
		if err := os.Unsetenv("TEST_KEY2"); err != nil {
			t.Fatalf("Error unsetting TEST_KEY2: %v", err)
		}
	}()

	if err := os.Setenv("TEST_KEY3", "value3"); err != nil {
		t.Fatalf("Error setting TEST_KEY3: %v", err)
	}
	defer func() {
		if err := os.Unsetenv("TEST_KEY3"); err != nil {
			t.Fatalf("Error unsetting TEST_KEY3: %v", err)
		}
	}()

	expectedEnvVars := map[string]string{
		"TEST_KEY1": "value1",
		"TEST_KEY2": "value2",
		"TEST_KEY3": "value3",
	}

	config, err := configSource.LoadConfig("")

	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if len(config.Fields) == 0 {
		t.Error("Expected non-empty Fields map, but got empty")
	}

	for key, value := range expectedEnvVars {
		if val, ok := config.Fields[key]; ok {
			if valStr, ok := val.(string); ok {
				if valStr != value {
					t.Errorf("Expected %s=%s, but got %s=%s", key, value, key, valStr)
				}
			} else {
				t.Errorf("Expected string value for %s, but got %v", key, val)
			}
		} else {
			t.Errorf("Expected key %s in Fields, but not found", key)
		}
	}
}
