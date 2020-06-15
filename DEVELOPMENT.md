# Development Plan

## Goal

This is where the thinking process behind the project took place.
I'm always looking for constructive criticism.
The main source of template is from [toptal/gitignore](https://github.com/toptal/gitignore)
The project is develop in [Go](https://github.com/golang/go)
[Cobra](https://github.com/spf13/cobra) library.

The goal is to provides a simple way to generate `.gitignore` file.
CLI is no stranger to any developers because of its simplicity and fast.
CLI takes a command, process the input, and returns an output.

Why Cobra?
As a junior, I really don't have much knowledge on why it is good or bad.
I looking for a well maintanced library, good documentation, and lot of examples.

Cobra built on a structure commands, arguments, and flags.
Commands represent actions, Args are things and Flags are modifiers for those actions.

```bash
discard node go react
```

## Process

- [] Fetch `.gitignore` templates
- [] Store templates files in somewhere? localstorage? db?
- [] Take user commands, combinees multiple template files into a single file
- [] Write `.gitignore` file into the current directory

## Minutes

- Use `cobra add list` to add an command `discard list`.
  - Return a list of all the available templates.

