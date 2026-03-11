package readfile

import (
	"fmt"
	"os"
)

func ReadFile(filename *string) []byte {
	// Returns a byte slice ([]byte) and an error
	data, err := os.ReadFile(*filename)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return nil
	}

	return data
}
