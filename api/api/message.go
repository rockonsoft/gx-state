package api

import (
	"encoding/json"
	"net/http"

	"rockonsoft.com/gx-state-api/db/models"
	"rockonsoft.com/gx-state-api/lib"
	"rockonsoft.com/gx-state-api/machine"

	"github.com/go-pg/pg/v10"
)

func createMessage(w http.ResponseWriter, r *http.Request) {
	//get the request body and decode it
	req := &lib.MessageRequest{}
	err := json.NewDecoder(r.Body).Decode(req)
	//if there's an error with decoding the information
	//send a response with an error
	if err != nil {
		handleMachineErr(w, err)
		return
	}
	//get the db from context
	pgdb, ok := r.Context().Value("DB").(*pg.DB)
	//if we can't get the db let's handle the error
	//and send an adequate response
	if !ok {
		handleDBFromContextErr(w)
		return
	}

	message, err := models.CreateMessage(pgdb, req)
	if err != nil {
		handleMachineErr(w, err)
		return
	}

	service, ok := r.Context().Value("Service").(*machine.MachineService)
	if !ok {
		handleDBFromContextErr(w)
		return
	}

	//get the intended machine and let it process the message
	machineModel, err := models.GetMachineInstanceById(pgdb, message.To)
	if err != nil {
		handleMachineErr(w, err)
		return
	}

	err = service.PostMessage(message, machineModel)
	if err != nil {
		handleMachineErr(w, err)
		return
	}

	succMessageResponse(message, w)
}
