# PC Load Letter - Git Repository Cloner

[PC Load Letter](https://en.wikipedia.org/wiki/PC_LOAD_LETTER)

A terminal user interface (TUI) application built with [Bubble Tea](https://github.com/charmbracelet/bubbletea) that allows users to select and clone multiple Git repositories.

## Features

- Interactive TUI with keyboard navigation
- Multi-select repository cloning
- Help menu with keybindings
- Stylish interface using [lipgloss](https://github.com/charmbracelet/lipgloss)

## Key Bindings

- `↑/k` - Move cursor up
- `↓/j` - Move cursor down
- `Space` - Select/deselect repository
- `Enter` - Clone selected repositories
- `?` - Toggle help menu
- `q/esc/ctrl+c` - Quit application

## Repository Structure

Repositories are defined with the following properties:
- Name
- URL
- Description
- Selection status

## Installation

1. Ensure you have Go installed
2. Clone this repository
3. Run:

```bash
go mod tidy
go build
```

## Usage

The application will clone selected repositories to:

```
~/repos/personal/clone/<repository-name>
```

## Code Structure

### Main Components

1. `keyMap` - Defines all keyboard shortcuts
2. `repository` - Structure for repository information
3. `model` - Main application state
4. `initialModel()` - Sets up initial application state
5. `Update()` - Handles all state updates
6. `View()` - Renders the TUI

### Styling

The interface uses lipgloss for styling with:
- Title in pink (#FF69B4)
- Centered alignment for title
- Left margin for repository list
- Cursor and checkbox indicators

## Dependencies

- github.com/charmbracelet/bubbles
- github.com/charmbracelet/bubbletea
- github.com/charmbracelet/lipgloss
- github.com/go-git/go-git/v5

## Example Output

```
PC Load Letter

  [ ] Dot_Files - All the dot files (kind of)
> [X] Tmux - Dots for tmux setup
  [ ] zsh - Dots for zsh setup

? help • q quit
```
