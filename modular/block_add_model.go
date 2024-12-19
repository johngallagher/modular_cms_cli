package modular

import (
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/thoas/go-funk"
)

type BlockAddModel struct {
	Blocks        []BlockInterface
	List          list.Model
	Parent        *MainModel
	BlockToAdd    BlockInterface
	CurrentBlocks []BlockInterface
}

func BlockAddModelFromMainModel(m *MainModel) *BlockAddModel {
	existingIds := funk.Map(m.LandingPage.Blocks, func(b BlockInterface) string {
		return b.ID()
	})
	items := []list.Item{}
	for _, block := range AllBlocks() {
		if !funk.Contains(existingIds, block.ID()) {
			items = append(items, block)
		}
	}

	l := list.New(items, list.NewDefaultDelegate(), 0, 0)
	l.SetShowStatusBar(true)
	l.SetFilteringEnabled(true)
	l.Title = "Add Block"
	l.SetSize(m.Width(), m.Height()-3)

	return &BlockAddModel{
		Blocks:        AllBlocks(),
		List:          l,
		Parent:        m,
		BlockToAdd:    AllBlocks()[0],
		CurrentBlocks: m.LandingPage.Blocks,
	}
}

func (m *BlockAddModel) Init() tea.Cmd {
	m.BlockToAdd = m.List.SelectedItem().(BlockInterface)
	m.Parent.LandingPage.Blocks = append(m.CurrentBlocks, m.BlockToAdd)
	m.Parent.LandingPage.Write()
	return nil
}

func (m *BlockAddModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if msg.String() == "esc" {
			m.Parent.LandingPage.Blocks = m.CurrentBlocks
			m.Parent.LandingPage.Write()
			m.Parent.NavigationCtx().Pop()
			m.Parent.ModelStack.Pop()
			return m.Parent.ModelStack.Current(), m.Parent.ModelStack.Current().Init()
		}

		if msg.String() == "enter" {
			m.Parent.NavigationCtx().Pop()
			m.Parent.ModelStack.Pop()
			return m.Parent.ModelStack.Current(), m.Parent.ModelStack.Current().Init()
		}
	}
	var cmd tea.Cmd
	m.List, cmd = m.List.Update(msg)
	m.BlockToAdd = m.List.SelectedItem().(BlockInterface)
	m.Parent.LandingPage.Blocks = append(m.CurrentBlocks, m.BlockToAdd)
	m.Parent.LandingPage.Write()

	return m, cmd
}

func (m *BlockAddModel) View() string {
	breadcrumb := m.Parent.NavigationCtx().Path[0]
	for i := 1; i < len(m.Parent.NavigationCtx().Path); i++ {
		breadcrumb += " > " + m.Parent.NavigationCtx().Path[i]
	}
	return breadcrumb + "\n\n" + m.List.View()
}
