package models

import (
	"encoding/json"
	"time"

	"github.com/go-pg/pg/v10"
	"rockonsoft.com/gx-state-api/lib"
)

type MessageModel struct {
	tableName      struct{}                   `pg:"gx_state.messages"`
	Id             int64                      `pg:"id,pk"`
	To             int64                      `pg:"recipient"`
	From           string                     `pg:"sender"`
	Message        string                     `pg:"message"`
	ProcessedState string                     `pg:"processed_state"`
	Args           map[string]json.RawMessage `pg:"context"`
	Created        time.Time                  `pg:"created"`
	Updated        time.Time                  `pg:"updated"`
}

func CreateMessage(db *pg.DB, message *lib.MessageRequest) (*lib.Message, error) {
	argsMap := map[string]json.RawMessage{}
	argsLen := len(message.Args)
	if argsLen > 0 {
		jsonData, err := json.Marshal(message.Args)
		if err != nil {
			return nil, err
		}
		jsonStr := string(jsonData)
		json.Unmarshal([]byte(jsonStr), &argsMap)
	}

	m := &MessageModel{
		To:             message.To,
		From:           message.From,
		Message:        message.Message,
		ProcessedState: lib.New.String(),
		Args:           argsMap,
		Created:        time.Now(),
		Updated:        time.Now(),
	}
	_, err := db.Model(m).Insert()
	if err != nil {
		return nil, err
	}

	return &lib.Message{
		Id:             m.Id,
		To:             m.To,
		From:           m.From,
		Message:        m.Message,
		ProcessedState: lib.FromString(m.ProcessedState),
		Args:           m.Args,
	}, nil
}

func UpdateMessageComplete(db *pg.DB, message *lib.Message) (*lib.Message, error) {
	msgModel := &MessageModel{
		Id: message.Id,
	}
	err := db.Model(msgModel).WherePK().Select()
	if err != nil {
		return nil, err
	}
	msgModel.ProcessedState = lib.Processed.String()
	msgModel.Updated = time.Now()
	_, err = db.Model(msgModel).WherePK().Update()
	if err != nil {
		return nil, err
	}
	return message, nil
}
