# Clarity

🌿 A 2D, story-driven, pixel art platformer. Written in [Go](https://golang.org) using [Ebiten](https://ebiten.org) with an [ECS](https://en.wikipedia.org/wiki/Entity_component_system) architecture.

## Requirements

- Go 1.17
- Aseprite 1.2
- GIMP (For Windows `.ico` file)

## Development

All commands should be run from the root directory.

Start the game:
```
go run .
```

Build executable:
```
go build
```

## Architecture

Game state is managed through the state file and it's functions. It is responsible to containing worlds and transitioning between them. Worlds are collections of systems and act like scenes. Systems contain all the games logic, which they run on all related entities. Entities are a collection of components that are purely data containers. All of these elements are initialised per world in the world package.

## Project structure

Folder | Description
--- | ---
Root directory | The main package is contained in the root. It contains the main entrypoint where everything is initialised, including the game state which is also in this folder. The state is passed to all Systems so data can be shared.
Asset | Package that contains built assets that are embedded into the final executable.
Component | Package holds all Component files. Components are data buckets that are included in an Entity.
Entity | ackage holds all Entity files and their constructor. Entities are a collection of Components with varying data. They make up all things in the game.
WIP | Contains all WIP files, such as for creating sprites. The subfolders should match corresponding folders in the Asset package.
System | Package holds all systems and their constructor. Systems will runs their logic on all Entities that have the required Components.
Utility | A utility package for various helpers used across multiple packages.
World | Package holds all World files. Each world contains Systems and Entities and acts as a way of separating game scenes.

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

## Exporting sprites

All sprites should be exported as PNG files using Aseprite. Animated sprites or sprites with multiple states (e.g. hover) should be exported as spritesheets with an accompanying JSON file. Here are the settings used:

Setting | Value
--- | ---
Sheet Type | By Rows
Constraints | None
Merge Duplicates | `false`
Ignore Empty | `false`
Layers | Visible Layers
Borders | Default settings
Output File | `true`
JSON Data | `true` with same name as output file
JSON Style | Array
JSON Meta | Tags, Slices
JSON File Name | {frame}
