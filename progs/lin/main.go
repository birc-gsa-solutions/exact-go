package main

import (
	"fmt"
	"os"

	"birc.au.dk/gsa/search"
)

// borderarray computes the border array over the string x. The border
// array ba will at index i have the length of the longest proper border
// of the string x[:i+1], i.e. the longest non-empty string that is both
// a prefix and a suffix of x[:i+1].
func borderarray(x string) []int {
	ba := make([]int, len(x))
	for i := 1; i < len(x); i++ {
		b := ba[i-1]

		for {
			if x[b] == x[i] {
				ba[i] = b + 1
				break
			}

			if b == 0 {
				ba[i] = 0
				break
			}

			b = ba[b-1]
		}
	}

	return ba
}

// strictBorderarray computes the strict border array over the string x.
// This is almost the same as the border array, but ba[i] will be the
// longest proper border of the string x[:i+1] such that x[ba[i]] != x[i].
func strictBorderarray(x string) []int {
	ba := borderarray(x)
	for i := 1; i < len(x)-1; i++ {
		if ba[i] > 0 && x[ba[i]] == x[i+1] {
			ba[i] = ba[ba[i]-1]
		}
	}

	return ba
}

// Kmp runs the O(n+m) time Knuth-Morris-Prat algorithm.
//
// Parameters:
//   - x: the string we search in.
//   - p: the string we search for
//   - callback: a function called for each occurrence
func kmp(x, p string, callback func(int)) {
	ba := strictBorderarray(p)

	var i, j int

	for i < len(x) {
		// Match...
		for i < len(x) && j < len(p) && x[i] == p[j] {
			i++
			j++
		}
		// Report...
		if j == len(p) {
			callback(i - len(p))
		}
		// Shift pattern...
		if j == 0 {
			i++
		} else {
			j = ba[j-1]
		}
	}
}

func main() {
	if len(os.Args) != 3 {
		fmt.Fprintf(os.Stderr, "Usage: genome-file reads-file\n")
		os.Exit(1)
	}
	search.SearchGenome(os.Args[1], os.Args[2], kmp)
}
