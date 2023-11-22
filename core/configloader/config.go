package configloader

import "fmt"

type Config struct {
	Fields map[string]interface{}
}

type ConfigSource interface {
	LoadConfig(basePath string) (*Config, error)
}

func (c *Config) ValueString(key string) (string, error) {
	val, ok := c.Fields[key]
	if !ok {
		return "", fmt.Errorf("key '%s' not found", key)
	}

	strVal, ok := val.(string)
	if !ok {
		return "", fmt.Errorf("value for key '%s' is not a string", key)
	}

	return strVal, nil
}

func (c *Config) ValueInt(key string) (int, error) {
	val, ok := c.Fields[key]
	if !ok {
		return 0, fmt.Errorf("key '%s' not found", key)
	}

	intVal, ok := val.(int)
	if !ok {
		return 0, fmt.Errorf("value for key '%s' is not an integer", key)
	}

	return intVal, nil
}

func (c *Config) ValueBool(key string) (bool, error) {
	val, ok := c.Fields[key]
	if !ok {
		return false, fmt.Errorf("key '%s' not found", key)
	}

	boolVal, ok := val.(bool)
	if !ok {
		return false, fmt.Errorf("value for key '%s' is not a boolean", key)
	}

	return boolVal, nil
}

func (c *Config) ValueStringDefault(key, defaultValue string) (string, error) {
	val, ok := c.Fields[key]
	if !ok {
		return defaultValue, fmt.Errorf("key '%s' not found", key)
	}

	strVal, ok := val.(string)
	if !ok {
		return defaultValue, fmt.Errorf("value for key '%s' is not a string", key)
	}

	return strVal, nil
}

func (c *Config) ValueIntDefault(key string, defaultValue int) (int, error) {
	val, ok := c.Fields[key]
	if !ok {
		return defaultValue, fmt.Errorf("key '%s' not found", key)
	}

	intVal, ok := val.(int)
	if !ok {
		return defaultValue, fmt.Errorf("value for key '%s' is not an integer", key)
	}

	return intVal, nil
}
