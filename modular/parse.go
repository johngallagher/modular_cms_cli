package modular

import (
	"fmt"

	"gopkg.in/yaml.v3"
)

func Parse(blockData map[string]interface{}) (BlockInterface, error) {
	typeStr, ok := blockData["type"].(string)
	if !ok {
		return nil, fmt.Errorf("block missing type field")
	}

	var block BlockInterface
	switch typeStr {
	case "MarketingHeroCoverImageWithCtas":
		block = &MarketingHeroCoverImageWithCtas{Type: typeStr}
	case "FeatureSectionsCtaList":
		block = &FeatureSectionsCtaList{Type: typeStr}
	case "FeatureSectionsIcons":
		block = &FeatureSectionsIcons{Type: typeStr}
	case "FeatureSectionsCardList":
		block = &FeatureSectionsCardList{Type: typeStr}
	case "PricingTable":
		block = &PricingTable{Type: typeStr}
	case "FaqSectionsAccordion":
		block = &FaqSectionsAccordion{Type: typeStr}
	case "ContactFormsDefault":
		block = &ContactFormsDefault{Type: typeStr}
	case "FeatureSection":
		block = &FeatureSection{Type: typeStr}
	case "HeroSectionsDefault":
		block = &HeroSectionsDefault{Type: typeStr}
	case "TestimonialSectionsBlockquote":
		block = &TestimonialSectionsBlockquote{Type: typeStr}
	case "SocialProofCardStatistics":
		block = &SocialProofCardStatistics{Type: typeStr}
	case "FeatureSectionsAlternate":
		block = &FeatureSectionsAlternate{Type: typeStr}
	case "StyledQuiz":
		block = &StyledQuiz{Type: typeStr}
	default:
		return nil, fmt.Errorf("unknown block type: %s", typeStr)
	}

	bytes, err := yaml.Marshal(blockData)
	if err != nil {
		return nil, fmt.Errorf("error marshaling block data: %v", err)
	}
	if err := yaml.Unmarshal(bytes, block); err != nil {
		return nil, fmt.Errorf("error unmarshaling block: %v", err)
	}

	return block, nil
}
