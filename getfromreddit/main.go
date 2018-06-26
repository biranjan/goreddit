package main

import (
	"fmt"
	"os"
	"log"
	"github.com/biranjan/reddit"
)

func main() {
	// get the input from command line 
	arg := os.Args[1]

	// Using the input as argument of Get function
	items, err := reddit.Get(arg)
	
	// error handling
	if err != nil {
		log.Fatal(err)
	}
	
	// pint the items
	for _, item := range items {
		fmt.Println(item)
	}
}

