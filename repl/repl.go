package repl

import (
	"bufio"
	"fmt"
	"io"
	"os"

	"wukong.com/evaluator"
	"wukong.com/lexer"
	"wukong.com/object"
	"wukong.com/parser"
)

const PROMPT = "-> "

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)
	env := object.NewEnvironment()

	for {
		fmt.Print(PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}
		line := scanner.Text()

		l := lexer.New(line)
		p := parser.New(l)

		program := p.ParseProgram()

		if len(p.Errors()) != 0 {
			printParseErrors(out, p.Errors())
			continue
		}

		evaluated := evaluator.Eval(program, env)
		if evaluated != nil {
			io.WriteString(out, evaluated.Inspect())
			io.WriteString(out, "\n")
		}
	}

}

func StartScript(in io.Reader, out io.Writer, filePath string) {
	env := object.NewEnvironment()

	fileContent, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Fprintf(out, "Error reading file: %s\n", err)
		return
	}

	l := lexer.New(string(fileContent))
	p := parser.New(l)

	program := p.ParseProgram()
	if len(p.Errors()) != 0 {
		printParseErrors(out, p.Errors())
		return
	}

	evaluated := evaluator.Eval(program, env)
	if evaluated != nil {
		io.WriteString(out, evaluated.Inspect())
		io.WriteString(out, "\n")
	}
}

func printParseErrors(out io.Writer, errors []string) {
	for _, msg := range errors {
		io.WriteString(out, "\t"+msg+"\n")
	}

}
