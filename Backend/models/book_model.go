package models

type BookStruct struct {
	Title  string `json:"title,omitempty" validate:"required"`
	Author string `json:"author,omitempty" validate:"required"`
	Year   string `json:"year,omitempty" validate:"required"`
	Pages  string `json:"pages,omitempty" validate:"required"`
	Genre  string `json:"genre,omitempty" validate:"required"`
}

type BooksArray struct {
	Books []BookStruct `json:"books"`
}

type BookUpdateStruct struct {
	Id     string `json:"id,omitempty" validate:"required"`
	Title  string `json:"title,omitempty" validate:"required"`
	Author string `json:"author,omitempty" validate:"required"`
	Year   string `json:"year,omitempty" validate:"required"`
	Pages  string `json:"pages,omitempty" validate:"required"`
	Genre  string `json:"genre,omitempty" validate:"required"`
}
