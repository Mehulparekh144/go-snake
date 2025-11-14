# 🐍 Snake Game

A classic Snake game written in Go with a terminal-based interface. Control the snake using WASD keys, eat apples to grow, and try to beat your best score!

## Features

- 🎮 Classic Snake gameplay
- 📊 Score tracking with best score persistence
- 🎯 Random apple generation
- 🖥️ Terminal-based ASCII graphics
- 💾 Automatic best score saving
- 🚀 Cross-platform support (Linux, macOS, Windows)

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

1. **Start the game**: Run the `snake` executable
2. **Control the snake**: Use WASD keys to move
   - `W` - Move up
   - `A` - Move left
   - `S` - Move down
   - `D` - Move right
3. **Eat apples**: Navigate the snake to the `*` (apple) to grow and increase your score
4. **Avoid collisions**: Don't hit the walls or your own body!
5. **Quit**: Press `CTRL+C` to exit the game

## Game Rules

- The snake starts with a length of 1
- Each apple eaten increases your score by 1 and makes the snake grow
- The game ends if the snake hits a wall or itself
- Your best score is automatically saved to `score.txt`

## Project Structure

```text
Snake/
├── cmd/
│   └── main.go          # Main game logic
├── go.mod               # Go module file
├── .goreleaser.yaml     # Release configuration
└── README.md            # This file
```

## Building for Multiple Platforms

This project uses [GoReleaser](https://goreleaser.com) for automated releases. To build for all platforms:

```bash
# Install GoReleaser (if not already installed)
brew install goreleaser  # macOS
# or download from https://goreleaser.com/install

# Build and release
goreleaser release --clean
```

## Development

### Running Locally

```bash
go run ./cmd
```

### Building

```bash
go build -o snake ./cmd
```

## Requirements

- Go 1.25.1+
- A terminal that supports ANSI escape codes (most modern terminals)

## License

This project is open source and available for personal use.

## Contributing

Contributions are welcome! Feel free to open an issue or submit a pull request.

## Acknowledgments

- Classic Snake game mechanics
- Built with Go for performance and cross-platform compatibility
