package context

import (
	"encoding/json"

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
	for call != nil {
		err := processActions(db, call)
		if err != nil {
			return err
		}
		call, err = models.GetCallForActor(db, actorName)
		if err != nil {
			return err
		}

	}

	return nil
}

func processActions(db *pg.DB, call *models.RemoteCall) error {
	machine, err := models.GetMachineInstanceById(db, call.MachineId)
	if err != nil {
		return err
	}
	context := machine.Context

	var path = call.Args["path"]

	var pathStr string
	json.Unmarshal(path, &pathStr)

	var value = context[pathStr]

	switch call.Action {
	case "increment":
		{
			//TODO - if these functions fails, it should not stop the whole process
			var valueInt int
			err := json.Unmarshal(value, &valueInt)
			if err != nil {
				return err
			}
			valueInt++
			j, err := json.Marshal(valueInt)
			if err != nil {
				return err
			}
			context[pathStr] = j
			machine.Context = context
			models.UpdateMachineInstance(db, machine)
		}

	case "decrement":
		{

			var valueInt int
			err := json.Unmarshal(value, &valueInt)
			if err != nil {
				return err
			}
			valueInt--
			j, err := json.Marshal(valueInt)
			if err != nil {
				return err
			}
			context[pathStr] = j
			machine.Context = context
			models.UpdateMachineInstance(db, machine)

		}
	}

	err = call.UpdateResult(db, context)
	return err
}
