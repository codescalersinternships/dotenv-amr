package pkg

import (
	"os"
	"reflect"
	"testing"
)

// TestParse tests the Parse function with different scenarios.
func TestParse(t *testing.T) {
	tmp := `
		# This is a comment
		VAR1=value1
		VAR2="value2 with spaces"
		export VAR3=value3
		VAR4=value4#withcomment
		VAR5=value5
		`

	tmpFile, err := os.CreateTemp("", ".env")
	if err != nil {
		t.Fatalf("error creating file: %v", err)
	}
	defer os.Remove(tmpFile.Name())

	_, err = tmpFile.WriteString(tmp)
	if err != nil {
		t.Fatalf("error writing to file: %v", err)
	}

	envVars, err := Parse(tmpFile.Name())
	if err != nil {
		t.Fatalf("parse failed: %v", err)
	}

	expected := map[string]string{
		"VAR1": "value1",
		"VAR2": "value2 with spaces",
		"VAR3": "value3",
		"VAR4": "value4",
		"VAR5": "value5",
	}

	if !reflect.DeepEqual(envVars, expected) {
		t.Errorf("expected %v, got %v", expected, envVars)
	}
}

// TestLoad tests the Load function with multiple files.
func TestLoadWithSpecificFile(t *testing.T) {
	tmp1 := `VAR1=value1`
	tmp2 := `VAR2=value2`

	tmpFile1, err := os.CreateTemp("", ".env1")
	if err != nil {
		t.Fatalf("error creating file: %v", err)
	}
	defer os.Remove(tmpFile1.Name())

	_, err = tmpFile1.WriteString(tmp1)
	if err != nil {
		t.Fatalf("error writing to file: %v", err)
	}

	tmpFile2, err := os.CreateTemp("", ".env2")
	if err != nil {
		t.Fatalf("error creating file: %v", err)
	}
	defer os.Remove(tmpFile2.Name())

	_, err = tmpFile2.WriteString(tmp2)
	if err != nil {
		t.Fatalf("error writing to file: %v", err)
	}

	err = Load(tmpFile1.Name(), tmpFile2.Name())
	if err != nil {
		t.Fatalf("load failed: %v", err)
	}

	if val := os.Getenv("VAR1"); val != "value1" {
		t.Errorf("expected VAR1 to be 'value1', got '%s'", val)
	}
	if val := os.Getenv("VAR2"); val != "value2" {
		t.Errorf("expected VAR2 to be 'value2', got '%s'", val)
	}
}

func TestLoadWithDefaultFile(t *testing.T) {
	tmp := `
		VAR1=default_value1
		VAR2=default_value2
	`

	tempFile, err := os.Create(".env")
	if err != nil {
		t.Fatalf("error creating .env file: %v", err)
	}
	defer os.Remove(tempFile.Name())

	_, err = tempFile.WriteString(tmp)
	if err != nil {
		t.Fatalf("error writing to .env file: %v", err)
	}

	err = Load()
	if err != nil {
		t.Fatalf("load failed: %v", err)
	}

	if val := os.Getenv("VAR1"); val != "default_value1" {
		t.Errorf("expected VAR1 to be 'default_value1', got '%s'", val)
	}
	if val := os.Getenv("VAR2"); val != "default_value2" {
		t.Errorf("expected VAR2 to be 'default_value2', got '%s'", val)
	}
}
