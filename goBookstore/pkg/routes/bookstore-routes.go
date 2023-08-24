package routes

import (
	"github.com/anuvratnehra/goProjectsFCC/goBookstore/pkg/controllers"
	"github.com/gorilla/mux"
)

var RegsiterBookStoreRoutes = func(router *mux.Router) {
	router.HandleFunc("/healthz", controllers.GetHealth).Methods("GET")
	router.HandleFunc("/book/", controllers.CreateBook).Methods("POST")
	router.HandleFunc("/book/", controllers.GetBook).Methods("GET")
	router.HandleFunc("/book/{bookId}", controllers.GetBookById).Methods("GET")
	router.HandleFunc("/book/{bookId}", controllers.UpdateBook).Methods("PUT")
	router.HandleFunc("/book/{bookId}", controllers.DeleteBook).Methods("DELETE")
}
