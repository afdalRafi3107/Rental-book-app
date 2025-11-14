package repository

import (
	"rental-book/internal/entity"

	"gorm.io/gorm"
)

type BookRepository struct {
	db *gorm.DB
}

func NewBookRepository(db *gorm.DB) *BookRepository {
	return &BookRepository{db}
}

func (r *BookRepository) Create(book *entity.Book) error {
	return r.db.Create(book).Error
}

func (r *BookRepository) FindAll() ([]entity.Book, error) {
	var books []entity.Book
	err := r.db.Find(&books).Error
	return books, err
}
