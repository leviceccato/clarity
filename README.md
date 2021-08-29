# Clarity

🍤 A game written in [Go](https://golang.org) using [Ebiten](https://ebiten.org) with an [ECS](https://en.wikipedia.org/wiki/Entity_component_system) architecture.

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