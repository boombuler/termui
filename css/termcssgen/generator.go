package main

import (
	"bufio"
	"bytes"
	"fmt"
	"github.com/boombuler/termui/css"
	"github.com/boombuler/termui/css/internal"
	"os"
)

const nl = "\n"
const styleParamDef = `struct {
     name string
     v interface{}
}`

const makeStyleFn = `
    makeStyle := func (p... ` + styleParamDef + `) css.Style {
        s := make(css.Style)
        for _, prop := range p {
            pp := css.MustFindProperty(prop.name)
            val, err := pp.Convert(prop.v)
            if err != nil {
                panic(err)
            }
            s[pp] = val
        }
        return s
    }` + nl

func generate(filename string, packagename string, methodname string, content internal.Rules) error {
	if packagename == "" {
		packagename = "style"
	}
	fOut, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer fOut.Close()

	w := bufio.NewWriter(fOut)

	w.WriteString(fmt.Sprintf("package %v", packagename) + nl + nl)
	w.WriteString(`import "github.com/boombuler/termui/css"` + nl + nl)
	w.WriteString(`func ` + methodname + `() css.SelectorStyles {` + nl)
	w.WriteString(makeStyleFn)
	w.WriteString(`    return css.SelectorStyles {` + nl)
	for _, rule := range content {
		generate_rule(w, rule)
	}
	w.WriteString(`    }` + nl)
	w.WriteString(`}` + nl)
	return w.Flush()
}

func write_value(w *bufio.Writer, val interface{}) {
	if strs, ok := val.([]string); ok {
		w.WriteString("[]string {" + nl)
		for _, s := range strs {
			w.WriteString("                    \"" + s + "\", " + nl)
		}
		w.WriteString("                }")
	} else {
		w.WriteString(fmt.Sprintf("%v", val))
	}
}

func generate_rule(w *bufio.Writer, r internal.Rule) {
	w.WriteString(`        css.SelectorStyle {` + nl)
	w.WriteString(`            Selector: `)
	w.WriteString(selector_to_str(r.Selector) + `,` + nl)
	w.WriteString(`            Style: makeStyle(` + nl)
	for _, p := range r.Values {
		w.WriteString(fmt.Sprintf("                "+styleParamDef+"{\"%v\", ", p.Name))
		write_value(w, p.Value)
		w.WriteString("}," + nl)
	}
	w.WriteString(`            ),` + nl)
	w.WriteString(`        },` + nl)
}

func selector_to_str(sel css.Selector) string {
	if sel == css.AnySelector {
		return "css.AnySelector"
	} else if sel == css.BodySelector {
		return "css.BodySelector"
	} else if name, ok := sel.(css.NameSelector); ok {
		return fmt.Sprintf("css.NameSelector(\"%v\")", string(name))
	} else if name, ok := sel.(css.ClassSelector); ok {
		return fmt.Sprintf("css.ClassSelector(\"%v\")", string(name))
	} else if name, ok := sel.(css.IdSelector); ok {
		return fmt.Sprintf("css.IdSelector(\"%v\")", string(name))
	} else if ps, ok := sel.(css.ParentSelector); ok {
		return fmt.Sprintf("css.ParentSelector{%v, %v}", selector_to_str(ps.Parent), selector_to_str(ps.Child))
	} else if name, ok := sel.(css.PseudoClassSelector); ok {
		return fmt.Sprintf("css.PseudoClassSelector(\"%v\")", string(name))
	} else if ps, ok := sel.(css.AnyParentSelector); ok {
		return fmt.Sprintf("css.AnyParentSelector{%v, %v}", selector_to_str(ps.Parent), selector_to_str(ps.Child))
	} else if as, ok := sel.(css.AndSelector); ok {
		if len(as) == 1 {
			return selector_to_str(as[0])
		} else {
			b := new(bytes.Buffer)
			b.WriteString("css.AndSelector{")
			b.WriteString(nl)
			for _, s := range as {
				b.WriteString(selector_to_str(s))
				b.WriteRune(',')
				b.WriteString(nl)
			}
			b.WriteString("}")
			return b.String()
		}
	}
	panic("Invalid Selector Type!")
}
