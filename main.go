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
  sail := flag.Bool("sail", false, "If set, sail will be installed during the setup process.")
  phpStanLvl := flag.Int("phpstanlvl", 9, "Set the phpstan level.")

  flag.Parse()

  var projectName string
  fmt.Println("Welcome to LaraGo!")
  fmt.Print("Please enter the project name (folder name) that will be created: ")
  _, err := fmt.Scan(&projectName)
  if err != nil {
    log.Fatal(err)
    os.Exit(1)
  }

  if err = InstallLaravel(projectName); err != nil {
    log.Fatal(err)
    os.Exit(1)
  }

  newPath := filepath.Join(".", projectName)

  if err := os.Chdir(newPath); err != nil {
    log.Fatal(err)
    os.Exit(1)
  }

  if err = InstallDependencies(*sail); err != nil {
    log.Fatal(err)
    os.Exit(1)
  }

  if err = WritePreCommitFile(); err != nil {
    log.Fatal(err)
    os.Exit(1)
  }

  _, err = os.Stat(".git")
  if err != nil {
    if errors.Is(err, os.ErrNotExist) {
      fmt.Println("No git repository initialized. Skip creating symbolic link for pre-commit hook.")
    } else {
      log.Fatal(err)
      os.Exit(1)
    }
  } else {
    if err = CreateSymbolicLink(); err != nil {
      log.Fatal(err)
      os.Exit(1)
    }
  }
  if err = WritePhpStanFile(*phpStanLvl); err != nil {
    log.Fatal(err)
    os.Exit(1)
  }
}
