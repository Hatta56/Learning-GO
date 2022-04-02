package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	_, err := IsFileExists("file.txt")
	if err != nil {
		fmt.Println(err)
		if errWrap := errors.Unwrap(err); errWrap != nil {
			fmt.Println(errWrap)
		}
	}
}
func IsFileExists(filename string) (bool, error) {
	file, err := os.Open(filename)
	defer file.Close()
	if err != nil {
		return false, fmt.Errorf("Error in IsFileExists, err: %w", err)
	}

	return true, nil
}
