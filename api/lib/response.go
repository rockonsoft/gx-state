package lib

type MachineErrorResponse struct {
	Success bool   `json:"success"`
	Error   string `json:"error"`
}

type MachineDefinitionResponse struct {
	Success    bool              `json:"success"`
	Error      string            `json:"error"`
	Definition MachineDefinition `json:"definition"`
}

type MachineResponse struct {
	Success bool    `json:"success"`
	Error   string  `json:"error"`
	Machine Machine `json:"machine"`
}

type MessageResponse struct {
	Success bool    `json:"success"`
	Error   string  `json:"error"`
	Message Message `json:"message"`
}
