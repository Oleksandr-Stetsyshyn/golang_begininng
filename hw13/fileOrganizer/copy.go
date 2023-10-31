package fileOrganaizer

import "fmt"

type Copy struct{}

func (cp *Copy) Perform(filesData PathFilesData) {
	err := CopyFile(filesData.Path, filesData.Name, filesData.NewPath)
	if err != nil {
		fmt.Printf("Error moving file: %s\n", err)
		return
	}

	fmt.Printf("Copyed file with path %s name %s to path %s\n", filesData.Path, filesData.Name, filesData.NewPath)
}
