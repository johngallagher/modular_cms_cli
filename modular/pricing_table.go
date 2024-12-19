package modular

type PricingTable struct {
	Type       string    `yaml:"type"`
	Library    string    `yaml:"library"`
	Heading    string    `yaml:"heading"`
	Subheading string    `yaml:"subheading"`
	Products   []Product `yaml:"products"`
}

type Product struct {
	ID            string `yaml:"id"`
	PaymentLinkID string `yaml:"payment_link_id"`
}

func (p *PricingTable) DisplayName() string {
	return "Pricing Table"
}

func (p PricingTable) FilterValue() string {
	return p.Heading
}

func (p PricingTable) Title() string {
	return p.Heading
}

func (p *PricingTable) TitlePointer() *string {
	return &p.Heading
}

func (p PricingTable) Description() string {
	return p.Subheading
}

func (p *PricingTable) DescriptionPointer() *string {
	return &p.Subheading
}

func (p PricingTable) GetFeatures() []Feature {
	return []Feature{}
}

func (p *PricingTable) SetFeatures(features []Feature) {
}

func (p *PricingTable) GetFieldDefinitions() []*FieldDefinition {
	return []*FieldDefinition{
		{Key: "Heading", Title: "Heading", Type: FieldTypeInput, ValuePointer: &p.Heading},
		{Key: "Subheading", Title: "Subheading", Type: FieldTypeInput, ValuePointer: &p.Subheading},
	}
}

func (p *PricingTable) ID() string {
	return "pricing_table"
}
