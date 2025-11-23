# Markdown Output Formatting

Scaphoid supports markdown-like formatting for console output, making it easy to integrate command output into documentation, reports, or simply improve readability.

## Features

### Automatic Prefixes

When markdown formatting is enabled (default), all output is automatically prefixed:

- **Success messages**: Prefixed with `- ` (unordered list item)
- **Error messages**: Prefixed with `> ` (blockquote)

### Section Headers

Use the `section` command to create markdown section headers:

```bash
scaphoid section "File Operations"
# Output: ## File Operations

scaphoid section "File Operations" "Testing file creation and deletion"
# Output:
# ## File Operations
# Testing file creation and deletion
```

### Separators

Use the `separator` command to add horizontal rules:

```bash
scaphoid separator
# Output: ---
```

## Configuration

Markdown formatting is controlled via `.scaphoid.yaml`:

```yaml
formatting:
  # Enable/disable markdown-like formatting
  markdown: true
  
  # Customize prefixes
  success_prefix: "- "
  error_prefix: "> "
```

### Disable Markdown Formatting

To use plain output without markdown prefixes:

```yaml
formatting:
  markdown: false
```

## Examples

### Basic Workflow

```bash
scaphoid section "File Operations" "Creating test files"
scaphoid file create test1.txt
scaphoid file create test2.txt
scaphoid separator
scaphoid section "Verification"
scaphoid file exists test1.txt
scaphoid file exists test2.txt
scaphoid separator
scaphoid section "Cleanup"
scaphoid file delete test1.txt -f
scaphoid file delete test2.txt -f
```

**Output:**
```markdown
## File Operations
Creating test files
- Successfully created file at "test1.txt"
- Successfully created file at "test2.txt"
---
## Verification
- "test1.txt" exists and is a file
- "test2.txt" exists and is a file
---
## Cleanup
- Successfully deleted "test1.txt"
- Successfully deleted "test2.txt"
```

### Error Handling

```bash
scaphoid section "Error Test"
scaphoid file exists /nonexistent/path.txt
```

**Output:**
```markdown
## Error Test
> "/nonexistent/path.txt" does not exist
```

### Mixed Success and Errors

```bash
scaphoid section "Mixed Operations"
scaphoid file create /tmp/success.txt
scaphoid file exists /tmp/success.txt
scaphoid file exists /tmp/failure.txt
scaphoid file delete /tmp/success.txt -f
```

**Output:**
```markdown
## Mixed Operations
- Successfully created file at "/tmp/success.txt"
- "/tmp/success.txt" exists and is a file
> "/tmp/failure.txt" does not exist
- Successfully deleted "/tmp/success.txt"
```

## Use Cases

### 1. Documentation Generation

Generate markdown documentation directly from command execution:

```bash
{
  scaphoid section "Setup" "Initialize project structure"
  scaphoid directory create -r ./project/src
  scaphoid directory create -r ./project/tests
  scaphoid file create ./project/README.md
  
  scaphoid separator
  
  scaphoid section "Verification" "Check created structure"
  scaphoid directory exists ./project/src
  scaphoid directory exists ./project/tests
  scaphoid file exists ./project/README.md
} > setup-log.md
```

### 2. Build Scripts

Create readable build logs:

```bash
#!/bin/bash
scaphoid section "Build Process" "Building application components"
scaphoid file copy src/main.go build/main.go
scaphoid directory copy -r assets/ build/assets/
scaphoid separator
scaphoid section "Build Status"
scaphoid file exists build/main.go
scaphoid directory exists build/assets
```

### 3. Testing Reports

Generate test execution reports:

```bash
scaphoid section "Integration Tests" "Running filesystem tests"
scaphoid file create /tmp/test1.txt && echo "- Test 1: PASS"
scaphoid file create /tmp/test2.txt && echo "- Test 2: PASS"
scaphoid separator
scaphoid section "Cleanup"
scaphoid file delete /tmp/test1.txt -f
scaphoid file delete /tmp/test2.txt -f
```

### 4. Deployment Logs

Track deployment steps:

```bash
scaphoid section "Deployment" "Deploying to production"
scaphoid file copy app.js /var/www/app.js
scaphoid directory copy -r static/ /var/www/static/
scaphoid separator
scaphoid section "Verification"
scaphoid file exists /var/www/app.js
scaphoid directory exists /var/www/static
```

## Custom Prefixes

You can customize the prefixes in your config file:

```yaml
formatting:
  markdown: true
  success_prefix: "✓ "  # Checkmark
  error_prefix: "✗ "    # X mark
```

Or use different markdown styles:

```yaml
formatting:
  markdown: true
  success_prefix: "* "  # Alternative list marker
  error_prefix: "! "    # Warning style
```

## Combining with Colors

Markdown formatting works seamlessly with colored output. When viewed in a terminal:

- Section headers appear in regular text
- Success messages (`- ` prefix) show with colored paths and actions
- Error messages (`> ` prefix) appear in red
- Separators (`---`) create visual breaks

When piped to a file or viewed in a markdown renderer, the formatting remains clean and readable.

## Tips

1. **Pipe to File**: Redirect output to create instant markdown documentation
   ```bash
   scaphoid section "Operations" > log.md
   scaphoid file create test.txt >> log.md
   ```

2. **Combine Commands**: Chain multiple operations for complete workflows
   ```bash
   scaphoid section "Workflow" && \
   scaphoid file create temp.txt && \
   scaphoid separator && \
   scaphoid file delete temp.txt -f
   ```

3. **Script Integration**: Use in shell scripts for self-documenting execution
   ```bash
   #!/bin/bash
   scaphoid section "Script Execution" "$(date)"
   # Your operations here
   scaphoid separator
   ```

4. **Disable When Needed**: Turn off markdown formatting for plain output
   ```yaml
   formatting:
     markdown: false
   ```

## Future Enhancements

Planned features include:

- Numbered list support for batch operations
- Code block formatting for verbose output
- Optional emoji/icon prefixes (✓, ✗, ℹ, ⚠)
- Info and warning message types
- Custom section header levels (###, ####)
