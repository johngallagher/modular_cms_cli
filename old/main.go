package modular

import (
	"fmt"
	"log"
	"os"

	"github.com/ktr0731/go-fuzzyfinder"
)

var filename = "../modular_cms/src/index.md"

// var filename = "index.md"

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: modular <command> <subcommand>")
		fmt.Println("Commands: block")
		os.Exit(1)
	}

	command := os.Args[1]
	if command == "block" {
		if len(os.Args) < 3 {
			fmt.Println("Usage: modular block <subcommand>")
			fmt.Println("Subcommands: add")
			os.Exit(1)
		}

		subcommand := os.Args[2]
		switch subcommand {
		case "add":
			addBlock()
		default:
			fmt.Printf("Unknown subcommand: %s\n", subcommand)
			os.Exit(1)
		}
	} else {
		fmt.Printf("Unknown command: %s\n", command)
		os.Exit(1)
	}
}

func addBlock() {
	mdData, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	page, err := PageFromMarkdown(mdData)
	if err != nil {
		log.Fatal(err)
	}

	allBlocks := AllBlocks()
	singleUseBlocks := page.SingleUseBlocks()

	blocks := removeBlocks(allBlocks, singleUseBlocks)

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
				return ""
			}
			page.Blocks[existingBlocksLength-1] = blocks[i]
			page.WriteToFile(filename)
			return fmt.Sprintf("Block: %s", blocks[i].DisplayName())
		}))
	if err != nil {
		// Remove the blank block if user cancels
		page.Blocks = page.Blocks[:existingBlocksLength-1]
		page.WriteToFile(filename)
		log.Fatal(err)
	}
	fmt.Printf("selected: %v\n", idx)
}
