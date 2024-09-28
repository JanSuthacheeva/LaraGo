package main

import (
  "os"
  "os/exec"
)

/**
 * Installs a new laravel project with the laravel new command.
 */
func InstallLaravel(projectName string) error {
  cmd := exec.Command("laravel", "new", projectName)
  cmd.Stdin = os.Stdin
  cmd.Stdout = os.Stdout
  cmd.Stderr = os.Stderr

  return cmd.Run()
}

/**
 *Installs all the dependencies into the new laravel project.
 */
func InstallDependencies(sail bool) error {
  dependencies := []string{
    "composer require --dev laravel/pint",
    "composer require --dev phpstan/phpstan",
    "composer require --dev larastan/larastan:^2.0",
  }
  if sail {
    dependencies = append(dependencies, "composer require laravel/sail --dev")
    dependencies = append(dependencies, "php artisan sail:install")
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
func WritePreCommitFile() error {
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
func CreateSymbolicLink() error {
  if err := os.Chmod("pre-commit.sample", 777); err != nil {
    return err
  }
  if err := os.Symlink("../../pre-commit.sample", ".git/hooks/pre-commit"); err != nil {
    return err
  }
  return nil
}

func WritePhpStanFile(phpstanLvl int) error {
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

