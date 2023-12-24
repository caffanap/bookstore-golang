package repositories

import (
	"caffanap/bookstore/api/config"
	"caffanap/bookstore/api/models"
	"errors"

	"gorm.io/gorm"
)

type BookRepository struct {
	db *gorm.DB
}

func NewBookRepository() *BookRepository {
	return &BookRepository{
		db: config.Database,
	}
}

func (r BookRepository) GetAllBooks() []models.Book {
	var books []models.Book
	r.db.Find(&books)
	return books
}

func (r BookRepository) CreateBook(book *models.Book) error {
	if err := r.db.Create(&book).Error; err != nil {
		return err
	}
	return nil
}

func (r BookRepository) GetOneBook(id string) (models.Book, error) {
	var book models.Book
	if err := r.db.First(&book, id).Error; err != nil {
		return book, err
	}
	return book, nil
}

func (r BookRepository) UpdateBook(id string, book *models.Book) (models.Book, error) {
	var newBook models.Book
	if err := r.db.First(&newBook, id).Error; err != nil {
		return newBook, err
	}
	if r.db.Where("id = ?", newBook.Id).Updates(&book).RowsAffected == 0 {
		return newBook, errors.New("tidak menemukan item yang diubah")
	}
	return newBook, nil
}

func (r BookRepository) DeleteBook(id string) (models.Book, error) {
	var book models.Book
	if err := r.db.First(&book, id).Error; err != nil {
		return book, err
	}
	if err := r.db.Delete(&book, book.Id).Error; err != nil {
		return book, err
	}
	return book, nil
}
