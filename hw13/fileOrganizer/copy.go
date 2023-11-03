package fileOrganizer

import (
	"fmt"
	"os"
	"path/filepath"
)

type Copy struct{}

func (cp *Copy) Perform(filesData PathFilesData) {
	err := copyFile(filesData.Path, filesData.Name, filesData.NewPath)
	if err != nil {
		fmt.Printf("Error moving file: %s\n", err)
		return
	}

	fmt.Printf("Copyed file with path %s name %s to path %s\n", filesData.Path, filesData.Name, filesData.NewPath)
}

func copyFile(path string, name string, newPath string) error {
	filePath := filepath.Join(path, name)
	newFilePath := filepath.Join(newPath, name)
	_, err := os.Stat(filePath)
	if err != nil {
		return fmt.Errorf("file %s does not exist", filePath)
	}

	err = os.MkdirAll(newPath, os.ModePerm)
	if err != nil {
		return fmt.Errorf("failed to create directory: %s", err)
	}

	err = os.Link(filePath, newFilePath)
	if err != nil {
		return err
	}

	return nil
}
