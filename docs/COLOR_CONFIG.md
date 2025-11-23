# Scaphoid Color Configuration

Scaphoid supports customizable colored output through a YAML configuration file.

## Configuration File Location

Scaphoid searches for `.scaphoid.yaml` in the following locations (in order):

1. Current directory (`./.scaphoid.yaml`)
2. User's home directory (`~/.scaphoid.yaml`)
3. System-wide config (`/etc/scaphoid/.scaphoid.yaml`)

The first configuration file found will be used.

## Configuration Format

The configuration file supports both color and formatting options:

```yaml
colors:
  # Color for action verbs (exists, create, delete, copy, move, etc.)
  action:
    color: "blue"
    bold: true
  
  # Color for filesystem object types (file, directory, link, archive)
  type:
    color: "green"
    bold: true
  
  # Color for error messages
  error:
    color: "red"
    bold: true
  
  # Color for success messages
  success:
    color: "green"
    bold: true
  
  # Color for paths
  path:
    color: "green"
    bold: true
  
  # Color for keywords
  keyword:
    color: "green"
    bold: true

# Output formatting options
formatting:
  # Enable markdown-like formatting (default: true)
  markdown: true
  
  # Prefix for successful operations
  success_prefix: "- "
  
  # Prefix for error messages
  error_prefix: "> "
```

See [MARKDOWN_OUTPUT.md](MARKDOWN_OUTPUT.md) for details on markdown formatting features.

## Available Colors

- `black`
- `red`
- `green`
- `yellow`
- `blue`
- `magenta`
- `cyan`
- `white`

## Bold Option

Set `bold: true` or `bold: false` for each color category to control text weight.

## Default Configuration

If no configuration file is found, Scaphoid uses these defaults:

- **Actions** (exists, create, delete, etc.): Blue, Bold
- **Types** (file, directory, link, archive): Green, Bold
- **Paths**: Green, Bold (shown in double quotes)
- **Errors**: Red, Bold
- **Success messages**: Green, Bold
- **Keywords**: Green, Bold

## Examples

### Minimal Configuration

Override just the colors you want to change:

```yaml
colors:
  action:
    color: "magenta"
  path:
    color: "yellow"
```

### Disable Bold

```yaml
colors:
  action:
    color: "blue"
    bold: false
  type:
    color: "green"
    bold: false
```

### High Contrast Theme

```yaml
colors:
  action:
    color: "cyan"
    bold: true
  type:
    color: "yellow"
    bold: true
  path:
    color: "white"
    bold: true
  error:
    color: "red"
    bold: true
```

## Testing Your Configuration

After creating your `.scaphoid.yaml` file, test it with:

```bash
scaphoid exists /path/to/some/file
scaphoid file create /tmp/test.txt
scaphoid file exists /tmp/test.txt
```

## Troubleshooting

- Make sure your YAML file is properly formatted
- Valid color names are lowercase
- The `bold` field must be `true` or `false` (not quoted)
- If colors don't appear, check your terminal supports ANSI colors

## Library Used

Scaphoid uses:
- `github.com/spf13/viper` for configuration management
- `github.com/fatih/color` for colored terminal output
