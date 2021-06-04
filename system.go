package main

type runner interface {
	run() error
	getComponents() []string
	addEntity(*entity)
	getEntities() []*entity
	sortEntities()
}

type system struct {
	components []string
	entities   []*entity
}
