package input

import (
	"time"

	"github.com/veandco/go-sdl2/sdl"
)

var (
	keyStates []uint8

	currentFrameTime  time.Time
	previousFrameTime time.Time
	deltaTime         time.Duration

	mouseX, mouseY int32
	mouseState     uint32

	listeners map[int][]func(*sdl.Event)
)

func Init() {
	keyStates = sdl.GetKeyboardState()
	listeners = make(map[int][]func(*sdl.Event))
}

//Adds a listener to a specific sdl event type
// input.AddListener(sdl.MOUSEBUTTONDOWN, func(arg *sdl.Event) {
// 	event := (*arg).(*sdl.MouseButtonEvent)
// 	fmt.Println(event.Button)
// })
func AddListener(eventType int, fn func(*sdl.Event)) {
	listeners[eventType] = append(listeners[eventType], fn)
}

//Returns the current mouse position
func GetMousePosition() (int32, int32) {
	return mouseX, mouseY
}

//Returns the current state of the mouse
//(which buttons are pressed)
func GetMouseState() uint32 {
	return mouseState
}

//Returns deltatime
func DeltaTime() time.Duration {
	return deltaTime
}

func IsKeyHeld(scancode sdl.Scancode) bool {
	return keyStates[scancode] == 1
}

func LockCursor() {
	sdl.SetRelativeMouseMode(true)
}

func UnlockCursor() {
	sdl.SetRelativeMouseMode(false)
}

//Returns false if the application should end
func Update() bool {
	currentFrameTime = time.Now()
	deltaTime = currentFrameTime.Sub(previousFrameTime)
	mouseX, mouseY, mouseState = sdl.GetMouseState()

	for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
		var functions []func(*sdl.Event)

		switch event.GetType() {
		case sdl.QUIT:
			functions = listeners[sdl.QUIT]
			//Exception
			//Needs to be done now because this case returns
			for i := 0; i < len(functions); i++ {
				functions[i](&event)
			}
			return false

		default:
			functions = listeners[int(event.GetType())]
		}

		//Call all listeners
		for i := 0; i < len(functions); i++ {
			functions[i](&event)
		}
	}

	previousFrameTime = currentFrameTime
	return true
}
