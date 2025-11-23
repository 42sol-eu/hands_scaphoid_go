# Scaphoid CLI Quick Reference

## Command Structure

```
scaphoid [COMMAND] [SUBCOMMAND] [FLAGS] [ARGS]
```

## Main Commands

### File Operations
```bash
scaphoid file create <path> [-f]
scaphoid file list <path>
scaphoid file copy <src> <dst> [-f]
scaphoid file move <src> <dst> [-f]
scaphoid file delete <path> [-f]
scaphoid file info <path> [-a]
scaphoid file compare <p1> <p2> [-d]
scaphoid file backup <src> <dst> [-t] [-c]
```

### Directory Operations
```bash
scaphoid directory create <path> [-f] [-r]
scaphoid directory list <path> [-a] [-l] [-r]
scaphoid directory copy <src> <dst> [-f] [-r]
scaphoid directory move <src> <dst> [-f]
scaphoid directory delete <path> [-f] [-r]
scaphoid directory info <path> [-a]
scaphoid directory compare <p1> <p2> [-d]
scaphoid directory backup <src> <dst> [-t] [-c]
```

### Link Operations
```bash
scaphoid link create <path> -t <target> [-f]
scaphoid link list <path>
scaphoid link copy <src> <dst> [-f]
scaphoid link move <src> <dst> [-f]
scaphoid link delete <path> [-f]
scaphoid link info <path> [-a]
scaphoid link compare <p1> <p2> [-d]
scaphoid link backup <src> <dst> [-t] [-c]
```

### Archive Operations
```bash
scaphoid archive create <path> -s <source> [-f]
scaphoid archive list <path>
scaphoid archive copy <src> <dst> [-f]
scaphoid archive move <src> <dst> [-f]
scaphoid archive delete <path> [-f]
scaphoid archive info <path> [-a]
scaphoid archive compare <p1> <p2> [-d]
scaphoid archive backup <src> <dst> [-t] [-c]
```

## General Commands

These work on any filesystem object type:

```bash
scaphoid create <type> <path> [flags]
  # Types: file, directory, link, archive
  # Flags: -f, -r, -t <target>, -s <source>

scaphoid list <path> [-a] [-l] [-r]
scaphoid copy <src> <dst> [-f] [-r] [-p]
scaphoid move <src> <dst> [-f]
scaphoid delete <path> [-f] [-r]
scaphoid info <path> [-a]
scaphoid compare <p1> <p2> [-d]
scaphoid backup <src> <dst> [-t] [-c]
```

## Common Flags

| Flag | Long Form | Description |
|------|-----------|-------------|
| `-f` | `--force` | Force operation, overwrite existing |
| `-r` | `--recursive` | Recursive operation |
| `-a` | `--all` | Show all/hidden items |
| `-l` | `--long` | Long format listing |
| `-d` | `--deep` | Deep comparison |
| `-t` | `--timestamp` | Add timestamp to backup |
| `-c` | `--compress` | Compress backup |
| `-p` | `--preserve` | Preserve metadata |
| `-s` | `--source` | Source directory (for archives) |
| `-t` | `--target` | Target path (for links) |

## Examples

### Create Operations
```bash
# Create a file
scaphoid file create myfile.txt

# Create a directory with parents
scaphoid directory create -r /path/to/deep/dir

# Create a symbolic link
scaphoid link create mylink --target /path/to/target

# Create an archive
scaphoid archive create backup.tar.gz --source /data

# Using general create command
scaphoid create file myfile.txt
scaphoid create directory mydir -r
```

### List Operations
```bash
# List file contents
scaphoid file list myfile.txt

# List directory with hidden files
scaphoid directory list -a /mydir

# List directory recursively in long format
scaphoid directory list -rl /mydir

# List archive contents
scaphoid archive list backup.tar.gz
```

### Copy/Move Operations
```bash
# Copy a file
scaphoid file copy source.txt dest.txt

# Copy directory recursively
scaphoid directory copy -r srcdir/ dstdir/

# Move/rename a file
scaphoid move oldname.txt newname.txt

# Copy with metadata preservation
scaphoid copy -p important.txt backup.txt
```

### Delete Operations
```bash
# Delete a file
scaphoid file delete myfile.txt

# Delete with force
scaphoid delete -f unwanted.txt

# Delete directory recursively with force
scaphoid directory delete -rf /tmpdir
```

### Info and Compare
```bash
# Get file info
scaphoid info myfile.txt

# Get detailed info
scaphoid file info -a myfile.txt

# Compare two files
scaphoid compare file1.txt file2.txt

# Deep comparison
scaphoid compare -d dir1/ dir2/
```

### Backup Operations
```bash
# Backup with timestamp
scaphoid backup myfile.txt /backups/

# Backup with compression
scaphoid backup -c largefile.txt /backups/

# Backup with both timestamp and compression
scaphoid backup -tc important/ /backups/
```

## Getting Help

```bash
# General help
scaphoid --help

# Command help
scaphoid file --help

# Subcommand help
scaphoid file create --help

# List all commands
scaphoid help
```

## Shell Completion

Generate completion for your shell:

```bash
# Bash
scaphoid completion bash > /etc/bash_completion.d/scaphoid

# Zsh
scaphoid completion zsh > "${fpath[1]}/_scaphoid"

# Fish
scaphoid completion fish > ~/.config/fish/completions/scaphoid.fish

# PowerShell
scaphoid completion powershell > scaphoid.ps1
```

## Tips

1. **Use tab completion**: After setting up shell completion, use Tab to autocomplete commands and paths
2. **Check help first**: Use `--help` on any command to see all options
3. **Force operations**: Use `-f` flag to skip confirmations and overwrite
4. **Recursive operations**: Always use `-r` when working with directory trees
5. **Combine flags**: Short flags can be combined: `-rf` = `-r -f`

## Common Patterns

### Safe file operations with backup
```bash
scaphoid backup important.txt /backups/
scaphoid copy -p important.txt important_modified.txt
# Make changes to important_modified.txt
scaphoid compare important.txt important_modified.txt
```

### Bulk operations
```bash
# Copy multiple files to a directory
for file in *.txt; do
    scaphoid copy "$file" /backup/
done

# Compare all files in two directories
scaphoid compare -d /dir1 /dir2
```

### Conditional operations
```bash
# Only create if doesn't exist (no -f flag)
scaphoid file create config.txt || echo "File exists"

# Force create (overwrite if exists)
scaphoid file create -f config.txt
```
