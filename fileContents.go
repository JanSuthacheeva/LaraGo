package main

import (
  "fmt"
)

func GetPreCommitFileContent() string {
 return `#!/bin/bash
# Get the list of staged files.
staged_files=$(git diff --cached --name-only)

# Run Laravel Pint only on staged files.
./vendor/bin/pint $staged_files
# Check if Pint failed
if [ $? -ne 0 ]; then
    echo "Laravel Pint failed. Commit aborted."
    exit 1
fi

# Restage the modified files.
git add $staged_files

# Run Larastan.
vendor/bin/phpstan analyse

# Check if Larastan failed.
if [ $? -ne 0 ]; then
    echo "Larastan found some issues. Commit aborted."
    exit 1
fi

# If everything is cool, let the commit proceed.
exit 0
`
}

func GetPhpStanFileContent(phpstanLvl *int) string {
  return fmt.Sprintf(`
includes:
    - vendor/larastan/larastan/extension.neon

parameters:

    paths:
      - app/

    # Level 9 is the highest level
    level: %d
    excludePaths:
	- ./app/Http/Controllers/Auth/*
	- ./app/Http/Controllers/ProfileController.php
	- ./app/Http/Requests/ProfileUpdateRequest.php
	- ./*/*/AuthServiceProvider.php
	- ./*/*/NewPasswordController.php
	- ./*/*/RegisteredUserController.php
	- ./*/*/LoginRequest.php

    checkGenericClassInNonGenericObjectType: false
`, phpstanLvl)
}
