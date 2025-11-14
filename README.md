# Snake Game

A simple Snake game written in Go. Use WASD to control the snake, eat apples to grow, and try to beat your best score.

## Features

- Score tracking with best score saved to file
- Terminal-based graphics
- Works on Linux, macOS, and Windows

## Installation

### From Releases

Download the latest release for your platform from the [Releases](https://github.com/Mehulparekh144/go-snake/releases) page:

- **Linux**: `snake_Linux_x86_64.tar.gz` or `snake_Linux_arm64.tar.gz`
- **macOS**: `snake_Darwin_x86_64.tar.gz` or `snake_Darwin_arm64.tar.gz`
- **Windows**: `snake_Windows_x86_64.zip`

Extract the archive and run the `snake` (or `snake.exe` on Windows) executable.

### From Source

Prerequisites:

- Go 1.25.1 or later

```bash
# Clone the repository
git clone https://github.com/Mehulparekh144/go-snake.git
cd go-snake

# Build the game
go build -o snake ./cmd

# Run the game
./snake
```

## How to Play

- Use WASD keys to move (W=up, A=left, S=down, D=right)
- Eat the `*` to grow and increase your score
- Don't hit walls or yourself
- Press CTRL+C to quit

## Game Rules

- Snake starts with length 1
- Each apple increases score by 1 and makes the snake grow
- Game ends if you hit a wall or yourself
- Best score is saved to `score.txt`

## Project Structure

```text
Snake/
├── cmd/
│   └── main.go          # Main game logic
├── go.mod               # Go module file
├── .goreleaser.yaml     # Release configuration
└── README.md            # This file
```

## Building

```bash
go build -o snake ./cmd
```

To build for multiple platforms, use [GoReleaser](https://goreleaser.com):

```bash
goreleaser release --clean
```

## Requirements

- Go 1.25.1+
- A terminal that supports ANSI escape codes
