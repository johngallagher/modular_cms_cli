package modular

import (
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/huh"
)

// Model definition
type FeatureImportModel struct {
	Block    BlockInterface
	Form     *huh.Form
	Parent   *MainModel
	width    int
	height   int
	Features *[]Feature
	List     *list.Model
}

func CreateFeatureImportModelFromMainModel(m *MainModel, b BlockInterface, features *[]Feature, list *list.Model) *FeatureImportModel {
	input := ""
	textInput := huh.NewText().
		Key("text").
		Title("Enter features (one per line)").
		CharLimit(1000000).
		Value(&input).
		WithWidth(m.Width()).
		WithHeight(m.Height() - 3)

	form := huh.NewForm(
		huh.NewGroup(textInput).WithShowHelp(false),
	)

	return &FeatureImportModel{
		Block:    b,
		Form:     form,
		Parent:   m,
		width:    m.Width(),
		height:   m.Height(),
		Features: features,
		List:     list,
	}
}

func (m *FeatureImportModel) Init() tea.Cmd {
	return m.Form.Init()
}

func (m *FeatureImportModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	model, cmd := m.Form.Update(msg)
	m.Form = model.(*huh.Form)

	if m.Form.State == huh.StateCompleted {

		input := m.Form.GetString("text")
		features := ParseFeatures(input)
		*m.Features = append(*m.Features, features...)
		m.Block.SetFeatures(*m.Features)
		items := make([]list.Item, len(*m.Features))
		for i := range *m.Features {
			items[i] = FeatureListItem{feature: &(*m.Features)[i]}
		}
		m.List.SetItems(items)
		// Update the list's internal state after setting items
		m.List.ResetSelected()
		m.List.ResetFilter()
		m.Parent.LandingPage.Write()

		m.Parent.NavigationCtx().Pop()
		m.Parent.ModelStack.Pop()
		return m.Parent.ModelStack.Current(), nil
	}

	return m, cmd
}

func (m *FeatureImportModel) View() string {
	breadcrumb := m.Parent.NavigationCtx().Path[0]
	for i := 1; i < len(m.Parent.NavigationCtx().Path); i++ {
		breadcrumb += " > " + m.Parent.NavigationCtx().Path[i]
	}
	return breadcrumb + "\n\n" + m.Form.View()
}
