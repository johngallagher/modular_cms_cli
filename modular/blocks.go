package modular

// Model definition
type Block struct {
	id          string
	title       string
	features    []Feature
	description string
}

type BlockInterface interface {
	ID() string
	Title() string
	Description() string
	DescriptionPointer() *string
	GetFieldDefinitions() []*FieldDefinition
	GetFeatures() []Feature
	SetFeatures([]Feature)
	FilterValue() string
}

func (b Block) FilterValue() string {
	return b.title
}

func (b Block) Title() string {
	return b.title
}

func (b Block) Description() string {
	return b.description
}

func (b Block) DescriptionPointer() *string {
	return &b.description
}

type CTA struct {
	Text string `yaml:"text"`
	URL  string `yaml:"url"`
}

type Side struct {
	Heading    string `yaml:"heading"`
	Subheading string `yaml:"subheading"`
	CTA        CTA    `yaml:"cta"`
}

type Image struct {
	URL string `yaml:"url"`
}

type NewFeature struct {
	Heading string `yaml:"heading"`
	Summary string `yaml:"summary"`
	Icon    string `yaml:"icon"`
}

type BlankBlock struct {
	Type string `yaml:"type"`
}
