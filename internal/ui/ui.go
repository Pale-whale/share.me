package ui

import (
	"github.com/pale-whale/share.me/internal/sharing"
	"github.com/rivo/tview"
)

type Ui struct {
	app     *tview.Application
	sharing *sharing.Server
	infos   *tview.TextView
	servers *tview.List
	preview PreviewPages
}

func CreateUI(server *sharing.Server) *Ui {
	u := &Ui{
		app:     tview.NewApplication(),
		sharing: server,
		infos:   tview.NewTextView(),
		servers: tview.NewList(),
		preview: PreviewPages{
			tview.NewPages(),
			tview.NewTextView(),
			tview.NewImage(),
		},
	}

	top := u.createTopBar()
	info := u.createInfos()
	preview := u.createPreview()
	tree := u.createTreeView()
	servers := u.createActiveServers()

	flex := tview.NewFlex().
		AddItem(tree, 0, 1, false).
		AddItem(tview.NewFlex().SetDirection(tview.FlexRow).
			AddItem(top, 3, 1, false).
			AddItem(preview, 0, 3, false).
			AddItem(info, 11, 1, false), 0, 4, false).
		AddItem(servers, 0, 1, false)

	u.app.SetRoot(flex, true).SetFocus(tree)

	return u
}

func (u *Ui) Run() {
	go u.sharing.Serve()
	if err := u.app.Run(); err != nil {
		panic(err)
	}
}
