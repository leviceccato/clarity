package main

type runner interface {
	run() error
	getComponents() []string
	addEntity(*entity)
	getEntities() []*entity
}

type system struct {
	components []string
	entities   []*entity
}
