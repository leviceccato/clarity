package main

import (
	r "github.com/gen2brain/raylib-go/raylib"
)

type cameraComponent struct{}

func newCameraComponent() *cameraComponent {
	return &cameraComponent{}
}

type uiComponent struct{}

func newUIComponent() *uiComponent {
	return &uiComponent{}
}

type zIndexComponent struct {
	value int32
}

func newZIndexComponent() *zIndexComponent {
	return &zIndexComponent{value: 0}
}

type appearanceComponent struct {
	texture r.Texture2D
}

func newAppearanceComponent() *appearanceComponent {
	return &appearanceComponent{}
}

type collisionComponent struct {
	min, max r.Vector2
}

func newCollisionComponent() *collisionComponent {
	return &collisionComponent{}
}

type velocityComponent struct {
	value r.Vector2
}

func newVelocityComponent() *velocityComponent {
	return &velocityComponent{}
}

type gravityComponent struct {
	value r.Vector2
}

func newGravityComponent() *gravityComponent {
	return &gravityComponent{}
}

type positionComponent struct {
	value r.Vector2
}

func newPositionComponent() *positionComponent {
	return &positionComponent{}
}

type controlsComponent struct{}

func newControlsComponent() *controlsComponent {
	return &controlsComponent{}
}
