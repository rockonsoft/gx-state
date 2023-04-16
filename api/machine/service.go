package machine

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/go-pg/pg/v10"
	"rockonsoft.com/gx-state-api/db/models"
	"rockonsoft.com/gx-state-api/lib"
	"rockonsoft.com/gx-state-api/system/context"
)

type RegisteredActor struct {
	Name  string
	Actor Actor
}

// MachineService is a service that manages machine instances
type MachineService struct {
	actors []RegisteredActor
	db     *pg.DB
}

func NewMachineService(actors []RegisteredActor, db *pg.DB) *MachineService {
	return &MachineService{
		actors: actors,
		db:     db,
	}
}

func StartMachineService(db *pg.DB) *MachineService {
	actors := []RegisteredActor{
		{"system.context.builder", NewActor("system.context.builder", db, context.Act)},
	}
	service := NewMachineService(actors, db)
	return service
}

func (service *MachineService) Close() error {
	fmt.Sprintln("Shutting down machine service")
	return nil
}

func (service *MachineService) GetActors() []RegisteredActor {
	return service.actors
}

func (service *MachineService) RegisterActor(name string, actor Actor) {
	service.actors = append(service.actors, RegisteredActor{name, actor})
}

func (service *MachineService) Create(machine *lib.Machine) (MemoryMachine, error) {
	memMachine := MemoryMachine{}
	memMachine.persistedMachine = machine
	return memMachine, nil
}

func (service *MachineService) PostMessage(msg *lib.Message, machineModel *lib.Machine) error {
	machine, err := service.Create(machineModel)
	if err != nil {
		return err
	}

	actions, err := machine.Send(msg)
	if err != nil {
		return err
	}

	_, err = machine.Save(service.db)
	if err != nil {
		return err
	}
	for _, action := range actions {
		service.NotifyActor(machine.persistedMachine, action.Actor, action.Action, action.Args)
	}
	_, err = models.UpdateMessageComplete(service.db, msg)
	if err != nil {
		return err
	}

	return nil
}

func (service *MachineService) NotifyActor(machine *lib.Machine, actor string, action string, args map[string]json.RawMessage) error {
	var found bool = false
	for _, a := range service.actors {
		if actor == a.Name {
			found = true
			a.Actor.SaveCall(action, args)
			a.Actor.RunCall()
		}
	}
	if !found {
		return errors.New(fmt.Sprintf("Actor:%s not registred", actor))
	}
	return nil
}

// func (service *MachineService) Start() {
// 	service.MachineInstance.Transition(service.MachineInstance.InitialStateName)
// }

// func (service *MachineService) Send(event *TransitionEvent) {
// 	transitions := service.MachineInstance.CurrentState.Transitions
// 	for _, transition := range transitions {
// 		if transition.OnMessage == event.Message {
// 			service.MachineInstance.Transition(transition.Target)
// 		}
// 	}

// }

// func (service *MachineService) Send(event *TransitionEvent) {
// 	transitions := service.MachineInstance.CurrentState.Transitions
// 	for _, transition := range transitions {
// 		if transition.OnMessage == event.Message {
// 			service.MachineInstance.Transition(transition.Target)
// 		}
// 	}

// }
