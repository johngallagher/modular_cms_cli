package modular

import (
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

type MainModel struct {
	navigationCtx *NavigationContext
	ModelStack    *ModelStack
	LandingPage   *LandingPage
	width         int
	height        int
}

func (m *MainModel) Width() int {
	return m.width
}

func (m *MainModel) Height() int {
	return m.height
}

func (m *MainModel) NavigationCtx() *NavigationContext {
	return m.navigationCtx
}

// Model definition
type NavigationContext struct {
	Path []string
}

func (nc *NavigationContext) Breadcrumb() string {
	return strings.Join(nc.Path, " > ")
}

func (nc *NavigationContext) Push(level string) {
	nc.Path = append(nc.Path, level)
}

func (nc *NavigationContext) Pop() {
	if len(nc.Path) > 1 {
		nc.Path = nc.Path[:len(nc.Path)-1]
	}
}

// Factory
func InitialModel(filePath string) *MainModel {
	navCtx := &NavigationContext{Path: []string{"Home"}}
	landingPage := LandingPageFromMarkdownAtPath(filePath)

	m := &MainModel{
		navigationCtx: navCtx,
		ModelStack:    NewModelStack(),
		LandingPage:   landingPage,
	}

	// Initialize with block list view
	blockList := CreateBlockListModelFromMainModel(m)
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
