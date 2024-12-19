package main

import (
	tea "github.com/charmbracelet/bubbletea"
	huh "github.com/charmbracelet/huh"
)

// FeatureEditModel for individual feature editing
type FeatureEditModel struct {
	Feature       *Feature
	Form          *huh.Form
	NavigationCtx *NavigationContext
	Parent        *MainModel
	Block         BlockInterface
}

func (m *FeatureEditModel) Init() tea.Cmd {
	// Initialize the form
	return m.Form.Init()
}

func (m *FeatureEditModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "esc":
			m.NavigationCtx.Pop()
			m.Parent.ModelStack.Pop()
			return m.Parent.ModelStack.Current(), nil
		}
	}

	var cmd tea.Cmd
	m.Parent.LandingPage.Write()
	model, cmd := m.Form.Update(msg)
	m.Form = model.(*huh.Form)

	// If the form is completed
	if m.Form.State == huh.StateCompleted {
		m.NavigationCtx.Pop()
		m.Parent.ModelStack.Pop()
		return m.Parent.ModelStack.Current(), nil
	}

	return m, cmd
}

func (m *FeatureEditModel) View() string {
	breadcrumb := m.NavigationCtx.Path[0]
	for i := 1; i < len(m.NavigationCtx.Path); i++ {
		breadcrumb += " > " + m.NavigationCtx.Path[i]
	}
	return breadcrumb + "\n\n" + m.Form.View()
}
