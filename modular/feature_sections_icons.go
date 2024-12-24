package modular

type FeatureSectionsIcons struct {
	Type     string    `yaml:"type"`
	Library  string    `yaml:"library"`
	Heading  string    `yaml:"heading"`
	Features []Feature `yaml:"features"`
}

func (f *FeatureSectionsIcons) DisplayName() string {
	return "Feature Sections Icons"
}
func (f FeatureSectionsIcons) FilterValue() string {
	return f.Heading
}

func (f FeatureSectionsIcons) Title() string {
	return "[" + f.DisplayName() + "] " + f.Heading
}

func (f *FeatureSectionsIcons) TitlePointer() *string {
	return &f.Heading
}

func (f FeatureSectionsIcons) Description() string {
	return ""
}

func (f *FeatureSectionsIcons) DescriptionPointer() *string {
	return nil
}

func (f FeatureSectionsIcons) GetFeatures() []Feature {
	return f.Features
}

func (f *FeatureSectionsIcons) SetFeatures(features []Feature) {
	f.Features = features
}

func (f *FeatureSectionsIcons) GetFieldDefinitions() []*FieldDefinition {
	return []*FieldDefinition{
		{Key: "Heading", Title: "Heading", Type: FieldTypeInput, ValuePointer: &f.Heading},
		{Key: "Features", Title: "Features", Type: FieldTypeFeaturesShow, FeaturesPointer: &f.Features},
	}
}

func (f *FeatureSectionsIcons) ID() string {
	return "feature_sections_icons"
}
