package engine

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
)

// 1/60th of a second in milliseconds. Ebiten ensures that the game is always run at 60 FPS.
const Delta = (1.0 / 60) * 1000

// The main game struct that contains all worlds and systems
type Game struct {
	worlds       map[string]*World
	activeWorlds []*World

	systems map[string]*System

	entityId int
	entities map[int]*Entity

	// Globals
	RenderWidth, RenderHeight int

	// Hooks
	BeforeUpdate func() error
	BeforeDraw   func(screen *ebiten.Image)
}

func NewGame() *Game {
	return &Game{
		worlds:   map[string]*World{},
		systems:  map[string]*System{},
		entities: map[int]*Entity{},
	}
}

func (g *Game) Run() error {
	return ebiten.RunGame(g)
}

// Entity IDs are unique to a Game
func (g *Game) NewEntity() *Entity {
	e := &Entity{Id: g.entityId}

	g.entityId++

	return e
}

func (g Game) GetEntity(id int) *Entity {
	return g.entities[id]
}

func (g Game) Update() error {
	if g.BeforeUpdate != nil {
		err := g.BeforeUpdate()
		if err != nil {
			return fmt.Errorf("before updating game: %w", err)
		}
	}

	for _, w := range g.activeWorlds {
		err := w.update(g)
		if err != nil {
			return fmt.Errorf("updating game: %w", err)
		}
	}

	return nil
}

func (g Game) Draw(screen *ebiten.Image) {
	if g.BeforeDraw != nil {
		g.BeforeDraw(screen)
	}

	for _, w := range g.activeWorlds {
		w.draw(g, screen)
	}
}

// Ebiten function to optionally change render size at runtime
func (g Game) Layout(_, _ int) (int, int) {
	return g.RenderWidth, g.RenderHeight
}

// Figure out which worlds need activation/deactivation then
// run their associated functions
func (g *Game) ActivateWorlds(worldNames ...string) {
	// Preallocate for performance
	exitingWorlds := make([]*World, 0, len(g.activeWorlds))
	enteringWorlds := make([]*World, 0, len(worldNames))
	newActiveWorlds := make([]*World, 0, len(worldNames))

	for _, w := range g.activeWorlds {
		for _, worldName := range worldNames {
			if w.name != worldName {
				exitingWorlds = append(exitingWorlds, w)
			}
		}
	}

	for _, worldName := range worldNames {
		newActiveWorlds = append(newActiveWorlds, g.worlds[worldName])

		for _, w := range g.activeWorlds {
			if worldName != w.name {
				enteringWorlds = append(enteringWorlds, g.worlds[worldName])
			}
		}
	}

	for _, w := range exitingWorlds {
		w.Exit()
	}

	g.activeWorlds = newActiveWorlds

	for _, w := range enteringWorlds {
		w.Enter()
	}
}

// Add entities to systems based on their components.
func (g Game) UpdateSystems(w World) {
	// Assume entity to be suitable by default
	hasComponents := true

	for _, systemName := range w.systemNames {
		s := g.systems[systemName]

		for _, e := range g.entities {
			// A system w/o components should have no entities
			if len(s.componentNames) == 0 {
				hasComponents = false
			}

			// Check if entity has required components
			for _, componentName := range s.componentNames {
				_, ok := e.components[componentName]
				if !ok {
					hasComponents = false
				}
			}

			// The entity is suitable and we can add it
			if hasComponents {
				s.AddEntity(e)
			}

			// Reset for next iteration
			hasComponents = true
		}
	}
}

func (g *Game) AddSystems(systems ...*System) {
	for _, s := range systems {
		g.systems[s.name] = s
	}
}

func (g *Game) AddWorlds(worlds ...*World) {
	for _, w := range worlds {
		g.worlds[w.name] = w
	}
}

// Add entities to a game and then update the worlds systems.
// This is an expensive function and should be used sparingly.
func (g *Game) AddEntities(w *World, entities ...*Entity) {
	for _, e := range entities {
		g.entities[e.Id] = e
	}

	g.UpdateSystems(*w)
}

type World struct {
	name string

	systemNames []string

	Enter func() error
	Exit  func()
}

func NewWorld(name string, systemNames []string) *World {
	return &World{
		name:        name,
		systemNames: systemNames,
	}
}

func (w World) update(g Game) error {
	for _, systemName := range w.systemNames {
		s := g.systems[systemName]

		if s.Update != nil {
			err := s.Update()
			if err != nil {
				return fmt.Errorf("updating world %s: %w", w.name, err)
			}
		}
	}
	return nil
}

func (w World) draw(g Game, screen *ebiten.Image) {
	for _, systemName := range w.systemNames {
		s := g.systems[systemName]

		if s.Draw != nil {
			s.Draw(screen)
		}
	}
}

type System struct {
	name           string
	EntityIds      []int
	componentNames []string
	Update         func() error
	Draw           func(screen *ebiten.Image)
}

func NewSystem(name string, componentNames []string) *System {
	return &System{
		name:           name,
		componentNames: componentNames,
	}
}

func (s *System) AddEntity(e *Entity) {
	s.EntityIds = append(s.EntityIds, e.Id)
}

type Entity struct {
	Id         int
	components map[string]component
}

func (e *Entity) AddComponent(c component) {
	e.components[c.Name()] = c
}

func GetComponent[T component](e *Entity, c T) (T, bool) {
	component, ok := e.components[c.Name()]
	if !ok {
		return component.(T), false
	}
	return component.(T), true
}

func CursorPosition() (float64, float64) {
	x, y := ebiten.CursorPosition()
	return float64(x), float64(y)
}

type component interface {
	Name() string
}
