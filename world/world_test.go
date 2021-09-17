package world

import (
	"testing"

	"github.com/leviceccato/clarity/component"
	"github.com/leviceccato/clarity/entity"
	"github.com/leviceccato/clarity/system"
)

func TestUpdateSystem(t *testing.T) {
	t.Run("defaults to no components", func(t *testing.T) {
		w := &world{}
		w.systems = []WorldSystem{
			system.NewDrawSystem(nil),
		}
		w.updateSystems()
		drawSystem := w.systems[0]
		got := drawSystem.EntityCount()
		want := 0
		if got != want {
			t.Errorf("Component count was incorrect, got: %d, wanted: %d", got, want)
		}
	})
	t.Run("adds components", func(t *testing.T) {
		w := &world{}
		w.systems = []WorldSystem{
			system.NewAnimationSystem(),
		}
		// Create player manually because path has to be relative to this dir
		player := entity.NewEntity()
		player.Position = &component.PositionComponent{}
		player.Size = &component.SizeComponent{}
		player.Playable = &component.PlayableComponent{}
		appearance, err := component.NewAppearanceComponent("../assets/player.png", "../assets/player.json")
		if err != nil {
			t.Errorf("Creating appearance component: %s", err)
		}
		player.Appearance = appearance

		w.entities = []*entity.Entity{
			player,
		}
		w.updateSystems()
		animationSystem := w.systems[0]
		got := animationSystem.EntityCount()
		want := 1
		if got != want {
			t.Errorf("Component count was incorrect, got: %d, wanted: %d, components: %s", got, want, animationSystem.Components())
		}
	})
	t.Run("handles incompatible entities", func(t *testing.T) {
		e := entity.NewEntity()
		e.Position = &component.PositionComponent{}
		w := &world{}
		w.systems = []WorldSystem{
			system.NewDrawSystem(nil),
		}
		w.entities = []*entity.Entity{
			e,
		}
		w.updateSystems()
		drawSystem := w.systems[0]
		got := drawSystem.EntityCount()
		want := 0
		if got != want {
			t.Errorf("Component count was incorrect, got: %d, wanted: %d", got, want)
		}
	})
	t.Run("doesn't add entities to systems w/o components", func(t *testing.T) {
		e := entity.NewEntity()
		e.Position = &component.PositionComponent{}
		w := &world{}
		w.systems = []WorldSystem{
			system.NewInputSystem(nil),
		}
		w.entities = []*entity.Entity{
			e,
		}
		w.updateSystems()
		drawSystem := w.systems[0]
		got := drawSystem.EntityCount()
		want := 0
		if got != want {
			t.Errorf("Component count was incorrect, got: %d, wanted: %d", got, want)
		}
	})
}
