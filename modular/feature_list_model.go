package modular

import (
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	huh "github.com/charmbracelet/huh"
	"github.com/charmbracelet/lipgloss"
)

type FeatureListItem struct {
	feature *Feature
}

func (i FeatureListItem) Title() string       { return i.feature.Heading }
func (i FeatureListItem) Description() string { return i.feature.Summary }
func (i FeatureListItem) FilterValue() string { return i.feature.Heading }

func CreateFeatureListModelFromMainModel(m *MainModel, block BlockInterface, features []Feature) *FeatureListModel {
	items := make([]list.Item, len(features))
	for i := range features {
		items[i] = FeatureListItem{feature: &features[i]}
	}

	l := list.New(items, list.NewDefaultDelegate(), 20, 14)
	l.Title = "Features"
	l.SetShowStatusBar(true)
	l.SetFilteringEnabled(true)
	l.SetShowHelp(true)
	l.SetSize(m.Width(), m.Height()-3)

	return &FeatureListModel{
		Block:         block,
		List:          l,
		NavigationCtx: m.NavigationCtx(),
		Parent:        m,
		Features:      &features,
	}
}

type FeatureListModel struct {
	List          list.Model
	NavigationCtx *NavigationContext
	Parent        *MainModel
	Block         BlockInterface
	Features      *[]Feature
}

func (m *FeatureListModel) Init() tea.Cmd {
	return nil
}

func (m *FeatureListModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if m.List.FilterState() == list.Filtering {
			break
		}
		if msg.String() == "esc" {
			m.NavigationCtx.Pop()
			m.Parent.ModelStack.Pop()
			return m.Parent.ModelStack.Current(), nil
		}
		if msg.String() == "e" {
			if i, ok := m.List.SelectedItem().(FeatureListItem); ok {
				m.List.ResetFilter()
				m.NavigationCtx.Push("Edit Feature")
				form := huh.NewForm(
					huh.NewGroup(
						huh.NewInput().
							Key("heading").
							Title("Heading").
							Value(&i.feature.Heading),
						huh.NewText().
							Key("summary").
							Title("Summary").
							Value(&i.feature.Summary),
					).WithShowHelp(false).WithHeight(m.Parent.height - 3).WithWidth(m.Parent.width),
				)

				featureEditModel := &FeatureEditModel{
					Feature:       i.feature,
					Block:         m.Block,
					Form:          form,
					NavigationCtx: m.NavigationCtx,
					Parent:        m.Parent,
				}
				m.Parent.ModelStack.Push(featureEditModel)
				return m.Parent.ModelStack.Current(), featureEditModel.Init()
			}
		}
		if msg.String() == "a" {
			newFeature := &Feature{
				Heading: "",
				Summary: "",
			}

			*m.Features = append(*m.Features, *newFeature)
			m.Block.SetFeatures(*m.Features)

			featureItem := FeatureListItem{feature: newFeature}
			m.List.InsertItem(len(m.List.Items()), featureItem)

			m.Parent.LandingPage.Write()
			form := huh.NewForm(
				huh.NewGroup(
					huh.NewInput().
						Key("heading").
						Title("Heading").
						Value(&newFeature.Heading),
					huh.NewText().
						Key("summary").
						Title("Summary").
						Value(&newFeature.Summary),
				).WithShowHelp(false),
			)

			m.NavigationCtx.Push("New Feature")
			featureEditModel := &FeatureEditModel{
				Feature:       newFeature,
				Block:         m.Block,
				Form:          form,
				NavigationCtx: m.NavigationCtx,
				Parent:        m.Parent,
			}
			m.Parent.ModelStack.Push(featureEditModel)
			return featureEditModel, featureEditModel.Init()
		}
		if msg.String() == "x" {
			if i, ok := m.List.SelectedItem().(FeatureListItem); ok {
				selectedIndex := -1
				for j := range m.List.Items() {
					if m.List.Items()[j] == i {
						selectedIndex = j
						break
					}
				}

				if selectedIndex != -1 {
					*m.Features = append((*m.Features)[:selectedIndex], (*m.Features)[selectedIndex+1:]...)
					m.Block.SetFeatures(*m.Features)
					m.List.RemoveItem(selectedIndex)
					m.Parent.LandingPage.Write()
				}
			}
		}

	case tea.WindowSizeMsg:
		h, v := lipgloss.NewStyle().Margin(2, 2).GetFrameSize()
		m.List.SetSize(msg.Width-h, msg.Height-v)
	}

	var cmd tea.Cmd
	m.List, cmd = m.List.Update(msg)
	return m, cmd
}

func (m *FeatureListModel) View() string {
	breadcrumb := m.NavigationCtx.Path[0]
	for i := 1; i < len(m.NavigationCtx.Path); i++ {
		breadcrumb += " > " + m.NavigationCtx.Path[i]
	}
	return breadcrumb + "\n\n" + m.List.View()
}
