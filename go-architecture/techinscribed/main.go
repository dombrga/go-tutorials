package main

import (
	"log"
	"net/http"
	"techinscribed-course/controllers"
	"techinscribed-course/repositories"
	"techinscribed-course/sqldb"
)

// https://techinscribed.com/different-approaches-to-pass-database-connection-into-controllers-in-golang/?source=post_page-----d22d3fa76d91--------------------------------
func main() {
	db := sqldb.ConnectDB()

	// each kind of repo gets a db (and other dependencies, if there is any),
	// because they probably need access to the db.
	// you can add another kind of repo. For example, bookRepo.
	// But the question for booksRepo is that, the newBaseHandlers only accepts
	// UserRepository interface type. How would you create a handler for bookRepo?
	// I think answer is create a new base handler for that new repo.
	userRepo := repositories.NewUserRepo(db)

	// BaseHandlers has repo field that has db as dependency. The repo is the one that interacts with the db.
	// The controller/handler uses the repo.
	h := controllers.NewBaseHandler(userRepo)
	u := controllers.NewUserBaseHandler(userRepo)

	router := http.NewServeMux()
	// router.Handle("/v1/", http.StripPrefix("/v1", router))
	router.HandleFunc("GET /{id}", h.FindById)
	router.HandleFunc("GET /posts/{id}", h.FindById)
	router.HandleFunc("GET /users/{id}", u.FindById)

	log.Println("Listening to port 3000")
	log.Fatalln(http.ListenAndServe(":3000", router))
}
