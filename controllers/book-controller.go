package controllers

import (
	"booksapi/modules"
	"booksapi/utils"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

var NewBook modules.Book

func GetBook(w http.ResponseWriter, r *http.Request) {

	newBooks := modules.GetAllBook()
	res, _ := json.Marshal(newBooks)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

func GetBookById(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	fmt.Println(vars, "-  vars")
	BookId := vars["bookId"]
	fmt.Println(BookId, "-bookid")

	ID, err := strconv.ParseInt(BookId, 0, 0)
	fmt.Printf("%T", ID)
	if err != nil {
		fmt.Println("err while parsing")
	}

	bookDetalis, BookError := modules.GetBookById(ID)
	if BookError != nil {
		fmt.Println("error while retriving")
	}

	res, _ := json.Marshal(bookDetalis)
	w.Header().Set("Content-type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	CrateBook := &modules.Book{}
	utils.ParseBody(r, CrateBook)
	b := CrateBook.CreateBook()
	res, _ := json.Marshal(b)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	Id, err := strconv.ParseInt(vars["bookId"], 0, 0)
	if err != nil {
		fmt.Println("failed to parse int")

	}

	bookToReturn := modules.DeleteBookById(Id)
	res, _ := json.Marshal(bookToReturn)
	w.Header().Set("content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	fmt.Println("in update")
	vars := mux.Vars(r)

	IDd := vars["bookId"]
	fmt.Println(IDd)
	bookId, err := strconv.ParseInt(IDd, 0, 0)
	if err != nil {
		fmt.Println("err while parsing int")

	}

	var bookIgotFromYou modules.Book
	utils.ParseBody(r, &bookIgotFromYou)
	// fmt.Println(bookIgotFromYou)
	BookIalredyHad, db := modules.GetBookById(bookId)
	fmt.Println(BookIalredyHad)
	if bookIgotFromYou.Name != "" && bookIgotFromYou.Author != "" && bookIgotFromYou.Publication != "" {

		modules.DeleteBookById(bookId)
		bookIgotFromYou.CreateBook()

	}

	if bookIgotFromYou.Name != BookIalredyHad.Name {
		BookIalredyHad.Name = bookIgotFromYou.Name

	}
	if bookIgotFromYou.Author != BookIalredyHad.Author {
		BookIalredyHad.Author = bookIgotFromYou.Author

	}
	if bookIgotFromYou.Publication != BookIalredyHad.Publication {
		BookIalredyHad.Publication = bookIgotFromYou.Publication

	}
	db.Save(&BookIalredyHad)

	res, _ := json.Marshal(BookIalredyHad)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}
