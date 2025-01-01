package modular

type StyledQuiz struct {
	Type        string         `yaml:"type"`
	HideFromNav bool           `yaml:"hide_from_nav"`
	Library     string         `yaml:"library"`
	Heading     string         `yaml:"heading"`
	Subheading  string         `yaml:"subheading"`
	Questions   []QuizQuestion `yaml:"questions"`
}

type QuizQuestion struct {
	ID          string   `yaml:"id"`
	Title       string   `yaml:"title"`
	Description string   `yaml:"description"`
	Type        string   `yaml:"type"`
	Fields      []Field  `yaml:"fields,omitempty"`
	Options     []Option `yaml:"options,omitempty"`
}

type Field struct {
	Type        string `yaml:"type"`
	Label       string `yaml:"label"`
	ID          string `yaml:"id"`
	Placeholder string `yaml:"placeholder,omitempty"`
}

type Option struct {
	ID    string `yaml:"id"`
	Label string `yaml:"label"`
}

func (q *StyledQuiz) DisplayName() string {
	return "Styled Quiz"
}

func (q *StyledQuiz) FilterValue() string {
	return q.Heading
}

func (q *StyledQuiz) Category() string {
	return "Quiz"
}

func (q *StyledQuiz) View() string {
	return "styled"
}

func (q *StyledQuiz) Title() string {
	return "[" + q.Category() + "] [" + q.View() + "] " + q.Heading
}

func (q *StyledQuiz) TitlePointer() *string {
	return &q.Heading
}

func (q *StyledQuiz) Description() string {
	return q.Subheading
}

func (q *StyledQuiz) DescriptionPointer() *string {
	return &q.Subheading
}

func (q *StyledQuiz) GetFeatures() []Feature {
	return []Feature{}
}

func (q *StyledQuiz) SetFeatures(features []Feature) {
	// No-op since this block type doesn't use features
}

func (q *StyledQuiz) ID() string {
	return "styled_quiz"
}

func (q *StyledQuiz) GetFieldDefinitions() []*FieldDefinition {
	return []*FieldDefinition{
		{Key: "Heading", Title: "Heading", Type: FieldTypeInput, ValuePointer: &q.Heading},
		{Key: "Subheading", Title: "Subheading", Type: FieldTypeInput, ValuePointer: &q.Subheading},
	}
}
