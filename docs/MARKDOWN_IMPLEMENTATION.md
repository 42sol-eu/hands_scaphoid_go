# Markdown Output Implementation Summary

## Date: November 23, 2025

## Overview

Implemented markdown-like console output formatting for Scaphoid, allowing commands to generate clean, structured output suitable for documentation, reports, and improved terminal readability.

## Features Implemented

### 1. Automatic Markdown Prefixes

All output messages are now automatically prefixed based on their type:

- **Success messages**: `- ` (unordered list item)
- **Error messages**: `> ` (blockquote)

This is controlled by the `formatting.markdown` setting in `.scaphoid.yaml`.

### 2. Section Command

Added `scaphoid section [title] [details]` command that outputs markdown headers:

```bash
scaphoid section "File Operations"
# Output: ## File Operations

scaphoid section "Build" "Compiling sources"
# Output:
# ## Build
# Compiling sources
```

### 3. Separator Command

Added `scaphoid separator` command that outputs horizontal rules:

```bash
scaphoid separator
# Output: ---
```

### 4. Configurable Formatting

Extended `.scaphoid.yaml` with formatting options:

```yaml
formatting:
  markdown: true              # Enable/disable markdown formatting
  success_prefix: "- "        # Prefix for success messages
  error_prefix: "> "          # Prefix for error messages
```

Users can customize prefixes, including emoji alternatives:
- `success_prefix: "✓ "` - Checkmark
- `error_prefix: "✗ "` - X mark
- `success_prefix: "* "` - Alternative markdown list

## Code Changes

### New Files

1. **cmd_new/formatting.go**
   - Created new file with `sectionCmd` and `separatorCmd`
   - Section command supports title and optional details
   - Separator command outputs `---`

2. **docs/MARKDOWN_OUTPUT.md**
   - Comprehensive documentation with examples
   - Use cases: documentation generation, build scripts, testing reports
   - Configuration options and customization tips

### Modified Files

1. **cmd_new/colors.go**
   - Added `getPrefix(messageType string)` function
   - Added `formatErrorMessage(errMsg string)` function
   - Added `formatSectionHeader(title, details string)` function
   - Updated all format functions to include markdown prefixes:
     - `formatExistsMessage()` - Adds `- ` or `> ` prefix
     - `formatTypeCheckMessage()` - Adds appropriate prefix
     - `formatSuccessMessage()` - Adds `- ` prefix
     - `formatCreateMessage()` - Adds `- ` prefix
     - `formatCopyMoveMessage()` - Adds `- ` prefix
   - Updated `setDefaultColors()` to include formatting defaults

2. **cmd_new/root.go**
   - Added `sectionCmd` and `separatorCmd` to root command

3. **cmd_new/file.go**
   - Updated error messages to use `formatErrorMessage()`
   - All success messages now have markdown prefixes

4. **cmd_new/exists.go**
   - Updated error messages to use `formatErrorMessage()`
   - All existence checks now have markdown prefixes

5. **.scaphoid.yaml**
   - Added `formatting` section with:
     - `markdown: true` - Enable markdown formatting
     - `success_prefix: "- "` - List item prefix
     - `error_prefix: "> "` - Blockquote prefix
   - Added comments with alternative emoji options

6. **README.md**
   - Added "Markdown Output Formatting" section
   - Included examples of section and separator commands
   - Added reference to MARKDOWN_OUTPUT.md

7. **docs/COLOR_CONFIG.md**
   - Updated to show formatting configuration
   - Added reference to MARKDOWN_OUTPUT.md

## Example Output

### Before
```
Successfully created file at "/tmp/test.txt"
"/tmp/test.txt" exists and is a file
"/tmp/nonexistent.txt" does not exist
```

### After (with markdown formatting)
```
- Successfully created file at "/tmp/test.txt"
- "/tmp/test.txt" exists and is a file
> "/tmp/nonexistent.txt" does not exist
```

### Complete Workflow Example
```bash
scaphoid section "File Operations" "Creating test files"
scaphoid file create /tmp/test1.txt
scaphoid file create /tmp/test2.txt
scaphoid separator
scaphoid section "Verification"
scaphoid file exists /tmp/test1.txt
scaphoid file exists /tmp/nonexistent.txt
scaphoid separator
scaphoid file delete /tmp/test1.txt -f
```

**Output:**
```markdown
## File Operations
Creating test files
- Successfully created file at "/tmp/test1.txt"
- Successfully created file at "/tmp/test2.txt"
---
## Verification
- "/tmp/test1.txt" exists and is a file
> "/tmp/nonexistent.txt" does not exist
---
- Successfully deleted "/tmp/test1.txt"
```

## Use Cases

1. **Documentation Generation**
   - Pipe command output directly to markdown files
   - Auto-generate execution logs

2. **Build Scripts**
   - Create readable build logs with sections
   - Track deployment steps with clear structure

3. **Testing Reports**
   - Generate test execution reports
   - Clear success/failure indicators

4. **CI/CD Integration**
   - Structured output for pipeline logs
   - Easy parsing of success/error states

## Benefits

1. **Readability**: Clear visual structure with markdown formatting
2. **Documentation**: Output can be directly used as markdown documentation
3. **Flexibility**: Configurable prefixes for different styles
4. **Integration**: Works seamlessly with colored output
5. **Scriptability**: Easy to parse success/error indicators
6. **Professional**: Clean, structured output for reports and logs

## Configuration Options

### Enable/Disable

```yaml
formatting:
  markdown: false  # Disable all markdown formatting
```

### Custom Prefixes

```yaml
formatting:
  markdown: true
  success_prefix: "✓ "  # Use checkmark
  error_prefix: "✗ "    # Use X mark
```

### Alternative Styles

```yaml
formatting:
  markdown: true
  success_prefix: "* "  # Alternative markdown list
  error_prefix: "! "    # Warning style prefix
```

## Testing

All features tested and verified:
- ✅ Section headers with and without details
- ✅ Separator command outputs `---`
- ✅ Success messages prefixed with `- `
- ✅ Error messages prefixed with `> `
- ✅ Configuration loading and defaults
- ✅ Custom prefix configuration
- ✅ Markdown output piped to files
- ✅ Integration with colored output
- ✅ Complete workflow examples

## Future Enhancements (in TODO)

1. **Numbered Lists**: For batch operations (`1. `, `2. `, etc.)
2. **Code Blocks**: Wrap verbose output in backticks/code blocks
3. **Emoji Support**: Optional emoji prefixes (✓, ✗, ℹ, ⚠)
4. **Info/Warning Types**: New message types with specific formatting
5. **Custom Section Levels**: Support for ### and #### headers
6. **Template Support**: User-defined output templates

## Backward Compatibility

- Fully backward compatible - existing commands work unchanged
- Markdown formatting can be disabled via config
- Default behavior provides enhanced output
- No breaking changes to command syntax

## Documentation

Created comprehensive documentation:
- `docs/MARKDOWN_OUTPUT.md` - Full guide with examples
- Updated `README.md` - Added markdown formatting section
- Updated `docs/COLOR_CONFIG.md` - Added formatting options
- Updated `.scaphoid.yaml` - Added formatting configuration with examples
