package htmlutils

import "os"

// fileGetContents reads entire file into a string
func fileGetContents(filename string) (string, error) {
	data, err := os.ReadFile(filename)
	return string(data), err
}
