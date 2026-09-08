# tactical-arena-go

A turn-based fantasy arena game built with **Go and Ebitengine**, developed as a hands-on project to learn game programming.

The goal is a tactical 1v1 game where positioning, abilities, and terrain shape each fight.

## Current Prototype

Two players share a keyboard and take turns controlling their heroes.

* Grid-based movement with 3 movement points per turn.
* Adjacent melee attacks costing 1 action point.
* Walls that block movement.
* Health tracking, victory detection, and match restart.

The project currently uses simple placeholder graphics. Online multiplayer is not implemented yet.

## Controls

| Key        | Action                           |
| ---------- | -------------------------------- |
| Arrow keys | Move one tile                    |
| F          | Strike an adjacent opponent      |
| Space      | End turn                         |
| R          | Restart after the match finishes |

## Run Locally

With Go and the platform dependencies for [Ebitengine](https://ebitengine.org/en/documents/install.html) installed, run these commands from the project directory:

```bash
go mod download
go run .
```

## Project Structure

* `main.go` — Window setup and application entry point.
* `game.go` — Keyboard input and rendering.
* `match.go` — Match state, movement rules, and combat.

## Next Steps

* Highlight valid movement options.
* Add ranged attacks and displacement abilities.
* Introduce terrain hazards.
* Explore online 1v1 multiplayer.

Development focuses on small, playable additions while learning how each system works.

