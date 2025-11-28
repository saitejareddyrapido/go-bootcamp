package main

import (
	"hello_world/rating"
	"log"
	"net/http"
)

func main() {
	// Sample data
	rating.AddProduct("product1")
	rating.AddProduct("product2")
	rating.AddRating("product1", "rating1", "user1", 5, "Excellent product!")

	// API endpoint
	http.HandleFunc("/api/products", rating.HandleGetAllProducts)
	http.ListenAndServe(":8080", nil)

	// Start server
	log.Println("Server starting on http://localhost:8080")
	// log.Fatal(http.ListenAndServe(":8080", nil))
}
