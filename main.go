package main

import (
	"SemiTrash_API/utils"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

const (
	apiPrefix string = "/api/v1"
)

var (
	port                  string
	BookResurcePrefix     string = apiPrefix + "/products"
	ManyBookResurcePrefix string = apiPrefix + "/products"
)

func init() {
	env := godotenv.Load()
	if env != nil {

		log.Fatal("Could not fiend")
	}
	port = os.Getenv("app_port")
}

func main() {
	log.Print("Start Server")
	router := mux.NewRouter()
	utils.BuilManyProductsResourcePrefix(router, ManyBookResurcePrefix)
	utils.BuildProductResourcePrefix(router, BookResurcePrefix)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
