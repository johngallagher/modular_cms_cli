package modular

import (
	"fmt"
	"io"
	"os"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	huh "github.com/charmbracelet/huh"
)

type FeaturesShowField struct {
	logger          io.WriteCloser
	focused         bool
	NavigationCtx   *NavigationContext
	Parent          *MainModel
	Block           BlockInterface
	FeaturesPointer *[]Feature
	Form            *huh.Form
}

func (f *FeaturesShowField) Blur() tea.Cmd {
	f.focused = false
	return nil
}

func (f *FeaturesShowField) Focus() tea.Cmd {
	f.focused = true
	return nil
}

func (f *FeaturesShowField) Error() error {
	return nil
}

func (f *FeaturesShowField) GetKey() string {
	return "Key"
}

func (f *FeaturesShowField) GetValue() any {
	return "Value"
}

func (f *FeaturesShowField) Run() error {
	return nil
}

func (f *FeaturesShowField) Skip() bool {
	return false
}
func (f *FeaturesShowField) activeStyles() *huh.FieldStyles {
	theme := huh.ThemeCharm()
	if f.focused {
		return &theme.Focused
	}
	return &theme.Blurred
}

func (f *FeaturesShowField) View() string {
	styles := f.activeStyles()
	return styles.Base.Render("Features")
}

func initLogger() io.WriteCloser {
	var f io.WriteCloser
	f, err := tea.LogToFile("debug.log", "debug")
	if err != nil {
		fmt.Println("fatal:", err)
		os.Exit(1)
	}
	return f
}

type FeaturesKeyMap struct {
}

func (f *FeaturesShowField) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case msg.String() == "shift+tab":
			return f, huh.PrevField
		case msg.String() == "enter":
			return f, huh.NextField
		case msg.String() == "tab":
			return f, huh.NextField
		case msg.String() == "e":
			return f, func() tea.Msg {
				f.NavigationCtx.Push("Features")
				f.Parent.ModelStack.Push(f.Parent.createFeatureListModel(f.Block, *f.FeaturesPointer))
				return f.Parent.ModelStack.Current().Init()
			}
		}
	}
	return f, nil
}

func (f *FeaturesShowField) WithAccessible(accessible bool) huh.Field {
	return f
}

func (f *FeaturesShowField) WithHeight(height int) huh.Field {
	return f
}

func (f *FeaturesShowField) WithKeyMap(keymap *huh.KeyMap) huh.Field {
	return f
}

func (f *FeaturesShowField) WithPosition(position huh.FieldPosition) huh.Field {
	return f
}

func (f *FeaturesShowField) WithWidth(width int) huh.Field {
	return f
}

func (f *FeaturesShowField) WithTheme(theme *huh.Theme) huh.Field {
	return f
}

func (f *FeaturesShowField) Zoom() bool {
	return false
}

func (f *FeaturesShowField) KeyBinds() []key.Binding {
	return []key.Binding{}
}

// Errors returns the current groups' errors.
func (f *FeaturesShowField) Errors() []error {
	return []error{}
}

// Help returns the current groups' help.
func (f *FeaturesShowField) Help() help.Model {
	return help.Model{}
}

// Get returns a result from the form.
func (f *FeaturesShowField) Get(key string) any {
	return ""
}

// GetString returns a result as a string from the form.
func (f *FeaturesShowField) GetString(key string) string {
	return ""
}

// GetInt returns a result as a int from the form.
func (f *FeaturesShowField) GetInt(key string) int {
	return 0
}

// GetBool returns a result as a string from the form.
func (f *FeaturesShowField) GetBool(key string) bool {
	return false
}

// NextGroup moves the form to the next group.
func (f *FeaturesShowField) NextGroup() tea.Cmd {
	return nil
}

// PrevGroup moves the form to the next group.
func (f *FeaturesShowField) PrevGroup() tea.Cmd {
	return nil
}

// NextField moves the form to the next field.
func (f *FeaturesShowField) NextField() tea.Cmd {
	return nil
}

// NextField moves the form to the next field.
func (f *FeaturesShowField) PrevField() tea.Cmd {
	return nil
}

// Init initializes the form.
func (f *FeaturesShowField) Init() tea.Cmd {
	return nil
}
