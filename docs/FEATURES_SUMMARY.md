# Scaphoid Features Summary

This document provides an overview of all completed features in Scaphoid.

## Markdown Output Formatting

### Overview
Scaphoid supports markdown-like formatting for console output, making command output easily integratable into documentation and reports.

### Features
- ✅ **Success prefixes**: `- ✓ ` for successful operations
- ✅ **Error prefixes**: `> ✗ ` for error messages
- ✅ **Numbered lists**: `1. `, `2. `, `3. ` for batch operations (list command)
- ✅ **Section headers**: `## Title` with optional details
- ✅ **Separators**: `---` for visual breaks
- ✅ **Configurable**: All formatting controlled via `.scaphoid.yaml`

### Commands
```bash
# Section header
scaphoid section "Title" "Optional details"

# Horizontal separator
scaphoid separator

# List with numbered output
scaphoid list /path/to/directory
```

### Configuration
```yaml
formatting:
  markdown: true           # Enable markdown formatting
  numbered_list: true      # Use numbered lists for batch ops
  success_prefix: "- ✓ "   # Customize success prefix
  error_prefix: "> ✗ "     # Customize error prefix
  heading_prefix: "\n## "  # Section heading format
```

### Documentation
- `docs/MARKDOWN_OUTPUT.md` - Comprehensive guide with examples
- `docs/MARKDOWN_IMPLEMENTATION.md` - Technical implementation details

---

## Interactive Prompts

### Overview
Scaphoid integrates the `survey/v2` library to provide Python Click-like interactive user prompts.

### Features
- ✅ **Confirmation**: Yes/No prompts
- ✅ **Selection**: Choose one option from a list
- ✅ **Multi-select**: Choose multiple options
- ✅ **Text input**: Enter text with optional default
- ✅ **Password**: Hidden password input

### Helper Functions
```go
// In cmd_new/interactive.go
Confirm(message, default) -> (bool, error)
Select(message, options) -> (string, error)
MultiSelect(message, options) -> ([]string, error)
Input(message, default) -> (string, error)
Password(message) -> (string, error)
```

### Integrated Commands
1. **file delete** - Asks for confirmation unless `--force` flag is used
2. **create** - Prompts for type if omitted (accepts 1 or 2 args)
3. **interactive** - Demo command showing all prompt types

### Usage Examples
```bash
# Delete with confirmation
scaphoid file delete myfile.txt
? Delete /tmp/myfile.txt? (y/N)

# Create with interactive type selection
scaphoid create myfile.txt
? Select object type to create: [file, directory, link, archive]

# Run interactive demo
scaphoid interactive
```

### Documentation
- `docs/INTERACTIVE_PROMPTS.md` - Comprehensive guide (300+ lines)
- `docs/INTERACTIVE_QUICKREF.md` - Quick reference for developers
- `demo_interactive.sh` - Executable demo script

---

## Color Configuration

### Overview
Customizable ANSI color output for enhanced terminal readability.

### Features
- ✅ **Colored output**: Actions, types, paths, errors, success messages
- ✅ **Bold support**: Optional bold styling for each color category
- ✅ **Configurable**: All colors defined in `.scaphoid.yaml`
- ✅ **Smart disabling**: Automatically disables when piped (non-TTY)

### Configuration
```yaml
colors:
  action:
    color: "blue"
    bold: true
  type:
    color: "green"
    bold: true
  error:
    color: "red"
    bold: true
  success:
    color: "green"
    bold: true
  path:
    color: "green"
    bold: true
  keyword:
    color: "green"
    bold: true
```

### Documentation
- `docs/COLOR_CONFIG.md` - Color configuration guide
- `docs/COLOR_IMPLEMENTATION.md` - Implementation details

---

## File Structure

