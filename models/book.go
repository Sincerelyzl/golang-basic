package models

type book struct {
	Title, Author string
}

func (b *book) Summary() string {
	return b.Title + " by " + b.Author
}

func CreatBook(title, author string) book {
	return book{Title: title, Author: author}
}
