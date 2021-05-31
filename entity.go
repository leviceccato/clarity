package main

import (
	r "github.com/gen2brain/raylib-go/raylib"
)

var entityId int32 = 0

type entity struct {
	id         int32
	appearance *appearanceComponent
	collision  *collisionComponent
	physics    *physicsComponent
	position   *positionComponent
	controls   *controlsComponent
}

func newEntity() *entity {
	e := &entity{
		id: entityId,
	}
	entityId += 1
	return e
}

func newPlayerEntity() *entity {
	e := newEntity()
	e.appearance = newAppearanceComponent()
	e.appearance.texture = r.LoadTexture("sprites/player.png")
	e.position = newPositionComponent()
	e.position.x = 20
	e.position.y = 20
	e.controls = newControlsComponent()
	return e
}

func newBoxEntity() *entity {
	e := newEntity()
	e.appearance = newAppearanceComponent()
	e.appearance.texture = r.LoadTexture("sprites/box.png")
	e.position = newPositionComponent()
	e.position.x = 40
	e.position.y = 40
	return e
}
