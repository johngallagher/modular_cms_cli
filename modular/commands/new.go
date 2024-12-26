package commands

import (
	"embed"
	"fmt"
	"github.com/spf13/cobra"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

//go:embed templates/*
var templateFiles embed.FS

func New(cmd *cobra.Command, args []string) {
	path := args[0]

	// Convert relative path to absolute
	absPath, err := filepath.Abs(path)
	if err != nil {
		fmt.Printf("Error resolving path: %v\n", err)
		os.Exit(1)
	}

	// Create the directory if it doesn't exist
	if err := os.MkdirAll(absPath, 0755); err != nil {
		fmt.Printf("Error creating directory: %v\n", err)
		os.Exit(1)
	}

	// Recursively walk through all embedded files
	var walkDir func(string) error
	walkDir = func(dir string) error {
		entries, err := templateFiles.ReadDir(dir)
		if err != nil {
			return fmt.Errorf("error reading directory %s: %v", dir, err)
		}

		for _, entry := range entries {
			path := filepath.Join(dir, entry.Name())

			if entry.IsDir() {
				if err := walkDir(path); err != nil {
					return err
				}
				continue
			}

			// Read template content
			content, err := templateFiles.ReadFile(path)
			if err != nil {
				return fmt.Errorf("error reading file %s: %v", path, err)
			}

			// Create output file path relative to absPath
			relPath, err := filepath.Rel("templates", path)
			if err != nil {
				return fmt.Errorf("error getting relative path: %v", err)
			}
			outPath := filepath.Join(absPath, relPath)

			// Create directory structure
			if err := os.MkdirAll(filepath.Dir(outPath), 0755); err != nil {
				return fmt.Errorf("error creating directories: %v", err)
			}

			// Create output file without .tmpl extension if present
			outPath = strings.TrimSuffix(outPath, ".tmpl")
			outFile, err := os.Create(outPath)
			if err != nil {
				return fmt.Errorf("error creating file %s: %v", outPath, err)
			}
			defer outFile.Close()

			if strings.HasSuffix(entry.Name(), ".tmpl") {
				// Parse and execute template
				tmpl, err := template.New(entry.Name()).Parse(string(content))
				if err != nil {
					return fmt.Errorf("error parsing template %s: %v", path, err)
				}

				// Convert first argument to title case and humanize for site name

				siteName := filepath.Base(absPath)
				siteName = strings.ReplaceAll(siteName, "_", " ")
				siteName = cases.Title(language.English).String(siteName)

				type Site struct {
					Name        string
					Product     string
					Company     string
					Description string
				}

				type Locals struct {
					Site Site
				}
				locals := Locals{
					Site: Site{
						Name:        siteName,
						Product:     "My Product",
						Company:     "My Company",
						Description: "My Description",
					},
				}
				// Execute template
				if err := tmpl.Execute(outFile, locals); err != nil {
					return fmt.Errorf("error executing template %s: %v", path, err)
				}
			} else {
				// Copy file
				content, err := templateFiles.ReadFile(path)
				if err != nil {
					return fmt.Errorf("error reading file %s: %v", path, err)
				}
				if err := os.WriteFile(outPath, content, 0644); err != nil {
					return fmt.Errorf("error writing file %s: %v", outPath, err)
				}
			}
		}
		return nil
	}

	if err := walkDir("templates"); err != nil {
		fmt.Printf("Error processing templates: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Created new Modular site at %s\n", absPath)
}
