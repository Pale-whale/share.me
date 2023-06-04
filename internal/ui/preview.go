package ui

import (
	"image"
	"image/jpeg"
	"image/png"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	"github.com/rivo/tview"
)

type PreviewPages struct {
	*tview.Pages

	text *tview.TextView
	img  *tview.Image
}

func (u *Ui) createPreview() tview.Primitive {
	u.preview.SetBorder(true).
		SetTitle("Preview").
		SetTitleAlign(tview.AlignCenter)

	u.preview.text.SetDynamicColors(true).
		SetRegions(true).
		SetTextAlign(tview.AlignLeft)
	u.preview.AddPage("text", u.preview.text, true, false)

	u.preview.img.SetColors(0).SetDithering(tview.DitheringFloydSteinberg)
	u.preview.AddPage("image", u.preview.img, true, false)

	u.preview.AddPage("modale", tview.NewModal().SetText("Loading"), false, false)
	return u.preview
}

func (u *Ui) SetPreview(path string) {
	f, _ := os.Open(path)
	defer u.app.Draw()

	body, _ := ioutil.ReadAll(f)
	f.Seek(0, 0)
	c := http.DetectContentType(body)
	defer f.Close()

	cType := strings.Split(c, ";")[0]
	tpe := strings.Split(cType, "/")

	u.preview.SwitchToPage(tpe[0])
	switch tpe[0] {
	case "text":
		u.preview.text.SetText(string(body))
		return
	case "image":
		var photo image.Image
		switch tpe[1] {
		case "jpeg":
			photo, _ = jpeg.Decode(f)
		case "png":
			photo, _ = png.Decode(f)
		}
		u.preview.img.SetImage(photo)
		return
	default:
	}
}
