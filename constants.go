package main

func GetPreCommitHook() string {
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
