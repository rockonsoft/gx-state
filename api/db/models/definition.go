package models

import (
	"encoding/json"
	"fmt"

	"github.com/go-pg/pg/v10"
	"rockonsoft.com/gx-state-api/lib"
)

type ActionDefinitionModel struct {
	Id        int64                      `pg:"id,pk" json:"id"`
	tableName struct{}                   `pg:"gx_state.action_definitions"`
	Name      string                     `json:"name"`
	Action    string                     `json:"action"`
	Actor     string                     `json:"actor"`
	Args      map[string]json.RawMessage `json:"args"`
	OwnerId   int64                      `json:"owner_id"`
}

type MessageActionModel struct {
	Id        int64    `pg:"id,pk" json:"id"`
	tableName struct{} `pg:"gx_state.message_actions"`
	Message   string   `json:"message"`
	Target    string   `json:"target"`
	StateId   int64    `json:"state_id"`
}

type EntryActionModel struct {
	Id        int64                      `pg:"id,pk" json:"id"`
	tableName struct{}                   `pg:"gx_state.entry_actions"`
	Action    string                     `json:"action"`
	Actor     string                     `json:"actor"`
	Args      map[string]json.RawMessage `json:"args"`
	StateId   int64                      `json:"state_id"`
}

type ExitActionModel struct {
	Id        int64                      `pg:"id,pk" json:"id"`
	tableName struct{}                   `pg:"gx_state.exit_actions"`
	Action    string                     `json:"action"`
	Actor     string                     `json:"actor"`
	Args      map[string]json.RawMessage `json:"args"`
	StateId   int64                      `json:"state_id"`
}

type ActivityModel struct {
	Id        int64                      `pg:"id,pk" json:"id"`
	tableName struct{}                   `pg:"gx_state.activities"`
	Name      string                     `json:"name"`
	Action    string                     `json:"action"`
	Actor     string                     `json:"actor"`
	Args      map[string]json.RawMessage `json:"args"`
	StateId   int64                      `json:"state_id"`
}

type MachineStateDefinitionModel struct {
	Id                  int64                      `pg:"id,pk" json:"id"`
	tableName           struct{}                   `pg:"gx_state.machine_states"`
	MachineDefinitionId int64                      `json:"machine_definition_id"`
	Name                string                     `json:"name"`
	FinalState          bool                       `default:"false" json:"final_state"`
	Context             map[string]json.RawMessage `json:"context"`
}

type MachineDefinitionModel struct {
	tableName     struct{}                   `pg:"gx_state.machine_definitions"`
	Id            int64                      `pg:"id,pk" json:"id"`
	TypeName      string                     `json:"machine_type"`
	Documentation string                     `json:"documentation"`
	InitialState  string                     `json:"initial_state"`
	Context       map[string]json.RawMessage `json:"context"`
}

func CreateMachineDefinition(db *pg.DB, req *lib.MachineDefinition) (*lib.MachineDefinition, error) {
	machineDef := &MachineDefinitionModel{
		TypeName:      req.TypeName,
		Documentation: req.Documentation,
		InitialState:  req.InitialState,
		Context:       req.Context,
	}

	_, err := db.Model(machineDef).Insert()
	if err != nil {
		return nil, err
	}

	machineDefId := machineDef.Id
	req.Id = machineDefId
	fmt.Println(fmt.Sprintf("Inserted Machine Definition %v with id %d", req.TypeName, machineDefId))

	err = db.Model(machineDef).
		Where("type_name = ?", req.TypeName).
		Select()
	// save the entry action
	// save the exit action

	// save the states
	for _, state := range req.States {
		// save the state
		stateModel := MachineStateDefinitionModel{
			Name:                state.Name,
			FinalState:          state.FinalState,
			Context:             state.Context,
			MachineDefinitionId: machineDefId,
		}
		_, err := db.Model(&stateModel).Insert()
		if err != nil {
			return nil, err
		}
		stateId := stateModel.Id
		// save the entry action
		if state.EntryAction.Actor != "" {
			entryAction := EntryActionModel{
				Action:  state.EntryAction.Action,
				Actor:   state.EntryAction.Actor,
				Args:    state.EntryAction.Args,
				StateId: stateId,
			}
			_, err := db.Model(&entryAction).Insert()
			if err != nil {
				return nil, err
			}
		}

		//save the exit action
		if state.ExitAction.Actor != "" {
			exitAction := ExitActionModel{
				Action:  state.ExitAction.Action,
				Actor:   state.ExitAction.Actor,
				Args:    state.ExitAction.Args,
				StateId: stateId,
			}
			_, err := db.Model(&exitAction).Insert()
			if err != nil {
				return nil, err
			}
		}
		// save the activities
		for _, activity := range state.Activities {
			activityModel := ActivityModel{
				Action:  activity.Action,
				Actor:   activity.Actor,
				Args:    activity.Args,
				StateId: stateId,
			}
			_, err := db.Model(&activityModel).Insert()
			if err != nil {
				return nil, err
			}
		}
		// save the action messages
		for _, message := range state.MessageActions {
			messageModel := MessageActionModel{
				Message: message.Message,
				Target:  message.Target,
				StateId: stateId,
			}
			_, err := db.Model(&messageModel).Insert()
			if err != nil {
				return nil, err
			}
			messageId := messageModel.Id
			// save the action messages
			for _, action := range message.Actions {
				actionModel := ActionDefinitionModel{
					Action:  action.Action,
					Actor:   action.Actor,
					Args:    action.Args,
					OwnerId: messageId,
				}
				_, err := db.Model(&actionModel).Insert()
				if err != nil {
					return nil, err
				}
			}
		}

	}

	return req, err
}

