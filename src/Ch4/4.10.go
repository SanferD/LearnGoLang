package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"Ch4/github"
)

func main() {
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%d issues:\n", result.TotalCount)
	currentUnixTime := time.Now().Unix()
	
	const unixFirstYear = 1970

	var differentYear, sameYear, sameMonth []*github.Issue

	for _, item := range result.Items {
		issueUnixTime := item.CreatedAt.Unix()
		timeDifference := time.Unix(currentUnixTime - issueUnixTime, 0)
		if timeDifference.Year() > unixFirstYear+1 {
			differentYear = append(differentYear, item)
		} else {
			sameYear = append(sameYear, item)
			if timeDifference.Month() == 1 {
				sameMonth = append(sameMonth, item)
			}
		}
	}

	if len(sameMonth) != 0 {
		fmt.Println("======== Within the Month ========")
		for _, issue := range sameMonth {
			printIssue(issue)
		}
	} else {
		fmt.Println("No issues within the month")
	}

	if len(sameYear) != 0 {
		fmt.Println("======== Within the Year ========")
		for _, issue := range sameYear {
			printIssue(issue)
		}
	} else {
		fmt.Println("No issues within the year")
	}

	if len(differentYear) != 0 {
		fmt.Println("======== Different Year ========")
		for _, issue := range differentYear {
			printIssue(issue)
		}
	} else {
		fmt.Println("No issues of another year")
	}
}

func abs(x int64) int64 {
	if x < 0 {
		x = -x
	}
	return x
}

func printIssue(issue *github.Issue) {
	fmt.Printf("#%-5d %9.9s %.55s\n", issue.Number, issue.User.Login, issue.Title)
}