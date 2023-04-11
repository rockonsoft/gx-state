package machine

import (
	"encoding/json"
	"fmt"
)

type MachineDefinition struct {
	Id               string
	States           []StateNode
	InitialStateName string
}

// func (def *MachineDefinition) MachineDefinition(id string) {
// 	def.Id = id
// }

type MachinePersistedState struct {
}

type Machine struct {
	Id               string
	StateNodes       []StateNode
	CurrentState     StateNode
	InitialStateName string
}

// create a persisted state machine
// func Persist() (persistedState []byte, err error) {

// }

// creates an instance of a machine
func Create(definition []byte, currentStates []byte) (machine *Machine, err error) {
	var machineDef MachineDefinition
	err = json.Unmarshal(definition, &machineDef)
	if err != nil {
		return nil, err
	}
	var machineState MachinePersistedState

	err = json.Unmarshal(currentStates, &machineState)
	if err != nil {
		return nil, err
	}

	return &Machine{Id: machineDef.Id, StateNodes: machineDef.States, InitialStateName: machineDef.InitialStateName}, nil

}

func (machine *Machine) Transition(newStateName string) {
	states := machine.StateNodes
	for _, state := range states {
		if state.Name == newStateName {
			machine.CurrentState = state
			fmt.Printf("Transitioned to:%s\n", newStateName)
		}
	}
	if machine.CurrentState.Name == "None" {
		panic("Machine does not have an initial state")
	}

}
