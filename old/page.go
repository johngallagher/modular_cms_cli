package modular

import (
	"bytes"
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type Page struct {
	Blocks  []Block `yaml:"blocks"`
	Content string
}

func rejectEmptyBlocks(blocks []Block) []Block {
	filteredBlocks := make([]Block, 0, len(blocks))
	for _, block := range blocks {
		if block.DisplayName() != "" {
			filteredBlocks = append(filteredBlocks, block)
		}
	}
	return filteredBlocks
}

func (p *Page) WriteToFile(filename string) error {
	// Marshal title and blocks to YAML frontmatter
	frontmatter := map[string]interface{}{
		"type":   "Page",
		"layout": "page.webc",
		"blocks": rejectEmptyBlocks(p.Blocks),
	}
	yaml, err := yaml.Marshal(frontmatter)
	if err != nil {
		return fmt.Errorf("error marshaling frontmatter: %v", err)
	}

	// Combine frontmatter and content
	content := fmt.Sprintf("---\n%s---\n%s", yaml, p.Content)

	// Write to file
	if err := os.WriteFile(filename, []byte(content), 0644); err != nil {
		return fmt.Errorf("error writing file: %v", err)
	}

	return nil
}

func PageFromMarkdown(mdData []byte) (*Page, error) {
	// Split frontmatter from content
	parts := bytes.Split(mdData, []byte("---\n"))
	if len(parts) < 3 {
		return nil, fmt.Errorf("invalid markdown file format - missing frontmatter (got %d parts, expected at least 3)", len(parts))
	}

	// Parse YAML frontmatter into intermediate map
	var frontmatter map[string]interface{}
	if err := yaml.Unmarshal(parts[1], &frontmatter); err != nil {
		return nil, fmt.Errorf("error parsing frontmatter: %v", err)
	}

	// Create page and handle blocks specially
	page := &Page{}
	if blocksData, ok := frontmatter["blocks"].([]interface{}); ok {
		blocks := make([]Block, 0, len(blocksData))
		for i, blockData := range blocksData {
			if blockMap, ok := blockData.(map[string]interface{}); ok {
				block, err := parseBlock(blockMap)
				if err != nil {
					return nil, fmt.Errorf("error parsing block %d: %v", i, err)
				}
				blocks = append(blocks, block)
			}
		}
		page.Blocks = blocks
	}

	page.Content = string(parts[2])

	return page, nil
}

func (p *Page) AppendBlankBlock() {
	p.Blocks = append(p.Blocks, &BlankBlock{Type: "BlankBlock"})
}

func (p *Page) SingleUseBlocks() []Block {
	return p.Blocks // All blocks are single use for now
}
