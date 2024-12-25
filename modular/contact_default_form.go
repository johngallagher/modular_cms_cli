package modular

type ContactDefaultForm struct {
	Type       string `yaml:"type"`
	Library    string `yaml:"library"`
	Heading    string `yaml:"heading"`
	Subheading string `yaml:"subheading"`
	SubmitText string `yaml:"submit_text"`
}

func (f *ContactDefaultForm) DisplayName() string {
	return "Contact Default Form"
}

func (f ContactDefaultForm) FilterValue() string {
	return f.Heading
}

func (f ContactDefaultForm) Title() string {
	return "[" + f.DisplayName() + "] " + f.Heading
}

func (f *ContactDefaultForm) TitlePointer() *string {
	return &f.Heading
}

func (f ContactDefaultForm) Description() string {
	return f.Subheading
}

func (f *ContactDefaultForm) DescriptionPointer() *string {
	return &f.Subheading
}

func (f ContactDefaultForm) GetFeatures() []Feature {
	return []Feature{}
}

func (f *ContactDefaultForm) SetFeatures(features []Feature) {
}

func (f *ContactDefaultForm) GetFieldDefinitions() []*FieldDefinition {
	return []*FieldDefinition{
		{Key: "Heading", Title: "Heading", Type: FieldTypeInput, ValuePointer: &f.Heading},
		{Key: "Subheading", Title: "Subheading", Type: FieldTypeInput, ValuePointer: &f.Subheading},
		{Key: "SubmitText", Title: "Submit Text", Type: FieldTypeInput, ValuePointer: &f.SubmitText},
	}
}

func (f *ContactDefaultForm) ID() string {
	return "contact_default_form"
}
