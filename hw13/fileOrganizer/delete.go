package fileOrganaizer

import "fmt"

type Delete struct{}

func (del *Delete) Perform(filesData PathFilesData) {

	err := DeleteFile(filesData.Path, filesData.Name)
	if err != nil {
		fmt.Printf("Error deleting file: %s\n", err)
		return
	}
	fmt.Printf("Deleting files with path %s and name pattern %s\n", filesData.Path, filesData.NamePattern)
}
