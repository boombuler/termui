## Introduction

[![License: MIT](https://img.shields.io/:license-MIT-blue.svg)](http://opensource.org/licenses/MIT)

TermUI is a styleable UI lib for console applications. It consists of two parts:
1) the elements which are the controls which build the UI
2) the css part which allows theming of the elements in a css like manner.

## Elements
For now there are only a few elements:
* `Text` is for basic text outpu
* `Border` draws a border around a given child
* `TextBorder` like a border but it also draws some text on the top left corner of the border.
* `TextBox` allows the user to enter text.
* `VPanel` stacks multiple children vertically
* `Grid` a table like layout for multiple children

### Example

```go
package main

import (
    "github.com/boombuler/termui"
    "github.com/nsf/termbox-go"
)

func main() {
    vPanel := termui.NewVPanel()               // create a new panel
    vPanel.AddChild(
        termui.NewText("Hello"),               // static text "Hello"
        termui.NewTextBox(),                   // and a textbox
        termui.NewText("World"),               // static text "World"
        termui.NewTextBox(),                   // and another textbox
    )

    termui.Start(vPanel)                       // Start the ui rendering.
    go func() {                                // Start a message loop for unhandled events.
        for ev := range termui.Events {
            if ev.Type == termbox.EventKey {
                if ev.Key == termbox.KeyEsc {  // If the user press `Esc`:
                    termui.Stop()              // Stop the UI
                }
            }
        }
    }()

    termui.Wait()                              // Wait for the ui lib to finish.
}
```

## Styles
The style system is a css like engine. As in browsers there are three possible style sources:
1) UserAgent:
   The default style for an element. Should be used if you write new elements for this lib.
2) Designer:
   The Program specific style. Should be used to set the default style of your program.
3) User:
   The User specific style of an element. Can be used to let the user theme the application.


### Generating styles
Program style can be pre-parsed and generated via `go generate`.
To enable this feature first run: `go install github.com/boombuler/termui/css/termcssgen`

After doing this you create a css file with the default style of your program:

```go
package main

//go:generate termcssgen -i=style.css -o=style_gen.go -p=main -m=getDefaultStyle

import (
    "github.com/boombuler/termui/css"
)

func main() {
    css.SetDesignerStyles(getDefaultStyle())
    // add more program code here...
}
```

to generate the go file with the compiled style you just need to invoke `go generate`


### Userdefined styles

It is also possible to load a style at runtime for example:

```go
package main

import (
    "github.com/boombuler/termui"
    "github.com/boombuler/termui/css"
    "github.com/boombuler/termui/css/parser"
    "github.com/nsf/termbox-go"
)

func main() {
    vPanel := termui.NewVPanel()               // create a new panel
    vPanel.AddChild(
        termui.NewText("Hello"),               // static text "Hello"
        termui.NewTextBox(),                   // and a textbox
        termui.NewText("World"),               // static text "World"
        termui.NewTextBox(),                   // and another textbox
    )

    styledef := `

text {
  background: red;
  color: black;
}

vpanel > * {
  gravity: right;
}

textbox:focused {
  background: magenta;
  color: black;
}`
    if parsedStyle, err := parser.Parse([]byte(styledef)); err != nil {
        panic(err)
    } else {
        css.SetUserStyles(parsedStyle)
    }

    termui.Start(vPanel)                       // Start the ui rendering.
    go func() {                                // Start a message loop for unhandled events.
        for ev := range termui.Events {
            if ev.Type == termbox.EventKey {
                if ev.Key == termbox.KeyEsc {  // If the user press `Esc`:
                    termui.Stop()              // Stop the UI
                }
            }
        }
    }()

    termui.Wait()                              // Wait for the ui lib to finish.
}
```


## Todo
* Examples and documentation
* Element Styles
* Mouse support
* Grid
  * Fix ColumnSpan and RowSpan to work with auto-width columns
* More Elements.
  * WrapPanel
  * HPanel
  * Button
  * TabControl
* Modal Dialogs