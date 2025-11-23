#!/bin/bash
# Interactive Scaphoid Demo Script
# This script demonstrates all interactive prompt features

cd "$(dirname "$0")"

echo "╔════════════════════════════════════════════════════╗"
echo "║   Scaphoid Interactive Prompts Demonstration       ║"
echo "╚════════════════════════════════════════════════════╝"
echo ""
echo "This demo will guide you through Scaphoid's interactive features."
echo "Please interact with the prompts as they appear."
echo ""
echo "Press Enter to begin..."
read

# Demo 1: Interactive demo command
echo ""
echo "=== Demo 1: Interactive Prompt Types ==="
echo ""
echo "This command demonstrates all available prompt types:"
echo "  • Confirmation (Yes/No)"
echo "  • Selection (Choose one option)"
echo "  • Multi-select (Choose multiple options)"
echo "  • Text input (Enter text)"
echo ""
echo "Running: ./bin/scaphoid interactive"
echo ""
./bin/scaphoid interactive
echo ""
echo "Press Enter to continue..."
read

# Demo 2: Delete with confirmation
echo ""
echo "=== Demo 2: Delete with Confirmation ==="
echo ""
echo "The 'file delete' command now asks for confirmation by default."
echo "You can skip confirmation with the --force flag."
echo ""
echo "Creating test file..."
./bin/scaphoid file create /tmp/demo_delete.txt --force
echo ""
echo "Now trying to delete it (will ask for confirmation):"
echo "Running: ./bin/scaphoid file delete /tmp/demo_delete.txt"
echo ""
./bin/scaphoid file delete /tmp/demo_delete.txt
echo ""
echo "Press Enter to continue..."
read

# Demo 3: Create with interactive type selection
echo ""
echo "=== Demo 3: Create with Type Selection ==="
echo ""
echo "The 'create' command can now prompt for the object type if omitted."
echo "This makes it easier to use interactively."
echo ""
echo "Running: ./bin/scaphoid create /tmp/demo_interactive_file.txt"
echo "(You'll be asked to select the object type)"
echo ""
./bin/scaphoid create /tmp/demo_interactive_file.txt
echo ""
echo "Cleaning up..."
./bin/scaphoid file delete /tmp/demo_interactive_file.txt --force 2>/dev/null
echo ""
echo "Press Enter to continue..."
read

# Summary
echo ""
echo "╔════════════════════════════════════════════════════╗"
echo "║              Demo Complete!                        ║"
echo "╚════════════════════════════════════════════════════╝"
echo ""
echo "Key Features Demonstrated:"
echo "  ✓ Confirmation prompts (Yes/No)"
echo "  ✓ Selection menus (Choose one)"
echo "  ✓ Multi-select (Choose multiple)"
echo "  ✓ Text input with defaults"
echo ""
echo "Integration Points:"
echo "  • file delete - asks confirmation unless --force"
echo "  • create - asks for type if omitted"
echo "  • interactive - demo command showing all prompt types"
echo ""
echo "Library Used: github.com/AlecAivazis/survey/v2"
echo ""
echo "Try these commands yourself:"
echo "  ./bin/scaphoid interactive          # Full demo"
echo "  ./bin/scaphoid file delete <path>   # Interactive delete"
echo "  ./bin/scaphoid create <path>        # Interactive create"
echo ""
echo "For more info: see docs/INTERACTIVE_PROMPTS.md"
