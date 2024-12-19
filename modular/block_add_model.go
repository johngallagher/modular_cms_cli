package modular

import (
	"slices"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
)

type BlockAddModel struct {
	Blocks []BlockInterface
	List   list.Model
	Parent *MainModel
}

func BlockAddModelFromMainModel(m *MainModel) *BlockAddModel {
	alreadyUsedBlocks := m.LandingPage.Blocks
	items := []list.Item{}
	for _, block := range AllBlocks() {
		if slices.Contains(alreadyUsedBlocks, block) {
			continue
		}
		items = append(items, block)
	}

	l := list.New(items, list.NewDefaultDelegate(), 0, 0)
	l.SetShowStatusBar(true)
	l.SetFilteringEnabled(true)
	l.Title = "Add Block"
	l.SetSize(m.Width(), m.Height()-3)

	return &BlockAddModel{
		Blocks: AllBlocks(),
		List:   l,
		Parent: m,
	}
}

func (m *BlockAddModel) Init() tea.Cmd {
	return nil
}

func (m *BlockAddModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	_, cmd := m.List.Update(msg)
	return m, cmd
}

func (m *BlockAddModel) View() string {
	breadcrumb := m.Parent.NavigationCtx().Path[0]
	for i := 1; i < len(m.Parent.NavigationCtx().Path); i++ {
		breadcrumb += " > " + m.Parent.NavigationCtx().Path[i]
	}
	return breadcrumb + "\n\n" + m.List.View()
}
