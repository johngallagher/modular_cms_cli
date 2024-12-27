package modular

type FeatureBlock struct {
	Type        string    `yaml:"type"`
	HideFromNav bool      `yaml:"hide_from_nav"`
	Heading     string    `yaml:"heading"`
	Subheading  string    `yaml:"subheading"`
	Features    []Feature `yaml:"features"`
	Library     string    `yaml:"library"`
	View        string    `yaml:"view"`
}

func (b FeatureBlock) DisplayName() string {
	return "Feature"
}

func (b FeatureBlock) FilterValue() string {
	return b.Heading
}

func (b FeatureBlock) Title() string {
	return "[" + b.DisplayName() + "] [" + b.View + "] " + b.Heading
}

func (b *FeatureBlock) TitlePointer() *string {
	return &b.Heading
}

func (b FeatureBlock) Description() string {
	return b.Subheading
}

func (b *FeatureBlock) DescriptionPointer() *string {
	return &b.Subheading
}

func (b FeatureBlock) GetFeatures() []Feature {
	return b.Features
}

func (b *FeatureBlock) SetFeatures(features []Feature) {
	b.Features = features
}

func (b FeatureBlock) ID() string {
	return "feature_section"
}

func (b *FeatureBlock) GetFieldDefinitions() []*FieldDefinition {
	return []*FieldDefinition{
		{Key: "Heading", Title: "Heading", Type: FieldTypeInput, ValuePointer: &b.Heading},
		{Key: "Subheading", Title: "Subheading", Type: FieldTypeInput, ValuePointer: &b.Subheading},
		{Key: "Features", Title: "Features", Type: FieldTypeFeaturesShow, FeaturesPointer: &b.Features},
	}
}
