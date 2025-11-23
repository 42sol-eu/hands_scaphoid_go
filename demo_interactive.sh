#!/bin/bash
# Interactive Scaphoid Demo Script
# This script demonstrates all interactive prompt features

cd "$(dirname "$0")"

echo "╔════════════════════════════════════════════════════╗"
echo "║   Scaphoid Interactive Prompts Demonstration       ║"
echo "╚════════════════════════════════════════════════════╝"
echo ""

# Demo 1: Interactive demo command
echo "=== Demo 1: Interactive Command ==="
echo "Running: scaphoid interactive"
echo ""
./bin/scaphoid interactive
echo ""
echo "Press Enter to continue..."
read

# Demo 2: Delete with confirmation
echo ""
echo "=== Demo 2: Delete with Confirmation ==="
echo "Creating test file..."
./bin/scaphoid file create /tmp/demo_delete.txt --force
echo ""
echo "Now trying to delete it (will ask for confirmation):"
echo "Running: scaphoid file delete /tmp/demo_delete.txt"
echo ""
./bin/scaphoid file delete /tmp/demo_delete.txt
echo ""
echo "Press Enter to continue..."
read

# Demo 3: Create with interactive type selection
echo ""
echo "=== Demo 3: Create with Type Selection ==="
echo "Running: scaphoid create /tmp/demo_interactive_file.txt"
echo "(You'll be asked to select the object type)"
echo ""
./bin/scaphoid create /tmp/demo_interactive_file.txt
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
echo "  ✓ Text input"
echo ""
echo "Integration Points:"
echo "  • file delete - asks confirmation unless --force"
echo "  • create - asks for type if omitted"
echo "  • interactive - demo command for all prompt types"
echo ""
echo "Library Used: github.com/AlecAivazis/survey/v2"
echo ""
echo "For more info: see docs/INTERACTIVE_PROMPTS.md"
