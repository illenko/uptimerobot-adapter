package main

import (
	"encoding/json"
	"log"
	"net/http"
	"uptimerobot-adapter/middleware"
	"uptimerobot-adapter/model"
)

func main() {
	mux := http.NewServeMux()
	mux.Handle("POST /uptime", middleware.ApiKeyMiddleware(middleware.RateLimit(http.HandlerFunc(uptimeHandler))))
	log.Println("Starting server on :8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}

func uptimeHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Received request for /uptime")

	product := model.ProductReq{
		Title: "BMW Pencil",
	}

	body, err := makePostRequest[model.ProductReq, model.ProductRes]("https://dummyjson.com/products/add", product)
	if err != nil {
		log.Printf("Error making post request: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(body); err != nil {
		log.Printf("Error encoding response: %v", err)
		http.Error(w, "Error encoding response", http.StatusInternalServerError)
		return
	}

	log.Println("Response sent successfully")
}
