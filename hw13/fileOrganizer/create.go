package fileOrganaizer

import (
	"fmt"
	"path/filepath"
)

type Create struct{}

func (cr *Create) Perform(filesData PathFilesData) {
	err := CreateDirectory(filesData.Path)
	if err != nil {
		fmt.Printf("Error creating directory: %s\n", err)
		return
	}

	err = CreateFile(filesData.Path, filesData.Name)
	if err != nil {
		fmt.Printf("Error creating file: %s\n", err)
		return
	}

	fmt.Printf("File created: %s\n", filepath.Join(filesData.Path, filesData.Name))
}
