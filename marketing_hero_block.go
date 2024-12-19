package main

// MarketingHeroCoverImageWithCtas represents a marketing hero block with CTAs
type MarketingHeroCoverImageWithCtas struct {
	_type       string    `yaml:"type"`
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

func (b MarketingHeroCoverImageWithCtas) Type() string {
	return b._type
}

func (b MarketingHeroCoverImageWithCtas) ID() string {
	return "marketing_hero_cover_image_with_ctas"
}

func (b *MarketingHeroCoverImageWithCtas) GetFieldDefinitions() []*FieldDefinition {
	return []*FieldDefinition{
		{Key: "Features", Title: "Features", Type: FieldTypeFeaturesShow, FeaturesPointer: &b.Features},
		{Key: "Heading", Title: "Heading", Type: FieldTypeInput, ValuePointer: &b.Heading},
		{Key: "Subheading", Title: "Subheading", Type: FieldTypeInput, ValuePointer: &b.Subheading},
		{Key: "Left.Heading", Title: "Left Heading", Type: FieldTypeInput, ValuePointer: &b.Left.Heading},
		{Key: "Left.Subheading", Title: "Left Subheading", Type: FieldTypeInput, ValuePointer: &b.Left.Subheading},
		{Key: "Left.CTA.Text", Title: "Left CTA Text", Type: FieldTypeInput, ValuePointer: &b.Left.CTA.Text},
		{Key: "Left.CTA.URL", Title: "Left CTA URL", Type: FieldTypeInput, ValuePointer: &b.Left.CTA.URL},
		{Key: "Right.Heading", Title: "Right Heading", Type: FieldTypeInput, ValuePointer: &b.Right.Heading},
		{Key: "Right.Subheading", Title: "Right Subheading", Type: FieldTypeInput, ValuePointer: &b.Right.Subheading},
		{Key: "Right.CTA.Text", Title: "Right CTA Text", Type: FieldTypeInput, ValuePointer: &b.Right.CTA.Text},
		{Key: "Right.CTA.URL", Title: "Right CTA URL", Type: FieldTypeInput, ValuePointer: &b.Right.CTA.URL},
		{Key: "Image.URL", Title: "Image URL", Type: FieldTypeInput, ValuePointer: &b.Image.URL},
	}
}

// NewMarketingHeroCoverImageWithCtas creates a new marketing hero block with default values
func NewMarketingHeroCoverImageWithCtas() *MarketingHeroCoverImageWithCtas {
	return &MarketingHeroCoverImageWithCtas{
		_type:       "marketing_hero_cover_image_with_ctas",
		Heading:     "Understand your Rails app in production",
		Subheading:  "Stop guessing and understand what your Rails app is actually doing in production.",
		HideFromNav: true,
		Image: Image{
			URL: "https://www.google.com/images/branding/googlelogo/1x/googlelogo_color_272x92dp.png",
		},
		Left: Side{
			Heading:    "3 February 2025",
			Subheading: "6 Weeks Workshop. Hands On And Practical. No DevOps Experience Needed. Less Than 4 Hours A Week. Designed For Seniors And Leads. Solo Or Team Workshops.",
			CTA: CTA{
				Text: "Book Now",
				URL:  "https://www.google.com",
			},
		},
		Right: Side{
			Heading:    "Understand your Rails app in production",
			Subheading: "Stop guessing and understand what your Rails app is actually doing in production.",
			CTA: CTA{
				Text: "Book Now",
				URL:  "https://www.google.com",
			},
		},
		Features: []Feature{
			{Name: "Feature 1", Description: "Description 1"},
			{Name: "Feature 2", Description: "Description 2"},
		},
	}
}
