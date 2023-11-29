package library

type Book struct {
    ID    int `gorm:"primary_key"`
    Title string
    IsBorrowed bool
    BorrowerID uint
}

type BookRepository interface {
    FindBookByID(ID uint) (*Book, error)
    SaveBook(book *Book) error
}