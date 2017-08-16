package main

import (
	"fmt"
	"flag"
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"
	"time"
)

var verbose = flag.Bool("v", false, "print progress periodically")

var wg sync.WaitGroup
var done chan struct{}

func main() {
	flag.Parse()
	roots := flag.Args()
	done = make(chan struct{})

	go func() {
		os.Stdin.Read(make([]byte, 1))
		close(done)
		fmt.Println("CLOSED")
	}()

	if len(roots) == 0 {
		roots = []string{"."}
	}

	var tick <- chan time.Time
	if *verbose {
		tick = time.Tick(500*time.Millisecond)
	}

	fileSizes := make(chan int64)
	for _, dir := range roots {
		wg.Add(1)
		go walkDir(dir, fileSizes)
	}
	
	go func() {
		wg.Wait()
		close(fileSizes)
	}()

	ticker := time.NewTicker(100*time.Millisecond)

	var numOfFiles, totalSize int64
loop:
	for {
		select {
		case size, ok := <- fileSizes:
			if !ok {
				break loop
			}
			numOfFiles++
			totalSize += size
		case <-tick:
			printDiskUsage(numOfFiles, totalSize)			
		case <-done:
			for range fileSizes {
				// let all existing go routines complete
			}
		case <-ticker.C:
			panic("check for goroutine leaks")
		}
	}

	fmt.Println("Finished:")
	printDiskUsage(numOfFiles, totalSize)			

}

func printDiskUsage(numOfFiles, totalSize int64) {
	fmt.Printf("%d files %.1f GB\n", numOfFiles, float64(totalSize)/1e9)
}

func walkDir(dir string, fileSizes chan <- int64) {
	defer wg.Done()
	if !pollCancellation() {
		return
	}
	for _, entry := range dirents(dir) {
		if entry.IsDir() {
			subdir := filepath.Join(dir, entry.Name())
			wg.Add(1)
			go walkDir(subdir, fileSizes)
		} else {
			fileSizes <- entry.Size()
		}
	}
}

var sema = make(chan struct{}, 20)
func dirents(dir string) []os.FileInfo {
	select {
		case sema <- struct{}{}:
			defer func() { <-sema }()
		case <- done:
			return nil
	}
	infos, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error in dirents: %s\n", err.Error())
		return nil
	}
	return infos
}

func pollCancellation() bool {
	select {
	case <-done:
		return false
	default:
		return true
	}
}
