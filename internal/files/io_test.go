package files

import (
	"os"
	"path/filepath"
	"testing"
)

func TestReadFile(t *testing.T) {
	var readResult string

	readResult = ReadFile("")
	if readResult != "" {
		t.Errorf("Case 1: Expected \"\", got \"%s\"", readResult);
	}

	readResult = ReadFile("/does/not/exist/path.txt")
	if readResult != "" {
		t.Errorf("Case 2: Expected \"\", got \"%s\"", readResult);
	}

	readResult = ReadFile("doesnotexist.txt")
	if readResult != "" {
		t.Errorf("Case 3: Expected \"\", got \"%s\"", readResult);
	}

	dir := t.TempDir()
	filename := filepath.Join(dir, "test.txt")

	readResult = ReadFile(filename)
	if readResult != "" {
		t.Errorf("Case 4: Expected \"\", got \"%s\"", readResult);
	}

	os.WriteFile(filename, []byte{}, 0644)
	readResult = ReadFile(filename)
	if readResult != "" {
		t.Errorf("Case 5: Expected \"\", got \"%s\"", readResult);
	}

	os.WriteFile(filename, []byte("Hello, World!"), 0644)
	readResult = ReadFile(filename)
	if readResult != "Hello, World!" {
		t.Errorf("Case 6: Expected \"Hello, World!\", got \"%s\"", readResult)
	}
}

