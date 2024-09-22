package main

import (
  "fmt"
  "os"
  "os/exec"
  "log"
)

func main() {

  if len(os.Args) < 2 {
    fmt.Println("Please provide a project name.")
    os.Exit(1)
  }

  projectName := os.Args[1]

  cmd := exec.Command("laravel", "new",projectName)
  cmd.Stdin = os.Stdin
  cmd.Stdout = os.Stdout
  cmd.Stderr = os.Stderr

  if err := cmd.Run(); err != nil {
    log.Fatal(err)
    os.Exit(1)
  }
  fmt.Print(cmd.Output())
}
