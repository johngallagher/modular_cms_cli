package main

import (
	"fmt"
	"log"
	"os"

	"github.com/ktr0731/go-fuzzyfinder"
)

// var filename = "../modular_cms/src/index.md"

var filename = "index.md"

func main() {
	mdData, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	page, err := PageFromMarkdown(mdData)
	if err != nil {
		log.Fatal(err)
	}

	blocks := AllBlocks()

	page.AppendBlankBlock()
	existingBlocks := page.Blocks
	existingBlocksLength := len(existingBlocks)

	idx, err := fuzzyfinder.FindMulti(
		blocks,
		func(i int) string {
			return blocks[i].DisplayName()
		},
		fuzzyfinder.WithPreviewWindow(func(i, w, h int) string {
			if i == -1 {
				page.Blocks[existingBlocksLength-1] = &BlankBlock{Type: "BlankBlock"}
				page.WriteToFile(filename)
				return ""
			}
			page.Blocks[existingBlocksLength-1] = blocks[i]
			page.WriteToFile(filename)
			return fmt.Sprintf("Block: %s", blocks[i].DisplayName())
		}))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("selected: %v\n", idx)
}
