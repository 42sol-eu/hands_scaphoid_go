# File-Type Specific Naming Rules

Scaphoid includes intelligent naming rules that apply to specific file types and well-known filenames.

## Log Files (*.log)

**Rule**: `log-files-iso-date-prefix`  
**Severity**: WANT (warning)  
**Description**: Log files should start with an ISO 8601 date prefix (YYYY-MM-DD)

### Examples

❌ Bad:
```
server.log
application.log
error.log
```

✅ Good:
```
2025-11-23-server.log
2025-11-23-application.log
2025-11-23-error.log
```

### Rationale

Using ISO 8601 date prefixes for log files provides several benefits:
- Automatic chronological sorting in directory listings
- Clear indication of when the log was created
- Easy to find logs from specific dates
- Consistent with the `prefer-iso-dates` rule

## README Files

**Rule**: `readme-uppercase`  
**Severity**: WANT (warning)  
**Description**: README files should use uppercase convention

### Examples

❌ Bad:
```
readme
readme.md
Readme.txt
ReadMe.md
```

✅ Good:
```
README
README.md
README.txt
README.rst
```

### Rationale

Uppercase README is a long-standing convention in software projects:
- Makes README files immediately visible in directory listings
- Universal convention across programming communities
- Follows the pattern set by LICENSE, CHANGELOG, etc.

## Common Project Files

**Rule**: `common-files-uppercase`  
**Severity**: WANT (warning)  
**Description**: Common project files should use uppercase naming

### Files Covered

The following files (with or without extensions) should use uppercase:

- LICENSE / LICENCE
- CHANGELOG
- CONTRIBUTING
- AUTHORS
- COPYING
- INSTALL
- MAKEFILE

### Examples

❌ Bad:
```
license.txt
changelog.md
contributing.md
authors
```

✅ Good:
```
LICENSE.txt
CHANGELOG.md
CONTRIBUTING.md
AUTHORS
```

### Rationale

- These are standard project documentation files
- Uppercase makes them stand out in repository root directories
- Follows established open-source conventions
- Improves discoverability for contributors

## Archive Files (*.zip)

**Rule**: `prefer-7z-archives`  
**Severity**: WANT (warning)  
**Description**: Suggests using .7z format instead of .zip for better compression

### Examples

❌ Warning:
```
backup.zip
data.zip
archive.ZIP
```

✅ Preferred:
```
backup.7z
data.7z
archive.7z
```

✅ Also Acceptable:
```
backup.tar.gz
backup.tar.bz2
backup.tar
```

### Rationale

- 7z format typically achieves better compression ratios than zip
- 7z supports stronger encryption
- Better suited for long-term archival
- Still widely supported across platforms

**Note**: This rule only suggests .7z as a preference. Using .zip does not cause an error, only a warning. Other archive formats (.tar.gz, .tar.bz2, etc.) are not flagged.

## Configuration

These rules are enabled by default when naming rules are enabled. They can be controlled through the naming rules configuration:

```yaml
naming_rules:
  enabled: true
  enforce_must: true
  show_warnings: true  # Set to false to hide these WANT-level warnings
```

## Integration with Check Command

All file-type-specific rules are automatically included when using the `check` command:

```bash
# Check current directory recursively
scaphoid check . -r

# Show only errors (skip WANT warnings)
scaphoid check . -r --errors-only

# Include valid files in output
scaphoid check . -r --show-valid
```

## Testing Examples

Create test files to see the rules in action:

```bash
# Log file without date prefix - generates warning
scaphoid create file server.log

# Log file with ISO date prefix - no warning
scaphoid create file 2025-11-23-server.log

# Lowercase readme - generates warning
scaphoid create file readme.md

# Uppercase README - no warning
scaphoid create file README.md

# ZIP archive - generates warning
scaphoid create archive backup.zip

# 7z archive - no warning
scaphoid create archive backup.7z

# Check all files in current directory
scaphoid check . -r
```

## Technical Details

### Implementation

File-type-specific rules are implemented in two helper functions in `pkg/naming/rules.go`:

1. `addExtensionRules()` - Rules based on file extensions (*.log, *.zip)
2. `addSpecialFilenameRules()` - Rules for specific filenames (README, LICENSE, etc.)

These functions are called from `NewValidator()` during initialization.

### Pattern Matching

- **Log files**: Checks if filename ends with `.log` (case-insensitive), then validates ISO date prefix
- **README files**: Checks if filename starts with "readme" (case-insensitive), then enforces uppercase
- **Common files**: Compares basename against list of common files (case-insensitive), then enforces uppercase
- **ZIP archives**: Checks if filename ends with `.zip` (case-insensitive), then suggests .7z alternative

### Test Coverage

Comprehensive test suites are available:

- `pkg/naming/extension_test.go` - Tests for extension-based rules
- `pkg/naming/filetype_test.go` - Tests for specific file-type rules

Run tests with:
```bash
go test ./pkg/naming/... -v -run "TestLogFile|TestReadme|TestCommonFiles|TestArchive"
```
