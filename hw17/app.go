package library

import (
	"fmt"
)

func Main() {
	state := NewGormState(NewGormConnection())

	library := &Library{
		BookRepo:    state,
		ManagerRepo: state,
	}

	// Create a book
	book1 := &Book{ID: 1, Title: "The Little Prince", IsBorrowed: false, BorrowerID:0 }

	err := library.BookRepo.SaveBook(book1)
	if err != nil {
		return
	}

	book2 := &Book{ID: 2, Title: "Троє поросят", IsBorrowed: false, BorrowerID:0 }
	err = library.BookRepo.SaveBook(book2)
	if err != nil {
		return
	}

	// Create managers
	manager1 := &Manager{ID: 1, Name: "Steve"}
	manager2 := &Manager{ID: 2, Name: "Harry"}
	library.ManagerRepo.SaveManager(manager1)
	library.ManagerRepo.SaveManager(manager2)

	// Create a client
	client := &Client{ID: 1, Name: "You"}

	// Borrow the book
	err = client.BorrowBook(library, 1, 1)
	if err != nil {
		fmt.Println(err)
	}

	// Return the wrong book
	err = client.RequestReturnBook(library, 2, 2)
	if err != nil {
		fmt.Println(err)
	}

	// Return the correct book
	err = client.RequestReturnBook(library, 1, 2)
	if err != nil {
		fmt.Println(err)
	}
}
