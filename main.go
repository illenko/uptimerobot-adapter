package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.Handle("POST /uptime", apiKeyMiddleware(rateLimitMiddleware(http.HandlerFunc(uptimeHandler))))
	log.Fatal(http.ListenAndServe(":8080", mux))
}

func uptimeHandler(w http.ResponseWriter, r *http.Request) {
	product := ProductReq{
		Title: "BMW Pencil",
	}

	body, err := makePostRequest[ProductReq, ProductRes]("https://dummyjson.com/products/add", product)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(body); err != nil {
		http.Error(w, "Error encoding response", http.StatusInternalServerError)
	}
}
