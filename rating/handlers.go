package rating

import (
	"encoding/json"
	"net/http"
)

type ProductResponse struct {
	ProductID string   `json:"productId"`
	Ratings   []Rating `json:"ratings"`
}

type AllProductsResponse struct {
	Products []ProductResponse `json:"products"`
}

func HandleGetAllProducts(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	allProducts := GetAllProducts()
	response := AllProductsResponse{
		Products: make([]ProductResponse, 0, len(allProducts)),
	}

	for productID, ratings := range allProducts {
		response.Products = append(response.Products, ProductResponse{
			ProductID: productID,
			Ratings:   ratings,
		})
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
