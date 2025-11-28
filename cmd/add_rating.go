package main

import (
	"fmt"
	"hello_world/rating"
	"strconv"

	"github.com/spf13/cobra"
)

var addRatingCmd = &cobra.Command{
	Use:   "add-rating",
	Short: "Add a rating for a product",
	Long:  "Usage: add-rating <productID> <ratingID> <userID> <rating> <comment>",
	Example: "add-rating product1 rating1 user1 5 \"Great product!\"",
	Args:  cobra.ExactArgs(5),
	Run: func(cmd *cobra.Command, args []string) {
		productID := args[0]
		ratingID := args[1]
		userID := args[2]
		ratingValue, err := strconv.Atoi(args[3])
		if err != nil {
			fmt.Printf("Error: rating must be a number\n")
			return
		}
		comment := args[4]

		rating.AddProduct(productID)
		err = rating.AddRating(productID, ratingID, userID, ratingValue, comment)
		if err != nil {
			fmt.Printf("Error: %s\n", err)
			return
		}

		fmt.Printf("✓ Rating added successfully!\n")
		fmt.Printf("  Product: %s\n", productID)
		fmt.Printf("  User: %s\n", userID)
		fmt.Printf("  Rating: %d/5\n", ratingValue)
	},
}

