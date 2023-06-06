package generatedproto

import (
	context "context"
	"test/service/database"
)

type ProtoService struct {
}

func (s *ProtoService) GetBooks(ctx context.Context, author *RequestAuthor) (*ResponseBook, error) {

	books, err := database.GetBooksByAuthor(author.Name)
	if err != nil {
		return nil, err
	}
	return &ResponseBook{
		Title: books,
	}, nil
}
func (s *ProtoService) GetAuthors(ctx context.Context, book *RequestBook) (*ResponseAuthor, error) {
	authors, err := database.GetAuthorsByBook(book.Title)
	if err != nil {
		return nil, err
	}
	return &ResponseAuthor{
		Name: authors,
	}, nil
}
func (s *ProtoService) mustEmbedUnimplementedSearchServer() {}
