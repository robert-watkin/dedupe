package main

import (
	"fmt"
	"io"
	"io/fs"
	"path/filepath"
)

type fileDets struct {
	name string
	size int64
}

// handles main running of the program and any error handling
func run(opts []string, in io.Reader, out io.Writer, stderr io.Writer) (int, error) {
	allFiles, err := walk(opts[0], stderr)
	if err != nil {
		return 1, fmt.Errorf("Error: A problem has occurred during the run function %w", err)
	}

	for _, file := range allFiles {
		fmt.Println(file.name, file.size)
	}

	return 0, nil
}

func walk(rootName string, stderr io.Writer) ([]fileDets, error) {
	var allFiles []fileDets

	err := filepath.WalkDir(rootName, func(path string, d fs.DirEntry, err error) error {
		// warning when the file is unreadable
		if err != nil {
			fmt.Fprintln(stderr, "WARNING: File is unreadable -", err)
			return nil

		}

		// non-regular file gets skipped
		if !d.Type().IsRegular() {
			return nil
		}

		// get file info
		info, err := d.Info()
		if err != nil {
			return err
		}

		allFiles = append(allFiles, fileDets{path, info.Size()})

		return nil
	})
	if err != nil {
		return allFiles, err
	}

	return allFiles, nil
}
