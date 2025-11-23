# README `scaphoid_go`

This repository contains the Go implementation of Scaphoid, a unified command-line interface for filesystem operations. Scaphoid provides a comprehensive and consistent interface for managing files, directories, symbolic links, and archives using the Cobra library.

## Features

### Unified CLI with Cobra

Scaphoid uses the powerful Cobra library to provide a clean, hierarchical command structure. All operations are accessible through a single `scaphoid` binary with intuitive subcommands.

### Object and Action Based Commands

You can use the filesystem commands in both forms:
1. **Object-specific commands**: `scaphoid [object] [action]`
2. **General action commands**: `scaphoid [action] [path]`

For example, all of the following commands are valid:

```bash
# Object-specific commands
scaphoid file create myfile.txt
scaphoid directory create mydir
scaphoid link create mylink --target /path/to/target

# General commands
scaphoid create file myfile.txt
scaphoid create directory mydir
scaphoid copy myfile.txt newfile.txt
scaphoid info myfile.txt
```

### Well Designed Options and Flags

Each command supports a variety of options and flags to customize its behavior. Both long and short forms are available for convenience:
- `--force, -f` - Force operation, overwrite existing
- `--recursive, -r` - Recursive operation for directories
- `--all, -a` - Show all/hidden items
- `--deep, -d` - Perform deep comparison

> [!Important]
> Similar flags are used across different commands to ensure consistency and ease of use.

### Comprehensive Command Set

The Scaphoid toolkit provides a wide range of commands to manage filesystem objects:

#### Main Commands
- `file` - File operations (create, list, copy, move, delete, info, compare, backup)
- `directory` - Directory operations with same subcommands
- `link` - Symbolic link operations with same subcommands
- `archive` - Archive file operations with same subcommands

#### General Commands
- `create` - Create any filesystem object (file, directory, link, archive)
- `list` - List contents of any object
- `copy` - Copy any object
- `move` - Move/rename any object
- `delete` - Delete any object
- `info` - Display information about any object
- `compare` - Compare two objects
- `backup` - Create backups of any object

## Installation

### Build from Source

```bash
# Clone the repository
git clone https://github.com/42sol-eu/hands_scaphoid_go
cd hands_scaphoid_go

# Build the binary using Task
task build

# Or build directly with Go
go build -o bin/scaphoid .

# Install to GOPATH/bin
task install

# Or install to /usr/local/bin (requires sudo)
task install:local
```

## Usage Examples

### File Operations

```bash
# Create a file
scaphoid file create myfile.txt

# List file contents
scaphoid file list myfile.txt

# Copy a file
scaphoid file copy source.txt dest.txt

# Get file information
scaphoid file info myfile.txt

# Delete a file
scaphoid file delete myfile.txt --force
```

### Directory Operations

```bash
# Create a directory (with parents)
scaphoid directory create -r /path/to/mydir

# List directory contents
scaphoid directory list /path/to/mydir

# List with hidden files
scaphoid directory list -a /path/to/mydir

# Copy directory recursively
scaphoid directory copy -r source/ dest/

# Delete directory recursively
scaphoid directory delete -rf /path/to/mydir
```

### Link Operations

```bash
# Create a symbolic link
scaphoid link create mylink --target /path/to/target

# Get link information
scaphoid link info mylink

# Copy a link
scaphoid link copy mylink newlink
```

### Archive Operations

```bash
# Create an archive from a directory
scaphoid archive create myarchive.tar.gz --source /path/to/dir

# List archive contents
scaphoid archive list myarchive.tar.gz

# Get archive information
scaphoid archive info myarchive.tar.gz
```

### General Commands

```bash
# Create any type of object
scaphoid create file myfile.txt
scaphoid create directory mydir -r
scaphoid create link mylink --target /target

# Copy any object
scaphoid copy source dest

# Move/rename any object
scaphoid move oldname newname

# Get info on any object
scaphoid info /path/to/object

# Compare two objects
scaphoid compare file1.txt file2.txt --deep

# Backup any object
scaphoid backup /path/to/object /backup/dir --timestamp --compress

# Delete any object
scaphoid delete /path/to/object -rf
```

