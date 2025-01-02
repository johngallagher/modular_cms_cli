package modular

type TestimonialSectionsBlockquote struct {
	Identifier  string      `yaml:"identifier"`
	Type        string      `yaml:"type"`
	HideFromNav bool        `yaml:"hide_from_nav"`
	Library     string      `yaml:"library"`
	Testimonial Testimonial `yaml:"testimonial"`
}

type Testimonial struct {
	Content string `yaml:"content"`
	Author  Author `yaml:"author"`
}

type Author struct {
	Name     string `yaml:"name"`
	Title    string `yaml:"title"`
	ImageSrc string `yaml:"image_src"`
}

func (t *TestimonialSectionsBlockquote) DisplayName() string {
	return "Testimonial Card"
}

func (t *TestimonialSectionsBlockquote) FilterValue() string {
	return t.Testimonial.Content
}

func (t *TestimonialSectionsBlockquote) Category() string {
	return "Testimonial"
}

func (t *TestimonialSectionsBlockquote) View() string {
	return "blockquote"
}

func (t *TestimonialSectionsBlockquote) Title() string {
	return "[" + t.Category() + "] [" + t.View() + "] " + t.Testimonial.Content
}

func (t *TestimonialSectionsBlockquote) TitlePointer() *string {
	return &t.Testimonial.Content
}

func (t *TestimonialSectionsBlockquote) Description() string {
	return t.Testimonial.Author.Name
}

func (t *TestimonialSectionsBlockquote) DescriptionPointer() *string {
	return &t.Testimonial.Author.Name
}

func (t *TestimonialSectionsBlockquote) GetFeatures() []Feature {
	return []Feature{}
}

func (t *TestimonialSectionsBlockquote) SetFeatures(features []Feature) {
	// No-op since this block type doesn't use features
}
func (t *TestimonialSectionsBlockquote) ID() string {
	return t.Identifier
}

func (t *TestimonialSectionsBlockquote) GetFieldDefinitions() []*FieldDefinition {
	return []*FieldDefinition{
		{Key: "Testimonial.Content", Title: "Testimonial Content", Type: FieldTypeInput, ValuePointer: &t.Testimonial.Content},
		{Key: "Testimonial.Author.Name", Title: "Testimonial Author Name", Type: FieldTypeInput, ValuePointer: &t.Testimonial.Author.Name},
		{Key: "Testimonial.Author.Title", Title: "Testimonial Author Title", Type: FieldTypeInput, ValuePointer: &t.Testimonial.Author.Title},
		{Key: "Testimonial.Author.ImageSrc", Title: "Testimonial Author Image", Type: FieldTypeInput, ValuePointer: &t.Testimonial.Author.ImageSrc},
	}
}
