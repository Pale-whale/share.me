package ui

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func (u *Ui) createActiveServers() *tview.List {
	u.servers.SetBorder(true).SetTitle("Active servers")

	return u.servers
}

func (u *Ui) AddServer(id, path string) {
	u.servers.AddItem(id, path, tcell.RuneRArrow, nil)
}
