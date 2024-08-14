package pkg

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Parse reads and parses the .env file, returning a map of environment variables.
func Parse(filename string) (map[string]string, error) {
	envVars := make(map[string]string)

	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if len(line) == 0 || strings.HasPrefix(line, "#") {
			continue
		}
		line = strings.TrimPrefix(line, "export ")
		keyVal := strings.SplitN(line, "=", 2)
		if len(keyVal) != 2 {
			return nil, fmt.Errorf("invalid line format: %s", line)
		}
		key := strings.TrimSpace(keyVal[0])
		val := strings.TrimSpace(keyVal[1])

		if strings.Contains(val, "#") {
			val = strings.SplitN(val, "#", 2)[0]
			val = strings.TrimSpace(val)
		}
		if strings.HasPrefix(val, "\"") && strings.HasSuffix(val, "\"") {
			val = strings.Trim(val, "\"")
		}

		envVars[key] = val
	}
	return envVars, nil
}

// Load loads environment variables from one or more .env files and overwrites existing variables.
func Load(filenames ...string) error {
	filenames = filenamesOrDefault(filenames)

	for _, filename := range filenames {
		envMap, err := Parse(filename)
		if err != nil {
			return err
		}

		for key, value := range envMap {
			os.Setenv(key, value)
		}
	}
	return nil
}

// filenamesOrDefault returns a default .env filename if none are provided.
func filenamesOrDefault(filenames []string) []string {
	if len(filenames) == 0 {
		return []string{".env"}
	}
	return filenames
}
