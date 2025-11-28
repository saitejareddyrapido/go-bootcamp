package rating

import "errors"

var products = make(map[string]*RatingInfo)

type Rating struct {
	RatingID string `json:"ratingId"`
	UserID   string `json:"userId"`
	Rating   int    `json:"rating"`
	Comment  string `json:"comment"`
}

type RatingInfo struct {
	Rating        []Rating
	AverageRating float64
}

func AddProduct(productID string) {
	if _, exists := products[productID]; !exists {
		products[productID] = &RatingInfo{
			Rating:        []Rating{},
			AverageRating: 0,
		}
	}
}

func AddRating(productID string, ratingID string, userID string, rating int, comment string) error {
	if rating <= 0 {
		return errors.New("rating must be greater than 0")
	}
	if product, exists := products[productID]; exists {
		newRating := Rating{
			RatingID: ratingID,
			UserID:   userID,
			Rating:   rating,
			Comment:  comment,
		}
		product.Rating = append(product.Rating, newRating)
		product.AverageRating = (product.AverageRating*float64(len(product.Rating)-1) + float64(rating)) / float64(len(product.Rating))
	}
	return nil
}

func GetProductRatingInfo(productID string) ([]Rating, float64) {
	if product, exists := products[productID]; exists {
		return product.Rating, product.AverageRating
	}
	return nil, 0.0
}

func GetAllProducts() map[string][]Rating {
	result := make(map[string][]Rating)
	for productID, ratingInfo := range products {
		result[productID] = ratingInfo.Rating
	}
	return result
}
