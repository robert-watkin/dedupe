package main

import (
	"cmp"
	"fmt"
	"io"
	"slices"
)

func sizeSort(unsortedFiles *[]fileDets, stderr io.Writer) error {
	// count sizes
	counts := make(map[int64]int)
	for _, file := range *unsortedFiles {
		counts[file.size]++
	}

	// return only files where the size matches with other files
	var sortedFiles []fileDets
	filesPerGroup := make(map[int64]int)
	for _, file := range *unsortedFiles {
		if counts[file.size] >= 2 {
			sortedFiles = append(sortedFiles, file)
			filesPerGroup[file.size]++
		}
	}

	// order based on size - not necessary for calculation but convenient for debugging loop
	slices.SortFunc(sortedFiles, func(a, b fileDets) int {
		return cmp.Compare(a.size, b.size)
	})
	// Correct way to update the caller's slice (using pointer to header)
	*unsortedFiles = sortedFiles

	// debugging loop
	// for _, file := range sortedFiles {
	// 	fmt.Fprintln(stderr, file.name, file.size)
	// }

	// stderr debug log
	fmt.Fprintf(stderr, "%v files scanned, %v candidates in %v size groups\n", len(*unsortedFiles), len(sortedFiles), len(filesPerGroup))
	return nil
}
