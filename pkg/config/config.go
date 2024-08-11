// package config - Contains methods related to configuration
// and manipulation of configuration files (properties, YAML, JSON).
//
// As an example, we have the LoadProperties function,
// which is responsible for recovering data from key-value files.
// It should also be considered that we are referring to generic
// functions, not functions that return specific values ​​from specific files.
package config

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"gopkg.in/yaml.v3"
)

// type Properties - Represents configuration values ​​stored
// in properties files that contain data in key-value format.
type Properties map[string]string

// LoadProperties - Loads data from a properties file that
// contains data in key-value format, then returns this mapped data.
func LoadProperties(filePath string) (Properties, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	properties := make(Properties)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 || strings.HasPrefix(line, "#") {
			continue
		}
		parts := strings.SplitN(line, "=", 2)
		if len(parts) != 2 {
			continue
		}
		properties[strings.TrimSpace(parts[0])] = strings.TrimSpace(parts[1])
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return properties, nil
}

// LoadJSON - Loads data from a JSON file, decodes this
// data and returns it based on the structure defined as a parameter.
func LoadJSON(filePath string, v interface{}) error {
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	return decoder.Decode(v)
}

// LoadYAML - Loads data from a YAML file, decodes this
// data and returns it based on the structure defined as a parameter.
func LoadYAML(filePath string, v interface{}) error {
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	decoder := yaml.NewDecoder(file)
	return decoder.Decode(v)
}

// SaveProperties - Unlike the LoadProperties function that
// receives data in key-value format, this function is responsible
// for overwriting the data using key-value. Can be used to
// create new properties files or modify existing files.
//
// CAUTION: This function overwrites data, so if you modify a file,
// the old data will be deleted.
func SaveProperties(filename string, properties map[string]string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	for key, value := range properties {
		_, err := fmt.Fprintf(writer, "%s=%s\n", key, value)
		if err != nil {
			return err
		}
	}

	return writer.Flush()
}
