package database_test

import (
	"testing"

	"github.com/ricomonster/kuripot/internal/platform/database"
)

type TestDatabaseStruct struct {
    Id string
    Name string
    Gender string
}

func TestCreate(t *testing.T) {
    db := database.NewDatabase[TestDatabaseStruct]("test")
    db.Create(TestDatabaseStruct{Id: "some-id", Name: "some-name", Gender: "some-gender"})
}
