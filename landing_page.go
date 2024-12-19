package main

// LandingPage represents the top-level structure
type LandingPage struct {
	Blocks []BlockInterface
}

func NewLandingPage() *LandingPage {
	return &LandingPage{
		Blocks: []BlockInterface{
			NewMarketingHeroCoverImageWithCtas(),
		},
	}
}
