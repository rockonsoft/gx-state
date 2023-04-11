package models

import (
	"encoding/json"
	"time"

	"github.com/go-pg/pg/v10"
	"rockonsoft.com/gx-state-api/lib"
)

type MachineModel struct {
	tableName        struct{}                   `pg:"gx_state.machine_instances"`
	Id               int64                      `pg:"id,pk"`
	TypeName         string                     `pg:"type_name"`
	Definition       map[string]json.RawMessage `pg:"definition"`
	CurrentStateName string                     `pg:"current_state_name"`
	Created          time.Time                  `pg:"created"`
	Updated          time.Time                  `pg:"updated"`
	Context          map[string]json.RawMessage `pg:"context"`
}

func CreateMachineInstance(db *pg.DB, machineDef *lib.MachineDefinition) (*lib.Machine, error) {
	jsonData, err := json.Marshal(machineDef)
	if err != nil {
		return nil, err
	}
	jsonStr := string(jsonData)
	x := map[string]json.RawMessage{}
	json.Unmarshal([]byte(jsonStr), &x)
	machineModel := &MachineModel{
		TypeName:         machineDef.TypeName,
		Definition:       x,
		CurrentStateName: machineDef.InitialState,
		Context:          machineDef.Context,
	}
	_, err = db.Model(machineModel).Insert()
	if err != nil {
		return nil, err
	}
	machine := &lib.Machine{
		Id:               machineModel.Id,
		TypeName:         machineModel.TypeName,
		CurrentStateName: machineModel.CurrentStateName,
		Context:          machineModel.Context,
	}
	return machine, nil
}
