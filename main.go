package main

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
)

func main() {
  sail := flag.Bool("sail", false, "If set, sail will be installed during the setup process.")
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

/**
 *Installs all the dependencies into the new laravel project.
 */
func installDependencies() error {
  dependencies := []string{
    "composer require --dev laravel/pint",
    "composer require --dev phpstan/phpstan",
    "composer require --dev larastan/larastan:^2.0",
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

/**
 * Creates and writes the pre-commit.sample file in the
 * new laravel project.
 */
func writePreCommitFile() error {
  file, err := os.Create("pre-commit.sample")
  if err != nil {
    return err
  }

  _, err = file.WriteString(GetPreCommitFileContent())
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

/**
 * Creates a symbolic link between the pre-commit.sample file
 * in the project directory and the pre-commit file in the .git/hooks directory.
 */
func createSymbolicLink() error {
  if err := os.Chmod("pre-commit.sample", 777); err != nil {
    return err
  }
  if err := os.Symlink("../../pre-commit.sample", ".git/hooks/pre-commit"); err != nil {
    return err
  }
  return nil
}

func writePhpStanFile(phpstanLvl *int) error {
  file, err := os.Create("phpstan.neon")
  if err != nil {
    return err
  }

  _, err = file.WriteString(GetPhpStanFileContent(phpstanLvl))
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

