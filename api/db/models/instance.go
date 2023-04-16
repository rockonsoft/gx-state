package models

import (
	"encoding/json"
	"fmt"
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

	var currentState lib.MachineState
	for _, state := range machineDef.States {
		if state.Name == machineModel.CurrentStateName {
			currentState = state
		}
	}
	machine := &lib.Machine{
		Id:               machineModel.Id,
		TypeName:         machineModel.TypeName,
		CurrentStateName: machineModel.CurrentStateName,
		Context:          machineModel.Context,
		CurrentState:     currentState,
	}
	return machine, nil
}

func UpdateMachineInstance(db *pg.DB, machine *lib.Machine) (*lib.Machine, error) {
	machineModel := &MachineModel{
		Id: machine.Id,
	}
	err := db.Model(machineModel).WherePK().Select()
	if err != nil {
		return nil, err
	}
	machineModel.CurrentStateName = machine.CurrentStateName
	machineModel.Context = machine.Context
	machineModel.Updated = time.Now()

	_, err = db.Model(machineModel).WherePK().Update()
	if err != nil {
		return nil, err
	}

	return machine, nil
}

func GetMachineInstanceById(db *pg.DB, id int64) (*lib.Machine, error) {
	fmt.Sprintln(fmt.Sprintf("Fetching machine instance with id %d", id))
	machineModel := &MachineModel{
		Id: id,
	}
	err := db.Model(machineModel).WherePK().Select()
	if err != nil {
		return nil, err
	}

	//TODO: get the machine definition from the the column definition of the machine instance
	machineDef, err := GetMachineDefinitionByName(db, machineModel.TypeName)
	if err != nil {
		return nil, err
	}

	var currentState lib.MachineState
	for _, state := range machineDef.States {
		if state.Name == machineModel.CurrentStateName {
			currentState = state
		}
	}
	machine := &lib.Machine{
		Id:               machineModel.Id,
		TypeName:         machineModel.TypeName,
		CurrentStateName: machineModel.CurrentStateName,
		Context:          machineModel.Context,
		CurrentState:     currentState,
		States:           machineDef.States,
	}
	//get the states

	return machine, nil

}
