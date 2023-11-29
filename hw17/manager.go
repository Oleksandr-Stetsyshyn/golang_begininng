package library

import (
	"errors"
	"log"
)

type Manager struct {
	ID   int `gorm:"primary_key"`
	Name string
}

type ManagerRepository interface {
	FindManagerByID(ID uint) (*Manager, error)
	SaveManager(manager *Manager) error
}

func (m *Manager) CheckAndBorrowBook(l *Library, bookID uint, clientID uint) error {
    book, err := l.BookRepo.FindBookByID(bookID)
    if err != nil {
        log.Printf("Failed to find book with ID %d: %v", bookID, err)
        return err
    }

    if book.IsBorrowed {
        log.Printf("Book %d is already borrowed", bookID)
        return errors.New("Book is already borrowed")
    }

    // Mark the book as borrowed and set the borrower ID
    book.IsBorrowed = true
    book.BorrowerID = clientID

    err = l.BookRepo.SaveBook(book)
    if err != nil {
        log.Printf("Failed to save book with ID %d: %v", bookID, err)
        return err
    }

    log.Printf("Manager %d successfully borrowed book %d for client %d", m.ID, bookID, clientID)
    return nil
}

func (m *Manager) ReturnBook(l *Library, bookID uint, clientID uint) error {
	log.Printf("Manager %d is trying to return book %d from client %d", m.ID, bookID, clientID)

	book, err := l.BookRepo.FindBookByID(bookID)
	if err != nil {
		log.Printf("Failed to find book with ID %d: %v", bookID, err)
		return err
	}

	// Check if the client is the one who borrowed the book
	if book.BorrowerID != clientID {
		log.Printf("Client %d did not borrow book %d", clientID, bookID)
		return errors.New("Cannot return a book that was not borrowed by the client")
	}

	// Mark the book as not borrowed
	book.IsBorrowed = false
	book.BorrowerID = 0

	err = l.BookRepo.SaveBook(book)
	if err != nil {
		log.Printf("Failed to save book with ID %d: %v", bookID, err)
		return err
	}

	log.Printf("Manager %d successfully returned book %d from client %d", m.ID, bookID, clientID)
	return nil
}
