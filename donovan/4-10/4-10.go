package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"exercises-go/donovan/4-10/github"
)

//!+
func main() {
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d issues:\n", result.TotalCount)
	for _, item := range result.Items {
		y, m, d := item.CreatedAt.Date()
		yn, mn, dn := time.Now().Date()
		if y == yn {
			if mn-m == 1 && dn-d < 0 || mn == m {
				fmt.Printf("#%-5d %9.9s %.55s less than month ago\n",
					item.Number, item.User.Login, item.Title)
			}
		} else if yn-y == 1 && mn-m < 0 && dn-d < 0 || yn == y {
			fmt.Printf("#%-5d %9.9s %.55s less than year ago\n",
				item.Number, item.User.Login, item.Title)
		} else {
			fmt.Printf("#%-5d %9.9s %.55s more than year ago\n",
				item.Number, item.User.Login, item.Title)
		}
	}
}
