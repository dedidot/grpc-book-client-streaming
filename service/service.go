package service

import (
	"github.com/dedidot/grpc-book-client-streaming/model"
	"github.com/dedidot/grpc-book-client-streaming/repository"
)

func AddBook(bookData model.Book) model.Book {
	return repository.AddBook(bookData)
}
