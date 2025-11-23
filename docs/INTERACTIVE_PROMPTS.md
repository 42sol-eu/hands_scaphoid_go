# Interactive User Prompts

Scaphoid supports interactive user prompts using the `survey/v2` library, providing a rich set of interactive elements similar to Python's Click or Inquirer.

## Available Prompt Types

### 1. Confirmation (Yes/No)

Ask for user confirmation:

```go
confirmed, err := Confirm("Delete this file?", false)
if err != nil {
    // Handle error
}
if confirmed {
    // User said yes
}
```

**Example in Scaphoid:**
```bash
$ scaphoid file delete myfile.txt
? Delete myfile.txt? (y/N) y
- ✓ Successfully deleted "myfile.txt"
```

### 2. Select (Single Choice)

Choose one option from a list:

```go
choice, err := Select("Choose an operation:", []string{
    "Create file",
    "Delete file", 
    "Copy file",
})
```

**Visual Example:**
```bash
? Choose an operation: 
  ▸ Create file
    Delete file
    Copy file
```

### 3. Multi-Select (Multiple Choices)

Choose multiple options from a list:

```go
choices, err := MultiSelect("Select file types:", []string{
    ".txt",
    ".md",
    ".go",
})
```

**Visual Example:**
```bash
? Select file types: [Use arrows to move, space to select, type to filter]
  ◉ .txt
  ◯ .md
  ◉ .go
```

### 4. Text Input

Collect text input from user:

```go
input, err := Input("Enter filename:", "default.txt")
```

**Visual Example:**
```bash
? Enter filename: (default.txt) myfile.txt
```

### 5. Password Input

Collect sensitive input (hidden):

```go
password, err := Password("Enter password:")
```

**Visual Example:**
```bash
? Enter password: ********
```

## Using Interactive Prompts in Commands

### Example 1: Confirmation Before Destructive Operation

```go
var fileDeleteCmd = &cobra.Command{
    Use:   "delete [path]",
    Short: "Delete a file",
    Run: func(cmd *cobra.Command, args []string) {
        force, _ := cmd.Flags().GetBool("force")
        
        // Ask for confirmation if not forced
        if !force {
            confirmed, err := Confirm(
                fmt.Sprintf("Delete %s?", args[0]), 
                false,
            )
            if err != nil || !confirmed {
                fmt.Println("Operation cancelled")
                os.Exit(1)
            }
        }
        
        // Proceed with deletion
        // ...
    },
}
```

### Example 2: Interactive File Type Selection

```go
var filterCmd = &cobra.Command{
    Use:   "filter",
    Short: "Filter files by type",
    Run: func(cmd *cobra.Command, args []string) {
        fileTypes, err := MultiSelect(
            "Select file types to process:",
            []string{".txt", ".md", ".go", ".yaml", ".json"},
        )
        if err != nil {
            fmt.Println("Error:", err)
            os.Exit(1)
        }
        
        fmt.Printf("Processing files: %v\n", fileTypes)
        // Process selected file types...
    },
}
```

### Example 3: Interactive Setup Wizard

```go
var setupCmd = &cobra.Command{
    Use:   "setup",
    Short: "Interactive setup wizard",
    Run: func(cmd *cobra.Command, args []string) {
        // Step 1: Choose project type
        projectType, _ := Select(
            "Select project type:",
            []string{"Web Application", "CLI Tool", "Library"},
        )
        
        // Step 2: Enter project name
        projectName, _ := Input("Project name:", "my-project")
        
        // Step 3: Select features
        features, _ := MultiSelect(
            "Select features:",
            []string{"Logging", "Configuration", "Testing", "Documentation"},
        )
        
        // Step 4: Confirm
        confirmed, _ := Confirm(
            fmt.Sprintf("Create %s with features %v?", projectName, features),
            true,
        )
        
        if confirmed {
            fmt.Println("Creating project...")
            // Create project with selected options
        }
    },
}
```

## Demo Command

Scaphoid includes an interactive demo command to try all prompt types:

```bash
scaphoid interactive
```

This will walk you through:
- Confirmation prompt
- Single selection
- Multiple selection
- Text input

## Advanced Features

### Custom Validation

```go
import "github.com/AlecAivazis/survey/v2"

var input string
prompt := &survey.Input{
    Message: "Enter email:",
}

err := survey.AskOne(prompt, &input, survey.WithValidator(func(ans interface{}) error {
    if str, ok := ans.(string); !ok || !strings.Contains(str, "@") {
        return fmt.Errorf("invalid email address")
    }
    return nil
}))
```

### Conditional Prompts

