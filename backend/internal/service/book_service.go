package service

import (
	"rental-book/internal/entity"
	"rental-book/internal/repository"
)

type BookService struct {
	repo *repository.BookRepository
}

func NewBookService(r *repository.BookRepository) *BookService {
	return &BookService{r}
}

func (s *BookService) Create(book *entity.Book) error {
	return s.repo.Create(book)
}

func (s *BookService) GetAll() ([]entity.Book, error) {
	return s.repo.FindAll()
}
