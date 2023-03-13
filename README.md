<p align="center">
    <img alt="Clarity" src="https://raw.githubusercontent.com/leviceccato/clarity/main/asset/sprite/banner.png">
</p>

# Clarity

🌿 A 2D, story-driven, pixel art platformer. Written in [Go](https://golang.org) using [Ebitengine](https://ebiten.org) with an [ECS](https://en.wikipedia.org/wiki/Entity_component_system) architecture. This project is in it's very early stages, it currently serves as a proof of concept but will be expanded into a full game in the future.

## Requirements

- Go 1.18
- Aseprite 1.2

## Development

All commands must be run from the root directory.

Start the game:

```
go run .
```

Start the game in debug mode:

```
go run . -debug
```

Build to test:

```
go build
```

Build the executable with included icon for Windows:

```
go run ./windows
```

Build the executable into a `.app` for macOS:

```
go run ./macos
```

## Architecture

The ECS portion of the game is managed in the engine package. The implementation is similar to that of most ECS's, there is a global Game struct responsible for containing game state and managing creation of other objects. Within this there are World's which act as containers for groups of System's and Entities. Systems are where all of the games logic is stored. Entities are a collection of Components which are themselves buckets of data which the Systems will be run on. Assemblages are groups of Entities represented as a single reusable thing. This package's functions are utilised by the game package where all of the Clarity-specific logic is contained and initialised.

## Project structure

| Folder   | Go Package | Description                                                                                                                                                                                                                                  |
| -------- | ---------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| /        | ✅         | The main package is contained in the root. It's purpose is to initialise the game and otherwise do as little as possible.                                                                                                                    |
| /engine  | ✅         | This where the ECS is implemented, it's purposefully decoupled from the game itself.                                                                                                                                                         |
| /game    | ✅         | This is where all of the Clarity-specific code resides, such Entity initilisation, System logic and Component definitions.                                                                                                                   |
| /asset   | ✅         | Contains built assets that are embedded into the final executable. This includes sprite, audio and translation files. The `icon.iconset` folder is named as such so the macOS `iconutil` program can use it to generate an `icon.icns` file. |
| /wip     |            | Contains all WIP files, such as for creating sprites. The subfolders should match corresponding folders in the asset package.                                                                                                                |
| /windows | ✅         | For building the project for Windows (`go run ./windows`).                                                                                                                                                                                   |
| /macos   | ✅         | For building the project for macOS (`go run ./macos`).                                                                                                                                                                                       |
| /util    | ✅         | Generic utility functions used across the project.                                                                                                                                                                                           |
| /logger  | ✅         | A custom logging implementation.                                                                                                                                                                                                             |

Files are prefixed with `0_` to ensure they are displayed first when sorted naturally. This is helpful for packages with many files.

## Exporting sprites

All sprites should be exported as `.png` files using Aseprite. Animated sprites or sprites with multiple states (e.g. hover) should be exported as spritesheets with an accompanying `.json` file. Here are the settings used:

| Setting          | Value                                |
| ---------------- | ------------------------------------ |
| Sheet Type       | By Rows                              |
| Constraints      | None                                 |
| Merge Duplicates | `false`                              |
| Ignore Empty     | `false`                              |
| Layers           | Visible Layers                       |
| Borders          | Default settings                     |
| Output File      | `true`                               |
| JSON Data        | `true` with same name as output file |
| JSON Style       | Array                                |
| JSON Meta        | Tags, Slices                         |
| JSON File Name   | {frame}                              |

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

@2x means 200% size. These icons must be manually exported from Aseprite. There is a dedicated 512 x 512 Aseprite icon file since there is a limit that the 32 x 32 version may be resized from the UI. Windows only utilises the 16 x 16 and 32 x 32 icons.
