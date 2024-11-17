![code coverage badge](https://github.com/Nukambe/go-nav/actions/workflows/tests.yml/badge.svg)

# Go-Nav

**Go-Nav** is an interactive CLI tool designed to simplify filesystem navigation. Go-Nav provides a user-friendly interface with features to make CLI navigation faster and more efficient.

---

## Features

- **Interactive Navigation**: Explore your filesystem interactively using arrow keys.
- **File and Directory Management**: Perform basic operations like opening files or changing directories.
- **Terminal-Friendly**: Lightweight and optimized for terminal environments.

---

## Installation

1. Ensure you have [Go](https://golang.org/) installed (version 1.23.2 or higher recommended).
2. Install Go-Nav:
   ```bash
   go install github.com/Nukambe/go-nav@latest
   ```
3. Add the installation directory (e.g., `$GOPATH/bin`) to your `PATH` environment variable if it's not already included.

---

## Usage

Run the program:
```bash
go-nav
```

### Navigation Controls
- **Up/Down Arrow Keys**: Move up and down through files and directories.
- **Left/Right Arrow Keys**: Move to parent or subdirectory.
- **Enter**: Open a directory or file.
- **Backspace**: Move up one directory level.
- **`q`**: Quit the program.

---

## Key Features

### Directory View
- Lists all files and directories in the current path.
- Highlights the current selection for easy tracking.

### File Preview
- Preview text files directly in the terminal.
- (Optional) Open non-text files with the system's default program.

### Search
- Quickly search for files and directories by name.

---

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
go build -o go-nav
```

Run locally:
```bash
./go-nav
```

---

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

---

## Support

For issues or feature requests, open an issue on the [GitHub repository](https://github.com/Nukambe/go-nav).

---

Enjoy a smoother filesystem navigation experience with Go-Nav! 🎉