package main

import (
	"cmp"
	"fmt"
	"io"
	"slices"
)

type fileDets struct {
	name string
	size int64
}

// handles main running of the program and any error handling
func run(opts []string, in io.Reader, out io.Writer, stderr io.Writer) (int, error) {
	// walk tree - gather files and sizes
	unsortedFiles, err := walk(opts[0], stderr)
	if err != nil {
		return 1, fmt.Errorf("Error: A problem has occurred during the run function %w", err)
	}

	// count sizes
	counts := make(map[int64]int)
	for _, file := range unsortedFiles {
		counts[file.size]++
	}

	// return only files where the size matches with other files
	var sortedFiles []fileDets
	filesPerGroup := make(map[int64]int)
	for _, file := range unsortedFiles {
		if counts[file.size] >= 2 {
			sortedFiles = append(sortedFiles, file)
			filesPerGroup[file.size]++
		}
	}

	// order based on size - not necessary for calculation but convenient for debugging loop
	slices.SortFunc(sortedFiles, func(a, b fileDets) int {
		return cmp.Compare(a.size, b.size)
	})

	// debugging loop
	for _, file := range sortedFiles {
		fmt.Fprintln(stderr, file.name, file.size)
	}

	// stderr debug log
	fmt.Fprintf(stderr, "%v files scanned, %v candidates in %v size groups\n", len(unsortedFiles), len(sortedFiles), len(filesPerGroup))

	return 0, nil
}
