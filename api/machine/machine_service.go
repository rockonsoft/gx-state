package machine

type MachineService struct {
	MachineInstance *Machine
}

func (service *MachineService) Start() {
	service.MachineInstance.Transition(service.MachineInstance.InitialStateName)
}

func (service *MachineService) Send(event *TransitionEvent) {
	transitions := service.MachineInstance.CurrentState.Transitions
	for _, transition := range transitions {
		if transition.OnMessage == event.Message {
			service.MachineInstance.Transition(transition.Target)
		}
	}

}
