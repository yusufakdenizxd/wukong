# Wukong Language

Wukong is a simple interpreter written in Go, inspired by the book "Writing an Interpreter in Go". This project aims to provide a basic understanding of how interpreters work and how to implement one from scratch.

## Features

- Basic arithmetic operations
- Variable bindings
- Functions and control flow
- Error handling

## Getting Started

### Prerequisites

- Go 1.16 or higher

### Installation

Clone the repository:

```sh
git clone https://github.com/yusufakdenizxd/wukong.git
cd wukong
```

Build the interpreter:

```sh
go build -o wukong main.go
```

### Usage

You can start the Wukong interpreter by running:

```sh
./wukong
```

You will be greeted with a prompt where you can start typing Wukong code.

### Example

```sh
>> let x = 5;
>> let y = 10;
>> let add = fn(a, b) { a + b };
>> add(x, y);
15
```

## Contributing

Contributions are welcome! Please fork the repository and submit a pull request.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Acknowledgments

- [Writing an Interpreter in Go](https://interpreterbook.com) by Thorsten Ball
- [Go Programming Language](https://golang.org)
