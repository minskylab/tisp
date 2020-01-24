package sql

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/minskylab/tisp"
)

func initDB(path string) (*gorm.DB, error) {
	db, err := gorm.Open("sqlite3", path)
	if err != nil {
		return nil, err
	}

	db.AutoMigrate(
		tisp.Client{},
		tisp.Project{},
		tisp.Task{},
		tisp.Resource{},
	)

	return db, nil
}
