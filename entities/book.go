package entities

// Book represents a book data in the database.
type Book struct {
	Id    string `json:"id,omitempty"`
	Name  string `json:"name"`
	ISBN  string `json:"isbn,omitempty"`
	Price int    `json:"price"`
}
