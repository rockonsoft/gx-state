package lib

import "encoding/json"

type ActorCall struct {
	ActorName string                     `json:"actor_name"`
	Action    string                     `json:"action"`
	Args      map[string]json.RawMessage `json:"args"`
}
