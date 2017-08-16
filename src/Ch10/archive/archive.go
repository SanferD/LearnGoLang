package main

import (
	"fmt"
	_ "Ch10/archive/zip"
	_ "Ch10/archive/tar"
	"image"
	"os"
)

func main() {
	_, kind, err := image.Decode(os.Stdin)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("kind = " + kind)
}

