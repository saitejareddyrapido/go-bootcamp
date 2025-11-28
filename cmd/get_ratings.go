package main

import (
	"fmt"
	"hello_world/rating"

	"github.com/spf13/cobra"
)

var getRatingsCmd = &cobra.Command{
	Use:     "get-ratings",
	Short:   "Get all ratings for a product",
	Long:    "Usage: get-ratings <productID>",
	Example: "get-ratings product1",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		productID := args[0]

		ratings, avgRating := rating.GetProductRatingInfo(productID)
		if ratings == nil {
			fmt.Printf("Product '%s' not found\n", productID)
			return
		}

		fmt.Printf("\nProduct: %s\n", productID)
		fmt.Printf("Average Rating: %.1f/5\n", avgRating)
		fmt.Printf("Total Ratings: %d\n\n", len(ratings))

		if len(ratings) == 0 {
			fmt.Println("No ratings yet")
			return
		}

		for i, r := range ratings {
			fmt.Printf("Rating #%d:\n", i+1)
			fmt.Printf("  User: %s\n", r.UserID)
			fmt.Printf("  Rating: %d/5\n", r.Rating)
			fmt.Printf("  Comment: %s\n", r.Comment)
			fmt.Println()
		}
	},
}
