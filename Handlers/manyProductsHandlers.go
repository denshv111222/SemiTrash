package handlers

import (
	"SemiTrash_API/models"
	"encoding/json"
	"net/http"
)

func initHeaders(writer http.ResponseWriter) {
	writer.Header().Set("Content-Type", "application/json")
}

func GetAllProdukts(writer http.ResponseWriter, request *http.Request) {
	initHeaders(writer)
	writer.WriteHeader(200)
	json.NewEncoder(writer).Encode(models.DB)
}
