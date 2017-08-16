package musicsort

import (
	"fmt"
	"os"
	"text/tabwriter"
	"time"
)

type Track struct {
	Title string
	Artist string
	Album string
	Year int
	Length time.Duration
}

func LessTitle(x, y *Track) bool { return x.Title < y.Title }
func LessAritst(x, y *Track) bool { return x.Artist < y.Artist }
func LessAlbum(x, y *Track) bool { return x.Album < y.Album }
func LessYear(x, y *Track) bool { return x.Year < y.Year }
func LessLength(x, y *Track) bool { return x.Length < y.Length }

func PrintTracks(tracks []*Track) {
	format := "%v\t%v\t%v\t%v\t%v\n"
	tw := new(tabwriter.Writer).Init(os.Stdout, 0, 8, 2, ' ', 0)
	fmt.Fprintf(tw, format, "Title", "Artist", "Album", "Year", "Length")
	fmt.Fprintf(tw, format, "-----", "------", "-----", "----", "------")
	for _, t := range tracks {
		fmt.Fprintf(tw, format, t.Title, t.Artist, t.Album, t.Year, t.Length)
	}
	tw.Flush()
}
