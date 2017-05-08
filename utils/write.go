package utils

import (
	"encoding/json"
	"os"
)

// WriteTo writes data to the given path
func WriteTo(path string, data []byte) error {
	f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE, 0644)

	if err != nil {
		return err
	}

	defer f.Close()
	f.Write(data)

	return nil
}

// WriteStructTo writes a struct of data to the given path
func WriteStructTo(path string, data interface{}) error {
	b, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return err
	}
	return WriteTo(path, b)
}

// ExistsPath validates if the paths exists
func ExistsPath(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}

	return false, err
}
