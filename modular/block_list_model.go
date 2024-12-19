package modular

import (
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// Model definition
type BlockListModel struct {
	List          list.Model
	NavigationCtx *NavigationContext
	Parent        *MainModel
	width         int
	height        int
}

// Factory
func CreateBlockListModelFromMainModel(m *MainModel) *BlockListModel {
	items := []list.Item{}
	for _, block := range m.LandingPage.Blocks {
		items = append(items, block)
	}

	// Set up list with default dimensions
	l := list.New(items, list.NewDefaultDelegate(), 0, 0)
	l.SetShowStatusBar(false)
	l.SetFilteringEnabled(false)
	l.Title = "Landing Page Blocks"
	l.SetSize(m.Width(), m.Height()-3)

	// Create model with stored dimensions
	model := &BlockListModel{
		List:          l,
		NavigationCtx: m.NavigationCtx(),
		Parent:        m,
	}

	return model
}

// Update
func (m *BlockListModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if msg.String() == "esc" {
			m.Parent.NavigationCtx().Pop()
			m.Parent.ModelStack.Pop()
			return m.Parent.ModelStack.Current(), nil
		}

		if msg.String() == "e" {
			if i, ok := m.List.SelectedItem().(BlockInterface); ok {
				var selectedBlock BlockInterface
				for j := range m.Parent.LandingPage.Blocks {
					if m.Parent.LandingPage.Blocks[j].ID() == i.ID() {
						selectedBlock = m.Parent.LandingPage.Blocks[j]
						break
					}
				}

				if selectedBlock != nil {
					m.Parent.NavigationCtx().Push("Edit Block")
					m.Parent.ModelStack.Push(createBlockEditModelFromMainModel(m.Parent, selectedBlock))
					return m.Parent.ModelStack.Current(), m.Parent.ModelStack.Current().Init()
				}
			}
		}

		if msg.String() == "a" {
			m.NavigationCtx.Push("Add Block")
			m.Parent.ModelStack.Push(BlockAddModelFromMainModel(m.Parent))
			return m.Parent.ModelStack.Current(), m.Parent.ModelStack.Current().Init()
		}
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
		h, v := lipgloss.NewStyle().Margin(2, 2).GetFrameSize()
		m.List.SetSize(msg.Width-h, msg.Height-v)
	}

	var cmd tea.Cmd
	m.List, cmd = m.List.Update(msg)
	return m, cmd
}

// View
func (m *BlockListModel) View() string {
	breadcrumb := m.NavigationCtx.Path[0]
	for i := 1; i < len(m.NavigationCtx.Path); i++ {
		breadcrumb += " > " + m.NavigationCtx.Path[i]
	}
	return breadcrumb + "\n\n" + m.List.View()
}

// Init
func (m *BlockListModel) Init() tea.Cmd {
	return nil
}
