package main

import (
	r "github.com/gen2brain/raylib-go/raylib"
)

var entityId int32 = 0

type entity struct {
	id         int32
	appearance *appearanceComponent
	collision  *collisionComponent
	velocity   *velocityComponent
	gravity    *gravityComponent
	position   *positionComponent
	controls   *controlsComponent
	camera     *cameraComponent
	zIndex     *zIndexComponent
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
	e.position.value.X = 20
	e.position.value.Y = 20
	e.controls = newControlsComponent()
	e.camera = newCameraComponent()
	return e
}

func newBoxEntity() *entity {
	e := newEntity()
	e.appearance = newAppearanceComponent()
	e.appearance.texture = r.LoadTexture("sprites/box.png")
	e.position = newPositionComponent()
	e.position.value.X = 40
	e.position.value.Y = 40
	e.camera = newCameraComponent()
	return e
}
