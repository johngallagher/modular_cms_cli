package modular

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
func createBlockEditModelFromMainModel(m *MainModel, b BlockInterface) *BlockEditModel {
	fields := make([]huh.Field, 0)
	for _, def := range b.GetFieldDefinitions() {
		fields = append(fields, def.CreateFormField(b, m))
	}

	group := huh.NewGroup(fields...).WithShowHelp(true)
	form := huh.NewForm(group).WithWidth(m.Width()).WithHeight(m.Height() - 3)

	return &BlockEditModel{
		Block:         b,
		Form:          form,
		NavigationCtx: m.NavigationCtx(),
		Parent:        m,
		width:         m.Width(),
		height:        m.Height(),
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

	m.Parent.LandingPage.Write()
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
