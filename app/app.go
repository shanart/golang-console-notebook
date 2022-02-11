package app

import (
	"fmt"
	"log"

	"github.com/jroimartin/gocui"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"notebook/app/config"
	"notebook/app/formatters"
	"notebook/app/models"
	"notebook/app/storage"
)

type App struct {
	DB    *gorm.DB
	notes []models.Note
}

var (
	viewArr = []string{"sidebar", "editor"}
	active  = 1
	app     = &App{}
)

func Run(config *config.Config) {

	db, err := gorm.Open(sqlite.Open(config.DB.DatabasePath), &gorm.Config{SkipDefaultTransaction: true})

	if err != nil {
		log.Panic()
	}
	app.DB = storage.Migrate(db)

	app.notes = storage.List(app.DB)

	g, err := gocui.NewGui(gocui.OutputNormal)
	if err != nil {
		log.Panicln(err)
	}
	defer g.Close()

	g.Highlight = true
	g.Cursor = true
	g.SelFgColor = gocui.ColorGreen

	g.SetManagerFunc(layout)

	// Global Keybingings
	if err := g.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, quit); err != nil {
		log.Panicln(err)
	}

	if err := g.SetKeybinding("", gocui.KeyTab, gocui.ModNone, nextView); err != nil {
		log.Panicln(err)
	}

	if err := g.SetKeybinding("sidebar", gocui.KeyArrowDown, gocui.ModNone, cursorDown); err != nil {
		log.Panicln(err)
	}

	if err := g.SetKeybinding("sidebar", gocui.KeyArrowUp, gocui.ModNone, cursorUp); err != nil {
		log.Panicln(err)
	}

	if err := g.SetKeybinding("sidebar", gocui.KeyEnter, gocui.ModNone, selectNote); err != nil {
		log.Panicln(err)
	}

	if err := g.MainLoop(); err != nil && err != gocui.ErrQuit {
		log.Panicln(err)
	}
}

func layout(g *gocui.Gui) error {

	maxX, maxY := g.Size()
	if v, err := g.SetView("sidebar", 0, 0, 35, maxY-1); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Title = "Sidebar"
		v.Highlight = true
		v.SelBgColor = gocui.ColorGreen
		v.SelFgColor = gocui.ColorBlack

		// handlers.ListHandler(v, app.DB)
		for _, item := range formatters.ListNotesItems(app.notes) {
			fmt.Fprintln(v, item)
		}
	}

	if v, err := g.SetView("editor", 36, 0, maxX-1, maxY-1); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Title = "Note title"
		v.Editable = true
		v.Highlight = true
		v.Wrap = true

		setCurrentViewOnTop(g, "editor")
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

	// TODO: Save current buffer
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

func cursorDown(g *gocui.Gui, v *gocui.View) error {
	if v != nil {
		cx, cy := v.Cursor()
		ox, oy := v.Origin()

		if err := v.SetCursor(cx, cy+1); err != nil && (cy+oy)+1 < len(app.notes) {
			if err := v.SetOrigin(ox, oy+1); err != nil {
				return err
			}
		}
	}
	return nil
}

func cursorUp(g *gocui.Gui, v *gocui.View) error {
	if v != nil {
		ox, oy := v.Origin()
		cx, cy := v.Cursor()
		if err := v.SetCursor(cx, cy-1); err != nil && oy > 0 {
			if err := v.SetOrigin(ox, oy-1); err != nil {
				return err
			}
		}
	}
	return nil
}

func selectNote(g *gocui.Gui, v *gocui.View) error {

	_, cy := v.Cursor()
	_, oy := v.Origin()

	note := app.notes[cy+oy]
	editor, _ := g.View("editor")
	editor.Clear()
	fmt.Fprintln(editor, note.Content)

	return nil
}
