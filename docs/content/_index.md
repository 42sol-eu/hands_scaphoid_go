---
title: "Scaphoid Filesystem Commands"
featured_image: ""
description: "A comprehensive toolkit for filesystem operations with dual command syntax support"
---

# Scaphoid Filesystem Commands

A powerful and flexible toolkit for filesystem operations built with Go. This project provides a comprehensive set of command-line utilities for managing files, directories, links, and archives with support for both action-first and object-first command syntax.

## Features

- **Dual Command Syntax**: Use either `create directory mydir` or `directory create mydir`
- **Comprehensive Operations**: Create, list, copy, move, delete, info, compare, and backup
- **Multiple Object Types**: Files, directories, symbolic links, and archives
- **Unix Philosophy**: Simple tools that work well together
- **Cross-Platform**: Built with Go for portability
- **Extensive Testing**: Unit and integration tests included

## Quick Start

```bash
# Clone the repository
git clone https://github.com/42sol-eu/hands_scaphoid_go.git
cd hands_scaphoid_go

# Build all commands
task build

# Run demo
task run:demo

# Install globally
task install
```

## Command Overview

### Action-Based Commands
- `create` - Create filesystem objects
- `list` - List contents of objects
- `copy` - Copy objects
- `move` - Move/rename objects
- `delete` - Delete objects
- `info` - Get object information
- `compare` - Compare objects
- `backup` - Backup objects

### Object-Based Commands
- `directory` - Directory operations
- `file` - File operations
- `link` - Symbolic link operations
- `archive` - Archive operations

## Examples

Create a directory:
```bash
# Action-first syntax
create directory /path/to/mydir

# Object-first syntax
directory create /path/to/mydir
```

List directory contents:
```bash
# Action-first syntax
list /path/to/directory

# Object-first syntax
directory list /path/to/directory
```

## Documentation Sections

- [Getting Started Guide](/getting-started/) - Installation and basic usage
- [Command Reference](/commands/) - Detailed command documentation
- [Examples](/examples/) - Common use cases and workflows