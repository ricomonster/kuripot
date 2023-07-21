package database

import (
	"fmt"

	scribble "github.com/nanobox-io/golang-scribble"
)

type Database[M any] struct {
	driver *scribble.Driver
	table  string
}

func NewDatabase[M any](table string) *Database[*M] {
	// TODO: Create a functionality to setup and determine the home directory
	// and application config
	driver, err := scribble.New("./", nil)
	if err != nil {
		fmt.Println("Error", err)
	}

	return &Database[*M]{driver, table}
}

func (db *Database[M]) Create(data *M) (int, error) {
	return 1, nil
}
