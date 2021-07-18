package system

type draw struct {
	system
}

func NewDrawSystem() *draw {
	s := &draw{}
	s.components = []string{
		"Appearance",
		"Position",
		"Size",
	}
	return s
}

func (s *draw) Load() {

}

func (s *draw) Update() {

}

func (s *draw) Draw() {

}

func (s *draw) Enter() {

}

func (s *draw) Exit() {

}
