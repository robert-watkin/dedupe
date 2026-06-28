package main

import (
	"fmt"
	"io"
	"io/fs"
	"path/filepath"
)

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
