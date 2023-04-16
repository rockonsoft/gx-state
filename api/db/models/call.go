package models

import (
	"encoding/json"

	"github.com/go-pg/pg/v10"
	"rockonsoft.com/gx-state-api/lib"
)

type RemoteCall struct {
	tableName      struct{}                   `pg:"gx_state.remote_call"`
	Id             int64                      `pg:"id,pk"`
	Actor          string                     `pg:"actor"`
	Action         string                     `pg:"action"`
	Args           map[string]json.RawMessage `pg:"args"`
	Result         map[string]json.RawMessage `pg:"result"`
	ProcessedState string                     `pg:"processed_state"`
}

func SaveCall(db *pg.DB, call lib.ActorCall) error {
	callModel := &RemoteCall{
		Actor:          call.ActorName,
		Action:         call.Action,
		ProcessedState: lib.New.String(),
		Args:           call.Args,
	}
	_, err := db.Model(callModel).Insert()
	//TODO: think about the return value
	return err
}

func GetCallForActor(db *pg.DB, actor string) (*RemoteCall, error) {
	call := &RemoteCall{}
	err := db.Model(call).
		Where("actor = ?", actor).
		Where("processed_state = ?", lib.New.String()).
		Order("id ASC").
		Limit(1).
		Select()
	if err != nil {
		return nil, err
	}
	return call, nil
}

func (r *RemoteCall) GetResult() map[string]interface{} {
	result := map[string]interface{}{}
	for k, v := range r.Result {
		var x interface{}
		json.Unmarshal(v, &x)
		result[k] = x
	}
	return result
}
func (r *RemoteCall) GetArgs() map[string]interface{} {
	args := map[string]interface{}{}
	for k, v := range r.Args {
		var x interface{}
		json.Unmarshal(v, &x)
		args[k] = x
	}
	return args
}
