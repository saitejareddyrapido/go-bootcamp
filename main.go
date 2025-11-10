package main

import (
	"fmt"
)

// rideRating struct stores rating information for a rider
type rideRating struct {
	userName string
	rating   int
	comment  string
}

// addRating adds a new rating to the map
func addRating(ratings map[string]rideRating, userName string, rating int, comment string) {
	newRating := rideRating{
		userName: userName,
		rating:   rating,
		comment:  comment,
	}
	ratings[userName] = newRating
	fmt.Printf("Rating added for %s!\n", userName)
}

// getRating retrieves a rating for a specific user
func getRating(ratings map[string]rideRating, userName string) {
	if rating, exists := ratings[userName]; exists {
		fmt.Printf("\nRating for %s:\n", userName)
		fmt.Printf("  Rating: %d/5\n", rating.rating)
		fmt.Printf("  Comment: %s\n", rating.comment)
		displayStars(rating.rating)
	} else {
		fmt.Printf("No rating found for %s\n", userName)
	}
}

// displayStars prints stars based on the rating
func displayStars(rating int) {
	fmt.Print("  Stars: ")
	for i := 0; i < rating; i++ {
		fmt.Print("*")
	}
	fmt.Println()
}

// displayAllRatings shows all ratings in the map
func displayAllRatings(ratings map[string]rideRating) {
	if len(ratings) == 0 {
		fmt.Println("No ratings available.")
		return
	}

	fmt.Println("\n=== All Rider Ratings ===")
	for userName, rating := range ratings {
		fmt.Printf("\nUser: %s\n", userName)
		fmt.Printf("  Rating: %d/5\n", rating.rating)
		fmt.Printf("  Comment: %s\n", rating.comment)
		displayStars(rating.rating)
	}
}

// checkRatingExists checks if a rating exists for a user
func checkRatingExists(ratings map[string]rideRating, userName string) bool {
	_, exists := ratings[userName]
	return exists
}

func main() {
	// Create a map to store rider ratings
	// Key: user name (string), Value: rideRating struct
	ratings := make(map[string]rideRating)

	// Add some initial ratings
	addRating(ratings, "John", 5, "Excellent ride! Very smooth and comfortable.")
	addRating(ratings, "Jane", 4, "Good ride, but could be faster.")
	addRating(ratings, "Bob", 3, "Average ride experience.")

	// Display all ratings
	displayAllRatings(ratings)

	// Check if a specific user has a rating
	userToCheck := "John"
	if checkRatingExists(ratings, userToCheck) {
		fmt.Printf("\n%s has a rating in the system.\n", userToCheck)
	} else {
		fmt.Printf("\n%s does not have a rating in the system.\n", userToCheck)
	}

	// Get a specific user's rating
	fmt.Println("\n--- Getting specific rating ---")
	getRating(ratings, "Jane")

	// Try to get a rating that doesn't exist
	fmt.Println("\n--- Getting non-existent rating ---")
	getRating(ratings, "Alice")
}
