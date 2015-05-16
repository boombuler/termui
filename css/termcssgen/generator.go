package main

import (
	"bufio"
	"bytes"
	"fmt"
	"github.com/boombuler/termui/css"
	"github.com/boombuler/termui/css/internal"
	"github.com/huin/goutil/codegen"
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
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	fOut, err := codegen.NewGofmtWriteCloser(file)
	if err != nil {
		file.Close()
		return err
	}
	defer fOut.Close()

	w := bufio.NewWriter(fOut)

	w.WriteString(fmt.Sprintf("package %v", packagename) + nl + nl)
	w.WriteString(`import "github.com/boombuler/termui/css"` + nl + nl)
	w.WriteString(`func ` + methodname + `() css.SelectorStyles {`)
	w.WriteString(makeStyleFn)
	w.WriteString(`return css.SelectorStyles {` + nl)
	for _, rule := range content {
		generateRule(w, rule)
	}
	w.WriteString(`}` + nl)
	w.WriteString(`}` + nl)
	return w.Flush()
}

func writeValue(w *bufio.Writer, val interface{}) {
	if strs, ok := val.([]string); ok {
		w.WriteString("[]string {" + nl)
		for _, s := range strs {
			w.WriteString("\"" + s + "\", " + nl)
		}
		w.WriteString("}")
	} else {
		w.WriteString(fmt.Sprintf("%v", val))
	}
}

func generateRule(w *bufio.Writer, r internal.Rule) {
	w.WriteString(`css.SelectorStyle {` + nl)
	w.WriteString(`Selector: `)
	w.WriteString(selectorToStr(r.Selector) + `,` + nl)
	w.WriteString(`Style: makeStyle(` + nl)
	for _, p := range r.Values {
		w.WriteString(fmt.Sprintf(styleParamDef+"{\"%v\", ", p.Name))
		writeValue(w, p.Value)
		w.WriteString("}," + nl)
	}
	w.WriteString(`),` + nl)
	w.WriteString(`},` + nl)
}

func selectorToStr(sel css.Selector) string {
	if sel == css.AnySelector {
		return "css.AnySelector"
	} else if sel == css.BodySelector {
		return "css.BodySelector"
	} else {
		switch s := sel.(type) {
		case css.NameSelector:
			return fmt.Sprintf("css.NameSelector(\"%v\")", string(s))
		case css.ClassSelector:
			return fmt.Sprintf("css.ClassSelector(\"%v\")", string(s))
		case css.IDSelector:
			return fmt.Sprintf("css.IDSelector(\"%v\")", string(s))
		case css.ParentSelector:
			return fmt.Sprintf("css.ParentSelector{%v, %v}", selectorToStr(s.Parent), selectorToStr(s.Child))
		case css.PseudoClassSelector:
			return fmt.Sprintf("css.PseudoClassSelector(\"%v\")", string(s))
		case css.AnyParentSelector:
			return fmt.Sprintf("css.AnyParentSelector{%v, %v}", selectorToStr(s.Parent), selectorToStr(s.Child))
		case css.AndSelector:
			if len(s) == 1 {
				return selectorToStr(s[0])
			}
			b := new(bytes.Buffer)
			b.WriteString("css.AndSelector{")
			b.WriteString(nl)
			for _, sc := range s {
				b.WriteString(selectorToStr(sc))
				b.WriteRune(',')
				b.WriteString(nl)
			}
			b.WriteString("}")
			return b.String()
		default:
			panic("Unknown Selector Type!")
		}
	}
}
