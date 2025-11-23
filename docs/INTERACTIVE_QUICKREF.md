# Quick Reference: Adding Interactive Prompts

## Installation

```bash
go get github.com/AlecAivazis/survey/v2
```

## Import

```go
import "github.com/AlecAivazis/survey/v2"
```

## Quick Examples

### Yes/No Confirmation

```go
var confirmed bool
prompt := &survey.Confirm{
    Message: "Continue?",
    Default: true,
}
survey.AskOne(prompt, &confirmed)
```

Or use the helper:
```go
confirmed, err := Confirm("Continue?", true)
```

### Single Selection

```go
var choice string
prompt := &survey.Select{
    Message: "Choose:",
    Options: []string{"Option 1", "Option 2", "Option 3"},
}
survey.AskOne(prompt, &choice)
```

Or use the helper:
```go
choice, err := Select("Choose:", []string{"Option 1", "Option 2"})
```

### Multiple Selection

```go
var choices []string
prompt := &survey.MultiSelect{
    Message: "Select all that apply:",
    Options: []string{"A", "B", "C"},
}
survey.AskOne(prompt, &choices)
```

Or use the helper:
```go
choices, err := MultiSelect("Select:", []string{"A", "B", "C"})
```

### Text Input

```go
var answer string
prompt := &survey.Input{
    Message: "Enter name:",
    Default: "default",
}
survey.AskOne(prompt, &answer)
```

Or use the helper:
```go
answer, err := Input("Enter name:", "default")
```

### Password

```go
var password string
prompt := &survey.Password{
    Message: "Password:",
}
survey.AskOne(prompt, &password)
```

Or use the helper:
```go
password, err := Password("Password:")
```

## Common Patterns

### Skip prompt with flag

```go
force, _ := cmd.Flags().GetBool("force")
if !force {
    confirmed, _ := Confirm("Continue?", false)
    if !confirmed {
        os.Exit(1)
    }
}
```

### Chain multiple prompts

```go
name, _ := Input("Project name:", "")
projectType, _ := Select("Type:", []string{"web", "cli"})
features, _ := MultiSelect("Features:", []string{"auth", "db"})
confirmed, _ := Confirm("Create project?", true)
if confirmed {
    // Create with name, projectType, features
}
```

### Validation

```go
prompt := &survey.Input{
    Message: "Email:",
}
survey.AskOne(prompt, &email, survey.WithValidator(survey.Required))
```

### Custom validation

```go
survey.AskOne(prompt, &answer, survey.WithValidator(func(ans interface{}) error {
    if str, ok := ans.(string); ok && len(str) < 3 {
        return fmt.Errorf("must be at least 3 characters")
    }
    return nil
}))
```

## Integration with Cobra

```go
var myCmd = &cobra.Command{
    Use: "my-command",
    Run: func(cmd *cobra.Command, args []string) {
        // Get flag
        interactive, _ := cmd.Flags().GetBool("interactive")
        
        if interactive {
            // Use prompts
            choice, _ := Select("Choose:", []string{"A", "B"})
            // ...
        } else {
            // Use args/flags only
        }
    },
}
```

## Helper Functions (Available in Scaphoid)

```go
// cmd_new/interactive.go provides these helpers:

Confirm(message string, defaultValue bool) (bool, error)
Select(message string, options []string) (string, error)
MultiSelect(message string, options []string) ([]string, error)
Input(message string, defaultValue string) (string, error)
Password(message string) (string, error)
```

## Styling

### Colors

```go
prompt := &survey.Select{
    Message: "Choose:",
    Options: []string{"Option 1", "Option 2"},
}
// Survey automatically handles terminal colors
```

### Help Text

```go
prompt := &survey.Input{
    Message: "Username:",
    Help:    "Enter your GitHub username",
}
```

### Custom Icons

```go
import "github.com/AlecAivazis/survey/v2/core"

core.SetIcon(survey.IconGood, "✓")
core.SetIcon(survey.IconBad, "✗")
```

## Error Handling

```go
answer, err := Input("Name:", "")
if err != nil {
    // User interrupted (Ctrl+C)
    fmt.Println("Operation cancelled")
    os.Exit(1)
}
```

## Testing

When writing tests, mock or bypass prompts:

```go
// In tests, use --force or yes
cmd := exec.Command("scaphoid", "delete", "file.txt", "--force")

// Or pipe yes
cmd := exec.Command("bash", "-c", "yes | scaphoid delete file.txt")
```

## Best Practices

1. ✅ Always provide defaults
2. ✅ Use --force to skip prompts
3. ✅ Handle Ctrl+C gracefully
4. ✅ Validate input
5. ✅ Provide clear messages
6. ✅ Group related prompts
7. ✅ Show confirmation summary before destructive ops

## Full Example

```go
var setupCmd = &cobra.Command{
    Use:   "setup",
    Short: "Interactive setup",
    Run: func(cmd *cobra.Command, args []string) {
        force, _ := cmd.Flags().GetBool("force")
        
        // Step 1: Get project name
        name, err := Input("Project name:", "my-project")
        if err != nil {
            fmt.Println("Setup cancelled")
            return
        }
        
        // Step 2: Select project type
        projectType, err := Select(
            "Project type:",
            []string{"Web App", "CLI Tool", "Library"},
        )
        if err != nil {
            fmt.Println("Setup cancelled")
            return
        }
        
        // Step 3: Select features
        features, err := MultiSelect(
            "Select features:",
            []string{"Logging", "Config", "Database", "Testing"},
        )
        if err != nil {
            fmt.Println("Setup cancelled")
            return
        }
        
        // Step 4: Confirm (unless forced)
        if !force {
            fmt.Printf("\nProject: %s\n", name)
            fmt.Printf("Type: %s\n", projectType)
            fmt.Printf("Features: %v\n", features)
            
            confirmed, err := Confirm("Create project?", true)
            if err != nil || !confirmed {
                fmt.Println("Setup cancelled")
                return
            }
        }
        
        // Create project
        fmt.Println("Creating project...")
        // Implementation here
    },
}

func init() {
    setupCmd.Flags().BoolP("force", "f", false, "Skip confirmation")
}
```

## Resources

- [Survey Documentation](https://github.com/AlecAivazis/survey)
- [Scaphoid Examples](docs/INTERACTIVE_PROMPTS.md)
- [Demo Script](demo_interactive.sh)
