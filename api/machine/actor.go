package machine

import (
	"encoding/json"

	"github.com/go-pg/pg/v10"
	"rockonsoft.com/gx-state-api/db/models"
	"rockonsoft.com/gx-state-api/lib"
)

type ActionResult struct {
}

type Actor struct {
	Name string
	db   *pg.DB
	Call interface{}
}

func NewActor(name string, db *pg.DB, call interface{}) Actor {
	return Actor{
		Name: name,
		db:   db,
		Call: call,
	}
}

func (actor *Actor) RunCall() error {
	return actor.Call.(func(db *pg.DB) error)(actor.db)
}

func (actor *Actor) SaveCall(machine *lib.Machine, action string, args map[string]json.RawMessage) (*ActionResult, error) {
	actionCall := lib.ActorCall{
		ActorName: actor.Name,
		Action:    action,
		Args:      args,
		MachineId: machine.Id,
	}
	models.CreateCall(actor.db, actionCall)
	return nil, nil
}
