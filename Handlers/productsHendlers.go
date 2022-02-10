package handlers

import (
	"SemiTrash_API/models"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func GetProductById(writer http.ResponseWriter, request *http.Request) {
	initHeaders(writer)
	id, err := strconv.ParseUint(mux.Vars(request)["id"], 0, 64)
	if err != nil {
		log.Print("Error", err)
		msg := models.Message{Message: "erroe while parsing happend:"}
		json.NewEncoder(writer).Encode(msg)
		return
	}
	product, ok := models.FindProductById(id)
	if !ok {
		writer.WriteHeader(404)
		msg := models.Message{Message: "dw"}
		json.NewEncoder(writer).Encode(msg)
	} else {
		writer.WriteHeader(200)

		json.NewEncoder(writer).Encode(product)

	}
}

func CreateProducts(writer http.ResponseWriter, request *http.Request) {
	initHeaders(writer)
	var product models.Produkt

	err := json.NewDecoder(request.Body).Decode(&product)
	if err != nil {
		msg := models.Message{Message: "error,provid file not ivalid:"}
		writer.WriteHeader(400)
		json.NewEncoder(writer).Encode(msg)
		return
	}
	var newId = len(models.DB) + 1
	product.Id = uint64(newId)
	models.DB = append(models.DB, product)

	json.NewEncoder(writer).Encode(product)
}

func UpdateBookById(writer http.ResponseWriter, request *http.Request) {
	/*initHeaders(writer)
	id, err := strconv.Atoi(mux.Vars(request)["id"])
	if err != nil {
		log.Println("error while parsing happend:", err)
		writer.WriteHeader(400)
		msg := models.Message{Message: "do not use parameter ID as uncasted to int type"}
		json.NewEncoder(writer).Encode(msg)
		return
	}
	product, ok := models.FindProductById(uint64(id))
	var newProdukt models.Products
	if !ok {
		log.Println("product not found in data base . id :", id)
		writer.WriteHeader(404)
		msg := models.Message{Message: "product with that ID does not exists in database"}
		json.NewEncoder(writer).Encode(msg)
		return
	}
	err = json.NewDecoder(request.Body).Decode(&newProdukt)
	if err != nil {
		msg := models.Message{Message: "provideed json file is invalid"}
		writer.WriteHeader(400)
		json.NewEncoder(writer).Encode(msg)
		return
	}
	//TODO:Нужно заменить oldBook на newBook в DB!*/
}

func DeleteProductById(writer http.ResponseWriter, request *http.Request) {
	initHeaders(writer)
	id, err := strconv.ParseUint(mux.Vars(request)["id"], 0, 64)
	if err != nil {
		log.Print("Error", err)
		msg := models.Message{Message: "erroe while parsing happend:"}
		json.NewEncoder(writer).Encode(msg)
		return
	}
	ok := models.DeleteBookById(id)
	if !ok {
		writer.WriteHeader(404)
		msg := models.Message{Message: "dw"}
		json.NewEncoder(writer).Encode(msg)
	} else {
		writer.WriteHeader(204)

		json.NewEncoder(writer).Encode("")

	}
}
