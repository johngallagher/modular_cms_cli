package modular

type FeatureSection struct {
	Type        string    `yaml:"type"`
	HideFromNav bool      `yaml:"hide_from_nav"`
	Heading     string    `yaml:"heading"`
	Subheading  string    `yaml:"subheading"`
	Features    []Feature `yaml:"features"`
	Library     string    `yaml:"library"`
	View        string    `yaml:"view"`
}

func (b FeatureSection) DisplayName() string {
	return "Feature"
}

func (b FeatureSection) FilterValue() string {
	return b.Heading
}

func (b FeatureSection) Title() string {
	return "[" + b.DisplayName() + "] [" + b.View + "] " + b.Heading
}

func (b *FeatureSection) TitlePointer() *string {
	return &b.Heading
}

func (b FeatureSection) Description() string {
	return b.Subheading
}

func (b *FeatureSection) DescriptionPointer() *string {
	return &b.Subheading
}

func (b FeatureSection) GetFeatures() []Feature {
	return b.Features
}

func (b *FeatureSection) SetFeatures(features []Feature) {
	b.Features = features
}

func (b FeatureSection) ID() string {
	return "feature_section_" + b.View
}

func (b *FeatureSection) GetFieldDefinitions() []*FieldDefinition {
	return []*FieldDefinition{
		{Key: "Heading", Title: "Heading", Type: FieldTypeInput, ValuePointer: &b.Heading},
		{Key: "Subheading", Title: "Subheading", Type: FieldTypeInput, ValuePointer: &b.Subheading},
		{Key: "Features", Title: "Features", Type: FieldTypeFeaturesShow, FeaturesPointer: &b.Features},
	}
}
