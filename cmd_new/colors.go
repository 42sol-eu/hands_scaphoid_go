package cmd_new

import (
	"fmt"
	"os"

	"github.com/fatih/color"
	"github.com/spf13/viper"
)

var (
	// Color spray functions - initialized from config
	actionColor  func(a ...interface{}) string
	typeColor    func(a ...interface{}) string
	errorColor   func(a ...interface{}) string
	successColor func(a ...interface{}) string
	pathColor    func(a ...interface{}) string
	keywordColor func(a ...interface{}) string
)

// ColorConfig holds the color configuration
type ColorConfig struct {
	Color string `mapstructure:"color"`
	Bold  bool   `mapstructure:"bold"`
}

func init() {
	loadColorConfig()
}

// loadColorConfig loads color configuration from YAML file
func loadColorConfig() {
	// Set default config file locations
	viper.SetConfigName(".scaphoid")
	viper.SetConfigType("yaml")
	
	// Search for config in: current directory, home directory, /etc/scaphoid/
	viper.AddConfigPath(".")
	if home, err := os.UserHomeDir(); err == nil {
		viper.AddConfigPath(home)
	}
	viper.AddConfigPath("/etc/scaphoid/")
	
	// Set defaults
	setDefaultColors()
	
	// Read config file (ignore if not found)
	if err := viper.ReadInConfig(); err != nil {
		// Use defaults if config file not found
		initializeColors()
		return
	}
	
	initializeColors()
}

// setDefaultColors sets default color configuration
func setDefaultColors() {
	viper.SetDefault("colors.action.color", "blue")
	viper.SetDefault("colors.action.bold", true)
	viper.SetDefault("colors.type.color", "green")
	viper.SetDefault("colors.type.bold", true)
	viper.SetDefault("colors.error.color", "red")
	viper.SetDefault("colors.error.bold", true)
	viper.SetDefault("colors.success.color", "green")
	viper.SetDefault("colors.success.bold", true)
	viper.SetDefault("colors.path.color", "green")
	viper.SetDefault("colors.path.bold", true)
	viper.SetDefault("colors.keyword.color", "green")
	viper.SetDefault("colors.keyword.bold", true)
}

// initializeColors creates color functions from configuration
func initializeColors() {
	actionColor = createColorFunc("colors.action")
	typeColor = createColorFunc("colors.type")
	errorColor = createColorFunc("colors.error")
	successColor = createColorFunc("colors.success")
	pathColor = createColorFunc("colors.path")
	keywordColor = createColorFunc("colors.keyword")
}

// createColorFunc creates a color function based on config
func createColorFunc(configKey string) func(a ...interface{}) string {
	colorName := viper.GetString(configKey + ".color")
	bold := viper.GetBool(configKey + ".bold")
	
	var attrs []color.Attribute
	
	// Map color names to color attributes
	switch colorName {
	case "black":
		attrs = append(attrs, color.FgBlack)
	case "red":
		attrs = append(attrs, color.FgRed)
	case "green":
		attrs = append(attrs, color.FgGreen)
	case "yellow":
		attrs = append(attrs, color.FgYellow)
	case "blue":
		attrs = append(attrs, color.FgBlue)
	case "magenta":
		attrs = append(attrs, color.FgMagenta)
	case "cyan":
		attrs = append(attrs, color.FgCyan)
	case "white":
		attrs = append(attrs, color.FgWhite)
	default:
		attrs = append(attrs, color.FgWhite)
	}
	
	if bold {
		attrs = append(attrs, color.Bold)
	}
	
	return color.New(attrs...).SprintFunc()
}

// colorizePath colorizes file paths
func colorizePath(path string) string {
	return pathColor(fmt.Sprintf(`"%s"`, path))
}

// colorizeType colorizes filesystem object types
func colorizeType(objType string) string {
	return typeColor(objType)
}

// colorizeAction colorizes action words (verbs)
func colorizeAction(action string) string {
	return actionColor(action)
}

// colorizeError colorizes error messages
func colorizeError(msg string) string {
	return errorColor(msg)
}

// colorizeSuccess colorizes success messages
func colorizeSuccess(msg string) string {
	return successColor(msg)
}

// colorizeKeyword colorizes keywords in output
func colorizeKeyword(keyword string) string {
	return keywordColor(keyword)
}

// colorizeNotExist colorizes "not exist" messages
func colorizeNotExist(msg string) string {
	return errorColor(msg)
}

// formatExistsMessage formats existence messages with colors
func formatExistsMessage(path string, exists bool, objType string) string {
	coloredPath := colorizePath(path)
	if exists {
		if objType != "" {
			return fmt.Sprintf("%s %s (type: %s)", coloredPath, colorizeAction("exists"), colorizeType(objType))
		}
		return fmt.Sprintf("%s %s", coloredPath, colorizeAction("exists"))
	}
	return fmt.Sprintf("%s does %s %s", coloredPath, colorizeNotExist("not"), colorizeAction("exist"))
}

// formatTypeCheckMessage formats type-specific existence messages with colors
func formatTypeCheckMessage(path string, exists bool, expectedType string, actualType string) string {
	coloredPath := colorizePath(path)
	if exists && actualType == expectedType {
		return fmt.Sprintf("%s %s and is a %s", coloredPath, colorizeAction("exists"), colorizeType(expectedType))
	} else if exists {
		return fmt.Sprintf("%s %s but is %s a %s (it's a %s)", 
			coloredPath,
			colorizeAction("exists"),
			colorizeNotExist("not"), 
			colorizeType(expectedType),
			colorizeType(actualType))
	}
	return fmt.Sprintf("%s does %s %s", coloredPath, colorizeNotExist("not"), colorizeAction("exist"))
}

// formatSuccessMessage formats success messages with colored action
func formatSuccessMessage(action string, path string) string {
	return fmt.Sprintf("Successfully %s %s", colorizeAction(action), colorizePath(path))
}

// formatCreateMessage formats create messages with colored action and path
func formatCreateMessage(objType string, path string) string {
	return fmt.Sprintf("Successfully %s %s at %s", colorizeAction("created"), colorizeType(objType), colorizePath(path))
}

// formatCopyMoveMessage formats copy/move messages with colored paths
func formatCopyMoveMessage(action string, source string, dest string) string {
	return fmt.Sprintf("Successfully %s %s to %s", colorizeAction(action), colorizePath(source), colorizePath(dest))
}
