package files

import (
	"errors"
	"os"
)

func ReadFile(filename string) string {
	content, err := os.ReadFile(filename)
	if err != nil {
		return ""
	}
	return string(content)
}

func WriteFile(filename, content string) error {
	return os.WriteFile(filename, []byte(content), 0644)
}

func WriteNewFile(filename, content string) error {
	if _, err := os.Stat(filename); err != nil {
		return WriteFile(filename, content)
	}
	return errors.New("file already exists")
}
