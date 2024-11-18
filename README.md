![code coverage badge](https://github.com/Nukambe/go-nav/actions/workflows/tests.yml/badge.svg)

# Go-Nav

**Go-Nav** is an interactive CLI tool designed to simplify filesystem navigation. Go-Nav provides a user-friendly interface with features to make CLI navigation faster and more efficient.

## Features

- **Interactive Navigation**: Explore your filesystem interactively using arrow keys.

## Installation

1. Ensure you have [Go](https://golang.org/) installed (version 1.23.2 or higher recommended).
2. Install Go-Nav:
   ```bash
   go install github.com/Nukambe/go-nav@latest
   ```
3. Add the installation directory (e.g., `$GOPATH/bin`) to your `PATH` environment variable if it's not already included.

## Usage

Run the program:
```bash
go-nav
```

### Navigation Controls
- **Up/Down Arrow Keys**: Move up and down to select a directories.
- **Left/Right Arrow Keys**: Move to parent or selected subdirectory.
- **Enter**: Open a directory in a new window.
- **`q`**: Quit the program.

## Key Features

### Directory View
- Lists all directories in the current path.
- Highlights the current selection for easy tracking.
- User-friendly scrolling for long lists.

### Directory Preview
- Show subdirectories of the current selected directory without opening or navigating to it.

## Development

### Prerequisites
- Go 1.23.2 or higher

### Building the Program
Clone the repository:
```bash
git clone https://github.com/Nukambe/go-nav.git
cd go-nav
```

Build the program:
```bash
go build
```

Run locally:
```bash
./go-nav
```

## Contributing

Contributions are welcome! To contribute:
1. Fork the repository.
2. Create a feature branch:
   ```bash
   git checkout -b feature-name
   ```
3. Commit your changes and push:
   ```bash
   git commit -m "Add feature description"
   git push origin feature-name
   ```
4. Open a pull request.

## Support

For issues or feature requests, open an issue on the [GitHub repository](https://github.com/Nukambe/go-nav).

Enjoy a smoother filesystem navigation experience with Go-Nav! 🎉