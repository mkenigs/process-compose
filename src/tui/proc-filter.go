package tui

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func (pv *pcView) showProcFilter() {
	const fieldWidth = 50
	f := tview.NewForm()
	f.SetCancelFunc(func() {
		pv.pages.RemovePage(PageDialog)
	})
	f.SetItemPadding(1)
	f.SetBorder(true)
	f.SetFieldBackgroundColor(tcell.ColorLightSkyBlue)
	f.SetFieldTextColor(tcell.ColorBlack)
	f.SetButtonsAlign(tview.AlignCenter)
	f.SetTitle("Search Process")
	f.AddInputField("Search For", pv.logsText.getSearchTerm(), fieldWidth, nil, nil)
	f.AddCheckbox("Case Sensitive", false, nil)
	f.AddCheckbox("Regex", false, nil)
	searchFunc := func() {
		pv.pages.RemovePage(PageDialog)
	}

	f.AddButton("Search", searchFunc)
	f.AddButton("Cancel", func() {
		pv.pages.RemovePage(PageDialog)
	})
	f.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch event.Key() {
		case tcell.KeyEnter:
			searchFunc()
		case tcell.KeyEsc:
			pv.pages.RemovePage(PageDialog)
		default:
			return event
		}
		return nil
	})
	f.SetFocus(0)
	// Display and focus the dialog
	pv.pages.AddPage(PageDialog, createDialogPage(f, fieldWidth+20, 11), true, true)
	pv.appView.SetFocus(f)
}
