package component

type InputComponent struct{}

func NewInputComponent() *InputComponent {
	return &InputComponent{}
}

type InputData struct {
	X, Y int
}

type Control int

const (
	ControlJump Control = iota + 1
	ControlUp
	ControlLeft
	ControlRight
	ControlDown
	ControlMenu
	ControlDebug
	ControlClick
)
