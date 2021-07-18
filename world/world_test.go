package world

import (
	"testing"

	"github.com/leviceccato/clarity/entity"
	"github.com/leviceccato/clarity/system"
)

func TestUpdateSystem(t *testing.T) {
	t.Run("defaults to no components", func(t *testing.T) {
		w := &world{}
		w.AddSystem(system.NewDrawSystem())
		w.updateSystems()
		drawSystem := w.systems[0]
		got := drawSystem.GetEntityCount()
		want := 0
		if got != want {
			t.Errorf("Component count was incorrect, got: %d, wanted: %d", got, want)
		}
	})
	t.Run("adds components", func(t *testing.T) {
		w := &world{}
		w.AddSystem(system.NewDrawSystem())
		w.AddEntity(entity.NewPlayerEntity())
		w.updateSystems()
		drawSystem := w.systems[0]
		got := drawSystem.GetEntityCount()
		want := 1
		if got != want {
			t.Errorf("Component count was incorrect, got: %d, wanted: %d", got, want)
		}
	})
}
