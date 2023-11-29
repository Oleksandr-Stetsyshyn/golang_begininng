package library

import (
	"fmt"

	nativeMysql "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type GormState struct {
	db *gorm.DB
}

func NewGormConnection() *gorm.DB {
	cfg := nativeMysql.Config{
		User:   "root",
		Passwd: "root",
		Net:    "tcp",
		Addr:   "localhost:8083",
		DBName: "librarydb",
	}
	gorm, err := gorm.Open(mysql.Open(cfg.FormatDSN()), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	return gorm
}
func NewGormState(db *gorm.DB) *GormState {
	db.AutoMigrate(&Book{})
	db.AutoMigrate(&Manager{})
	db.AutoMigrate(&Client{})
	return &GormState{db}
}

func (state *GormState) FindBookByID(ID uint) (*Book, error) {
	var book Book
	if err := state.db.First(&book, ID).Error; err != nil {
		return nil, err
	}
	return &book, nil
}

func (state *GormState) SaveBook(book *Book) error {
	fmt.Println(book.Title)
    return state.db.Save(book).Error
}

func (state *GormState) FindManagerByID(ID uint) (*Manager, error) {
	var manager Manager
	if err := state.db.First(&manager, ID).Error; err != nil {
		return nil, err
	}
	return &manager, nil
}

func (state *GormState) SaveManager(manager *Manager) error {
	return state.db.Save(manager).Error
}
