package machine

type MessagePump struct {
}

// func (pump *MessagePump) PostMessage(msg lib.Message, machine lib.Machine) error {
// 	memMachine, err := pump.service.Create(machine)
// 	if err != nil {
// 		return err
// 	}
// 	actions, err := memMachine.Send(msg)
// 	for _, act := range actions {
// 		pump.service.NotifyActor(act.Actor, act.Action, act.Args)
// 	}
// 	return nil
// }
