package utils

import (
	handlers "SemiTrash_API/Handlers"

	"github.com/gorilla/mux"
)

func BuilManyProductsResourcePrefix(router *mux.Router, prefix string) {
	router.HandleFunc(prefix, handlers.GetAllProdukts).Methods("GET")
}
func BuildProductResourcePrefix(router *mux.Router, prefix string) {
	router.HandleFunc(prefix+"/{id}", handlers.GetProductById).Methods("GET")
	router.HandleFunc(prefix, handlers.CreateProducts).Methods("POST")
	router.HandleFunc(prefix+"/{id}", handlers.UpdateBookById).Methods("PUT")
	router.HandleFunc(prefix+"/{id}", handlers.DeleteProductById).Methods("DELETE")
}
