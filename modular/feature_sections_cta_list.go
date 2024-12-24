package modular

type FeatureSectionsCtaList struct {
	Type        string    `yaml:"type"`
	HideFromNav bool      `yaml:"hide_from_nav"`
	Heading     string    `yaml:"heading"`
	Subheading  string    `yaml:"subheading"`
	Features    []Feature `yaml:"features"`
	Library     string    `yaml:"library"`
}

func (b FeatureSectionsCtaList) DisplayName() string {
	return "Feature Sections Cta List"
}

func (b FeatureSectionsCtaList) FilterValue() string {
	return b.Heading
}

func (b FeatureSectionsCtaList) Title() string {
	return "[" + b.DisplayName() + "] " + b.Heading
}

func (b *FeatureSectionsCtaList) TitlePointer() *string {
	return &b.Heading
}

func (b FeatureSectionsCtaList) Description() string {
	return b.Subheading
}

func (b *FeatureSectionsCtaList) DescriptionPointer() *string {
	return &b.Subheading
}

func (b FeatureSectionsCtaList) GetFeatures() []Feature {
	return b.Features
}

func (b *FeatureSectionsCtaList) SetFeatures(features []Feature) {
	b.Features = features
}

func (b FeatureSectionsCtaList) ID() string {
	return "feature_sections_cta_list"
}

func (b *FeatureSectionsCtaList) GetFieldDefinitions() []*FieldDefinition {
	return []*FieldDefinition{
		{Key: "Features", Title: "Features", Type: FieldTypeFeaturesShow, FeaturesPointer: &b.Features},
		{Key: "Heading", Title: "Heading", Type: FieldTypeInput, ValuePointer: &b.Heading},
		{Key: "Subheading", Title: "Subheading", Type: FieldTypeInput, ValuePointer: &b.Subheading},
	}
}
