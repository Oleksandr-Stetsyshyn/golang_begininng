package library

// Database represents the library's database
type Database struct {
    Books map[string]Book
    BorrowedBooks map[uint]bool
}