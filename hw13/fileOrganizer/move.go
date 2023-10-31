package fileOrganaizer

import "fmt"

type Move struct{}

func (mv *Move) Perform(filesData PathFilesData) {
	err := MoveFile(filesData.Path, filesData.Name, filesData.NewPath)
	if err != nil {
		fmt.Printf("Error moving file: %s\n", err)
		return
	}

	fmt.Printf("Moving files with path %s name %s to path %s\n", filesData.Path, filesData.Name, filesData.NewPath)
}
