package main

import (
	"bufio"
	"fmt"
	"hello_world/rating"
	"os"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

var interactiveCmd = &cobra.Command{
	Use:   "interactive",
	Short: "Start interactive mode (CLI stays running)",
	Run: func(cmd *cobra.Command, args []string) {
		scanner := bufio.NewScanner(os.Stdin)

		fmt.Println("Rating CLI - Interactive Mode")
		fmt.Println("Type 'exit' to quit")
		fmt.Println()

		for {
			fmt.Print("> ")
			scanner.Scan()
			input := strings.TrimSpace(scanner.Text())

			if input == "exit" {
				break
			}

			words := strings.Fields(input)
			if len(words) == 0 {
				continue
			}

			command := words[0]

			if command == "add" && len(words) == 6 {
				productID := words[1]
				ratingID := words[2]
				userID := words[3]
				ratingValue, _ := strconv.Atoi(words[4])
				comment := words[5]

				rating.AddProduct(productID)
				rating.AddRating(productID, ratingID, userID, ratingValue, comment)
				fmt.Println("Rating added!")
			} else if command == "get" && len(words) == 2 {
				productID := words[1]
				ratings, avgRating := rating.GetProductRatingInfo(productID)

				if ratings == nil {
					fmt.Println("Product not found")
					continue
				}

				fmt.Printf("Product: %s\n", productID)
				fmt.Printf("Average: %.1f/5\n", avgRating)
				fmt.Printf("Total: %d ratings\n", len(ratings))

				for i, r := range ratings {
					fmt.Printf("%d. User: %s, Rating: %d/5, Comment: %s\n", i+1, r.UserID, r.Rating, r.Comment)
				}
			} else {
				fmt.Println("Commands:")
				fmt.Println("  add <productID> <ratingID> <userID> <rating> <comment>")
				fmt.Println("  get <productID>")
				fmt.Println("  exit")
			}
		}
	},
}
