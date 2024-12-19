package modular

import (
	"os"

	"gopkg.in/yaml.v3"
)

func AllBlocks() []BlockInterface {
	yamlBytes, err := os.ReadFile("modular/all_blocks.yml")
	if err != nil {
		panic(err)
	}

	var blocks []map[string]interface{}
	if err := yaml.Unmarshal(yamlBytes, &blocks); err != nil {
		panic(err)
	}
	var parsedBlocks []BlockInterface
	for _, blockData := range blocks {
		block, err := Parse(blockData)
		if err != nil {
			panic(err)
		}
		parsedBlocks = append(parsedBlocks, block)
	}

	return parsedBlocks
}
