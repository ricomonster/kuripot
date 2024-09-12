package database

import (
	"fmt"

	scribble "github.com/nanobox-io/golang-scribble"
)

type Database[M any] struct {
	driver *scribble.Driver
	table  string
}

var (
    COLLECTION = "kuripot"
)

func NewDatabase[M any](table string) *Database[*M] {
	// TODO: Create a functionality to setup and determine the home directory
	// and application config
	driver, err := scribble.New("./", nil)
	if err != nil {
		fmt.Println("Error", err)
	}

	return &Database[*M]{driver, table}
}

func (db *Database[M]) Create(data interface{}) (string, error) {
    err := db.driver.Write(COLLECTION, db.table, data)
    if err != nil {
        return "", err
    }

    return "string", nil 
}
