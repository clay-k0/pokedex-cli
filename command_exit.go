package main

import (
	"fmt"
	"os"
)

func commandExit(cfg *config) error {
	fmt.Println("\nexiting...")
	os.Exit(0)
	return nil
}
