package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"reflect"

	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

// type book struct {
// 	Id     string  `json:"Id"`
// 	UnqId  string  `json:"unqid"`
// 	Name   string  `json:"name"`
// 	Author *Author `json:"author"`
// }

// type Author struct {
// 	FirstName string `json:"firstName"`
// 	LastName  string `json:"lastName"`
// }

type book struct {
	Id        string
	UnqId     string
	Name      string
	FirstName string
	LastName  string
}

var books []book
var mysqlCon *sql.DB

 var mysqlHostString string = "testUser:test@/testdb"
//var mysqlHostString string = "testUser:asdfG!@345@/testdb"

func main() {
	r := mux.NewRouter()
	getMysqlCon()
	// fmt.Println(mysqlCon)
	// // qur := "select title,status from tasks2 limit 1"
	// rows, err := mysqlCon.Query(qur)
	// if err != nil {
	// 	fmt.Println("Error whie querying in mysql : ", err)
	// }
	// var title string
	// var status string
	// for rows.Next() {
	// 	err := rows.Scan(&title, &status)
	// 	if err != nil {
	// 		fmt.Println("Error while scaning rows : ", err)
	// 	}
	// 	fmt.Println("Row retrived : ", title, status)

	// }

	// books = append(books, book{Id: "1", UnqId: "1234", Name: "Book1", Author: &Author{FirstName: "Jane", LastName: "Doe"}})
	// books = append(books, book{Id: "2", UnqId: "1235", Name: "Book2", Author: &Author{FirstName: "Peter", LastName: "Smith"}})

	r.HandleFunc("/books", getBooks).Methods("GET")
	r.HandleFunc("/book/{id}", getBook).Methods("GET")
	// r.HandleFunc("/addBook", addBook).Methods("POST")
	// r.HandleFunc("/updatebook/{id}", updateBook).Methods("POST")
	log.Fatal(http.ListenAndServe(":8090", r))
}

func getBooks(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Getting all books")
	id := ""
	books := fetchbooks(id)
	// fmt.Fprint(w, "In get books")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

func getBook(w http.ResponseWriter, r *http.Request) {
	fmt.Println("In get single book")
	param := mux.Vars(r)
	books := fetchbooks(param["id"])
	if !reflect.ValueOf(books).IsZero() {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(books[0])
		return
	} else {
		str := "No book found with id : " + param["id"]
		json.NewEncoder(w).Encode(str)
	}
}

// func addBook(w http.ResponseWriter, r *http.Request) {
// 	var addBook book
// 	_ = json.NewDecoder(r.Body).Decode(&addBook)
// 	addBook.Id = strconv.Itoa(rand.Intn(10000000))
// 	books = append(books, addBook)
// 	json.NewEncoder(w).Encode(books)
// }

// func updateBook(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	params := mux.Vars(r)
// 	for index, item := range books {
// 		if item.Id == params["id"] {
// 			books = append(books[:index], books[index+1:]...)
// 			var book book
// 			_ = json.NewDecoder(r.Body).Decode(&book)
// 			book.Id = params["id"]
// 			books = append(books, book)
// 			json.NewEncoder(w).Encode(book)
// 			return
// 		}
// 	}
// }

func getMysqlCon() {
	var err error
	if mysqlCon, err = sql.Open("mysql", mysqlHostString); err != nil {
		log.Println("Error while connecting to mysql :", err)
	}

}

func fetchbooks(id string) []book {
	var books []book
	var queryCond string
	if id == "" {
		queryCond = ""
	} else {
		queryCond = "where a.id=" + id
	}
	query := "select a.Id,a.UnqId,a.Name,b.firstName,b.lastName from books a join Authors b using (authorId)" + queryCond
	// query := "select a.Id from books a join Authors b using (authorId)" + queryCond
	rows, err := mysqlCon.Query(query)
	if err != nil {
		fmt.Println("Error whie querying in mysql : ", err)
	}
	log.Println("mysql query done")
	for rows.Next() {
		var b book
		// var Id string
		err := rows.Scan(&b.Id, &b.UnqId, &b.Name, &b.FirstName, &b.LastName)
		if err != nil {
			fmt.Println("Error while scaning rows : ", err)
		}
		fmt.Println("Row retrived : ", b)
		books = append(books, b)
	}
	return books
}
