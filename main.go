package main

import (
	"fmt"
	"hello_world/rating"
)

func main() {

	rating.AddProduct("product1")
	rating.AddProduct("product2")
	rating.AddProduct("product3")

	rating.AddRating("product1", "rating1", "user1", 5, "Excellent ride! Very smooth and comfortable.")
	rating.AddRating("product1", "rating2", "user2", 4, "Good ride, but could be faster.")
	rating.AddRating("product1", "rating3", "user3", 3, "Average ride experience.")

	fmt.Println(rating.GetProductRatingInfo("product1"))

}
