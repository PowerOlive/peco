package peco

import "time"

// View handles the drawing/updating the screen
type View struct {
	*Ctx
	layout Layout
}

// PagingRequest can be sent to move the selection cursor
type PagingRequest int

const (
	// ToLineAbove moves the selection to the line above
	ToLineAbove PagingRequest = iota
	// ToNextPage moves the selection to the next page
	ToNextPage
	// ToLineBelow moves the selection to the line below
	ToLineBelow
	// ToPrevPage moves the selection to the previous page
	ToPrevPage
)

// Loop receives requests to update the screen
func (v *View) Loop() {
	defer v.ReleaseWaitGroup()
	for {
		select {
		case <-v.LoopCh():
			return
		case m := <-v.StatusMsgCh():
			v.printStatus(m.DataString())
			m.Done()
		case m := <-v.ClearStatusCh():
			v.clearStatus(m.DataInterface().(time.Duration))
			m.Done()
		case r := <-v.PagingCh():
			v.movePage(r.DataInterface().(PagingRequest))
			r.Done()
		case lines := <-v.DrawCh():
			v.drawScreen(lines.DataInterface().([]Match))
			lines.Done()
		}
	}
}

func (v *View) printStatus(m string) {
	v.layout.PrintStatus(m)
}

func (v *View) clearStatus(d time.Duration) {
	v.layout.ClearStatus(d)
}

func (v *View) drawScreenNoLock(targets []Match) {
	if targets == nil {
		if current := v.current; current != nil {
			targets = v.current
		} else {
			targets = v.lines
		}
	}

	v.layout.DrawScreen(targets)
	// FIXME
	v.current = targets
}

func (v *View) drawScreen(targets []Match) {
	v.mutex.Lock()
	defer v.mutex.Unlock()
	v.drawScreenNoLock(targets)
}

func (v *View) movePage(p PagingRequest) {
	v.mutex.Lock()
	defer v.mutex.Unlock()

	v.layout.MovePage(p)
	v.drawScreenNoLock(nil)
}
