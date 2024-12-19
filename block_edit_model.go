package main

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/huh"
)

// Model definition
type BlockEditModel struct {
	Block         BlockInterface
	Form          *huh.Form
	NavigationCtx *NavigationContext
	Parent        *MainModel
	width         int
	height        int
}

// Factory
func (m *MainModel) createBlockEditModel(block BlockInterface) *BlockEditModel {
	fields := make([]huh.Field, 0)
	for _, def := range block.GetFieldDefinitions() {
		fields = append(fields, def.CreateFormField(block, m))
	}

	group := huh.NewGroup(fields...).WithShowHelp(true)
	form := huh.NewForm(group).WithWidth(m.width).WithHeight(m.height - 3)

	return &BlockEditModel{
		Block:         block,
		Form:          form,
		NavigationCtx: m.NavigationCtx,
		Parent:        m,
		width:         m.width,
		height:        m.height,
	}
}

// Update
func (m *BlockEditModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {

		case "esc":
			m.NavigationCtx.Pop()
			m.Parent.ModelStack.Pop()
			return m.Parent.ModelStack.Current(), nil
		}
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height

		return m, nil
	}

	var cmd tea.Cmd

	m.Parent.LandingPage.WriteToFile("index.md")
	model, cmd := m.Form.Update(msg)
	m.Form = model.(*huh.Form)

	if m.Form.State == huh.StateCompleted {
		// Return to block list after saving
		m.NavigationCtx.Pop()
		m.Parent.ModelStack.Pop()
		return m.Parent.ModelStack.Current(), nil
	}

	if m.Parent.ModelStack.Current() != m {
		return m.Parent.ModelStack.Current(), nil
	}

	return m, cmd
}

// View
func (m *BlockEditModel) View() string {
	breadcrumb := m.NavigationCtx.Path[0]
	for i := 1; i < len(m.NavigationCtx.Path); i++ {
		breadcrumb += " > " + m.NavigationCtx.Path[i]
	}
	return breadcrumb + "\n\n" + m.Form.View()
}

// Init
func (m *BlockEditModel) Init() tea.Cmd {
	return m.Form.Init()
}
