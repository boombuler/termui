package parser

//go:generate gocc -p=github.com/boombuler/termui/css/internal -o=../internal css.bnf

import (
	"bytes"
	"fmt"
	"github.com/boombuler/termui/css"
	"github.com/boombuler/termui/css/internal"
	er "github.com/boombuler/termui/css/internal/errors"
	"github.com/boombuler/termui/css/internal/lexer"
	"github.com/boombuler/termui/css/internal/parser"
)

func parseBase(data []byte) (internal.Rules, error) {
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

func Parse(data []byte) (css.SelectorStyles, error) {
	rules, err := parseBase(data)
	if err != nil {
		return nil, err
	}
	result := make(css.SelectorStyles, 0)
	for _, rule := range rules {
		selector := rule.Selector
		style := make(css.Style)

		for _, prop := range rule.Values {
			p := css.FindProperty(prop.Name)
			if p == nil {
				return nil, fmt.Errorf("Invalid Property: %v", prop.Name)
			}
			style[p], err = p.Convert(prop.Value)
			if err != nil {
				return nil, err
			}
		}
		result = append(result, css.SelectorStyle{selector, style})
	}

	return nil, nil
}
