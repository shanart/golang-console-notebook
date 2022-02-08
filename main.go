package main

import (
	"fmt"
	"log"

	"github.com/jroimartin/gocui"
)

var (
	viewArr = []string{"sidebar", "editor"}
	active  = 0
)

func main() {
	// config := config.GetConfig()

	// app := &app.App{}
	// app.Initialize(config)
	// app.Run()

	g, err := gocui.NewGui(gocui.OutputNormal)
	if err != nil {
		log.Panicln(err)
	}
	defer g.Close()

	g.Highlight = true
	g.Cursor = true
	g.SelFgColor = gocui.ColorGreen

	g.SetManagerFunc(layout)

	if err := g.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, quit); err != nil {
		log.Panicln(err)
	}

	if err := g.SetKeybinding("", gocui.KeyTab, gocui.ModNone, nextView); err != nil {
		log.Panicln(err)
	}

	if err := g.MainLoop(); err != nil && err != gocui.ErrQuit {
		log.Panicln(err)
	}
}

func layout(g *gocui.Gui) error {
	maxX, maxY := g.Size()
	if v, err := g.SetView("sidebar", 0, 0, 30, maxY-1); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Title = "Sidebar"
		fmt.Fprintln(v, "Sidebar notes tree")
		// TODO: list of notes
	}

	if v, err := g.SetView("editor", 31, 0, maxX-1, maxY-1); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Title = "Note title"
		v.Editable = true
		v.Highlight = true
		v.Wrap = true

		fmt.Fprintln(v, "Note detail")
		// TODO: add note detail ls like title, date created, size, number of characters
	}
	return nil
}

func quit(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}

func setCurrentViewOnTop(g *gocui.Gui, name string) (*gocui.View, error) {
	if _, err := g.SetCurrentView(name); err != nil {
		return nil, err
	}
	return g.SetViewOnTop(name)
}

func nextView(g *gocui.Gui, v *gocui.View) error {
	nextIndex := (active + 1) % len(viewArr)
	name := viewArr[nextIndex]

	if _, err := setCurrentViewOnTop(g, name); err != nil {
		return err
	}

	if nextIndex == 0 {
		g.Cursor = false
	} else {
		g.Cursor = true
	}

	active = nextIndex
	return nil
}
