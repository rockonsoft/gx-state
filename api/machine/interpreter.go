package machine

//
type Interpreter struct {
}

func Interpret(machine *Machine) (*MachineService, error) {
	return &MachineService{MachineInstance: machine}, nil
}
