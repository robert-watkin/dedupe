package main

import (
	"fmt"
	"io"
)

type fileDets struct {
	name string
	size int64
	hash []byte
}

// handles main running of the program and any error handling
func run(opts []string, in io.Reader, out io.Writer, stderr io.Writer) (int, error) {
	// walk tree - gather files and sizes
	files, err := walk(opts[0], stderr)
	if err != nil {
		return 1, fmt.Errorf("Error: A problem has occurred during the run function %w", err)
	}

	// sort by size, only keep files where size occurs >2 times
	err = sizeSort(&files, stderr)
	if err != nil {
		return 1, fmt.Errorf("Error: A problem has occurred during the sizeSort function %w", err)
	}

	// sequential hashing
	err = hashSort(&files, stderr)
	if err != nil {
		return 1, fmt.Errorf("Error: A problem has occurred during the hash function %w", err)
	}

	return 0, nil
}
