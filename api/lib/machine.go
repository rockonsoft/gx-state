package lib

import (
	"encoding/json"
)

type ActionDefinition struct {
	Name   string                     `json:"name"`
	Actor  string                     `json:"actor"`
	Action string                     `json:"action"`
	Args   map[string]json.RawMessage `json:"args"`
}
type MessageAction struct {
	Message string             `json:"message"`
	Actions []ActionDefinition `json:"actions"`
	Target  string             `json:"target"`
}
type MachineState struct {
	Name           string                     `json:"name"`
	FinalState     bool                       `default:"false" json:"final_state"`
	Context        map[string]json.RawMessage `json:"context"`
	MessageActions []MessageAction            `json:"on"`
	EntryAction    ActionDefinition           `json:"entry_actions"`
	ExitAction     ActionDefinition           `json:"exit_actions"`
	Activities     []ActionDefinition         `json:"activities"`
}

type MachineDefinition struct {
	Id            int64                      `json:"id"`
	TypeName      string                     `json:"machine_type"`
	Documentation string                     `json:"documentation"`
	InitialState  string                     `json:"initial_state"`
	Context       map[string]json.RawMessage `json:"context"`
	States        []MachineState             `json:"states"`
}

type MachineDefinitionRequest struct {
	Definition MachineDefinition `json:"definition"`
}

type MachineDefinitionResponse struct {
	Success    bool              `json:"success"`
	Error      string            `json:"error"`
	Definition MachineDefinition `json:"definition"`
}
