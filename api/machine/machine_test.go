package machine

var pendingState = StateNode{Name: "pending", Transitions: []StateTransition{{Target: resolvedState.Name, OnMessage: "RESOLVE"}, {Target: rejectedState.Name, OnMessage: "REJECT"}}}
var resolvedState = StateNode{Name: "resolved"}
var rejectedState = StateNode{Name: "rejected"}
var testStates = []StateNode{pendingState, resolvedState, rejectedState}

const testMachineName = "javascript_promise"

// func GetJavaScriptPromiseMachine() *Machine {

// 	def := MachineDefinition{Id: testMachineName, States: testStates, InitialStateName: pendingState.Name}
// 	currentState := MachinePersistedState{}
// 	defs, _ := json.Marshal(def)
// 	currentStateInBytes, _ := json.Marshal(currentState)

// 	machine, _ := Create(defs, currentStateInBytes)
// 	return machine

// }

// When creating a machine with empty definition a error is returned
// for a valid return value.
// func TestCreateWhenEmptyDefinition(t *testing.T) {
// 	var empty []byte
// 	machine, err := Create(empty, empty)
// 	if machine != nil {
// 		t.Errorf("Expect machine to be nil")
// 	}
// 	if err == nil {
// 		t.Errorf("Expect err to not be nil")
// 	}
// }

// This testcase uses the simple JavaScript promise state machine
// func TestWhenDefinitionJsonIsProvideMachineStateAreDefined(t *testing.T) {
// 	pendingState := StateNode{Name: "pending"}
// 	resolvedState := StateNode{Name: "resolved"}
// 	rejectedState := StateNode{Name: "rejected"}
// 	var testStates = []StateNode{pendingState, resolvedState, rejectedState}
// 	const testMachineName = "test_machine"
// 	def := MachineDefinition{Id: testMachineName, States: testStates}
// 	currentState := MachinePersistedState{}
// 	defs, err := json.Marshal(def)
// 	currentStateInBytes, err := json.Marshal(currentState)

// 	machine, err := Create(defs, currentStateInBytes)
// 	if machine == nil {
// 		t.Errorf("Expect machine to exist")
// 	}
// 	if err != nil {
// 		t.Errorf("Expect err to be nil")
// 	}
// 	if machine.Id != testMachineName {
// 		t.Errorf("Expect machine id to be:%s, but found:%s", testMachineName, machine.Id)

// 	}
// 	states := machine.StateNodes
// 	if l := len(states); l == 0 {
// 		t.Errorf("Expect machine to have:%d states, but found:%d", l, l)
// 	}

// }

// func TestTestMachineInstance(t *testing.T) {
// 	machine := GetJavaScriptPromiseMachine()
// 	if machine.InitialStateName != pendingState.Name {
// 		t.Errorf("Expect initial state name to be set")

// 	}
// 	if machine.InitialStateName != pendingState.Name {
// 		t.Errorf("Expect initial state name to be set")

// 	}
// }

// TestHelloEmpty calls greetings.Hello with an empty string,
// checking for an error.
// func TestHelloEmpty(t *testing.T) {
// 	msg, err := Hello("")
// 	if msg != "" || err == nil {
// 		t.Fatalf(`Hello("") = %q, %v, want "", error`, msg, err)
// 	}
// }
