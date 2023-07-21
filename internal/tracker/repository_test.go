package tracker_test

import (
	"testing"

	"github.com/ricomonster/kuripot/internal/tracker"
)

func TestTrackerRepositoryCreate(t *testing.T) {
	tr := tracker.NewTrackerRepository()
	tr.Create(tracker.CreateTrackerOptions{Name: "Test"})
}
