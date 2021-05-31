package main

import (
	r "github.com/gen2brain/raylib-go/raylib"
)

type component struct {
	name string
}

type appearanceComponent struct {
	component
	texture r.Texture2D
}

func newAppearanceComponent() *appearanceComponent {
	c := &appearanceComponent{}
	c.name = "appearance"
	return c
}

type collisionComponent struct {
	component
	min r.Vector2
	max r.Vector2
}

func newCollisionComponent() *collisionComponent {
	c := &collisionComponent{}
	c.name = "collision"
	return c
}

type physicsComponent struct {
	component
	velocty r.Vector2
	gravity r.Vector2
}

func newPhysicsComponent() *physicsComponent {
	c := &physicsComponent{}
	c.name = "physics"
	return c
}

type positionComponent struct {
	component
	x, y float32
}

func newPositionComponent() *positionComponent {
	c := &positionComponent{}
	c.name = "position"
	return c
}

type controlsComponent struct {
	component
}

func newControlsComponent() *controlsComponent {
	return &controlsComponent{}
}
