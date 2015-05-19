package termui

import (
	"github.com/boombuler/termui/css"
	"github.com/nsf/termbox-go"
	"sync"
)

var (
	quitChan   chan struct{}
	updateChan chan struct{}
	newBody    chan Element
	setFocus   chan FocusElement
	wgexit     *sync.WaitGroup
	// Events can be used to recieve unhandled termbox events.
	Events <-chan termbox.Event
)

// Start the UI with the given element as root.
func Start(body Element) {
	wgexit = new(sync.WaitGroup)
	wgexit.Add(1)
	termbox.Init()
	eventChan := make(chan termbox.Event)
	updateChan = make(chan struct{})
	quitChan = make(chan struct{})
	newBody = make(chan Element)
	setFocus = make(chan FocusElement)
	go func() {
	loop:
		for {
			select {
			case <-quitChan:
				break loop
			default:
				eventChan <- termbox.PollEvent()
			}
		}

		close(eventChan)
	}()

	events := make(chan termbox.Event)
	Events = events

	input := newInputManager(body)

	go func() {
		defer func() {
			close(events)
			wgexit.Done()
			termbox.Close()

			close(eventChan)
			close(updateChan)
			close(newBody)
			close(setFocus)
			css.ClearCache()
		}()

		body.Arrange(termbox.Size())
		for {
			css.ClearCache()
			termbox.Clear(ForegroundProperty.Get(nil), BackgroundProperty.Get(nil))
			w, h := termbox.Size()
			rootrenderer.RenderChild(body, w, h, 0, 0)

			err := termbox.Flush()
			if err != nil {
				panic(err)
			}

			select {
			case <-quitChan:
				return
			case ev := <-eventChan:
				if ev.Type == termbox.EventResize {
					body.Arrange(termbox.Size())
				}
				if !input.DispatchEvent(ev) {
					events <- ev
				}
			case <-updateChan:
				body.Arrange(termbox.Size())
			case nb := <-newBody:
				body = nb
				input = newInputManager(body)
				body.Arrange(termbox.Size())
			case fe := <-setFocus:
				input.TrySetFocusTo(fe)
				body.Arrange(termbox.Size())
			}
		}
	}()
}

func SetFocus(e FocusElement) {
	go func() {
		setFocus <- e
	}()
}

// SetBody replaces the root element with the new body.
func SetBody(e Element) {
	go func() {
		newBody <- e
	}()
}

// Update tells the UI to update measurement and render again.
func Update() {
	go func() {
		updateChan <- struct{}{}
	}()
}

// Wait for the UI to finish.
func Wait() {
	wgexit.Wait()
}

// Stop shutsdown the UI.
func Stop() {
	go func() {
		close(quitChan)
	}()
}
