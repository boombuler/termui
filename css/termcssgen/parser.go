package main

import (
	"bytes"
	"fmt"

	"github.com/boombuler/termui/css/internal"
	er "github.com/boombuler/termui/css/internal/errors"
	"github.com/boombuler/termui/css/internal/lexer"
	"github.com/boombuler/termui/css/internal/parser"
)

func parse(data []byte) (internal.Rules, error) {
	lex := lexer.NewLexer(data)
	res, err := parser.NewParser().Parse(lex)
	if err != nil {
		return nil, err
	}
	if E, isError := res.(*er.Error); isError {
		w := new(bytes.Buffer)
		fmt.Fprintf(w, "Error")
		if E.Err != nil {
			fmt.Fprintf(w, " %s\n", E.Err)
		} else {
			fmt.Fprintf(w, "\n")
		}
		fmt.Fprintf(w, "Invalid token: \"%s\"", E.ErrorToken.Lit)
		fmt.Fprintf(w, " at Line %d / Column %d", E.ErrorToken.Pos.Line, E.ErrorToken.Pos.Column)

		return nil, fmt.Errorf("%v", w.String())
	}
	if c, ok := res.(internal.Rules); ok {
		return c, nil
	}
	return nil, fmt.Errorf("Invalid parsing result %v", res)
}
