package services

import (
	"fmt"
	"gowhere/models"
)


type LibraryManager interface {
	AddBook(book *models.Book)
	RemoveBook(bookID int)
	BorrowBook(bookID int,memberID int) error
	ReturnBook(bookID int, memberID int) error
	ListAvailableBooks() []models.Book
	ListBorrowedBooks(memberId int) []models.Book
}

type libraryService struct {
	books map[int]*models.Book
	members map[int]*models.Member
}

// this function is use for add book

func (libServ libraryService) AddBook(book *models.Book ) error {
	BookToAdd := *book
	_,ok := libServ.book[BookToAdd.ID]
	if ok {
		return errors.new("Book already exists")
	} else {
		libServ.books[BookToAdd.ID] = book
		return nil
	}
	
}

func(libServ *libraryService) RemoveBook(bookID int) error {
	_,ok := libServ.books[bookID]
	if !ok {
		return errors.New("Book not found")
	}

	delete(libServ.books,bookID)
	return nil

}

func (libServ *libraryService) BorrowBook(bookID int ,memberID int) error {
	book,ok := libServ.books[bookID]

	if !ok {
		error.New("Book not found")
	}

	if book.Status == "Borowed" {
		return errors.New("Book already borrowed")
	}

	member ,ok := libServ.members[memberID]
	if !ok {
		return errors.New(fmt.Printf("no member with id %d",memberID))
	}

	//this code dose make a slice including list of borrowedbook + new book
	newBorrow := append(member.BorrowedBook,*book)

	libServ.books[bookID] =&models.Member{
		ID: member.ID,
		Name: member.Name,
		BorrowedBook: newlist,
	}

	libServ.books[bookID] = &models.Book{
		ID: book.ID,
		Title : book.Title,
		Author: book.Author,
		Status : "Borrowed",
	}
	return nil
}

func (libServ *libraryService) ReturnBook(bookID int ,memberID int) error {
	book,ok := libServ.books[bookId]

	if !ok {
		return errors.New("Book not found")
	}

	if book.Status != "Borrowed" {
		return errors.New("Not borrowed Book")
	}

	member,ok := libServ.members[memberID]
	if !ok {
		return errors.New("NO such Member")
	}

	for i,borrowed_Book := range member.BorrowedBooks {
		if borrowed_Book.ID == bookID {
			Remaining_Book := append(member.BorrowedBooks[:i],member.listBorrowedBooks[i+1]...)
			Status := "Available"
		}

		libServ.books[bookID] = & models.book{
			ID: book.ID,
			Title:book.Title,
			Author: bottowedBook.Author,
			Status: Status,
		}

		libServ.members[memberID] = &models.Member {
			ID: member.ID,
			Name: member.Name,
			borrowedBooks: Remaining_Book,
		}

		return nil
			
	}

	return errors.New("Book not borrowed by current member")

}

func (libServ *libraryService) ListAvailableBooks() []models.Book {
	var all_book []models.Book
	for _, book :=range libserv.books {
		all_book = append(all_book,*book)
	}

	return allBooks
}

func (libServ *libraryService) ListBorrowedBooks(memberID int) ([]models.Book,error) {
	member,ok := libServ.members[memberID]
	if !ok {
		return errors.New("No Such Member")
	}
	
	return member.BorrowedBook ,nil
}

func new_library_starter() libraryService {
	books_in := map[int]*models.Book{
		1:{ID: 1,Title: "Book 1" ,Author:"Author 1",Status: "Available"},
		2:{ID: 2,Title: "Book 2" ,Author:"Author 2",Status: "Borrowed"},
		3:{ID: 3,Title: "Book 3" ,Author:"Author 3",Status: "Available"},
		4:{ID: 4,Title: "Book 4" ,Author:"Author 4",Status: "Borrowed"},
		5:{ID: 5,Title: "Book 5" ,Author:"Author 5",Status: "Available"},
		6:{ID: 6,Title: "Book 6" ,Author:"Author 6",Status: "Borrowed"},
		7:{ID: 7,Title: "Book 7" ,Author:"Author 7",Status: "Available"},
		8:{ID: 8,Title: "Book 8" ,Author:"Author 8",Status: "Borrowed"},

	}

	members_in := map[int]*models.Member{
		1:{ID:1,Name: "Member 1",BorrowedBooks : []models.Book{*(books_in[8])}},
		2:{ID:2,Name: "Member 2"},
		3:{ID:3,Name: "Member 3",BorrowedBooks : []models.Book{*(books_in[2])*(books_in[4]),*(books_in[6])}},
	}

	return &libraryService{
		books: books_in,
		members: members_in 
	}
	
}


