package api

import (
	"encoding/json"
	"log"
	"net/http"

	"rockonsoft.com/gx-state-api/lib"
)

func handleMachineErr(w http.ResponseWriter, err error) {
	res := &lib.MachineErrorResponse{
		Success: false,
		Error:   err.Error(),
	}
	err = json.NewEncoder(w).Encode(res)
	//if there's an error with encoding handle it
	if err != nil {
		log.Printf("error sending response %v\n", err)
	}
	//return a bad request and exist the function
	w.WriteHeader(http.StatusBadRequest)
}

func handleDBFromContextErr(w http.ResponseWriter) {
	res := &lib.MachineDefinitionResponse{
		Success:    false,
		Error:      "could not get the DB from context",
		Definition: lib.MachineDefinition{},
	}
	err := json.NewEncoder(w).Encode(res)
	//if there's an error with encoding handle it
	if err != nil {
		log.Printf("error sending response %v\n", err)
	}
	//return a bad request and exist the function
	w.WriteHeader(http.StatusBadRequest)
}

func succMachineDefResponse(machineDef *lib.MachineDefinition, w http.ResponseWriter) {
	//return successful response
	res := &lib.MachineDefinitionResponse{
		Success:    true,
		Error:      "",
		Definition: *machineDef,
	}
	//send the encoded response to responsewriter
	err := json.NewEncoder(w).Encode(res)
	if err != nil {
		log.Printf("error encoding comment: %v\n", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	//send a 200 response
	w.WriteHeader(http.StatusOK)
}

func succMachineResponse(machine *lib.Machine, w http.ResponseWriter) {
	//return successful response
	res := &lib.MachineResponse{
		Success: true,
		Error:   "",
		Machine: *machine,
	}
	//send the encoded response to responsewriter
	err := json.NewEncoder(w).Encode(res)
	if err != nil {
		log.Printf("error encoding comment: %v\n", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	//send a 200 response
	w.WriteHeader(http.StatusOK)
}

func succMessageResponse(message *lib.Message, w http.ResponseWriter) {
	//return successful response
	res := &lib.MessageResponse{
		Success: true,
		Error:   "",
		Message: *message,
	}
	//send the encoded response to responsewriter
	err := json.NewEncoder(w).Encode(res)
	if err != nil {
		log.Printf("error encoding comment: %v\n", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	//send a 200 response
	w.WriteHeader(http.StatusOK)
}
