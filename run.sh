#!/bin/sh
# Launcher script for VioletNotes
# Ensures correct locale settings for Fyne on all Linux environments

# Force English locale for Fyne to avoid "Error parsing user locale C"
export FYNE_LANG=en-US
export LANG=en_US.UTF-8
export LC_ALL=en_US.UTF-8

# Get the directory of this script
DIR="$(cd "$(dirname "$0")" && pwd)"

# Run the binary
exec "$DIR/notes-app" "$@"
