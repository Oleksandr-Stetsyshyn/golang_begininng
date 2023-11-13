package fileChangedEvent

import (
	"log"
	"os"
	"time"
)

type Publisher struct {
	Broker *Broker
	Path   string
}

func (p *Publisher) FolderWatcher() {
	fileState := make(map[string]bool)

	for {
		files, err := os.ReadDir(p.Path)
		if err != nil {
			log.Fatal(err)
		}

		newFileState := make(map[string]bool)
		for _, file := range files {
			newFileState[file.Name()] = true
			if !fileState[file.Name()] {
				p.Broker.NotifyAll("changed or added file: " + file.Name())
			}
		}

		for filename := range fileState {
			if !newFileState[filename] {
				p.Broker.NotifyAll("deleted file: " + filename)
			}
		}

		fileState = newFileState
		time.Sleep(3 * time.Second)
	}
}
