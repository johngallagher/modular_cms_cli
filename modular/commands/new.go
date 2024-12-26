package commands

import (
	"embed"
	"fmt"
	"io/fs"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"text/template"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

//go:embed templates/*
var templateFiles embed.FS

func CreateNewSite(path string) {
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

	fs.WalkDir(templateFiles, "templates", func(filePath string, d fs.DirEntry, err error) error {
		if err != nil {
			fmt.Printf("Error walking template files: %v\n", err)
			os.Exit(1)
		}

		relPath := strings.TrimPrefix(filePath, "templates/")
		destPath := filepath.Join(absPath, relPath)

		if d.IsDir() {
			// Create directory
			if err := os.MkdirAll(destPath, 0755); err != nil {
				fmt.Printf("Error creating directory: %v\n", err)
				os.Exit(1)
			}
			return nil
		}

		content, err := templateFiles.ReadFile(filePath)
		if err != nil {
			fmt.Printf("Error reading template file: %v\n", err)
			os.Exit(1)
		}

		if strings.HasSuffix(relPath, ".tmpl") {
			if err := executeTemplate(filePath, absPath, destPath); err != nil {
				fmt.Printf("Error executing template: %v\n", err)
				os.Exit(1)
			}
		} else {
			// Copy file
			if err := os.WriteFile(destPath, content, 0644); err != nil {
				fmt.Printf("Error writing file: %v\n", err)
				os.Exit(1)
			}
		}

		return nil
	})

	// Change to the new directory
	if err := os.Chdir(absPath); err != nil {
		fmt.Printf("Error changing directory: %v\n", err)
		os.Exit(1)
	}

	createPublicDir(absPath)
	runCommand("yarn", "install")
	runCommand("chmod", "+x", "bin/serve", "bin/build")
	runCommand("touch", ".env")
	runCommand("git", "init")
	runCommand("git", "add", "-A")
	runCommand("git", "commit", "-m", "Initial commit")
	fmt.Printf("Created new Modular site at %s\n", absPath)
	fmt.Println("\nNext Steps:")
	fmt.Printf(" 1. cd %s\n", absPath)
	fmt.Println(" 2. bin/serve")
}

func createPublicDir(path string) {
	publicDir := filepath.Join(path, "public")
	if err := os.MkdirAll(publicDir, 0755); err != nil {
		fmt.Printf("error creating public directory: %v", err)
		os.Exit(1)
	}
	if err := os.WriteFile(filepath.Join(publicDir, ".gitkeep"), []byte{}, 0644); err != nil {
		fmt.Printf("error creating .gitkeep file: %v", err)
		os.Exit(1)
	}
}

func executeTemplate(path string, absPath string, destPath string) error {
	content, err := templateFiles.ReadFile(path)
	if err != nil {
		return fmt.Errorf("error reading file %s: %v", path, err)
	}
	// Parse and execute template
	tmpl, err := template.New(path).Parse(string(content))
	if err != nil {
		return fmt.Errorf("error parsing template %s: %v", path, err)
	}

	// Convert path to title case and humanize for site name
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

	destPath = strings.TrimSuffix(destPath, ".tmpl")
	file, err := os.Create(destPath)
	if err != nil {
		return fmt.Errorf("error creating file %s: %v", destPath, err)
	}
	defer file.Close()

	// Execute template
	if err := tmpl.Execute(file, locals); err != nil {
		return fmt.Errorf("error executing template %s: %v", path, err)
	}
	return nil
}

func runCommand(name string, args ...string) {
	cmd := exec.Command(name, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		fmt.Printf("Error running command %s: %v\n", name, err)
		os.Exit(1)
	}
}
