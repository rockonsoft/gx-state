package machine

import (
	"errors"

	"github.com/go-pg/pg/v10"
	"rockonsoft.com/gx-state-api/db/models"
	"rockonsoft.com/gx-state-api/lib"
)

type MemoryMachine struct {
	persistedMachine *lib.Machine
}

func (machine *MemoryMachine) Save(db *pg.DB) (*lib.Machine, error) {

	//save the machine

	_, err := models.UpdateMachineInstance(db, machine.persistedMachine)

	if err != nil {
		return nil, err
	}

	// update the event history
	for _, message := range machine.persistedMachine.EventHistory {
		models.UpdateMessageComplete(db, &message)
	}

	return machine.persistedMachine, nil
}

func (machine *MemoryMachine) Send(msg *lib.Message) ([]lib.ActionDefinition, error) {

	var actions []lib.ActionDefinition
	var target string
	for _, messageAction := range machine.persistedMachine.CurrentState.MessageActions {
		if messageAction.Message == msg.Message {
			actions = messageAction.Actions
			target = messageAction.Target
		}
	}
	//if target - transition to target state
	// var transitionActions []lib.ActionDefinition
	if target != "" {
		for _, state := range machine.persistedMachine.States {
			if state.Name == target {
				//This is a state transition
				transitionActions, err := machine.Transition(state)
				machine.persistedMachine.EventHistory = append(machine.persistedMachine.EventHistory, *msg)
				return transitionActions, err
			}
		}
	}
	//no target - must execute actions

	if len(actions) == 0 {
		return []lib.ActionDefinition{}, errors.New("No actions found for message")
	}

	return actions, nil
	// transitions := machine.persistedMachine.CurrentState.Transitions
	// for _, transition := range transitions {
	// 	if transition.OnMessage == event.Message {
	// 		machine.persistedMachine.Transition(transition.Target)
	// 	}
	// }
	// return []lib.ActionDefinition{}, nil
}

func (machine *MemoryMachine) Transition(targetState lib.MachineState) ([]lib.ActionDefinition, error) {
	machine.persistedMachine.CurrentState = targetState
	machine.persistedMachine.CurrentStateName = targetState.Name

	//start the actions and activities of the new state
	var actions []lib.ActionDefinition

	actions = append(actions, targetState.EntryAction)
	for _, act := range targetState.Activities {
		actions = append(actions, act)
	}
	return actions, nil

}

// create a persisted state machine
// func Persist() (persistedState []byte, err error) {

// }

// creates an instance of a machine
// func Create(definition []byte, currentStates []byte) (machine *Machine, err error) {
// 	var machineDef MachineDefinition
// 	err = json.Unmarshal(definition, &machineDef)
// 	if err != nil {
// 		return nil, err
// 	}
// 	var machineState MachinePersistedState

// 	err = json.Unmarshal(currentStates, &machineState)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return &Machine{Id: machineDef.Id, StateNodes: machineDef.States, InitialStateName: machineDef.InitialStateName}, nil

// }

// func (machine *Machine) Transition(newStateName string) {
// 	states := machine.StateNodes
// 	for _, state := range states {
// 		if state.Name == newStateName {
// 			machine.CurrentState = state
// 			fmt.Printf("Transitioned to:%s\n", newStateName)
// 		}
// 	}
// 	if machine.CurrentState.Name == "None" {
// 		panic("Machine does not have an initial state")
// 	}

// }
