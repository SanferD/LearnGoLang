package main

import (
	"fmt"
	"os"
)

func main() {
	_, err := os.Open("does_not_exist");
	if err != nil {
		fmt.Println(err.Error())
	}
}