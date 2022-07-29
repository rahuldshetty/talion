package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/rahuldshetty/talion/eval"
	"github.com/rahuldshetty/talion/lexer"
	"github.com/rahuldshetty/talion/parser"
)

const ERROR_MESSAGE = ` 
                                    		\ / _
                                      ___,,,
                                      \_[o o]
     There is an Error !              C\  _\/
             /                     _____),_/__
        ________                  /     \/   /
      _|       .|                /      o   /
     | |       .|               /          /
      \|       .|              /          /
       |________|             /_        \/
       __|___|__             _//\        \
 _____|_________|____       \  \ \        \
                    _|       ///  \        \
                   |               \       /
                   |               /      /
                   |              /      /
 ________________  |             /__    /_
 b'ger        ...|_|.............. /______\.......
`

const PROMPT = ">> "

func Start(in io.Reader, out io.Writer){
	scanner := bufio.NewScanner(in)
	for {
		fmt.Print(PROMPT)
		scanned := scanner.Scan()
		if !scanned{
			return
		}

		line := scanner.Text()
		l := lexer.New(line)
		p := parser.New(l)

		program := p.ParseProgram()
		if len(p.Errors()) != 0 {
			printParserErrors(out, p.Errors())
			continue
		}
		
		evaluated := eval.Eval(program)
		if evaluated != nil{
			io.WriteString(out, evaluated.Inspect())
			io.WriteString(out, "\n")
		}
	}
}

func printParserErrors(out io.Writer,errors []string){
	io.WriteString(out, ERROR_MESSAGE)
	io.WriteString(out, "Parser Errors:\n")
	for _, msg := range errors{
		io.WriteString(out, "\t"+msg+"\n")
	}
}	