```go
// Only ask for password if authentication is enabled
if authEnabled {
    password, _ := Password("Enter password:")
    // Use password...
}
```

### Transform Input

```go
import "strings"

var input string
prompt := &survey.Input{
    Message: "Enter filename:",
}

err := survey.AskOne(prompt, &input, survey.WithTransform(func(ans interface{}) interface{} {
    return strings.ToLower(ans.(string))
}))
```

## Integration with Flags

You can combine interactive prompts with command-line flags:

```go
var deleteCmd = &cobra.Command{
    Use: "delete [path]",
    Run: func(cmd *cobra.Command, args []string) {
        force, _ := cmd.Flags().GetBool("force")
        interactive, _ := cmd.Flags().GetBool("interactive")
        
        if !force && interactive {
            // Ask for confirmation
            confirmed, _ := Confirm("Delete?", false)
            if !confirmed {
                return
            }
        }
        
        // Proceed with deletion
    },
}
```

## Disabling Interactive Mode

For scripts and CI/CD, you can disable interactive prompts:

```bash
# Using force flag bypasses confirmation
scaphoid file delete myfile.txt --force

# Or use yes command
yes | scaphoid file delete myfile.txt
```

## Best Practices

1. **Always provide a default**: Make it easy for users to hit Enter
   ```go
   Input("Filename:", "default.txt")  // Good
   Input("Filename:", "")              // Less user-friendly
   ```

2. **Use confirmation for destructive operations**:
   - File/directory deletion
   - Data overwriting
   - Permanent changes

3. **Skip prompts with --force flag**:
   - Allow automation and scripting
   - Respect the force flag

4. **Provide clear messages**:
   ```go
   Confirm("Delete file.txt? This cannot be undone.", false)  // Good
   Confirm("Delete?", false)                                   // Less clear
   ```

5. **Handle errors gracefully**:
   ```go
   choice, err := Select("Choose:", options)
   if err != nil {
       fmt.Println("Operation cancelled")
       os.Exit(1)
   }
   ```

6. **Group related prompts**:
   ```go
   fmt.Println("## Configuration Setup")
   name, _ := Input("Name:", "")
   email, _ := Input("Email:", "")
   // ...
   ```

## Library Reference

Scaphoid uses [`github.com/AlecAivazis/survey/v2`](https://github.com/AlecAivazis/survey)

Features:
- ✅ Cross-platform (Windows, macOS, Linux)
- ✅ Keyboard navigation (arrows, vim keys)
- ✅ Type-to-filter in select prompts
- ✅ Custom validation
- ✅ Help text and descriptions
- ✅ Icons and styling
- ✅ Password masking
- ✅ Multi-select with space bar

## Alternative Libraries

Other popular Go interactive prompt libraries:

1. **promptui** (`github.com/manifoldco/promptui`)
   - Simpler API
   - Good for basic prompts

2. **go-prompt** (`github.com/c-bata/go-prompt`)
   - More complex, shell-like experience
   - Auto-completion support

3. **huh** (`github.com/charmbracelet/huh`)
   - Modern, beautiful TUIs
   - Part of Charm ecosystem

Survey was chosen for Scaphoid because of its:
- Rich feature set
- Clean API
- Wide adoption
- Active maintenance
- Python Click/Inquirer similarity

## Examples in the Wild

Real-world usage of interactive prompts:

```bash
# Git-like confirmation
scaphoid file delete important.txt
? Delete important.txt? This cannot be undone. (y/N) 

# Package manager-like multi-select
scaphoid backup create
? Select items to backup: [Use arrows to move, space to select]
  ◉ Documents/
  ◉ Pictures/
  ◯ Downloads/
  ◉ .config/

# CLI wizard
scaphoid project init
? Project name: (my-app) awesome-project
? Project type: 
  ▸ Web Application
    CLI Tool
    Library
? Enable features: [space to select]
  ◉ Logging
  ◉ Configuration  
  ◯ Database
  ◉ Testing
? Create project 'awesome-project'? (Y/n) y
- ✓ Successfully created project "awesome-project"
```

## Testing Interactive Commands

When testing, bypass interactive prompts:

```bash
# Using force flag
scaphoid file delete test.txt --force

# Using yes
yes | scaphoid file delete test.txt

# In Go tests, mock the survey functions
```

## Summary

Interactive prompts make CLI tools more user-friendly while maintaining scriptability through flags. Scaphoid integrates survey/v2 to provide:

- ✅ Confirmation dialogs for safety
- ✅ Selection menus for choices
- ✅ Multi-select for batch operations
- ✅ Text/password input for data collection
- ✅ Seamless integration with Cobra commands
- ✅ Scriptable with --force and piping
