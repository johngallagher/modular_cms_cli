package modular

type FeatureSectionsAlternate struct {
	Type        string     `yaml:"type"`
	HideFromNav bool       `yaml:"hide_from_nav"`
	Library     string     `yaml:"library"`
	Sentiment   string     `yaml:"sentiment"`
	Left        BulletList `yaml:"left"`
	Right       BulletList `yaml:"right"`
}

type BulletList struct {
	Heading    string `yaml:"heading"`
	Subheading string `yaml:"subheading"`
	Footer     string `yaml:"footer"`
	Feature1   string `yaml:"feature_1"`
	Feature2   string `yaml:"feature_2"`
	Feature3   string `yaml:"feature_3"`
	Feature4   string `yaml:"feature_4"`
	Feature5   string `yaml:"feature_5"`
	Feature6   string `yaml:"feature_6"`
	Sentiment  string `yaml:"sentiment"`
}

func (t *FeatureSectionsAlternate) ID() string {
	return "feature_sections_alternate"
}

func (f *FeatureSectionsAlternate) DisplayName() string {
	return "Feature Sections Alternate"
}
func (f FeatureSectionsAlternate) FilterValue() string {
	return f.Left.Heading + " | " + f.Right.Heading
}

func (f *FeatureSectionsAlternate) Category() string {
	return "Feature"
}

func (f FeatureSectionsAlternate) View() string {
	return "alternate"
}

func (f FeatureSectionsAlternate) Title() string {
	return "[" + f.Category() + "] [" + f.View() + "] " + f.Left.Heading + " | " + f.Right.Heading
}

func (f *FeatureSectionsAlternate) TitlePointer() *string {
	return &f.Left.Heading
}

func (f FeatureSectionsAlternate) Description() string {
	return ""
}

func (f *FeatureSectionsAlternate) DescriptionPointer() *string {
	return nil
}

func (f FeatureSectionsAlternate) GetFeatures() []Feature {
	return []Feature{}
}

func (f *FeatureSectionsAlternate) SetFeatures(features []Feature) {
}

func (t *FeatureSectionsAlternate) GetFieldDefinitions() []*FieldDefinition {
	return []*FieldDefinition{
		{Key: "Left.Heading", Title: "Left Heading", Type: FieldTypeInput, ValuePointer: &t.Left.Heading},
		{Key: "Left.Subheading", Title: "Left Subheading", Type: FieldTypeInput, ValuePointer: &t.Left.Subheading},
		{Key: "Left.Sentiment", Title: "Left Sentiment", Type: FieldTypeInput, ValuePointer: &t.Left.Sentiment},
		{Key: "Left.Feature1", Title: "Left Feature 1", Type: FieldTypeInput, ValuePointer: &t.Left.Feature1},
		{Key: "Left.Feature2", Title: "Left Feature 2", Type: FieldTypeInput, ValuePointer: &t.Left.Feature2},
		{Key: "Left.Feature3", Title: "Left Feature 3", Type: FieldTypeInput, ValuePointer: &t.Left.Feature3},
		{Key: "Left.Feature4", Title: "Left Feature 4", Type: FieldTypeInput, ValuePointer: &t.Left.Feature4},
		{Key: "Left.Feature5", Title: "Left Feature 5", Type: FieldTypeInput, ValuePointer: &t.Left.Feature5},
		{Key: "Left.Feature6", Title: "Left Feature 6", Type: FieldTypeInput, ValuePointer: &t.Left.Feature6},
		{Key: "Right.Heading", Title: "Right Heading", Type: FieldTypeInput, ValuePointer: &t.Right.Heading},
		{Key: "Right.Subheading", Title: "Right Subheading", Type: FieldTypeInput, ValuePointer: &t.Right.Subheading},
		{Key: "Right.Sentiment", Title: "Right Sentiment", Type: FieldTypeInput, ValuePointer: &t.Right.Sentiment},
		{Key: "Right.Feature1", Title: "Right Feature 1", Type: FieldTypeInput, ValuePointer: &t.Right.Feature1},
		{Key: "Right.Feature2", Title: "Right Feature 2", Type: FieldTypeInput, ValuePointer: &t.Right.Feature2},
		{Key: "Right.Feature3", Title: "Right Feature 3", Type: FieldTypeInput, ValuePointer: &t.Right.Feature3},
		{Key: "Right.Feature4", Title: "Right Feature 4", Type: FieldTypeInput, ValuePointer: &t.Right.Feature4},
		{Key: "Right.Feature5", Title: "Right Feature 5", Type: FieldTypeInput, ValuePointer: &t.Right.Feature5},
		{Key: "Right.Feature6", Title: "Right Feature 6", Type: FieldTypeInput, ValuePointer: &t.Right.Feature6},
	}
}
