---
title: "Examples"
date: 2025-11-23
draft: false
---

# Examples and Common Workflows

This page provides practical examples and common workflows using the Scaphoid filesystem commands.

## Basic File Operations

### Creating and Managing Files

```bash
# Create a new file
create file notes.txt
# or
file create notes.txt

# Get information about the file
info notes.txt

# Copy the file
copy notes.txt notes_backup.txt
# or
file copy notes.txt notes_backup.txt

# Rename the file
move notes.txt important_notes.txt
# or
file move notes.txt important_notes.txt
```

### Working with Directories

```bash
# Create a project directory structure
create directory project
create directory project/src
create directory project/tests
create directory project/docs

# Or using object-first syntax
directory create project
directory create project/src
directory create project/tests
directory create project/docs

# List the structure
list project --recursive
# or
directory list project --recursive

# Copy entire directory
copy project project_backup --recursive
# or
directory copy project project_backup --recursive
```

## Advanced Operations

### Working with Symbolic Links

```bash
# Create a symbolic link to a file
create link current_config.json --target /etc/app/config.json
# or
link create current_config.json --target /etc/app/config.json

# Create a link to a directory
create link current_data --target /var/lib/app/data
# or
link create current_data --target /var/lib/app/data

# List the link (shows target)
info current_config.json
list current_data
```

### Archive Management

```bash
# Create an archive from a directory
create archive project_backup.tar.gz --source project
# or
archive create project_backup.tar.gz --source project

# List archive contents
list project_backup.tar.gz
# or
archive list project_backup.tar.gz

# Compare original with extracted version (after extraction)
compare project extracted_project
```

## Common Workflows

### Project Setup Workflow

```bash
#!/bin/bash
# Setup a new Go project structure

PROJECT_NAME="myproject"

# Create main directories
create directory $PROJECT_NAME
create directory $PROJECT_NAME/cmd
create directory $PROJECT_NAME/pkg
create directory $PROJECT_NAME/internal
create directory $PROJECT_NAME/api
create directory $PROJECT_NAME/web
create directory $PROJECT_NAME/scripts
create directory $PROJECT_NAME/deployments
create directory $PROJECT_NAME/test
create directory $PROJECT_NAME/docs

# Create initial files
create file $PROJECT_NAME/README.md
create file $PROJECT_NAME/go.mod
create file $PROJECT_NAME/Makefile
create file $PROJECT_NAME/.gitignore

# Create symbolic link to current project
link create current_project --target $PROJECT_NAME

echo "Project structure created!"
list $PROJECT_NAME --recursive
```

### Backup and Archive Workflow

```bash
#!/bin/bash
# Complete backup workflow

SOURCE_DIR="/important/data"
BACKUP_DIR="/backups"
DATE=$(date +%Y%m%d)

# Create backup directory for today
create directory $BACKUP_DIR/$DATE

# Create full backup archive
create archive $BACKUP_DIR/$DATE/full_backup.tar.gz --source $SOURCE_DIR

# Create individual file backups for critical files
backup $SOURCE_DIR/config.json $BACKUP_DIR/$DATE/
backup $SOURCE_DIR/database.sqlite $BACKUP_DIR/$DATE/

# Verify backup integrity
info $BACKUP_DIR/$DATE/full_backup.tar.gz
list $BACKUP_DIR/$DATE/full_backup.tar.gz

echo "Backup completed for $DATE"
list $BACKUP_DIR/$DATE --long
```

### File Organization Workflow

```bash
#!/bin/bash
# Organize downloads folder

DOWNLOADS="/Users/$(whoami)/Downloads"
ORGANIZED="/Users/$(whoami)/Documents/Organized"

# Create organized structure
create directory $ORGANIZED
create directory $ORGANIZED/images
create directory $ORGANIZED/documents  
create directory $ORGANIZED/archives
create directory $ORGANIZED/others

# Move files by type (simplified example)
# In real usage, you'd want more sophisticated file type detection

# Get list of files
list $DOWNLOADS > /tmp/downloads_list.txt

# Example of moving specific file types
# Note: This is a simplified example - real implementation would need
# proper file type detection and error handling

echo "Organizing files in $DOWNLOADS"
echo "Created organized structure in $ORGANIZED"
list $ORGANIZED --recursive
```

### Synchronization Workflow

```bash
#!/bin/bash
# Synchronize two directories

SOURCE="/source/directory"
TARGET="/target/directory"

echo "Synchronizing $SOURCE to $TARGET"

# Compare directories first
compare $SOURCE $TARGET

# If different, backup target and copy source
if [ $? -ne 0 ]; then
    echo "Directories differ, creating backup and syncing"
    
    # Create backup of target
    backup $TARGET /backups/
    
    # Remove target and copy source
    delete $TARGET --recursive --force
    copy $SOURCE $TARGET --recursive --preserve
    
    echo "Synchronization complete"
else
    echo "Directories are already synchronized"
fi
```

## Complex Examples

### Multi-step Data Processing

```bash
#!/bin/bash
# Process data files with staging

WORKSPACE="/tmp/processing"
DATA_DIR="/data/input"
RESULTS_DIR="/data/output"

# Setup workspace
create directory $WORKSPACE
create directory $WORKSPACE/staging
create directory $WORKSPACE/processed
create directory $WORKSPACE/failed

# Create links to data sources
link create $WORKSPACE/input --target $DATA_DIR

# Stage files for processing
list $DATA_DIR | while read -r file; do
    if [[ $file == *.csv ]]; then
        copy "$DATA_DIR/$file" "$WORKSPACE/staging/"
    fi
done

echo "Staged files:"
list $WORKSPACE/staging

# After processing (placeholder for actual processing)
# move processed files to results
# move failed files to failed directory

# Cleanup
echo "Processing complete, cleaning up workspace"
delete $WORKSPACE --recursive --force
```

### Configuration Management

```bash
#!/bin/bash
# Manage application configurations

CONFIG_BASE="/etc/myapp"
CONFIG_ENV="production"  # or development, staging, etc.

# Create environment-specific config
create directory $CONFIG_BASE/environments
create directory $CONFIG_BASE/environments/$CONFIG_ENV

# Copy base config
copy $CONFIG_BASE/base-config.json $CONFIG_BASE/environments/$CONFIG_ENV/config.json

# Create link to active config
link create $CONFIG_BASE/active-config.json --target $CONFIG_BASE/environments/$CONFIG_ENV/config.json

# Backup current configuration
backup $CONFIG_BASE/environments/$CONFIG_ENV /backups/configs/

echo "Configuration set to $CONFIG_ENV"
info $CONFIG_BASE/active-config.json
```

## Tips and Best Practices

### Using with Scripts

1. **Always check exit codes:**
```bash
if ! create directory /important/path; then
    echo "Failed to create directory"
    exit 1
fi
```

2. **Use the `--force` flag carefully:**
```bash
# Safe: only force when you're sure
delete /tmp/safe-to-delete --force

# Be careful with recursive deletes
delete /path/to/directory --recursive --force
```

3. **Prefer explicit object types:**
```bash
# More explicit and clear
create directory mydir
create file myfile.txt

# Less clear what type will be created
create myitem
```

### Performance Considerations

- Use `--recursive` carefully with large directory trees
- For large archives, consider using the appropriate compression level
- When comparing large files, the `--deep` flag will be slower but more accurate

### Error Handling

Most commands return non-zero exit codes on failure:

```bash
if copy large-file.dat backup/; then
    echo "Backup successful"
    delete large-file.dat
else
    echo "Backup failed, keeping original"
fi
```