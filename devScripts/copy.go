package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	// Prompt user for choice
	fmt.Println("Do you want to copy from outside to inside or inside to outside? (Type 'out' or 'in'):")
	reader := bufio.NewReader(os.Stdin)
	choice, _ := reader.ReadString('\n')
	choice = strings.TrimSpace(choice)

	switch choice {
	case "out":
		copyFromOutside()
	case "in":
		copyFromInside()
	default:
		fmt.Println("Invalid choice. Please type 'out' or 'in'.")
	}
}

func copyFromOutside() {
	// Source directory
	srcDir := "/storage/code/weekOne"

	// Destination directory
	destDir := "/storage/internal/weekOne"

	// Call the copy function
	err := copyDir(srcDir, destDir)
	if err != nil {
		fmt.Printf("Error copying directory: %v\n", err)
		return
	}

	fmt.Println("Directory copied from outside to inside successfully!")
}

func copyFromInside() {
	// Source directory
	srcDir := "/storage/internal/weekOne"

	// Destination directory
	destDir := "/storage/code/weekOne"

	// Call the copy function
	err := copyDir(srcDir, destDir)
	if err != nil {
		fmt.Printf("Error copying directory: %v\n", err)
		return
	}

	fmt.Println("Directory copied from inside to outside successfully!")
}

// Function to recursively copy a directory
func copyDir(src, dest string) error {
	// Create destination directory if it does not exist
	err := os.MkdirAll(dest, os.ModePerm)
	if err != nil {
		return err
	}

	// Get the content of the source directory
	entries, err := os.ReadDir(src)
	if err != nil {
		return err
	}

	for _, entry := range entries {
		// Ignore .git directory
		if entry.IsDir() && entry.Name() == ".git" {
			continue
		}

		srcPath := filepath.Join(src, entry.Name())
		destPath := filepath.Join(dest, entry.Name())

		if entry.IsDir() {
			// If the entry is a directory, recursively copy it
			err := copyDir(srcPath, destPath)
			if err != nil {
				return err
			}
		} else {
			// If the entry is a file, copy it
			err := copyFile(srcPath, destPath)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

// Function to copy a file
func copyFile(src, dest string) error {
	srcFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer srcFile.Close()

	destFile, err := os.Create(dest)
	if err != nil {
		return err
	}
	defer destFile.Close()

	_, err = io.Copy(destFile, srcFile)
	if err != nil {
		return err
	}

	return nil
}
