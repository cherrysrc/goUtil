package events

var (
	listeners map[string][]eventCall
)

type eventCall struct {
	parameter interface{}
	function  func(interface{})
}

func Init() {
	listeners = make(map[string][]eventCall)
}

//Trigger calls a given events listeners.
func Trigger(identifier string) {
	var eventCalls []eventCall
	if eventCalls = listeners[identifier]; eventCalls == nil {
		return
	}

	for _, eCall := range eventCalls {
		eCall.function(eCall.parameter)
	}
}

//Adds an event listener.
//Param should be a pointer to a resource that will be needed in your listener.
//It will be passed as the parameter to your callback function.
func AddListener(identifier string, param interface{}, function func(interface{})) {
	listeners[identifier] = append(listeners[identifier], eventCall{param, function})
}
