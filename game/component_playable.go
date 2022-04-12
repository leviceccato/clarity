package game

// Playable component
type playableComponent struct{}

func (_ playableComponent) Name() string {
	return "playable"
}
