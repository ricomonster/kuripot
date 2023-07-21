package tracker

import (
	"fmt"
	"time"

	scribble "github.com/nanobox-io/golang-scribble"
)

type TrackerRepository struct {
	db *scribble.Driver
}

type CreateTrackerOptions struct {
	Name string
}

// Use to instantiate the usage of the Tracker Repository
func NewTrackerRepository() *TrackerRepository {
	db, err := scribble.New("./", nil)
	if err != nil {
		fmt.Println("Error", err)
	}

	return &TrackerRepository{db}
}

func (tr *TrackerRepository) Create(options CreateTrackerOptions) {
	err := tr.db.Write("kuripot", "tracker", Tracker{
		Name:    options.Name,
		Source:  "credit-card",
		Created: time.Now().Unix(),
		Updated: time.Now().Unix(),
	})

	if err != nil {
		fmt.Println("Error", err)
	}
}
