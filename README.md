# go-fire
A forest fire cellular automaton simulation in Go, rendered with [raylib](https://github.com/gen2brain/raylib-go).

## How it works
The simulation runs on a 640x480 grid of cells, each one of four states:

| State | Colour | Description |
|---|---|---|
| Empty | Black | Bare ground |
| Tree | Green | Burnable vegetation |
| Fire | Red | Actively burning |
| BurnedOut | Black | Ash — can regrow |

Each frame, two update passes run:

- **Fire spread** — a burning cell has a 30% chance to ignite each neighbouring tree
- **Burn out** — a burning cell has a 70% chance to become burned out per neighbour check
- **Regrowth** — burned out cells have a 10% chance to regrow as a tree each frame

## Dependencies
- [raylib-go](https://github.com/gen2brain/raylib-go)

## Run
```
go run .
```

> [!NOTE] AI generated readme
> This README was generated using AI
