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

func (lp *LandingPage) RemoveBlockAtIndex(index int) {
	lp.Blocks = append(lp.Blocks[:index], lp.Blocks[index+1:]...)
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
		b := make([]BlockInterface, 0, len(blocksData))
		for i, blockData := range blocksData {
			if blockMap, ok := blockData.(map[string]interface{}); ok {
				block, err := Parse(blockMap)
				if err != nil {
					panic(fmt.Errorf("error parsing block %d: %v", i, err))
				}
				b = append(b, block)
			}
		}
		return b, string(parts[2])
	}

	panic(fmt.Errorf("invalid markdown file format - missing blocks"))
}
