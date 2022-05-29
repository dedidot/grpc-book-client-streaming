package repository

import "github.com/dedidot/grpc-book-client-streaming/model"

var storage []model.Book = []model.Book{}

func AddBook(bookData model.Book) model.Book {
	storage = append(storage, bookData)
	return bookData
}
