package routes

import (
	"log"
	"net/http"

	"github.com/chagasduarte/go-rest-api/controllers"
	"github.com/chagasduarte/go-rest-api/middleware"
	"github.com/gorilla/mux"
)

func HandleRequest() {
	r := mux.NewRouter()
	r.Use(middleware.ContentTypeMiddleware)
	r.HandleFunc("/", controllers.Home)
	r.HandleFunc("/api/contas", controllers.TodasContas).Methods("Get")
	r.HandleFunc("/api/contas/{id}", controllers.Conta).Methods("Get")
	r.HandleFunc("/api/contas", controllers.NovaConta).Methods("Post")
	r.HandleFunc("/api/contas/{id}", controllers.DeletaConta).Methods("Delete")
	r.HandleFunc("/api/contas/{id}", controllers.EditaConta).Methods("Put")

	log.Fatal(http.ListenAndServe(":8000", r))
}
