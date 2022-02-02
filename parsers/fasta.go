package parsers

import (
	"fmt"
	"io"
	"os"
	"strings"
)

// MapFasta reads a fasta file from f and maps the function f
// over all records. f is called with the record name and the
// record sequence
func mapFasta(r io.Reader, f func(string, string)) error {
	bytes, err := io.ReadAll(r)
	if err != nil {
		return err
	}

	records := strings.Split(string(bytes), ">")
	if len(records) == 0 {
		// empty, it shouldn't happen, but we will consider
		// it valid...
		return nil
	}

	if records[0] != "" {
		return fmt.Errorf("Expected an empty string before first header")
	}

	for i := 1; i < len(records); i++ {
		lines := strings.Split(records[i], "\n")
		header := strings.TrimSpace(lines[0])
		seq := strings.Join(lines[1:], "")
		f(header, seq)
	}

	return nil
}

// LoadFasta loads a fasta file into a map that maps
// from record names to sequences.
func SafeLoadFasta(r io.Reader) (map[string]string, error) {
	m := map[string]string{}
	err := mapFasta(r, func(name, seq string) {
		m[name] = seq
	})

	if err != nil {
		return nil, err
	}

	return m, nil
}

// This one just does the same as SafeLoadFasta, but it will terminate the
// program if we can't load the file.
func LoadFasta(fname string) map[string]string {
	f, err := os.Open(fname)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error opening file: %s", err.Error())
		os.Exit(1)
	}
	genome, err := SafeLoadFasta(f)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error loading fasta file: %s", err.Error())
		os.Exit(1)
	}
	f.Close()
	return genome
}
