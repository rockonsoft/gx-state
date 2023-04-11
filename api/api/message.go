package api

import (
	"encoding/json"
	"net/http"

	"rockonsoft.com/gx-state-api/db/models"
	"rockonsoft.com/gx-state-api/lib"

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

	succMessageResponse(message, w)

}
