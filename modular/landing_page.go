package modular

import (
	"bytes"
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

// LandingPage represents the top-level structure
type LandingPage struct {
	Blocks  []BlockInterface
	Content string
	Path    string
}

func NewLandingPage() *LandingPage {
	return &LandingPage{
		Blocks: []BlockInterface{
			NewMarketingHeroCoverImageWithCtas(),
		},
	}
}

func LandingPageFromMarkdownAtPath(path string) *LandingPage {
	markdown, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}

	blocks, content := ParseBlocksAndContent(markdown)

	return &LandingPage{
		Blocks:  blocks,
		Content: content,
		Path:    path,
	}
}

func (lp LandingPage) Write() error {
	frontmatter := map[string]interface{}{
		"type":   "Page",
		"layout": "page.webc",
		"blocks": lp.Blocks,
	}
	yaml, err := yaml.Marshal(frontmatter)
	if err != nil {
		return fmt.Errorf("error marshaling frontmatter: %v", err)
	}

	// Combine frontmatter and content
	content := fmt.Sprintf("---\n%s---\n%s", yaml, lp.Content)

	// Write to file
	if err := os.WriteFile(lp.Path, []byte(content), 0644); err != nil {
		return fmt.Errorf("error writing file: %v", err)
	}

	return nil

}

func ParseBlocksAndContent(markdown []byte) ([]BlockInterface, string) {
	// Split frontmatter from content
	parts := bytes.Split(markdown, []byte("---\n"))
	if len(parts) < 3 {
		panic(fmt.Errorf("invalid markdown file format - missing frontmatter (got %d parts, expected at least 3)", len(parts)))
	}

	// Parse YAML frontmatter into intermediate map
	var frontmatter map[string]interface{}
	if err := yaml.Unmarshal(parts[1], &frontmatter); err != nil {
		panic(fmt.Errorf("error parsing frontmatter: %v", err))
	}

	// Create page and handle blocks specially
	if blocksData, ok := frontmatter["blocks"].([]interface{}); ok {
		blocks := make([]BlockInterface, 0, len(blocksData))
		for i, blockData := range blocksData {
			if blockMap, ok := blockData.(map[string]interface{}); ok {
				block, err := ParseBlock(blockMap)
				if err != nil {
					panic(fmt.Errorf("error parsing block %d: %v", i, err))
				}
				blocks = append(blocks, block)
			}
		}
		return blocks, string(parts[2])
	}

	panic(fmt.Errorf("invalid markdown file format - missing blocks"))
}

func ParseBlock(blockData map[string]interface{}) (BlockInterface, error) {
	typeStr, ok := blockData["type"].(string)
	if !ok {
		return nil, fmt.Errorf("block missing type field")
	}

	var block BlockInterface
	switch typeStr {
	// case "FeatureSectionsCtaList":
	// 	block = &FeatureSectionsCtaList{Type: typeStr}
	case "MarketingHeroCoverImageWithCtas":
		block = &MarketingHeroCoverImageWithCtas{Type: typeStr}
	// case "FeatureSectionsIcons":
	// 	block = &FeatureSectionsIcons{Type: typeStr}
	// case "FeatureSectionsCardList":
	// 	block = &FeatureSectionsCardList{Type: typeStr}
	// case "PricingTable":
	// 	block = &PricingTable{Type: typeStr}
	// case "FaqSectionsAccordion":
	// 	block = &FaqSectionsAccordion{Type: typeStr}
	// case "BlankBlock":
	// 	block = &BlankBlock{Type: typeStr}
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
