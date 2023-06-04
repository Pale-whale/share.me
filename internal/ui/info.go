package ui

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/rivo/tview"
)

func (u *Ui) createInfos() *tview.TextView {
	u.infos.SetDynamicColors(true).
		SetRegions(true).
		SetChangedFunc(func() {
			u.app.Draw()
		}).
		SetBorder(true).
		SetTitle("Infos")
	return u.infos
}

func (u *Ui) FindInfos(path string) {
	fi, err := os.Stat(path)
	if err != nil {
		return
	}

	infos := fmt.Sprintf(`name:						%s
permissions:			%v
size:						%v
date modified:		%v
`, fi.Name(), fi.Mode(), fi.Size, fi.ModTime())

	file, err := exec.Command("file", path).Output()
	if err != nil {
		u.infos.SetText(infos)
		return
	}

	u.infos.SetText(fmt.Sprintf("%s\n\n%s", infos, file))
}
