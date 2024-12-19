package modular

import (
	tea "github.com/charmbracelet/bubbletea"
)

// Feature represents a detailed item within a block
type Feature struct {
	Name        string
	Description string
}

// MainModel is the root model managing app state
type MainModel struct {
	NavigationCtx *NavigationContext
	ModelStack    *ModelStack
	LandingPage   *LandingPage
	width         int
	height        int
}

// Factory
func InitialModel(filePath string) *MainModel {
	navCtx := &NavigationContext{Path: []string{"Home"}}
	landingPage := LandingPageFromMarkdownAtPath(filePath)

	m := &MainModel{
		NavigationCtx: navCtx,
		ModelStack:    NewModelStack(),
		LandingPage:   landingPage,
	}

	// Initialize with block list view
	blockList := m.createBlockListModel()
	m.ModelStack.Push(blockList)
	return m
}

// Init
func (m *MainModel) Init() tea.Cmd {
	return nil
}

// Update
func (m *MainModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
	}

	if m.ModelStack.IsEmpty() {
		return m, tea.Quit
	}

	current := m.ModelStack.Current()
	updatedModel, cmd := current.Update(msg)

	// If the updated model is different from the current one,
	// it means we need to switch models
	if updatedModel != current {
		if updatedModel == nil {
			// Pop the current model if nil is returned
			m.ModelStack.Pop()
		} else {
			// Push the new model
			m.ModelStack.Push(updatedModel)
		}
	}

	return m.ModelStack.Current(), cmd
}

// View
func (m *MainModel) View() string {
	if m.ModelStack.IsEmpty() {
		return "No active views"
	}
	return m.ModelStack.Current().View()
}
