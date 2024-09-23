package main

import (
  "fmt"
  "log"
  "os"
  "os/exec"
  "path/filepath"
)

func main() {

  if len(os.Args) < 2 {
    fmt.Println("Please provide a project name.")
    os.Exit(1)
  }

  projectName := os.Args[1]


  if err := installLaravel(projectName); err != nil {
    log.Fatal(err)
    os.Exit(1)
  }

  newPath := filepath.Join(".", projectName)

  if err := os.Chdir(newPath); err != nil {
    log.Fatal(err)
    os.Exit(1)
  }
}

func installLaravel(projectName string) error {
  cmd := exec.Command("laravel", "new", projectName)
  cmd.Stdin = os.Stdin
  cmd.Stdout = os.Stdout
  cmd.Stderr = os.Stderr

  return cmd.Run()
}
