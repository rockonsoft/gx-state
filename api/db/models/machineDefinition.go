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

// func GetComment(db *pg.DB, commentID string) (*Comment, error) {
// 	comment := &Comment{}

// 	err := db.Model(comment).
// 		Relation("User").
// 		Where("comment.id = ?", commentID).
// 		Select()

// 	return comment, err
// }

// func GetComments(db *pg.DB) ([]*Comment, error) {
// 	comments := make([]*Comment, 0)

// 	err := db.Model(&comments).
// 		Relation("User").
// 		Select()

// 	return comments, err
// }

// func UpdateComment(db *pg.DB, req *Comment) (*Comment, error) {
// 	_, err := db.Model(req).
// 		WherePK().
// 		Update()
// 	if err != nil {
// 		return nil, err
// 	}

// 	comment := &Comment{}

// 	err = db.Model(comment).
// 		Relation("User").
// 		Where("comment.id = ?", req.ID).
// 		Select()

// 	return comment, err
// }

// func DeleteComment(db *pg.DB, commentID int64) error {
// 	comment := &Comment{}

// 	err := db.Model(comment).
// 		Relation("User").
// 		Where("comment.id = ?", commentID).
// 		Select()
// 	if err != nil {
// 		return err
// 	}

// 	_, err = db.Model(comment).WherePK().Delete()

// 	return err
// }
