package fileChangedEvent

import (
	"log"
	"os"
	"time"
)

type Publisher struct {
	Subscribers []*Subscriber
}

func (p *Publisher) Subscribe(s *Subscriber) {
	p.Subscribers = append(p.Subscribers, s)
}

func (p *Publisher) Notify(event string) {
	for _, s := range p.Subscribers {
		s.Events <- event
	}
}

func FolderWatcher(path string, publisher *Publisher) {
    fileState := make(map[string]bool)

    for {
        files, err := os.ReadDir(path)
        if err != nil {
            log.Fatal(err)
        }

        newFileState := make(map[string]bool)
        for _, file := range files {
            newFileState[file.Name()] = true
            if !fileState[file.Name()] {
                publisher.Notify("changed or added file: " + file.Name())
            }
        }

        for filename := range fileState {
            if !newFileState[filename] {
                publisher.Notify("deleted file: " + filename)
            }
        }

        fileState = newFileState

        time.Sleep(3 * time.Second)
    }
}
