package fileOrganizer

import (
	"flag"
	"golang_beginning/hw13/fileOrganizer"
)

func Main() {
	// go run . -operation=Create -Path="./hw13/files/" -name="text.txt"
	// go run . -operation=Move -Path="./hw13/files/" -name="text.txt" -newPath="./hw13/files/new/"
	var filesData = fileOrganizer.PathFilesData{}
	var operation string

	flag.StringVar(&operation, "operation", "Create", "Operation to perform: Delete, Create, Copy, Move, Rename")

	flag.StringVar(&filesData.Name, "name", "", "Name of file")
	flag.StringVar(&filesData.Path, "Path", "", "Path to operate on")

	flag.StringVar(&filesData.NewName, "newName", "", "New name for files")

	flag.StringVar(&filesData.NewPath, "newPath", "", "New path for files")

	flag.StringVar(&filesData.NamePattern, "pattern", ".*_", "Regular expression pattern for file names")
	flag.Parse()

	selectedOp := &fileOrganizer.SelectedOperation{}
	selectedOp.SetOperation(operation)
	selectedOp.ExecuteOperation(filesData)
}
