package machine

import (
	"testing"
)

// When creating a machine with empty definition a error is returned
// for a valid return value.
func TestServiceSendMovesStateAlong(t *testing.T) {
	testMachine := GetJavaScriptPromiseMachine()
	service, _ := Interpret(testMachine)
	service.Start() //Todo not sure why this is yet

	resolve := TransitionEvent{Message: "RESOLVE"}
	service.Send(&resolve)

	newState := service.MachineInstance.CurrentState
	if newState.Name != "resolved" {
		t.Errorf("Expect machine state be:%s, but got:%s", "resolved", newState.Name)

	}
}
