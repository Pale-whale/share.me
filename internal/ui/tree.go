package ui

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

var g_selected_node *tview.TreeNode

func addNode(target *tview.TreeNode, path string) {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		panic(err)
	}
	for _, file := range files {
		node := tview.NewTreeNode(file.Name()).
			SetReference(filepath.Join(path, file.Name())).
			SetSelectable(true)
		if file.IsDir() {
			node.SetColor(tcell.ColorBlue)
		}
		target.AddChild(node)
	}
}

func pathFromNode(node *tview.TreeNode) string {
	reference := node.GetReference()
	if reference == nil {
		return "" // Selecting the root node does nothing.
	}

	return reference.(string)
}

func tryExpand(node *tview.TreeNode) {
	path := pathFromNode(node)
	if path == "" {
		return
	}
	fi, err := os.Stat(path)
	if err != nil {
		log.Println("error to fix share.go 190")
	}
	if fi.IsDir() {
		children := node.GetChildren()
		if len(children) == 0 {
			addNode(node, path)
		} else {
			// Collapse if visible, expand if collapsed.
			node.SetExpanded(!node.IsExpanded())
		}
	}
}

func (u *Ui) createTreeView() *tview.TreeView {
	rootDir := "."
	root := tview.NewTreeNode(rootDir).
		SetColor(tcell.ColorRed)
	tree := tview.NewTreeView().
		SetRoot(root).
		SetCurrentNode(root).
		SetGraphics(true)

	tree.SetBorder(true).
		SetTitle("Files")

	tree.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Key() == tcell.KeyRight {
			tryExpand(g_selected_node)
			return nil
		}
		return event
	})

	addNode(root, rootDir)

	tree.SetChangedFunc(func(node *tview.TreeNode) {
		g_selected_node = node
		reference := node.GetReference()
		if reference == nil {
			return
		}
		path := reference.(string)
		u.FindInfos(path)
		go u.SetPreview(path)
	})

	tree.SetSelectedFunc(func(node *tview.TreeNode) {
		path := pathFromNode(node)
		if path != "" {
			id := u.sharing.ServeFile(path, true)
			u.AddServer(id, path)
		}
		return
	})
	return tree
}
