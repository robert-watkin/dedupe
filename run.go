package main

import (
	"fmt"
	"io"
)

type fileDets struct {
	name string
	size int64
}

// handles main running of the program and any error handling
func run(opts []string, in io.Reader, out io.Writer, stderr io.Writer) (int, error) {
	// walk tree - gather files and sizes
	files, err := walk(opts[0], stderr)
	if err != nil {
		return 1, fmt.Errorf("Error: A problem has occurred during the run function %w", err)
	}

	err = sizeSort(&files, stderr)
	if err != nil {
		return 1, fmt.Errorf("Error: A problem has occurred during the run function %w", err)
	}

	return 0, nil
}
