package main

import (
	"fmt"
)

type rideRating struct {
	user    string
	rating  int
	comment []string
}

func main() {
	// var rideRating1 rideRating

	map1 := make(map[string]rideRating)
	map1["Jhon"] = rideRating{
		user:    "Jhon",
		rating:  5,
		comment: []string{"good ride!", "good ride2!"},
	}
	map1["Jane"] = rideRating{
		user:    "Jane",
		rating:  3,
		comment: []string{"good ride!", "good ride2!"},
	}

	if _, ok := map1["Jhon"]; ok {
		fmt.Println("Jhon is in the map")
	} else {
		fmt.Println("Jhon is not in the map")
	}

	for key, _ := range map1 {
		fmt.Print("the rating is ")
		for i := 0; i < map1[key].rating; i++ {
			fmt.Print("*")
		}
		fmt.Println()
	}

	// fmt.Println(map1)

	// rideRating1.user = "Jhon"
	// rideRating1.rating = 5
	// rideRating1.comment = append(rideRating1.comment, "good ride!")
	// rideRating1.comment = append(rideRating1.comment, "good ride2!")

	// fmt.Println(rideRating1)

}
