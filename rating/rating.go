package rating

import "errors"

/*
Data Structure:
{
(productID, []RatingInfo{})
(productID, []RatingInfo{})
}
*/

var products = make(map[string]*RatingInfo)

type Rating struct {
	RatingID string
	UserID   string
	Rating   int
	Comment  string
}

type RatingInfo struct {
	Rating []Rating
}

func AddProduct(productID string) {
	if _, exists := products[productID]; !exists {
		products[productID] = &RatingInfo{
			Rating: []Rating{},
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
	}
	return nil
}

func GetProductRatingInfo(productID string) []Rating {
	if product, exists := products[productID]; exists {
		return product.Rating
	}
	return nil
}
