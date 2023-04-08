# XTran

## Objectives
1. Serialize and deserialize state machine definitions to database

## Machines

## Actions
Actions are triggered when a state is entered or exited 
All actions are treated as messages to Actors, the message is defined as a path string that consist of:{actor}.{module}.{method}
and all messages will contain:
1. The current StateNode
2. The event that caused the transition
3. The context

## Activities
Activities are actors that will do non trivial work while the machine is in a current state
On entering a state, all activity actors will receive a start message.
On exiting a state, all activity actors will receive a stop message.
Activity actors will respond on a start message with a STARTED message.
Activity actors will respond on a stop message with a STOPPED message.
Activity actors could be running the same activity for many machines and will start or stop the activity for a particular machine on the start and stop messages.



When running the service should behave as follows:

New transaction requests:
1. Create a new instance of the state machine requested
2. Set the initial state and persist the instance

Background processing

The server will be in an infinite loop, checking the input table for new records
When a new input is received, it will do the following
1. Find the instance of the state machine 