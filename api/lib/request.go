package lib

type MachineDefinitionRequest struct {
	Definition MachineDefinition `json:"definition"`
}

type MachineInstanceRequest struct {
	MachineType string `json:"machine_type"`
}

type MessageArgs struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type MessageRequest struct {
	To      int64         `json:"to"`
	From    string        `json:"from"`
	Message string        `json:"message"`
	Args    []MessageArgs `json:"args"`
}
