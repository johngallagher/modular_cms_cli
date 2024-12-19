package modular

import tea "github.com/charmbracelet/bubbletea"

// ModelStack manages a stack of tea.Models
type ModelStack struct {
	models []tea.Model
}

// NewModelStack creates a new empty model stack
func NewModelStack() *ModelStack {
	return &ModelStack{
		models: make([]tea.Model, 0),
	}
}

// Push adds a model to the top of the stack
func (s *ModelStack) Push(m tea.Model) {
	s.models = append(s.models, m)
}

// Pop removes and returns the top model from the stack
func (s *ModelStack) Pop() tea.Model {
	if len(s.models) == 0 {
		return nil
	}
	if len(s.models) == 1 {
		return s.models[0]
	}

	lastIdx := len(s.models) - 1
	model := s.models[lastIdx]
	s.models = s.models[:lastIdx]
	return model
}

// Current returns the current (top) model without removing it
func (s *ModelStack) Current() tea.Model {
	if len(s.models) == 0 {
		return nil
	}
	return s.models[len(s.models)-1]
}

// IsEmpty returns true if the stack has no models
func (s *ModelStack) IsEmpty() bool {
	return len(s.models) == 0
}
