package htmlutils

import (
	"os"
	"testing"

	"github.com/gouniverse/utils"
)

func TestFileExists(t *testing.T) {
	filePath := "FileExistsTest.txt"

	if fileExists(filePath) {
		utils.Unlink(filePath)
	}

	if fileExists(filePath) {
		t.Error("Non-existing file exists")
	}

	utils.FilePutContents(filePath, "Hello", os.FileMode(os.O_APPEND))

	if fileExists(filePath) == false {
		t.Error("File DOES NOT exist")
	}
}
