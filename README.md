# README `scaphoid_go`

This repository contains the Go implementation of the Scaphoid command line tools, a secure and efficient communication protocol designed for distributed systems. The Scaphoid protocol focuses on providing low-latency, high-throughput messaging with strong security guarantees.

## Features
### Object and action based command line tools

You can use the filesystem commands in both forms `object action` and `action object`. For example, both of the following commands are valid and equivalent:

```bash
create directory fortytwo
directory create fortytwo
```

### Well designed options and flags (long and short form)

Each command supports a variety of options and flags to customize its behavior. Both long and short forms are available for convenience. For example, the `create` command supports the following flags:
- `--force, -f` - Force creation, overwrite existing
- `--recursive, -r` - Create parent directories if needed

> [!Important]
> Similar flags are used across different commands to ensure consistency and ease of use.

### Comprehensive command set

The Scaphoid filesystem toolkit provides a wide range of commands to manage files, directories, symbolic links, and archives. The available commands include:
- `create` - Create new filesystem objects
- `list` - List contents or properties
- `copy` - Copy objects to new locations
- `move` - Move or rename objects
- `delete` - Remove objects
- `info` - Display object information
- `compare` - Compare two objects
- `backup` - Create backups of objects