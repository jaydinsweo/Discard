# Discard

Discard is a Go library for quick and easy .gitignore files for your project.

## ✨ Features

- Clean and simple CLI
- Pulls from the largest collection of `.gitignore` templates, [https://github.com/toptal/gitignore](gitignore.io).
- Generates a single `.gitignore` file from multiple languages

## 📦 Installation

🍺 **Brew**

```zsh
brew install discard
```

🏗 **Go**

```go
go get -u github.com/jaynguyens/Discard
```

## 💻  Usage

```bash
# Generates a gitignore file for node,react, and go project
discard node react go
```

```bash
# Shows full list of available templates
discard list
```

```bash
# Don't know how to use?
# There is help command for you to read
discard -h
```

## Contributing

Pull requests are welcome.
For major changes, please open an issue first to discuss what you would like to change.

Please make sure to update tests as appropriate.

## License

[MIT](https://choosealicense.com/licenses/mit/)