## Command Structure

```
scaphoid
├── file [subcommands]
│   ├── create
│   ├── list
│   ├── copy
│   ├── move
│   ├── delete
│   ├── info
│   ├── compare
│   └── backup
├── directory [subcommands]
│   └── (same as file)
├── link [subcommands]
│   └── (same as file)
├── archive [subcommands]
│   └── (same as file)
├── create [type] [path]
├── list [path]
├── copy [source] [dest]
├── move [source] [dest]
├── delete [path]
├── info [path]
├── compare [path1] [path2]
└── backup [source] [dest]
```

## Development

### Prerequisites

- Go 1.24.2 or higher
- Task (optional, for task runner)

### Building

```bash
# Build the binary
task build

# Build with race detection
task build:race

# Run tests
task test

# Run tests with coverage
task test:coverage

# Format code
task fmt

# Run linters
task lint
```

### Testing

```bash
# Run all tests
task test

# Run integration tests
task test:integration

# Run benchmarks
task bench
```

## Interactive Prompts

Scaphoid supports **interactive user prompts** for a better user experience:

- **Confirmations**: Ask yes/no questions (e.g., "Delete this file?")
- **Selections**: Choose from a list of options
- **Multi-select**: Choose multiple options
- **Text input**: Collect user input
- **Password**: Secure password entry

Example:
```bash
# Delete with confirmation prompt
$ scaphoid file delete important.txt
? Delete important.txt? (y/N) y

# Interactive object type selection
$ scaphoid create myfile.txt
? Select object type: 
  ▸ file
    directory
    link
    archive
```

See [docs/INTERACTIVE_PROMPTS.md](docs/INTERACTIVE_PROMPTS.md) for full details.

## Color Customization

Scaphoid supports **colored output** with bold text that can be customized via a `.scaphoid.yaml` configuration file. The tool automatically colorizes:

- **Action verbs** (exists, create, delete, etc.) - Blue by default
- **Filesystem types** (file, directory, link, archive) - Green by default
- **Paths** - Green by default, shown in double quotes
- **Error messages** - Red by default

Place a `.scaphoid.yaml` file in your current directory, home directory, or `/etc/scaphoid/` to customize the colors. See [docs/COLOR_CONFIG.md](docs/COLOR_CONFIG.md) for full configuration details.

Example output:
```bash
$ scaphoid file exists README.md
"README.md" exists and is a file
# "README.md" appears in green, "exists" in blue, "file" in green
```

## Markdown Output Formatting

Scaphoid supports **markdown-like console output** for better readability and documentation generation:

- **Success messages**: Automatically prefixed with `- ` (unordered list)
- **Error messages**: Automatically prefixed with `> ` (blockquote)
- **Section headers**: `scaphoid section "Title" ["details"]` outputs `## Title\ndetails\n`
- **Separators**: `scaphoid separator` outputs `---`

This makes it easy to generate documentation directly from command execution. See [docs/MARKDOWN_OUTPUT.md](docs/MARKDOWN_OUTPUT.md) for examples.

Example:
```bash
$ scaphoid section "File Operations"
## File Operations

$ scaphoid file create test.txt
- Successfully created file at "test.txt"

$ scaphoid file exists /nonexistent.txt
> "/nonexistent.txt" does not exist

$ scaphoid separator
---
```

## Project Structure

```
.
├── main.go              # Entry point
├── cmd_new/             # Cobra command definitions
│   ├── root.go          # Root command
│   ├── file.go          # File subcommands
│   ├── directory.go     # Directory subcommands
│   ├── link.go          # Link subcommands
│   ├── archive.go       # Archive subcommands
│   ├── general.go       # General commands
│   ├── exists.go        # Exists commands
│   └── colors.go        # Color configuration
├── pkg/
│   ├── actions/         # Action implementations
│   └── fsobj/           # Filesystem object abstractions
├── tests/               # Tests
├── docs/                # Documentation
│   └── COLOR_CONFIG.md  # Color customization guide
├── .scaphoid.yaml       # Default color config
└── Taskfile.yaml        # Task definitions
```