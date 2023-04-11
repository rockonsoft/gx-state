package lib

import (
	"encoding/json"
)

type MessageStatus int64

const (
	Unknown        MessageStatus = 0
	New            MessageStatus = 1
	Processed      MessageStatus = 2
	ReturnToSender MessageStatus = 3
)

func (m MessageStatus) String() string {
	switch m {
	case New:
		return "New"
	case Processed:
		return "Processed"
	case ReturnToSender:
		return "ReturnToSender"
	default:
		return "Unknown"
	}
}

func FromString(s string) MessageStatus {
	switch s {
	case "New":
		return New
	case "Processed":
		return Processed
	case "ReturnToSender":
		return ReturnToSender
	default:
		return Unknown
	}
}

type Message struct {
	Id             int64                      `json:"id"`
	To             int64                      `json:"to"`
	From           string                     `json:"from"`
	Message        string                     `json:"message"`
	ProcessedState MessageStatus              `json:"processed_state"`
	Args           map[string]json.RawMessage `json:"args"`
}
