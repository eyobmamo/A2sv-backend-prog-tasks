package models

type Member struct {
	ID            int
	Name          string
	BorrowedBooks []string
}
