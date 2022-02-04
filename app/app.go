package app

import (
	"flag"
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"notebook/app/config"
	"notebook/app/handlers"
	"notebook/app/storage"
)

type App struct {
	DB *gorm.DB
}

func (a *App) Initialize(config *config.Config) {

	db, err := gorm.Open(sqlite.Open(config.DB.DatabasePath), &gorm.Config{SkipDefaultTransaction: true})

	if err != nil {
		log.Panic()
	}
	a.DB = storage.Migrate(db)
}

func (a *App) Run() {

	flagList := flag.Bool("l", false, "List notes")
	flagDelete := flag.Bool("d", false, "Delete note by id entered next")
	flagSearch := flag.Bool("s", false, "Search note by string entered next")

	flag.Parse()

	if *flagList {
		handlers.ListHandler(a.DB)
	} else if *flagDelete {
		handlers.DeleteHandler()
	} else if *flagSearch {
		handlers.SearchHandler()
	} else {
		handlers.CreateHandler(a.DB)
	}
}
