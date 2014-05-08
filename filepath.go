package main

import (
  "fmt"
  "os"
  "path/filepath"
)

// list all file paths for . directory
// walks into sub-directories recursively
func main() {
  filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
    fmt.Println(path)
    return nil
  })
}
