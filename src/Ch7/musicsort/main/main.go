package main

import (
	"fmt"
	"sort"
	"time"
	"Ch7/musicsort"
)

var tracks = []*musicsort.Track{
	{"Go", "Delilah", "From the Roots Up", 2012, length("3m38s")},
	{"Go", "Moby", "Moby", 1992, length("3m37s")},
	{"Go Ahead", "Alicia Keys", "As I Am", 2007, length("4m36s")},
	{"Ready 2 Go", "Martin Solveig", "Smash", 2011, length("4m24s")},
}


func length(s string) time.Duration {
	d, err := time.ParseDuration(s)
	if err != nil {
		panic(s)
	}
	return d
}

func main() {
	forwardSorter := MakeForwardTrackSorter(tracks)
	reverseSorter := MakeReverseTrackSorter(tracks)
	printTracks := MakeTrackPrinter(tracks)

	PrintForwardTitle("Artist")
	forwardSorter(musicsort.LessAritst)
	printTracks()
	
	PrintReverseTitle("Artist")
	reverseSorter(musicsort.LessAritst)
	printTracks()

	PrintForwardTitle("Album")
	forwardSorter(musicsort.LessAlbum)
	printTracks()
}

func MakeForwardTrackSorter(tracks []*musicsort.Track) func( func(x, y *musicsort.Track) bool ) {
	return func(comp func(x, y *musicsort.Track) bool) {
		sort.Stable(musicsort.MakeCustomSort(tracks, comp))
	}
}

func MakeReverseTrackSorter(tracks []*musicsort.Track) func( func(x, y *musicsort.Track) bool ) {
	return func(comp func(x, y *musicsort.Track) bool) {
		sort.Stable(sort.Reverse(musicsort.MakeCustomSort(tracks, comp)))
	}
}

func PrintForwardTitle(s string) {
	fmt.Println(fmt.Sprintf("===Custom Stateful Sort: %s Forward===", s))
}

func PrintReverseTitle(s string) {
	fmt.Println(fmt.Sprintf("===Custom Stateful Sort: %s Reverse===", s))
}

func MakeTrackPrinter(tracks []*musicsort.Track) func() {
	return func() {
		musicsort.PrintTracks(tracks)
		fmt.Println()	
	}
}