---
title: "Getting Started"
date: 2025-11-23
draft: false
---

# Getting Started with Scaphoid Filesystem Commands

This guide will help you install, build, and start using the Scaphoid filesystem commands toolkit.

## Prerequisites

- Go 1.19 or later
- Git
- Task runner (optional but recommended)

### Installing Task Runner

If you don't have Task installed, you can install it using:

**macOS (Homebrew):**
```bash
brew install go-task/tap/go-task
```

**Other platforms:**
```bash
go install github.com/go-task/task/v3/cmd/task@latest
```

## Installation

### Option 1: Clone and Build

```bash
# Clone the repository
git clone https://github.com/42sol-eu/hands_scaphoid_go.git
cd hands_scaphoid_go

# Install dependencies and build
task deps
task build

# Optionally install to system PATH
task install
```

### Option 2: Go Install (coming soon)

```bash
go install github.com/42sol-eu/hands_scaphoid_go/cmd/...@latest
```

## Quick Verification

After building, verify the installation by running:

```bash
# List all available tasks
task

# Run demo to test functionality
task run:demo

# Run tests
task test
```

## Basic Usage

The toolkit provides two equivalent command syntaxes:

### Action-First Syntax
```bash
# Pattern: [action] [object-type] [arguments] [flags]
create directory mydir
list mydir
copy file source.txt dest.txt
delete file unwanted.txt --force
```

### Object-First Syntax
```bash
# Pattern: [object-type] [action] [arguments] [flags]
directory create mydir
directory list mydir
file copy source.txt dest.txt
file delete unwanted.txt --force
```

## Available Commands

### Action Commands
- **create** - Create new filesystem objects
- **list** - List contents or properties
- **copy** - Copy objects to new locations
- **move** - Move or rename objects
- **delete** - Remove objects
- **info** - Display object information
- **compare** - Compare two objects
- **backup** - Create backups of objects

### Object Commands
- **directory** - Operations on directories
- **file** - Operations on files
- **link** - Operations on symbolic links
- **archive** - Operations on archive files

## Common Flags

Most commands support these common flags:

- `--force, -f` - Force operation, overwrite existing
- `--recursive, -r` - Operate recursively on directories
- `--help, -h` - Show command help

## Examples

### Creating Objects

```bash
# Create a directory
create directory /tmp/testdir
directory create /tmp/testdir

# Create a file
create file /tmp/testfile.txt
file create /tmp/testfile.txt

# Create a symbolic link
create link /tmp/mylink --target /tmp/testfile.txt
link create /tmp/mylink --target /tmp/testfile.txt

# Create an archive
create archive /tmp/backup.tar.gz --source /tmp/testdir
archive create /tmp/backup.tar.gz --source /tmp/testdir
```

### Listing Contents

```bash
# List directory contents
list /tmp/testdir
directory list /tmp/testdir

# Long format listing
list /tmp/testdir --long
directory list /tmp/testdir --long

# Include hidden files
list /tmp/testdir --all
directory list /tmp/testdir --all
```

### Getting Information

```bash
# Get file information
info /tmp/testfile.txt
file info /tmp/testfile.txt

# Get detailed information
info /tmp/testfile.txt --all
```

## Next Steps

- Read the [Command Reference](/commands/) for detailed documentation
- Check out [Examples](/examples/) for common workflows
- Explore the source code on [GitHub](https://github.com/42sol-eu/hands_scaphoid_go)