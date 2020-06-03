package practice

import (
	"fmt"
	"testing"
)

type Book struct {
	id      int
	title   string
	author  string
	subject string
	price   float32
}

func (b *Book) bookToString() {
	fmt.Printf("Book[%d] id=%d, tilte=%s, author=%s, subject=%s, price=%f\n", &b, b.id, b.title, b.author, b.subject, b.price)
}

func TestCreateBook(t *testing.T) {
	book1 := Book{1, "Effective Go", "Jack", "CT", 87.32}
	t.Log(book1)
	book1.bookToString()

	book2 := Book{id: 2, title: "信息技术部", price: 34.32}
	t.Log(book2)
	book2.bookToString()
}
