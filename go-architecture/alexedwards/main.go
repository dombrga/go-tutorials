package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/dombrga/alexedwards-tutorial-organizing/models"
	_ "github.com/lib/pq"
)

// Create a custom Env struct which holds a connection pool.
type Env struct {
	// books models.BookModel
	books interface {
		All() ([]models.Book, error)
	}
}

// https://www.alexedwards.net/blog/organising-database-access
func main() {
	// Initialise the connection pool.
	db, err := sql.Open("postgres", "postgres://user:pass@localhost/bookstore")
	if err != nil {
		log.Fatal(err)
	}

	// Create an instance of Env containing the connection pool.
	env := &Env{
		books: models.BookModel{
			DB: db,
		},
	}

	// Use env.booksIndex as the handler function for the /books route.
	http.HandleFunc("/books", env.booksIndex)
	http.ListenAndServe(":3000", nil)
}

// Define booksIndex as a method on Env.
func (env *Env) booksIndex(w http.ResponseWriter, r *http.Request) {
	// We can now access the connection pool directly in our handlers.
	bks, err := env.books.All()
	if err != nil {
		log.Print(err)
		http.Error(w, http.StatusText(500), 500)
		return
	}

	for _, bk := range bks {
		fmt.Fprintf(w, "%s, %s, %s, £%.2f\n", bk.Isbn, bk.Title, bk.Author, bk.Price)
	}
}
