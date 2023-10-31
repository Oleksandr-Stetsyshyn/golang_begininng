package fileOrganaizer

import "fmt"

type OperationToPerform interface {
	Perform(pahtName PathFilesData)
}

type SelectedOperation struct {
	operation OperationToPerform
}

func (sp *SelectedOperation) SetOperation(operation string) {
	switch operation {
	case "Delete":
		sp.operation = &Delete{}
	case "Create":
		sp.operation = &Create{}
	case "Copy":
		sp.operation = &Copy{}
	case "Move":
		sp.operation = &Move{}
	case "Rename":
		sp.operation = &Rename{}
	default:
		fmt.Println("Invalid operation")
		return
	}
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
