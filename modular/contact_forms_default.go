package modular

type ContactFormsDefault struct {
	Type        string `yaml:"type"`
	HideFromNav bool   `yaml:"hide_from_nav"`
	Library     string `yaml:"library"`
	Heading     string `yaml:"heading"`
	Subheading  string `yaml:"subheading"`
	SubmitText  string `yaml:"submit_text"`
}

func (f *ContactFormsDefault) DisplayName() string {
	return "Contact Default Form"
}

func (f ContactFormsDefault) FilterValue() string {
	return f.Heading
}

func (f *ContactFormsDefault) Category() string {
	return "Contact"
}

func (f ContactFormsDefault) View() string {
	return "default"
}

func (f ContactFormsDefault) Title() string {
	return "[" + f.Category() + "] [" + f.View() + "] " + f.Heading
}

func (f *ContactFormsDefault) TitlePointer() *string {
	return &f.Heading
}

func (f ContactFormsDefault) Description() string {
	return f.Subheading
}

func (f *ContactFormsDefault) DescriptionPointer() *string {
	return &f.Subheading
}

func (f ContactFormsDefault) GetFeatures() []Feature {
	return []Feature{}
}

func (f *ContactFormsDefault) SetFeatures(features []Feature) {
}

func (f *ContactFormsDefault) GetFieldDefinitions() []*FieldDefinition {
	return []*FieldDefinition{
		{Key: "Heading", Title: "Heading", Type: FieldTypeInput, ValuePointer: &f.Heading},
		{Key: "Subheading", Title: "Subheading", Type: FieldTypeInput, ValuePointer: &f.Subheading},
		{Key: "SubmitText", Title: "Submit Text", Type: FieldTypeInput, ValuePointer: &f.SubmitText},
	}
}

func (f *ContactFormsDefault) ID() string {
	return "contact_default_form"
}
