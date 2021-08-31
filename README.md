# Clarity

🍤 A game written in [Go](https://golang.org) using [Ebiten](https://ebiten.org) with an [ECS](https://en.wikipedia.org/wiki/Entity_component_system) architecture.

## Architecture

Game state is managed through the state file and it's functions. It is responsible to containing worlds and transitioning between them. Worlds are collections of systems and act like scenes. Systems contain all the games logic, which they run on all related entities. Entities are a collection of components that are purely data containers. All of these elements are initialised per world in the world package.

## Requirements

- Go 1.15

## Development

All commands should be run from the root directory.

Start the game:
```
go run .
```

Build executable:
```
go build .
```

## Tests

Run tests:
```
go test ./...
```

Generate coverage and then view it in a browser:
```
go test ./... -coverprofile='coverage.out'
go tool cover -html='coverage.out'
```

## Tasks

A list of complete and incomplete tasks to track the progress of the game.

- [x] Create ECS
- [x] Add translation support
- [x] Make event system remove events after actioning them
- [x] Add player movement
- [x] Have system effect no entities when no components are specified
- [ ] Add systems and components for basic UI interactions
- [ ] Get world switching working