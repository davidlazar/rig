// rig is a random identity generator.
package main

import (
	"crypto/rand"
	"flag"
	"fmt"
	"math/big"
)

//go:generate go run data/togo.go data

var male = flag.Bool("m", false, "generate male name")
var female = flag.Bool("f", false, "generate female name")

func randint(max int) int {
	x, err := rand.Int(rand.Reader, big.NewInt(int64(max)))
	if err != nil {
		panic(err)
	}
	return int(x.Int64())
}

func randitem(items []string) string {
	i := randint(len(items))
	return items[i]
}

type place struct {
	city, state, zip, areacode string
}

func parseLocData(loc string) (p place) {
	n, err := fmt.Sscan(loc, &p.city, &p.state, &p.areacode, &p.zip)
	if n != 4 {
		panic(err)
	}
	return
}

func main() {
	flag.Parse()

	if *male == *female {
		coin := randint(2)
		if coin == 0 {
			*male = !*male
		} else {
			*female = !*female
		}
	}

	var firstNames []string
	if *male {
		firstNames = mnames
	} else if *female {
		firstNames = fnames
	}

	fmt.Printf("%s %s\n", randitem(firstNames), randitem(lnames))
	fmt.Printf("%d %s\n", randint(1024)+1, randitem(street))
	place := parseLocData(randitem(locdata))
	fmt.Printf("%s, %s %s\n", place.city, place.state, place.zip)
	fmt.Printf("(%s) xxx-xxxx\n", place.areacode)
}
