package files

import "os"

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
