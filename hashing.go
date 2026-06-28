package main

import (
	"crypto/sha256"
	"fmt"
	"io"
	"os"
)

func hashSort(files *[]fileDets, stderr io.Writer) error {
	hashCount := make(map[string]int)

	// Use index so we can take address of the real element in the slice the caller owns.
	for i := range *files {
		if err := hashFile(&(*files)[i], hashCount); err != nil {
			fmt.Fprintf(stderr, "WARNING: %v\n", err)
			continue
		}
	}

	var keepers []fileDets
	count := 0
	for _, f := range *files {
		if hashCount[string(f.hash)] >= 2 {
			count++
			keepers = append(keepers, f)
		}
	}
	*files = keepers

	// debugging loop
	// for _, keeper := range keepers {
	// 	fmt.Println(keeper.name, keeper.hash, keeper.size)
	// }

	fmt.Fprintln(stderr, count, "duplicates identified")
	return nil
}

func hashFile(file *fileDets, hashCount map[string]int) error {
	f, err := os.Open(file.name)
	if err != nil {
		return fmt.Errorf("opening %s: %w", file.name, err)
	}
	defer f.Close()

	h := sha256.New()
	if _, err := io.Copy(h, f); err != nil {
		return fmt.Errorf("hashing %s: %w", file.name, err)
	}

	digest := h.Sum(nil)
	file.hash = digest

	key := string(digest)
	hashCount[key]++

	return nil
}
