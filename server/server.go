package server

import (
	"fmt"
	"github.com/dedidot/grpc-book-client-streaming/model"
	"github.com/dedidot/grpc-book-client-streaming/proto"
	"github.com/dedidot/grpc-book-client-streaming/service"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"io"
)

type Server struct {
	proto.BookServiceServer
}

func (s *Server) AddBook(stream proto.BookService_AddBookServer) error {
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&proto.GetBookResponse{
				Status:  true,
				Message: "All book data inserted succesfully",
			})
		}

		if err != nil {
			return status.Errorf(codes.Internal, fmt.Sprintf("Internal error %v", err))
		}

		var bookData *proto.Book = req.GetBook()
		fmt.Println(bookData.Title)
		var book model.Book = model.Book{
			Id:     bookData.Id,
			Title:  bookData.Title,
			Author: bookData.Author,
			IsRead: bookData.IsRead,
		}

		service.AddBook(book)
	}

	return nil
}
