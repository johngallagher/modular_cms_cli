package modular

type HeroSectionsDefault struct {
	Type        string      `yaml:"type"`
	HideFromNav bool        `yaml:"hide_from_nav"`
	Library     string      `yaml:"library"`
	Heading     string      `yaml:"heading"`
	Subheading  string      `yaml:"subheading"`
	Left        SideJustCTA `yaml:"left"`
	Right       SideJustCTA `yaml:"right"`
}

type SideJustCTA struct {
	CTA CTA `yaml:"cta"`
}

func (h *HeroSectionsDefault) DisplayName() string {
	return "Hero Sections Default"
}

func (h *HeroSectionsDefault) FilterValue() string {
	return h.Heading
}

func (h *HeroSectionsDefault) Category() string {
	return "Hero"
}

func (h *HeroSectionsDefault) View() string {
	return "default"
}

func (h *HeroSectionsDefault) Title() string {
	return "[" + h.Category() + "] [" + h.View() + "] " + h.Heading
}

func (h *HeroSectionsDefault) TitlePointer() *string {
	return &h.Heading
}

func (h *HeroSectionsDefault) Description() string {
	return h.Subheading
}

func (h *HeroSectionsDefault) DescriptionPointer() *string {
	return &h.Subheading
}

func (h *HeroSectionsDefault) GetFeatures() []Feature {
	return []Feature{}
}

func (h *HeroSectionsDefault) SetFeatures(features []Feature) {
	// No-op since this block type doesn't use features
}
func (h *HeroSectionsDefault) ID() string {
	return "hero_sections_default"
}

func (b *HeroSectionsDefault) GetFieldDefinitions() []*FieldDefinition {
	return []*FieldDefinition{
		{Key: "Heading", Title: "Heading", Type: FieldTypeInput, ValuePointer: &b.Heading},
		{Key: "Subheading", Title: "Subheading", Type: FieldTypeInput, ValuePointer: &b.Subheading},
		{Key: "Left.CTA.Text", Title: "Left CTA Text", Type: FieldTypeInput, ValuePointer: &b.Left.CTA.Text},
		{Key: "Left.CTA.URL", Title: "Left CTA URL", Type: FieldTypeInput, ValuePointer: &b.Left.CTA.URL},
		{Key: "Right.CTA.Text", Title: "Right CTA Text", Type: FieldTypeInput, ValuePointer: &b.Right.CTA.Text},
		{Key: "Right.CTA.URL", Title: "Right CTA URL", Type: FieldTypeInput, ValuePointer: &b.Right.CTA.URL},
	}
}
