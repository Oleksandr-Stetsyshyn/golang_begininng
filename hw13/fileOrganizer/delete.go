package fileOrganizer

import (
	"fmt"
	"os"
	"path/filepath"
)

type Delete struct{}

func (del *Delete) Perform(filesData PathFilesData) {

	err := deleteFile(filesData.Path, filesData.Name)
	if err != nil {
		fmt.Printf("Error deleting file: %s\n", err)
		return
	}
	fmt.Printf("Deleting files with path %s and name pattern %s\n", filesData.Path, filesData.Name)
}
func deleteFile(path string, name string) error {
	filePath := filepath.Join(path, name)

	_, err := os.Stat(filePath)
	if err != nil {
		return fmt.Errorf("file %s does not exist", filePath)
	}

	err = os.Remove(filePath)
	if err != nil {
		return err
	}

	return nil
}
