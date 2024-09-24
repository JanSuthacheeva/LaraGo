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
