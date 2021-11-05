<p align="center">
    <img width="472" height="232" alt="Clarity" src="https://raw.githubusercontent.com/leviceccato/clarity/main/banner.png">
</p>

# Clarity

🌿 A 2D, story-driven, pixel art platformer. Written in [Go](https://golang.org) using [Ebiten](https://ebiten.org) with an [ECS](https://en.wikipedia.org/wiki/Entity_component_system) architecture.

## Requirements

- Go 1.17
- Aseprite 1.2

## Development

All commands should be run from the root directory.

Start the game:
```
go run .
```

Build to test:
```
go build
```

Build the executable with included icon for Windows:
```
go run ./windows
```

Build the executable into a `.app` for macOS.
```
go run ./macos
```

## Architecture

Game state is managed through the state file and it's functions. It is responsible for containing worlds and transitioning between them. Worlds are collections of systems and act like scenes. Systems contain all the games logic, which they run on all related entities. Entities are a collection of components that are purely data containers. All of these elements are initialised per world in the world package.

## Project structure

Folder | Description
--- | ---
Root directory | The main package is contained in the root. It contains the main entrypoint where everything is initialised, including the game state which is also in this folder. The state is passed to all Systems so data can be shared.
asset | Package that contains built assets that are embedded into the final executable. The `icon.iconset` folder is named as such so the macOS `iconutil` program can use it to generate an `icon.icns` file.
config | Global config is contained within this package and exposed as constants and variables.
component | Package holds all Component files. Components are data buckets that are included in an Entity.
entity | Package holds all Entity files and their constructor. Entities are a collection of Components with varying data. They make up all things in the game.
wip | Contains all WIP files, such as for creating sprites. The subfolders should match corresponding folders in the Asset package.
system | Package holds all systems and their constructor. Systems will runs their logic on all Entities that have the required Components.
util | A util package for various helpers used across multiple packages.
world | Package holds all World files. Each world contains Systems and Entities and acts as a way of separating game scenes.
windows | Package specifically for generating Windows specific resource file.
macos | Package specifically for generating macOS `.app` folder.

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

## Creating icons for macOS

The macOS `.app` folder requires an array of specific icon sizes named the following:

- icon_16x16.png
- icon_16x16<span>@</span>2x.png
- icon_32x32.png
- icon_32x32<span>@</span>2x.png
- icon_64x64.png
- icon_64x64<span>@</span>2x.png
- icon_128x128.png
- icon_128x128<span>@</span>2x.png
- icon_256x256.png
- icon_256x256<span>@</span>2x.png
- icon_512x512.png
- icon_512x512<span>@</span>2x.png

The **@2x** simply means 200% size. These must be manually exported from Aseprite. There is a dedicated 512 x 512 Aseprite icon file since there is a limit that the 32 x 32 version may be resized from the UI. Windows currently only utilises the 16 x 16 and 32 x 32 icons.