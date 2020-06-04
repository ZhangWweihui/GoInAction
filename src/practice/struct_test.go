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

type Rectangle struct {
	width, height float32
}

func (r *Rectangle) Area() float32 {
	return r.width * r.height
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

	book2.author = "Lily"
	book2.subject = "Agriculture"
	book2.bookToString()

	change1(book2)
	t.Log(book2)

	/*
		& 取址符，获取变量在内存中的地址
	*/
	change2(&book2)
	t.Log(book2)
}

func change1(b Book) {
	b.title += "_change1"
}

func change2(b *Book) {
	b.title += "_change2"
}

func TestRectangle(t *testing.T) {
	rect := Rectangle{23.4, 22.55}
	t.Logf("area : %f", rect.Area())
}
