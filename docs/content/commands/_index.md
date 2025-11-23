---
title: "Command Reference"
date: 2025-11-23
draft: false
---

# Command Reference

Complete reference for all Scaphoid filesystem commands. Each command supports both action-first and object-first syntax.

## Action Commands

### create

Create new filesystem objects.

**Syntax:**
```bash
create [object-type] [path] [flags]
```

**Object types:** `directory`, `file`, `link`, `archive`

**Flags:**
- `--force, -f` - Force creation, overwrite existing
- `--recursive, -r` - Create parent directories if needed
- `--target, -t` - Target path for links (required for link creation)
- `--source, -s` - Source directory for archives

**Examples:**
```bash
create directory /path/to/newdir
create file /path/to/newfile.txt
create link /path/to/link --target /path/to/target
create archive backup.tar.gz --source /path/to/backup
```

---

### list

List contents of filesystem objects.

**Syntax:**
```bash
list [path] [flags]
```

**Flags:**
- `--all, -a` - Show hidden files
- `--long, -l` - Use long listing format
- `--recursive, -r` - List subdirectories recursively
- `--sort [name|size|time]` - Sort by specified field
- `--reverse` - Reverse sort order

**Examples:**
```bash
list /path/to/directory
list /path/to/archive.tar.gz
list /path/to/directory --long --all
```

---

### copy

Copy filesystem objects.

**Syntax:**
```bash
copy [source] [destination] [flags]
```

**Flags:**
- `--recursive, -r` - Copy directories recursively
- `--force, -f` - Force copy, overwrite existing
- `--preserve, -p` - Preserve metadata (permissions, timestamps)

**Examples:**
```bash
copy file1.txt file2.txt
copy /path/to/source/dir /path/to/dest/dir --recursive
copy important.doc backup/ --preserve
```

---

### move

Move or rename filesystem objects.

**Syntax:**
```bash
move [source] [destination] [flags]
```

**Flags:**
- `--force, -f` - Force move, overwrite existing

**Examples:**
```bash
move oldname.txt newname.txt
move /old/path /new/path
move file.txt /different/directory/
```

---

### delete

Delete filesystem objects.

**Syntax:**
```bash
delete [path] [flags]
```

**Flags:**
- `--force, -f` - Force deletion without confirmation
- `--recursive, -r` - Delete directories recursively

**Examples:**
```bash
delete unwanted.txt
delete /path/to/directory --recursive --force
delete broken_link
```

---

### info

Get detailed information about filesystem objects.

**Syntax:**
```bash
info [path] [flags]
```

**Flags:**
- `--all, -a` - Show all available information

**Examples:**
```bash
info document.pdf
info /path/to/directory --all
info symlink
```

---

### compare

Compare two filesystem objects.

**Syntax:**
```bash
compare [path1] [path2] [flags]
```

**Flags:**
- `--deep, -d` - Perform deep comparison (content-based)

**Examples:**
```bash
compare file1.txt file2.txt
compare /dir1 /dir2
compare original.jpg copy.jpg --deep
```

---

### backup

Create backups of filesystem objects.

**Syntax:**
```bash
backup [source] [backup-directory] [flags]
```

**Flags:**
- `--timestamp, -t` - Add timestamp to backup name (default: true)
- `--compress, -c` - Compress backup

**Examples:**
```bash
backup important.doc /backups/
backup /project /backups/ --compress
backup config.ini /safe/location/
```

## Object Commands

### directory

Directory-specific operations.

**Syntax:**
```bash
directory [action] [arguments] [flags]
```

**Actions:** `create`, `list`, `copy`, `move`, `delete`

**Examples:**
```bash
directory create /new/path
directory list /existing/path --long
directory copy /source /destination --recursive
directory move /old/path /new/path
directory delete /unwanted/path --recursive --force
```

---

### file

File-specific operations.

**Syntax:**
```bash
file [action] [arguments] [flags]
```

**Actions:** `create`, `list`, `copy`, `move`, `delete`

**Examples:**
```bash
file create newfile.txt
file copy source.txt destination.txt
file move oldname.txt newname.txt
file delete unwanted.txt --force
```

---

### link

Symbolic link operations.

**Syntax:**
```bash
link [action] [arguments] [flags]
```

**Actions:** `create`, `list`, `copy`, `move`, `delete`

**Special flags for create:**
- `--target, -t` - Target path for the symbolic link (required)

**Examples:**
```bash
link create mylink --target /path/to/target
link list mylink
link copy mylink mylink_copy
link delete broken_link
```

---

### archive

Archive file operations.

**Syntax:**
```bash
archive [action] [arguments] [flags]
```

**Actions:** `create`, `list`, `copy`, `move`, `delete`

**Special flags for create:**
- `--source, -s` - Source directory to archive

**Examples:**
```bash
archive create backup.tar.gz --source /data
archive list backup.tar.gz
archive copy backup.tar.gz backup_copy.tar.gz
archive delete old_backup.tar.gz
```

## Global Flags

These flags are available for most commands:

- `--help, -h` - Show command help
- `--version, -v` - Show version information (where applicable)

## Exit Codes

- `0` - Success
- `1` - General error
- `2` - Usage error (invalid arguments)

## Configuration

Currently, the commands use default settings. Future versions may support configuration files.