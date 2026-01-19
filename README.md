# VioletNotes

A modern, lightweight, and beautiful note-taking application built with [Go](https://go.dev/) and the [Fyne](https://fyne.io/) GUI toolkit.

Designed for efficiency and aesthetics, VioletNotes features a custom dark purple theme and is optimized for Linux environments, including tiling window managers like BSPWM.

## Features

*   **Modern UI:** Sleek interface with a custom purple theme (`VioletTheme`) and deep gradient background.
*   **Fast & Native:** Compiled to native machine code for maximum performance and low resource usage.
*   **Auto-Persistence:** Notes are automatically saved to `~/.notes_app/notes.json` in real-time.
*   **Smart Shortcuts:**
    *   `Ctrl+S`: Instantly save the current note from anywhere in the editor.
*   **Dual-Pane Layout:**
    *   **Sidebar:** Scrollable list of your notes.
    *   **Editor:** Distraction-free title and markdown-friendly content area.

## Prerequisites

*   **Go:** Version 1.20 or higher.
*   **System Dependencies (Debian/Kali/Ubuntu):**
    ```bash
    sudo apt-get install libgl1-mesa-dev xorg-dev libxxf86vm-dev
    ```

## Installation

1.  Clone the repository:
    ```bash
    git clone https://github.com/yourusername/violet-notes.git
    cd violet-notes
    ```

2.  Install dependencies:
    ```bash
    go mod tidy
    ```

3.  Build the application:
    ```bash
    go build -o notes-app .
    ```

4.  Run it:
    ```bash
    ./run.sh
    ```
    *Note: We recommend using `./run.sh` to ensure compatibility with all Linux system locales.*

## Project Structure

*   `main.go`: Application entry point and Clean Architecture orchestration.
*   `components.go`: Custom UI widgets and event handling.
*   `theme.go`: `VioletTheme` definition and color palette.
*   `storage.go`: JSON persistence layer.
*   `note.go`: Domain entities.
*   `run.sh`: Production-ready launcher script.

## License

This project is licensed under the MIT License.