package modular

// MarketingHeroCoverImageWithCtas represents a marketing hero block with CTAs
type MarketingHeroCoverImageWithCtas struct {
	Type        string    `yaml:"type"`
	HideFromNav bool      `yaml:"hide_from_nav"`
	Heading     string    `yaml:"heading"`
	Subheading  string    `yaml:"subheading"`
	Left        Side      `yaml:"left"`
	Right       Side      `yaml:"right"`
	Image       Image     `yaml:"image"`
	Features    []Feature `yaml:"features"`
}

func (b MarketingHeroCoverImageWithCtas) DisplayName() string {
	return "Marketing Hero Cover Image With Ctas"
}

func (b MarketingHeroCoverImageWithCtas) FilterValue() string {
	return b.Heading
}

func (b MarketingHeroCoverImageWithCtas) Title() string {
	return b.Heading
}

func (b *MarketingHeroCoverImageWithCtas) TitlePointer() *string {
	return &b.Heading
}

func (b MarketingHeroCoverImageWithCtas) Description() string {
	return b.Subheading
}

func (b *MarketingHeroCoverImageWithCtas) DescriptionPointer() *string {
	return &b.Subheading
}

func (b MarketingHeroCoverImageWithCtas) GetFeatures() []Feature {
	return b.Features
}

func (b *MarketingHeroCoverImageWithCtas) SetFeatures(features []Feature) {
	b.Features = features
}

func (b MarketingHeroCoverImageWithCtas) ID() string {
	return "marketing_hero_cover_image_with_ctas"
}

func (b *MarketingHeroCoverImageWithCtas) GetFieldDefinitions() []*FieldDefinition {
	return []*FieldDefinition{
		{Key: "Features", Title: "Features", Type: FieldTypeFeaturesShow, FeaturesPointer: &b.Features},
		{Key: "Heading", Title: "Heading", Type: FieldTypeInput, ValuePointer: &b.Heading},
		{Key: "Subheading", Title: "Subheading", Type: FieldTypeInput, ValuePointer: &b.Subheading},
		{Key: "LeftHeading", Title: "Left Heading", Type: FieldTypeInput, ValuePointer: &b.Left.Heading},
		{Key: "LeftSubheading", Title: "Left Subheading", Type: FieldTypeTextArea, ValuePointer: &b.Left.Subheading},
		{Key: "Left.CTA.Text", Title: "Left CTA Text", Type: FieldTypeInput, ValuePointer: &b.Left.CTA.Text},
		{Key: "Left.CTA.URL", Title: "Left CTA URL", Type: FieldTypeInput, ValuePointer: &b.Left.CTA.URL},
		{Key: "RightHeading", Title: "Right Heading", Type: FieldTypeInput, ValuePointer: &b.Right.Heading},
		{Key: "RightSubheading", Title: "Right Subheading", Type: FieldTypeTextArea, ValuePointer: &b.Right.Subheading},
		{Key: "Right.CTA.Text", Title: "Right CTA Text", Type: FieldTypeInput, ValuePointer: &b.Right.CTA.Text},
		{Key: "Right.CTA.URL", Title: "Right CTA URL", Type: FieldTypeInput, ValuePointer: &b.Right.CTA.URL},
		{Key: "Image.URL", Title: "Image URL", Type: FieldTypeInput, ValuePointer: &b.Image.URL},
	}
}
