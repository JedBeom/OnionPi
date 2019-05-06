package models

import "github.com/go-pg/pg"

func CreateTables(db *pg.DB) error {
	for _, model := range []interface{}{&Post{}, &Session{}, &Vote{}} {

		err := db.CreateTable(model, nil)
		if err != nil {
			return err
		}

	}

	return nil
}
