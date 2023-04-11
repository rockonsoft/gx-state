package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"

	"github.com/go-pg/pg/v10"
	"rockonsoft.com/gx-state-api/db/models"
	"rockonsoft.com/gx-state-api/lib"
)

func createMachineDef(w http.ResponseWriter, r *http.Request) {
	//get the request body and decode it
	req := &lib.MachineDefinitionRequest{}
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
	fmt.Println("req.Definition: ", req.Definition.TypeName)
	//if we can get the db then
	// var states []*models.MachineStateDefinition
	// for _, state := range req.States {
	// 	states = append(states, &models.MachineStateDefinition{
	// 		Name:       state.Name,
	// 		FinalState: state.FinalState,
	// 		Context:    state.Context,
	// 	})
	// }

	machineDef, err := models.CreateMachineDefinition(pgdb, &req.Definition)
	if err != nil {
		handleMachineErr(w, err)
		return
	}
	//everything is good
	//let's return a positive response
	succMachineDefResponse(machineDef, w)
}

// getMachineDefs
func getMachineDefsByTyneName(w http.ResponseWriter, r *http.Request) {
	//get the id from the URL parameter
	//alternatively you could use a URL query

	machineTypeName := chi.URLParam(r, "machineName")
	fmt.Println(fmt.Sprintf("machineTypeName: %s", machineTypeName))

	// req := &lib.MachineDefinitionRequest{}
	// err := json.NewDecoder(r.Body).Decode(req)
	// if err != nil {
	// 	handleMachineErr(w, err)
	// 	return
	// }
	// machineTypeName := req.Definition.TypeName

	//get db from ctx
	pgdb, ok := r.Context().Value("DB").(*pg.DB)
	if !ok {
		handleDBFromContextErr(w)
		return
	}

	machineDef, err := models.GetMachineDefinitionByName(pgdb, machineTypeName)

	if err != nil {
		handleMachineErr(w, err)
		return
	}
	//positive response
	succMachineDefResponse(machineDef, w)
}

func getMachine(w http.ResponseWriter, r *http.Request) {
	//get the id from the URL parameter
	//alternatively you could use a URL query
	slugParam := chi.URLParam(r, "slug")

	fmt.Println(fmt.Sprintf("fetching machine: %s", slugParam))

	// machineTypeName := chi.URLParam(r, "slug")
	// fmt.Println(fmt.Sprintf("machineTypeName: %s", machineTypeName))

	// req := &lib.MachineDefinitionRequest{}
	// err := json.NewDecoder(r.Body).Decode(req)
	// if err != nil {
	// 	handleMachineErr(w, err)
	// 	return
	// }
	// machineTypeName := req.Definition.TypeName

	//get db from ctx
	// pgdb, ok := r.Context().Value("DB").(*pg.DB)
	// if !ok {
	// 	handleDBFromContextErr(w)
	// 	return
	// }

	// machineDef, err := models.GetMachineDefinitionByName(pgdb, machineTypeName)

	// if err != nil {
	// 	handleMachineErr(w, err)
	// 	return
	// }
	// //positive response
	// succMachineDefResponse(machineDef, w)
}
