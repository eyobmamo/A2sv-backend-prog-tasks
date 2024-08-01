package controllers

import (
	"bufio"
	"fmt"
	"gowhere/models"
	"gowhere/services"
	"os"
	"strconv"
	"strings"
	"time"
)

type LibraryControllerType struct {
	libraryService services.LibraryService
}

var LibraryController *LibraryControllerType

func InitLibraryController(libraryService services.LibraryService) {
	LibraryController = &LibraryControllerType{
		libraryService: libraryService,
	}
}

// helper function
func getStringInput(prompt string) string {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print(prompt)
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		if input != "" {
			return input
		}
		fmt.Println("Invalid input. Please enter a non-empty string.")
	}
}

// main  consol interface
func (lc *LibraryControllerType) Console_control() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("================================================================================")
		fmt.Println("Welcome to the Library Management System!")
		fmt.Println("Please select an option:")
		fmt.Println("1. Add a book")
		fmt.Println("2. Remove a book")
		fmt.Println("3. Borrow a book")
		fmt.Println("4. Return a book")
		fmt.Println("5. List available books")
		fmt.Println("6. List borrowed books")
		fmt.Println("0. Exit")
		choice, _ := reader.ReadString('\n')
		choice = string(choice[0])
		switch choice {
		case "1":
			lc.AddBook()
		case "2":
			lc.RemoveBook()
		case "3":
			lc.BorrowBook()
		case "4":
			lc.ReturnBook()
		case "5":
			lc.ListAvailableBooks()
			break
		case "6":
			lc.ListBorrowedBooks()
		case "0":
			fmt.Println("Exiting...")
			reader.ReadString('\n')
			return
		default:
			fmt.Println("Invalid input")
		}
		fmt.Println("================================================================================")
		time.Sleep(2 * time.Second)
	}
}

// helper function
func getIntInput(prompt string) int {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print(prompt)
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		num, err := strconv.Atoi(input)
		if err == nil {
			return num
		}
		fmt.Println("Invalid input. Please enter an integer.")
	}
}

func (libCont *LibraryControllerType) AddBook() {
	bookID := getIntInput("Enter book ID: ")
	bookTitle := getStringInput("Enter book title: ")
	bookAuthor := getStringInput("Enter book author: ")
	book := &models.Book{
		ID:     bookID,
		Title:  bookTitle,
		Author: bookAuthor,
		Status: "Available",
	}

	error := libCont.libraryService.AddBook(book)

	if error != nil {
		fmt.Println("Error:", error)
	} else {
		fmt.Println("Book added successfully", *book)
	}
}

// impliment add member control

func (libCont *LibraryControllerType) RemoveBook() {
	bookID := getIntInput("Enter book ID: ")
	error := libCont.LibraryService.RemoveBook(bookID)

	if error != nil {
		fmt.Println("Error:", error)
	} else {
		fmt.Println("Book removed successfully")
	}
}

func (libCont *LibraryControllerType) BorrowBook() {
	memberID := getIntInput("Enter member ID: ")
	bookID := getIntInput("Enter book ID: ")
	err := libCont.LibraryService.BorrowBook(bookID, memberID)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Book borrowed successfully")
	}
}

func (libCont *LibraryControllerType) ReturnBook() {
	memberID := getIntInput("Enter member ID: ")
	bookID := getIntInput("Enter book ID: ")
	error := libCont.LibraryService.ReturnBook(bookID, memberID)
	if error != nil {
		fmt.Println("Error:", error)
	} else {
		fmt.Println("Book returned successfully")
	}
}

func (libCont *LibraryControllerType) ListAvailableBooks() {
	books := libCont.LibraryService.ListAvailableBooks()
	for _, book := range books {
		fmt.Println(book)
	}
}

func (libCont *LibraryControllerType) ListBorrowedBooks() {

	memberID := getIntInput("Enter member ID: ")
	books, error := libCont.LibraryService.ListBorrowedBooks(memberID)
	if error == nil {
		for _, book := range books {
			fmt.Println(book)
		}
	} else {
		fmt.Println("Error:", error)
	}
}
