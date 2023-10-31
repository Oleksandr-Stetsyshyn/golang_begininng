package fileOrganaizer

import "fmt"

type Rename struct{}

func (rn *Rename) Perform(filesData PathFilesData) {
	fmt.Printf("Renaming files with path %s and name pattern %s to %s\n", filesData.Path, filesData.NamePattern, filesData.NewName)
}
