package modular

type SocialProofCardStatistics struct {
	Identifier  string     `yaml:"identifier"`
	Type        string     `yaml:"type"`
	HideFromNav bool       `yaml:"hide_from_nav"`
	Library     string     `yaml:"library"`
	Heading     string     `yaml:"heading"`
	Subheading  string     `yaml:"subheading"`
	Note        string     `yaml:"note"`
	Sentiment   string     `yaml:"sentiment"`
	Statistics  Statistics `yaml:"statistics"`
}

type Statistics struct {
	Left   Statistic `yaml:"left"`
	Center Statistic `yaml:"center"`
	Right  Statistic `yaml:"right"`
}

type Statistic struct {
	Value       string `yaml:"value"`
	Title       string `yaml:"title"`
	Description string `yaml:"description"`
	Note        string `yaml:"note"`
}

func (t *SocialProofCardStatistics) DisplayName() string {
	return "Social Proof Card Statistics"
}

func (t *SocialProofCardStatistics) FilterValue() string {
	return t.Heading
}

func (t *SocialProofCardStatistics) Category() string {
	return "Social Proof"
}

func (t *SocialProofCardStatistics) View() string {
	return "card_statistics"
}

func (t *SocialProofCardStatistics) Title() string {
	return "[" + t.Category() + "] [" + t.View() + "] " + t.Heading
}

func (t *SocialProofCardStatistics) TitlePointer() *string {
	return &t.Heading
}

func (t *SocialProofCardStatistics) Description() string {
	return t.Subheading
}

func (t *SocialProofCardStatistics) DescriptionPointer() *string {
	return &t.Subheading
}

func (t *SocialProofCardStatistics) GetFeatures() []Feature {
	return []Feature{}
}

func (t *SocialProofCardStatistics) SetFeatures(features []Feature) {
	// No-op since this block type doesn't use features
}
func (t *SocialProofCardStatistics) ID() string {
	return t.Identifier
}

func (t *SocialProofCardStatistics) GetFieldDefinitions() []*FieldDefinition {
	return []*FieldDefinition{
		{Key: "Heading", Title: "Heading", Type: FieldTypeInput, ValuePointer: &t.Heading},
		{Key: "Subheading", Title: "Subheading", Type: FieldTypeInput, ValuePointer: &t.Subheading},
		{Key: "Note", Title: "Note", Type: FieldTypeInput, ValuePointer: &t.Note},
		{Key: "Sentiment", Title: "Sentiment", Type: FieldTypeInput, ValuePointer: &t.Sentiment},
		{Key: "Statistics.Left.Value", Title: "Left Value", Type: FieldTypeInput, ValuePointer: &t.Statistics.Left.Value},
		{Key: "Statistics.Left.Title", Title: "Left Title", Type: FieldTypeInput, ValuePointer: &t.Statistics.Left.Title},
		{Key: "Statistics.Left.Description", Title: "Left Description", Type: FieldTypeInput, ValuePointer: &t.Statistics.Left.Description},
		{Key: "Statistics.Center.Value", Title: "Center Value", Type: FieldTypeInput, ValuePointer: &t.Statistics.Center.Value},
		{Key: "Statistics.Center.Title", Title: "Center Title", Type: FieldTypeInput, ValuePointer: &t.Statistics.Center.Title},
		{Key: "Statistics.Center.Description", Title: "Center Description", Type: FieldTypeInput, ValuePointer: &t.Statistics.Center.Description},
		{Key: "Statistics.Right.Value", Title: "Right Value", Type: FieldTypeInput, ValuePointer: &t.Statistics.Right.Value},
		{Key: "Statistics.Right.Title", Title: "Right Title", Type: FieldTypeInput, ValuePointer: &t.Statistics.Right.Title},
		{Key: "Statistics.Right.Description", Title: "Right Description", Type: FieldTypeInput, ValuePointer: &t.Statistics.Right.Description},
	}
}
