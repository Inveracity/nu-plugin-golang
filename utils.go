package main

import (
	"fmt"
	"os"
)

// Print to stderr to avoid NuShell parsing it
func Debug(out []byte) {
	fmt.Fprintln(os.Stderr, string(out))
}

// Print to stdout for NuShell to pickup
func Return(out []byte) {
	fmt.Print(string(out))
	os.Stdout.Sync()
}

func Sendencoding() {
	fmt.Printf("%c", 4)
	for _, ch := range "json" {
		fmt.Printf("%c", ch)
	}
	os.Stdout.Sync()
}
