package main

import (
	"fmt"
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

func parseBlock(blockData map[string]interface{}) (Block, error) {
	typeStr, ok := blockData["type"].(string)
	if !ok {
		return nil, fmt.Errorf("block missing type field")
	}

	var block Block
	switch typeStr {
	case "FeatureSectionsCtaList":
		block = &FeatureSectionsCtaList{Type: typeStr}
	case "MarketingHeroCoverImageWithCtas":
		block = &MarketingHeroCoverImageWithCtas{Type: typeStr}
	case "BlankBlock":
		block = &BlankBlock{Type: typeStr}
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

func AllBlocks() []Block {
	yamlData, err := os.ReadFile("all_blocks.yml")
	if err != nil {
		log.Fatal(err)
	}

	var raw []map[string]interface{}
	if err := yaml.Unmarshal(yamlData, &raw); err != nil {
		log.Fatal(err)
	}

	blocks := make([]Block, len(raw))
	for i, blockData := range raw {
		block, err := parseBlock(blockData)
		if err != nil {
			log.Fatalf("error parsing block %d: %v", i, err)
		}
		blocks[i] = block
	}

	return blocks
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

type MarketingHeroCoverImageWithCtas struct {
	Type        string `yaml:"type"`
	HideFromNav bool   `yaml:"hide_from_nav"`
	Heading     string `yaml:"heading"`
	Subheading  string `yaml:"subheading"`
	Left        Side   `yaml:"left"`
	Right       Side   `yaml:"right"`
	Image       Image  `yaml:"image"`
}

type Feature struct {
	Heading string `yaml:"heading"`
	Summary string `yaml:"summary"`
	Icon    string `yaml:"icon"`
}

type FeatureSectionsCtaList struct {
	Type        string    `yaml:"type"`
	HideFromNav bool      `yaml:"hide_from_nav"`
	Heading     string    `yaml:"heading"`
	Subheading  string    `yaml:"subheading"`
	Features    []Feature `yaml:"features"`
	Library     string    `yaml:"library"`
}

type BlankBlock struct {
	Type string `yaml:"type"`
}

func (b *BlankBlock) DisplayName() string {
	return ""
}

type Block interface {
	DisplayName() string
}

func (m *MarketingHeroCoverImageWithCtas) DisplayName() string {
	return "Marketing Hero Cover Image With Ctas"
}

func (f *FeatureSectionsCtaList) DisplayName() string {
	return "Feature Sections Cta List"
}