func GetMachineDefinitionByName(db *pg.DB, machineTypeName string) (*lib.MachineDefinition, error) {
	fmt.Println(fmt.Sprintf("Getting Machine Definition %v", machineTypeName))
	machineModel := &MachineDefinitionModel{}
	err := db.Model(machineModel).
		Where("type_name = ?", machineTypeName).
		Select()
	if err != nil {
		return nil, err
	}
	//compile the machine definition
	machineDef := &lib.MachineDefinition{
		Id:            machineModel.Id,
		TypeName:      machineModel.TypeName,
		Documentation: machineModel.Documentation,
		InitialState:  machineModel.InitialState,
		Context:       machineModel.Context,
	}
	// get the states
	err = GetMachineStates(db, machineModel, machineDef)
	if err != nil {
		return nil, err
	}
	return machineDef, nil

}

func GetMachineStates(db *pg.DB, machineModel *MachineDefinitionModel, machineDef *lib.MachineDefinition) error {
	fmt.Println(fmt.Sprintf("Getting Machine States for %d", machineModel.Id))
	states := make([]*MachineStateDefinitionModel, 0)
	err := db.Model(&states).Where("machine_definition_id = ?", machineModel.Id).Select()
	if err != nil {
		return err
	}
	fmt.Println(fmt.Sprintf("Found %d Machine States for %d", len(states), machineModel.Id))
	machineDef.States = make([]lib.MachineState, len(states))

	for idx, state := range states {
		fmt.Println(fmt.Sprintf("Getting pieces for state %d", state.Id))
		// get the entry action
		outState := &lib.MachineState{
			Name:       state.Name,
			FinalState: state.FinalState,
			Context:    state.Context,
			// EntryAction:         &lib.ActionDefinition{},
			// ExitAction:          &lib.ActionDefinition{},
			// Activities:     make([]*lib.ActionDefinition, 0),
			// MessageActions: make([]*lib.MessageActionDefinition, 0),
		}
		entryAction := &EntryActionModel{}
		err := db.Model(entryAction).Where("state_id = ?", state.Id).Select()
		if err == nil {
			outState.EntryAction = lib.ActionDefinition{
				Action: entryAction.Action,
				Actor:  entryAction.Actor,
				Args:   entryAction.Args,
			}
		}
		// get the exit action
		exitAction := &ExitActionModel{}
		err = db.Model(exitAction).Where("state_id = ?", state.Id).Select()
		if err == nil {
			outState.ExitAction = lib.ActionDefinition{
				Action: exitAction.Action,
				Actor:  exitAction.Actor,
				Args:   exitAction.Args,
			}
		}
		// get the activities
		activities := make([]*ActivityModel, 0)
		err = db.Model(&activities).Where("state_id = ?", state.Id).Select()
		if err == nil {
			outState.Activities = make([]lib.ActionDefinition, len(activities))
			for idx, activity := range activities {
				outState.Activities[idx] = lib.ActionDefinition{
					Action: activity.Action,
					Actor:  activity.Actor,
					Args:   activity.Args,
				}
			}
		}
		// get the message actions
		messageActions := make([]*MessageActionModel, 0)
		err = db.Model(&messageActions).Where("state_id = ?", state.Id).Select()
		fmt.Println(fmt.Sprintf("Found %d Message Actions for state %d", len(messageActions), state.Id))

		if err == nil {
			outState.MessageActions = make([]lib.MessageAction, len(messageActions))
			for idx, messageAction := range messageActions {
				outState.MessageActions[idx] = lib.MessageAction{
					Message: messageAction.Message,
					Target:  messageAction.Target,
				}
				// get the actions
				actions := make([]*ActionDefinitionModel, 0)
				err = db.Model(&actions).Where("owner_id = ?", messageAction.Id).Select()
				if err != nil {
					return err
				}
				outState.MessageActions[idx].Actions = make([]lib.ActionDefinition, len(actions))
				for actCount, action := range actions {
					outState.MessageActions[idx].Actions[actCount] = lib.ActionDefinition{
						Action: action.Action,
						Actor:  action.Actor,
						Args:   action.Args,
					}
				}
			}
		}
		machineDef.States[idx] = *outState
	}
	return nil
}
