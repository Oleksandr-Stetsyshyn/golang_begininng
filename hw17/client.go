package library

import (

	"log"
)

type Client struct {
	ID   uint `gorm:"primary_key"`
	Name string
}

func (c *Client) BorrowBook(l *Library, bookID uint, managerID uint) error {
    log.Printf("Client %d is trying to borrow book %d from manager %d", c.ID, bookID, managerID)

    manager, err := l.ManagerRepo.FindManagerByID(managerID)
    if err != nil {
        log.Printf("Failed to find manager with ID %d: %v", managerID, err)
        return err
    }

    err = manager.CheckAndBorrowBook(l, bookID, c.ID)
    if err != nil {
        log.Printf("Failed to borrow book with ID %d: %v", bookID, err)
        return err
    }

    log.Printf("Client %d successfully borrowed book %d from manager %d", c.ID, bookID, managerID)
    return nil
}

func (c *Client) RequestReturnBook(l *Library, bookID uint, managerID uint) error {
    log.Printf("Client %d is requesting to return book %d to manager %d", c.ID, bookID, managerID)

    manager, err := l.ManagerRepo.FindManagerByID(managerID)
    if err != nil {
        log.Printf("Failed to find manager with ID %d: %v", managerID, err)
        return err
    }

    return manager.ReturnBook(l, bookID, c.ID)
}
