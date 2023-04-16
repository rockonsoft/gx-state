package context

import (
	"github.com/go-pg/pg/v10"
	"rockonsoft.com/gx-state-api/db/models"
)

const actorName = "system.context.builder"

func Act(db *pg.DB) error {

	// get the next call from the database
	call, err := models.GetCallForActor(db, actorName)
	if err != nil {
		return nil
	}
	var path = call.Args["path"]
	switch call.Action {
	case "increment":

	case "decrement":
	}
	// run the call
	// save the result

	return nil
}
