package fileOrganaizer

import (
	"fmt"
	"os"
	"path/filepath"
)

func CreateDirectory(path string) error {
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

func CreateFile(path string, name string) error {
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

func DeleteFile(path string, name string) error {
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

func MoveFile(path string, name string, newPath string) error {
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

	err = os.Rename(filePath, newFilePath)
	if err != nil {
		return err
	}

	return nil
}

func CopyFile(path string, name string, newPath string) error {
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
