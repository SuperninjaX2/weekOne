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
    fmt.Println("Do you want to sync from outside to inside or inside to outside? (Type 'out' or 'in'):")
    reader := bufio.NewReader(os.Stdin)
    choice, _ := reader.ReadString('\n')
    choice = strings.TrimSpace(choice)

    // Source directory
    var srcDir, destDir string
    if choice == "out" {
        srcDir = "/storage/code/weekOne"
        destDir = "/storage/internal/weekOne"
    } else if choice == "in" {
        srcDir = "/storage/internal/weekOne"
        destDir = "/storage/code/weekOne"
    } else {
        fmt.Println("Invalid choice. Please type 'out' or 'in'.")
        return
    }

    // Sync directories
    err := syncDirs(srcDir, destDir)
    if err != nil {
        fmt.Printf("Error syncing directories: %v\n", err)
        return
    }

    fmt.Println("Directories synchronized successfully!")
}

// Function to sync directories
func syncDirs(src, dest string) error {
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
        srcPath := filepath.Join(src, entry.Name())
        destPath := filepath.Join(dest, entry.Name())

        if entry.IsDir() {
            // If the entry is a directory, recursively sync it
            err := syncDirs(srcPath, destPath)
            if err != nil {
                return err
            }
        } else {
            // If the entry is a file, sync it
            err := syncFile(srcPath, destPath)
            if err != nil {
                return err
            }
        }
    }

    return nil
}

// Function to sync a file
func syncFile(src, dest string) error {
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

    // Copy the content of the source file to the destination file
    _, err = io.Copy(destFile, srcFile)
    if err != nil {
        return err
    }

    // Update the modification time of the destination file to match the source file
    fileInfo, err := os.Stat(src)
    if err != nil {
        return err
    }
    err = os.Chtimes(dest, fileInfo.ModTime(), fileInfo.ModTime())
    if err != nil {
        return err
    }

    return nil
}
