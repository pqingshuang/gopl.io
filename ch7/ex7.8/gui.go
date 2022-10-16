
package main
import (
	// "fmt"
	// "os"
	"sort"
	// "text/tabwriter"
	"time"
)
type Track struct {
	Title  string
	Artist string
	Album  string
	Year   int
	Length time.Duration
}

type multier struct {
    t         []*Track
    primary   string
    secondary string
    third     string
}

func (x *multier) Len() int      { return len(x.t) }
func (x *multier) Swap(i, j int) { x.t[i], x.t[j] = x.t[j], x.t[i] }

func (x *multier) Less(i, j int) bool {
    key := x.primary
    for k := 0; k < 3; k++ {
        switch key {
        case "Title":
            if x.t[i].Title != x.t[j].Title {
                return x.t[i].Title < x.t[j].Title
            }
        case "Year":
            if x.t[i].Year != x.t[j].Year {
                return x.t[i].Year < x.t[j].Year
            }
        case "Length":
            if x.t[i].Length != x.t[j].Length {
                return x.t[i].Length < x.t[j].Length
            }
        }
        if k == 0 {
            key = x.secondary
        } else if k == 1 {
            key = x.third
        }
    }
    return false
}

// update primary sorting key
func setPrimary(x *multier, p string) {
    x.primary, x.secondary, x.third = p, x.primary, x.secondary
}

// if x is *multiple type, then update ordering keys
func SetPrimary(x sort.Interface, p string) {
    if x, ok := x.(*multier); ok {
        setPrimary(x, p)
    }
}

// return a new multier
func NewMultier(t []*Track, p, s, th string) sort.Interface {
    return &multier{
        t:         t,
        primary:   p,
        secondary: s,
        third:     th,
    }
}