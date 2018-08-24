package main

import (
	"time"
	"fmt"
	"github.com/elycruz/go-programming-lang-book/chp12/format"
	"github.com/elycruz/go-programming-lang-book/chp12/display"
	"os"
	"reflect"
)

func formatAnyExample () {
	var (
		x int64 = 1
		d time.Duration = 1 * time.Nanosecond
	)
	fmt.Println("x int64 = 1")
	fmt.Println("d time.Duration = 1 * time.Nanosecond")
	fmt.Println("format.Any(x)== ", format.Any(x)) 	// "1"
	fmt.Println("format.Any(d) == ", format.Any(d)) 	// "1"
	fmt.Println("format.Any([]int64{x}) == ", format.Any([]int64{x}))	// "[]int64 0x..."
	fmt.Println("format.Any([]time.Duration{d}) == ", format.Any([]time.Duration{d}))  // "[]time.Duration 0x..."
}

func DisplayExample1 () {

	type MetaDataKey struct {
		Name string
	}

	type Movie struct {
		Title, Subtitle string
		Year int
		Color bool
		Actor map[string]string
		Oscars []string
		Sequel *string
		MetaData map[MetaDataKey]string
	}

	metaData := make(map[MetaDataKey]string)
	metaData[MetaDataKey{"description"}] = "Some description here"
	metaData[MetaDataKey{"keywords"}] = "some,keywords,here"

	strangelove := Movie{
		Title: "Dr. Strangelove",
		Subtitle: "How I Learned to Stop Worrying and Love the Bomb",
		Year: 1964,
		Color: false,
		Actor: map[string]string{
			"Dr. Strangelove":				"Peter Sellers",
			"Grp. Capt. Lionel Mandrake":	"Peter Sellers",
			"Pres. Merkin Muffley":			"Peter Sellers",
			"Gen. Buck Turgidson":			"George C. Scott",
			"Brig. Gen. Jack D. Ripper":	"Sterling Hayden",
			"Maj. T.J. \"King\" Kong": 		"Slim Pickens",
		},

		Oscars: []string{
			"Best Actor (Nomin.)",
			"Best Adapted Screenplay (Nomin.)",
			"Best Director (Nomin.)",
			"Best Picture (Nomin.)",
		},
		MetaData: metaData,
	}

	fmt.Println("-----")
	display.Display("strangelove", strangelove, 10)

	fmt.Println("-----")
	display.Display("rV(os.Stderr)", reflect.ValueOf(os.Stderr), 10)

	fmt.Println("-----")
	//Has cyclic reference (prints indefinitely)
	display.Display("os.Stderr", os.Stderr, 10)

	var i interface{} = 3

	fmt.Println("-----")
	display.Display("i", i, 10)

	fmt.Println("-----")
	display.Display("&i", &i, 10)
}

func main () {
	fmt.Println("`format.Any` example:")
	fmt.Println("-----")
	formatAnyExample()
	fmt.Println("`display.Display` example:")
	fmt.Println("-----")
	DisplayExample1()
}
