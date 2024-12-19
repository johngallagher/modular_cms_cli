package modular

type FeatureSectionsCardList struct {
	Type     string    `yaml:"type"`
	Library  string    `yaml:"library"`
	Heading  string    `yaml:"heading"`
	Features []Feature `yaml:"features"`
}

func (f *FeatureSectionsCardList) DisplayName() string {
	return "Feature Sections Card List"
}
func (f FeatureSectionsCardList) FilterValue() string {
	return f.Heading
}

func (f FeatureSectionsCardList) Title() string {
	return f.Heading
}

func (f *FeatureSectionsCardList) TitlePointer() *string {
	return &f.Heading
}

func (f FeatureSectionsCardList) Description() string {
	return ""
}

func (f *FeatureSectionsCardList) DescriptionPointer() *string {
	return nil
}

func (f FeatureSectionsCardList) GetFeatures() []Feature {
	return f.Features
}

func (f *FeatureSectionsCardList) SetFeatures(features []Feature) {
	f.Features = features
}

func (f *FeatureSectionsCardList) GetFieldDefinitions() []*FieldDefinition {
	return []*FieldDefinition{
		{Key: "Heading", Title: "Heading", Type: FieldTypeInput, ValuePointer: &f.Heading},
		{Key: "Features", Title: "Features", Type: FieldTypeFeaturesShow, FeaturesPointer: &f.Features},
	}
}

func (f *FeatureSectionsCardList) ID() string {
	return "feature_sections_card_list"
}
