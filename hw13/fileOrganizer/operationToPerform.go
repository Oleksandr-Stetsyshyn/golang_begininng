package fileOrganizer

import "fmt"

type operationToPerform interface {
	Perform(pathName PathFilesData)
}

type SelectedOperation struct {
	operation operationToPerform
}

func (sp *SelectedOperation) SetOperation(operation string) {
	operations := map[string]operationToPerform{
		"Delete": &Delete{},
		"Create": &Create{},
		"Copy":   &Copy{},
		"Move":   &Move{},
		"Rename": &Rename{},
	}

	op, ok := operations[operation]
	if !ok {
		fmt.Println("Invalid operation")
		return
	}

	sp.operation = op
}

func (sp *SelectedOperation) ExecuteOperation(filesData PathFilesData) {
	sp.operation.Perform(filesData)
}

type PathFilesData struct {
	Name string
	Path string

	NewName string
	NewPath string

	NamePattern string
}
