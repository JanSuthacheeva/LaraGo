package main

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func main() {
  // sail := flag.Bool("sail", false, "If set, sail will be installed during the setup process.")
  phpStanLvl := flag.Int("phpstan-lvl", 9, "Set the phpstan level. Defaults to 9.")

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

  if err := installDependencies(); err != nil {
    log.Fatal(err)
    os.Exit(1)
  }

  if err := writePreCommitFile(); err != nil {
    log.Fatal(err)
    os.Exit(1)
  }

  _, err := os.Stat(".git")
  if err != nil {
    if errors.Is(err, os.ErrNotExist) {
      fmt.Println("No git repository initialized. Skip creating symbolic link for pre-commit hook.")
    } else {
      log.Fatal(err)
      os.Exit(1)
    }
  } else {
    if err := createSymbolicLink(); err != nil {
      log.Fatal(err)
      os.Exit(1)
    }
  }
  if err := writePhpStanFile(phpStanLvl); err != nil {
    log.Fatal(err)
    os.Exit(1)
  }
}
