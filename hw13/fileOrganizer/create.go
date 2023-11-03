package fileOrganizer

import (
	"fmt"
	"os"
	"path/filepath"
)

type Create struct{}

func (cr *Create) Perform(filesData PathFilesData) {
	err := createDirectory(filesData.Path)
	if err != nil {
		fmt.Printf("Error creating directory: %s\n", err)
		return
	}

	err = createFile(filesData.Path, filesData.Name)
	if err != nil {
		fmt.Printf("Error creating file: %s\n", err)
		return
	}

	fmt.Printf("File created: %s\n", filepath.Join(filesData.Path, filesData.Name))
}
func createDirectory(path string) error {
	_, err := os.Stat(path)
	if err == nil {
		fmt.Printf("Directory %s already exists\n", path)
		return nil
	}

	err = os.MkdirAll(path, os.ModePerm)
	if err != nil {
		return err
	}
	return nil
}

func createFile(path string, name string) error {
	filePath := filepath.Join(path, name)
	_, err := os.Stat(filePath)
	if err == nil {
		return fmt.Errorf("file %s already exists", filePath)
	}

	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Printf("Error closing file %s: %s\n", filePath, err)
		}
	}(file)

	return nil
}
