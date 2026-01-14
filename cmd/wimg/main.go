package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/nhdfr/wimg/internal/imgtool"
)

func main() {
	removeOriginal := flag.Bool("r", false, "Remove original file after processing")
	flag.Parse()

	args := flag.Args()

	if len(args) == 0 {
		fmt.Println("Usage:")
		fmt.Println("  wimg source.jpg .webp        # Convert")
		fmt.Println("  wimg source.jpg              # Compress")
		fmt.Println("  wimg -r source.jpg           # Compress & remove original")
		os.Exit(1)
	}

	source := args[0]
	var target string

	if _, err := os.Stat(source); os.IsNotExist(err) {
		log.Fatalf("Error: Source file '%s' does not exist", source)
	}

	if len(args) == 2 {
		targetExt := args[1]
		if !strings.HasPrefix(targetExt, ".") {
			log.Fatalf("Invalid format: must start with '.' (e.g., .png, .webp)")
		}
		target = strings.TrimSuffix(source, filepath.Ext(source)) + targetExt
		err := imgtool.ConvertImage(source, target, targetExt)
		if err != nil {
			log.Fatalf("Error converting image: %v", err)
		}
		fmt.Printf("Converted %s → %s\n", source, target)
	} else {
		compressedPath, err := imgtool.CompressImage(source)
		if err != nil {
			log.Fatalf("Error compressing image: %v", err)
		}
		fmt.Printf("Compressed %s → %s\n", source, compressedPath)

		if *removeOriginal {
			err = os.Remove(source)
			if err != nil {
				log.Printf("Warning: Could not remove original file: %v", err)
			} else {
				err = os.Rename(compressedPath, source)
				if err != nil {
					log.Printf("Warning: Could not rename compressed file: %v", err)
				} else {
					fmt.Printf("Replaced original with compressed version\n")
				}
			}
		}
	}
}
