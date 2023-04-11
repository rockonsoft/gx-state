package machine

type TransitionEvent struct {
	Message string
}

type TransitionAction struct {
	Message string
	Target  string
}

type StateTransition struct {
	OnMessage string
	Target    string
}

type StateContext struct {
}

type StateInstance struct {
}
type NodeState struct {
}

type StateNode struct {
	Name          string `default:"Nome"`
	Documentation string
	Transitions   []StateTransition
	// value - the current state value (e.g., {red: 'walk'})
	Value NodeState

	// context - the current context of this state
	Context StateContext

	// event - the event object that triggered the transition to this state
	Event TransitionEvent

	// actions - an array of actions to be executed
	Actions []TransitionAction

	// activities - a mapping of activities to true if the activity started, or false if stopped.
	Activities []string

	// history - the previous State instance
	History []StateInstance

	// meta - any static meta data defined on the meta property of the state node
	Meta string

	// done - whether the state indicates a final state
	Done bool
}