### Core Implementation
```
cmd_new/
├── colors.go           # Color and formatting functions
├── formatting.go       # Section and separator commands
├── interactive.go      # Interactive prompt helpers
├── file.go            # File operations (with confirmation)
├── directory.go       # Directory operations
├── general.go         # General commands (create, list, etc.)
└── root.go            # Root command and registration

docs/
├── MARKDOWN_OUTPUT.md           # Markdown formatting guide
├── MARKDOWN_IMPLEMENTATION.md   # Markdown technical docs
├── INTERACTIVE_PROMPTS.md       # Interactive prompts guide
├── INTERACTIVE_QUICKREF.md      # Quick reference
├── COLOR_CONFIG.md              # Color configuration
├── COLOR_IMPLEMENTATION.md      # Color technical docs
└── FEATURES_SUMMARY.md          # This file

demo_interactive.sh              # Interactive demo script
.scaphoid.yaml                   # Configuration file
```

---

## Key Functions

### Formatting Functions (cmd_new/colors.go)
```go
// Prefixes
getPrefix(messageType) -> string
getNumberedPrefix(index, total) -> string

// Wrappers
wrapSuccessMessage(msg) -> string
wrapErrorMessage(msg) -> string
wrapSuccessMessageNumbered(msg, index, total) -> string
wrapErrorMessageNumbered(msg, index, total) -> string

// Formatters
formatSectionHeader(title, details) -> string
formatSuccessMessage(action, path) -> string
formatErrorMessage(errMsg) -> string
formatCreateMessage(objType, path) -> string
```

### Interactive Functions (cmd_new/interactive.go)
```go
Confirm(message string, defaultValue bool) -> (bool, error)
Select(message string, options []string) -> (string, error)
MultiSelect(message string, options []string) -> ([]string, error)
Input(message string, defaultValue string) -> (string, error)
Password(message string) -> (string, error)
```

---

## Testing

### Manual Testing
```bash
# Test markdown formatting
./bin/scaphoid section "Test" "Details"
./bin/scaphoid separator
./bin/scaphoid list docs | head -5

# Test interactive prompts
./bin/scaphoid interactive
./bin/scaphoid create /tmp/test.txt
./bin/scaphoid file delete /tmp/test.txt

# Test with config changes
# Edit .scaphoid.yaml and re-run commands
```

### Demo Script
```bash
# Run the comprehensive interactive demo
./demo_interactive.sh
```

---

## Dependencies

### Go Modules
```
github.com/spf13/cobra v1.10.1        # CLI framework
github.com/spf13/viper v1.21.0        # Configuration
github.com/fatih/color v1.18.0        # ANSI colors
github.com/AlecAivazis/survey/v2 v2.3.7  # Interactive prompts
```

### Installation
```bash
go get github.com/spf13/cobra@latest
go get github.com/spf13/viper@latest
go get github.com/fatih/color@latest
go get github.com/AlecAivazis/survey/v2@latest
```

---

## Future Enhancements

### Not Yet Implemented
- Code block formatting for verbose output
- Info/warning message types with dedicated prefixes
- Apply interactive prompts to more commands
- Additional prompt types (date, time, autocomplete)

### Potential Improvements
- Theme support (predefined color schemes)
- Plugin system for custom formatters
- Markdown table output for structured data
- Progress bars for long operations
- Spinner animations for background tasks

---

## Quick Reference

### Common Commands
```bash
# Markdown formatting
scaphoid section "Title" "Details"
scaphoid separator

# List with numbers
scaphoid list /path

# Interactive create
scaphoid create myfile.txt

# Delete with confirmation
scaphoid file delete myfile.txt

# Force without prompts
scaphoid file delete myfile.txt --force

# Demo all features
scaphoid interactive
./demo_interactive.sh
```

### Configuration Locations
1. `./.scaphoid.yaml` (current directory)
2. `~/.scaphoid.yaml` (home directory)
3. `/etc/scaphoid/.scaphoid.yaml` (system-wide)

### Disable Features
```yaml
formatting:
  markdown: false       # Disable markdown formatting
  numbered_list: false  # Use bullets instead of numbers
```

---

## Resources

- **GitHub**: [42sol-eu/hands_scaphoid_go](https://github.com/42sol-eu/hands_scaphoid_go)
- **Survey Library**: [AlecAivazis/survey](https://github.com/AlecAivazis/survey)
- **Cobra CLI**: [spf13/cobra](https://github.com/spf13/cobra)
- **Color Library**: [fatih/color](https://github.com/fatih/color)

---

*Last Updated: November 23, 2025*
