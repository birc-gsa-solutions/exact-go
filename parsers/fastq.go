package parsers

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

type FastqRecord struct {
	Name string
	Read string
}

var ScanError = fmt.Errorf("Premature end of file inside record")

func scanRecord(s *bufio.Scanner) (*FastqRecord, error) {
	name := s.Text()[1:]
	if !s.Scan() {
		return nil, ScanError
	}
	read := s.Text()

	return &FastqRecord{name, read}, nil
}

type ReadCallback func(*FastqRecord)

func SafeScanFastq(r io.Reader, fn ReadCallback) error {
	s := bufio.NewScanner(r)

	for s.Scan() {
		if rec, err := scanRecord(s); err != nil {
			return err
		} else {
			fn(rec)
		}
	}

	return nil
}

func ScanFastq(fname string, fn ReadCallback) {
	f, err := os.Open(fname)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error opening file:", err.Error())
	}
	err = SafeScanFastq(f, fn)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error scanning file:", err.Error())
	}
	f.Close()
}
