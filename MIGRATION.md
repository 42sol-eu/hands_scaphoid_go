# Migration Guide: From Multiple Binaries to Unified Scaphoid CLI

## Overview

The Scaphoid project has been refactored to use a single unified CLI binary powered by the Cobra library. This guide will help you transition from the old multi-binary structure to the new unified interface.

## What Changed?

### Old Structure (Multiple Binaries)
Previously, each command was a separate binary:
- `./bin/file` - File operations
- `./bin/directory` - Directory operations
- `./bin/link` - Link operations
- `./bin/archive` - Archive operations
- `./bin/copy` - Copy operations
- `./bin/move` - Move operations
- `./bin/delete` - Delete operations
- etc.

### New Structure (Single Binary)
Now, all operations are accessible through a single `scaphoid` binary with subcommands:
- `scaphoid file [subcommand]` - File operations
- `scaphoid directory [subcommand]` - Directory operations
- `scaphoid [action] [args]` - General actions

## Command Migration

### File Operations

| Old Command | New Command |
|------------|-------------|
| `./bin/file create myfile.txt` | `scaphoid file create myfile.txt` |
| `./bin/file list myfile.txt` | `scaphoid file list myfile.txt` |
| `./bin/file copy src.txt dst.txt` | `scaphoid file copy src.txt dst.txt` |
| `./bin/file move old.txt new.txt` | `scaphoid file move old.txt new.txt` |
| `./bin/file delete myfile.txt -f` | `scaphoid file delete myfile.txt -f` |
| `./bin/file info myfile.txt` | `scaphoid file info myfile.txt` |
| `./bin/file compare f1.txt f2.txt` | `scaphoid file compare f1.txt f2.txt` |
| `./bin/file backup src.txt /backup` | `scaphoid file backup src.txt /backup` |

### Directory Operations

| Old Command | New Command |
|------------|-------------|
| `./bin/directory create mydir` | `scaphoid directory create mydir` |
| `./bin/directory list mydir` | `scaphoid directory list mydir` |
| `./bin/directory copy -r src/ dst/` | `scaphoid directory copy -r src/ dst/` |
| `./bin/directory move olddir newdir` | `scaphoid directory move olddir newdir` |
| `./bin/directory delete -rf mydir` | `scaphoid directory delete -rf mydir` |
| `./bin/directory info mydir` | `scaphoid directory info mydir` |

### Link Operations

| Old Command | New Command |
|------------|-------------|
| `./bin/link create mylink --target /path` | `scaphoid link create mylink --target /path` |
| `./bin/link info mylink` | `scaphoid link info mylink` |
| `./bin/link delete mylink` | `scaphoid link delete mylink` |

### Archive Operations

| Old Command | New Command |
|------------|-------------|
| `./bin/archive create arch.tar.gz --source /dir` | `scaphoid archive create arch.tar.gz --source /dir` |
| `./bin/archive list arch.tar.gz` | `scaphoid archive list arch.tar.gz` |
| `./bin/archive info arch.tar.gz` | `scaphoid archive info arch.tar.gz` |

### General Commands

The old standalone action binaries now work as general commands:

| Old Command | New Command |
|------------|-------------|
| `./bin/create directory mydir` | `scaphoid create directory mydir` |
| `./bin/create file myfile.txt` | `scaphoid create file myfile.txt` |
| `./bin/copy src dst` | `scaphoid copy src dst` |
| `./bin/move src dst` | `scaphoid move src dst` |
| `./bin/delete path -rf` | `scaphoid delete path -rf` |
| `./bin/info path` | `scaphoid info path` |
| `./bin/list path` | `scaphoid list path` |
| `./bin/compare p1 p2` | `scaphoid compare p1 p2` |
| `./bin/backup src /backup` | `scaphoid backup src /backup` |

## Building and Installation

### Old Way
```bash
# Build all binaries
task build

# Install multiple binaries
for cmd in file directory link archive copy move delete info list compare backup; do
    cp ./bin/$cmd $GOPATH/bin/fs-$cmd
done
```

### New Way
```bash
# Build single binary
task build
# or
go build -o bin/scaphoid .

# Install single binary
task install
# or
cp ./bin/scaphoid $GOPATH/bin/scaphoid
```

## Benefits of the New Structure

### 1. **Single Binary**
- One binary to install and maintain
- Smaller overall disk footprint
- Easier distribution

### 2. **Consistent Interface**
- All commands use the same flag parsing and help system
- Cobra provides auto-completion support
- Better error messages and help text

### 3. **Hierarchical Organization**
- Logical grouping of commands
- Easier to discover related commands
- Better namespace management

### 4. **Enhanced Features**
- Built-in shell completion generation
- Consistent flag behavior across commands
- Better help and documentation

### 5. **Simpler Development**
- Shared code is easier to manage
- Single build target
- Easier to add new commands

## Shell Completion

The new CLI supports auto-completion for bash, zsh, fish, and PowerShell:

```bash
# Generate bash completion
scaphoid completion bash > /etc/bash_completion.d/scaphoid

# Generate zsh completion
scaphoid completion zsh > "${fpath[1]}/_scaphoid"

# Generate fish completion
scaphoid completion fish > ~/.config/fish/completions/scaphoid.fish

# Generate PowerShell completion
scaphoid completion powershell > scaphoid.ps1
```

## Updating Scripts

If you have scripts using the old commands, you can use these approaches:

### Option 1: Update Scripts Directly
Replace all old command calls with new ones:
```bash
# Old
./bin/file create myfile.txt

# New
scaphoid file create myfile.txt
```

### Option 2: Create Wrapper Scripts (Temporary)
Create wrapper scripts for backward compatibility:
```bash
#!/bin/bash
# file wrapper script
scaphoid file "$@"
```

### Option 3: Use Aliases
Add aliases to your shell configuration:
```bash
alias file='scaphoid file'
alias directory='scaphoid directory'
alias link='scaphoid link'
alias archive='scaphoid archive'
```

## Troubleshooting

### Issue: Command not found
**Solution**: Make sure `scaphoid` is in your PATH or use the full path:
```bash
# Add to PATH
export PATH="$PATH:/path/to/scaphoid/bin"

# Or use full path
/path/to/scaphoid/bin/scaphoid file create myfile.txt
```

### Issue: Old binaries still in PATH
**Solution**: Remove old binaries:
```bash
# Remove old binaries from GOPATH/bin
rm $GOPATH/bin/fs-*

# Or from /usr/local/bin
sudo rm /usr/local/bin/fs-*
```

### Issue: Flags not working as expected
**Solution**: Check the help for the specific command:
```bash
scaphoid file create --help
```

## Getting Help

### Command Help
```bash
# General help
scaphoid --help

# Help for a specific command
scaphoid file --help

# Help for a subcommand
scaphoid file create --help
```

### Documentation
- Check the main README.md for usage examples
- View inline help with `--help` flag
- Consult the docs/ directory for detailed documentation

## Reporting Issues

If you encounter any issues during migration:
1. Check this migration guide
2. Run commands with `--help` for correct usage
3. Open an issue on GitHub with:
   - Old command that worked
   - New command you tried
   - Error message or unexpected behavior
   - Your environment (OS, Go version)

## Timeline

- **Old structure**: Available in the `cmd/` directory (deprecated)
- **New structure**: Available in `cmd_new/` directory and `main.go`
- **Future**: The old `cmd/` directory will be removed in a future release

## Feedback

We welcome your feedback on the new structure! Please open an issue or discussion on GitHub if you have:
- Suggestions for improvements
- Questions about the new structure
- Issues with migration
- Ideas for new features
