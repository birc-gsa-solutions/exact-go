package main

import (
	"fmt"
	"os"

	"birc.au.dk/gsa/search"
)

func naive(x, p string, callback func(int)) {
	var i, j int
	for i = 0; i < len(x)-len(p)+1; i++ {
		for j = 0; j < len(p); j++ {
			if x[i+j] != p[j] {
				break
			}
		}

		if j == len(p) {
			callback(i)
		}
	}
}

func main() {
	if len(os.Args) != 3 {
		fmt.Fprintf(os.Stderr, "Usage: genome-file reads-file\n")
		os.Exit(1)
	}
	search.SearchGenome(os.Args[1], os.Args[2], naive)
}
