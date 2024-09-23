package main

import (
	"errors"
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
}

/**
 * Installs a new laravel project with the laravel new command.
 */
func installLaravel(projectName string) error {
  cmd := exec.Command("laravel", "new", projectName)
  cmd.Stdin = os.Stdin
  cmd.Stdout = os.Stdout
  cmd.Stderr = os.Stderr

  return cmd.Run()
}

func installDependencies() error {
  dependencies := []string{
    "composer require --dev laravel/pint",
    "composer require --dev phpstan/phpstan",
  }

  for _, dep := range dependencies {
    cmd := exec.Command("sh", "-c", dep)
    cmd.Stdin = os.Stdin
    cmd.Stdout = os.Stdout
    cmd.Stderr = os.Stderr

    if err := cmd.Run(); err != nil {
      return err
    }
  }

  return nil
}

// Make this rely of an argument
func writePreCommitFile() error {
  file, err := os.Create("pre-commit.sample")
  if err != nil {
    return err
  }

  _, err = file.WriteString(GetPreCommitHook())
  if err != nil {
    file.Close()
    return err
  }

  err = file.Close()
  if err != nil {
    return err
  }

  return nil
}

func createSymbolicLink() error {
  if err := os.Chmod("pre-commit.sample", 777); err != nil {
    return err
  }
  if err := os.Symlink("../../pre-commit.sample", ".git/hooks/pre-commit"); err != nil {
    return err
  }
  return nil
}

