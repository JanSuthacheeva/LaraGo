# LaraGo - Faster setup, smarter code.

LaraGo is a small script that sets up your Laravel project just like the `laravel new` command
does. No really, it actually uses `laravel new` under the hood. However, it adds some features
for cleaner code to the project automatically.

Writing clean code is always a thing to strive for. But in order to install the tools that help
you to write clean code for every project manually, you can simply use LaraGo. On top of the
basic laravel project it also installs `laravel pint`, `phpstan` and `larastan` to the project as dev
dependencies. On top of that, it creates a `phpstan.neon` file in the project directory and if
you selected "initialize git repository" during the installation, it also creates a `pre-commit.sample`
file and a symbolic link in the `.git/hooks` folder. This way, each time you commit new changes
laravel pint formats the `.php` files in your directoy and phpstan checks the changed files for any
"not clean" code.

You can adjust the `phpstan.neon` and `pre-commit.sample` files by yourself and adjust them to your
needs after the project is set up.

## Quick Start

If you just want to use the script to create your laravel projects, download the executable in the
`executables` folder (`larago` for Linux/Mac and `larago.exe` for Windows). After you downloaded the
executable, you need to add it to your `path` in order to execute it from everywhere.

On Mac or Linux, you can type `echo $PATH` in your terminal to see what is in your `PATH`. Simply place the
executable in one of the listed directories or create a new directory, place the executable in it and
add the directory to your `$PATH`.

On Windows, please refer to the countless tutorials in the internet on how to add a file to your `path`.
Simply google something like `How to add something to path on Windows` and countless tutorials pop up.

## Options


## Adjust the script

The script is written in Golang and is provided in this repository. So in case you want to adjust it to your
needs, feel free to create your own fork and develop it further. I assume that you are a little familiar with
Golang if you want to develop this script. It is not that hard and I am learning Golang myself currently so
you should be able to pick it up pretty quickly.

## Upcoming Features

- Adjust phpstan level through option parameter
- Add option to install Laravel sail

