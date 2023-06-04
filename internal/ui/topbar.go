package ui

import "github.com/rivo/tview"

func (u *Ui) createTopBar() tview.Primitive {
	v := tview.NewTextView().
		SetDynamicColors(false).
		SetRegions(false).
		SetTextAlign(tview.AlignCenter).
		SetText("bite")
	v.SetBorder(true).
		SetTitle("address")
	v.SetText(u.sharing.Addr())
	return v
}
