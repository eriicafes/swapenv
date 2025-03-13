# swapenv

### Manage .env files locally and effortlessly swap between .env files

`swapenv` is a command-line tool designed to simplify the management of .env files that vary per environment. You can create, list and switch between different .env files which are associated with your git branches. Switching git branches will automatically create a copy of your .env file for that branch.

## Why Use swapenv?

Managing multiple .env files for different environments can be error-prone and time-consuming. `swapenv` centralizes and automates this process by associating .env files with `git` branches. With `swapenv`, you can:
  
- **Quickly switch environments:** Easily toggle between different .env files.
- **Keep environments organized:** Create, list, and remove envs as needed.
- **Switch envs with git:** Automatically switch envs by changing git branches.

## Installation

```sh
go install github.com/eriicafes/swapenv@latest
```

## Usage

`swapenv` provides several commands to manage your .env files. Below is a brief overview of the available commands:

> You can also use it as a git subcommand `git env` after initialization.

```sh
Manage .env files locally. Create and swap between .env files

Usage:
  swapenv [command]

Available Commands:
  commit      Commit .env file contents to the current env
  create      Create new env preset
  help        Help about any command
  init        Initialize swapenv
  list        List all envs
  rm          Remove env
  show        Show the current env
  use         Switch env

Flags:
  -h, --help      help for swapenv
  -v, --version   version for swapenv

Use "swapenv [command] --help" for more information about a command.
```

## Commands Overview

### init
Initializes `swapenv` in your project by creating necessary configuration files, git hook and git subcommand.
`swapenv` stores its data in the `.git/envs` directory of your project.

```sh
swapenv init
```

After initialization you can also use `swapenv` as a `git` subcommand.
```sh
git env version
```

### create
Creates a new env using the .env file or a base env.

```sh
swapenv create dev

# create and use a new env in one command.
swapenv create dev -u
```

### list
Displays all available envs.

```sh
swapenv ls

# list and switch env in interactive mode
swapenv ls -i
```

### use
Switches the .env file contents to one of the available envs.

```sh
swapenv use dev

# switch env creating it if it doesn't already exist
swapenv use dev -b
```

### show
Displays the contents of the current env.

```sh
swapenv show dev

# show without any arguments will display the name of the current env
swapenv show
```

### commit
Commits the .env file contents to the current env, ensuring that any changes are immediately saved.

> `swapenv` automatically commits the env before switching so commit is not necessary.

```sh
  swapenv commit
```

### rm
Removes existing envs.

```sh
swapenv rm dev

# remove multiple envs
swapenv rm prod staging
```