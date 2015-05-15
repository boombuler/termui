package termui

import (
	"github.com/nsf/termbox-go"
	"sync"
)

var (
	quitChan   chan struct{}
	updateChan chan struct{}
	wgexit     *sync.WaitGroup
	Events     <-chan termbox.Event
)

func Start(body Element) {
	wgexit = new(sync.WaitGroup)
	wgexit.Add(1)
	termbox.Init()
	eventChan := make(chan termbox.Event)
	updateChan = make(chan struct{})
	go func() {
		for e := range termbox.PollEvent() {
			eventChan <- e
		}
		close(eventChan)
	}()

	quitChan = make(chan struct{})
	events := make(chan termbox.Event)
	Events = events

	input := newInputManager(body)

	go func() {
		defer func() {
			close(events)
			wgexit.Done()
			termbox.Close()
		}()

		body.Arrange(body.Measure(termbox.Size()))
		for {
			termbox.Clear(termbox.ColorWhite, termbox.ColorBlue)
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
					body.Arrange(body.Measure(termbox.Size()))
				}
				if !input.DispatchEvent(ev) {
					events <- ev
				}
			case <-updateChan:
				body.Arrange(body.Measure(termbox.Size()))
			}

		}
	}()
}

func Update() {
	go func() {
		updateChan <- struct{}{}
	}()
}

func Wait() {
	wgexit.Wait()
}

func Stop() {
	go func() {
		quitChan <- struct{}{}
	}()
}
