package main

import (
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	huh "github.com/charmbracelet/huh"
	"github.com/charmbracelet/lipgloss"
)

type FeatureItem struct {
	feature *Feature
}

func (i FeatureItem) Title() string       { return i.feature.Name }
func (i FeatureItem) Description() string { return i.feature.Description }
func (i FeatureItem) FilterValue() string { return i.feature.Name }

func (m *MainModel) createFeatureListModel(block BlockInterface, features []Feature) *FeatureListModel {
	items := make([]list.Item, len(features))
	for i := range features {
		items[i] = FeatureItem{feature: &features[i]}
	}

	l := list.New(items, list.NewDefaultDelegate(), 20, 14)
	l.Title = "Features"
	l.SetShowStatusBar(true)
	l.SetFilteringEnabled(true)
	l.SetShowHelp(true)

	return &FeatureListModel{
		Block:         block,
		List:          l,
		NavigationCtx: m.NavigationCtx,
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
			if i, ok := m.List.SelectedItem().(FeatureItem); ok {
				m.List.ResetFilter()
				m.NavigationCtx.Push("Edit Feature")
				form := huh.NewForm(
					huh.NewGroup(
						huh.NewInput().
							Key("name").
							Title("Name").
							Value(&i.feature.Name),
						huh.NewText().
							Key("description").
							Title("Description").
							Value(&i.feature.Description),
					).WithShowHelp(false),
				)

				featureEditModel := &FeatureEditModel{
					Feature:       i.feature,
					Block:         m.Block,
					Form:          form,
					NavigationCtx: m.NavigationCtx,
					Parent:        m.Parent,
				}
				m.Parent.ModelStack.Push(featureEditModel)
				return featureEditModel, featureEditModel.Init()
			}
		}
		if msg.String() == "a" {
			newFeature := &Feature{
				Name:        "New Feature",
				Description: "New Feature Description",
			}

			*m.Features = append(*m.Features, *newFeature)
			m.Block.SetFeatures(*m.Features)

			featureItem := FeatureItem{feature: newFeature}
			m.List.InsertItem(len(m.List.Items()), featureItem)

			m.Parent.LandingPage.Write()
			form := huh.NewForm(
				huh.NewGroup(
					huh.NewInput().
						Key("name").
						Title("Name").
						Value(&newFeature.Name),
					huh.NewText().
						Key("description").
						Title("Description").
						Value(&newFeature.Description),
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
