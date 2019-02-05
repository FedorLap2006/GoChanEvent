package GoChanEvent

type func(Event*) Hfunc

type Event struct {
	Body string
//	CChan *chan Event
	args []interface{}	
//	Handler Hfunc
}
type hEvent {
	Handle Hfunc
	CChan *chan	Event
	Events []string
}

func ELoop(ch *chan Event) {
	for {
		// Do Something
	}
}