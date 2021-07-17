package world

type start struct {
	world
}

func NewStartWorld() *start {
	w := &start{}
	w.name = "start"
	return w
}

func (w *start) Load() {

}

func (w *start) Update() {

}

func (w *start) Draw() {

}

func (w *start) Enter() {

}

func (w *start) Exit() {

}
