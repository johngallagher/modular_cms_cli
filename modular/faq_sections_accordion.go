package modular

type FaqSectionsAccordion struct {
	Type       string     `yaml:"type"`
	Library    string     `yaml:"library"`
	Heading    string     `yaml:"heading"`
	Subheading string     `yaml:"subheading"`
	Questions  []Question `yaml:"questions"`
}

func (f *FaqSectionsAccordion) DisplayName() string {
	return "Faq Sections Accordion"
}

type Question struct {
	Question string `yaml:"question"`
	Answer   string `yaml:"answer"`
}

func (f FaqSectionsAccordion) FilterValue() string {
	return f.Heading
}

func (f FaqSectionsAccordion) Title() string {
	return "[" + f.DisplayName() + "] " + f.Heading
}

func (f *FaqSectionsAccordion) TitlePointer() *string {
	return &f.Heading
}

func (f FaqSectionsAccordion) Description() string {
	return f.Subheading
}

func (f *FaqSectionsAccordion) DescriptionPointer() *string {
	return &f.Subheading
}

func (f FaqSectionsAccordion) GetFeatures() []Feature {
	return []Feature{}
}

func (f *FaqSectionsAccordion) SetFeatures(features []Feature) {
}

func (f *FaqSectionsAccordion) GetFieldDefinitions() []*FieldDefinition {
	return []*FieldDefinition{
		{Key: "Heading", Title: "Heading", Type: FieldTypeInput, ValuePointer: &f.Heading},
		{Key: "Subheading", Title: "Subheading", Type: FieldTypeInput, ValuePointer: &f.Subheading},
	}
}

func (f *FaqSectionsAccordion) ID() string {
	return "faq_sections_accordion"
}
