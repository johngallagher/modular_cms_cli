package modular

import (
	"fmt"
	"reflect"

	"github.com/charmbracelet/huh"
)

// FieldType represents the type of form field
type FieldType int

const (
	FieldTypeInput FieldType = iota
	FieldTypeTextArea
	FieldTypeFeaturesShow
)

// FieldDefinition defines how a field should be rendered and bound
type FieldDefinition struct {
	Type            FieldType
	Key             string   // The field key (e.g., "title" or "left.cta.text")
	Title           string   // Display title
	Help            string   // Optional help text
	Required        bool     // Whether the field is required
	Path            []string // Parsed path for nested fields (e.g., ["left", "cta", "text"])
	ValuePointer    *string
	FeaturesPointer *[]Feature
}

// NewFieldDefinition creates a new field definition
// func NewFieldDefinition(fieldType FieldType, key string, title string) *FieldDefinition {
// 	return &FieldDefinition{
// 		Type:  fieldType,
// 		Key:   key,
// 		Title: title,
// 		Path:  strings.Split(key, "."),
// 	}
// }

// CreateFormField creates the appropriate huh.Field based on the definition
func (fd FieldDefinition) CreateFormField(block BlockInterface, m *MainModel) huh.Field {
	switch fd.Type {
	case FieldTypeInput:
		return huh.NewInput().
			Key(fd.Key).
			Title(fd.Title).
			Value(fd.ValuePointer)
	case FieldTypeTextArea:
		return huh.NewText().
			Key(fd.Key).
			Title(fd.Title).
			Value(fd.ValuePointer)
	case FieldTypeFeaturesShow:
		return &FeaturesShowField{
			Block:           block,
			NavigationCtx:   m.NavigationCtx(),
			Parent:          m,
			FeaturesPointer: fd.FeaturesPointer,
		}
	default:
		panic(fmt.Sprintf("unknown field type: %d", fd.Type))
	}
}

// getValue retrieves the value from the block using reflection for nested fields
func (fd *FieldDefinition) getValue(block BlockInterface) *string {
	val := reflect.ValueOf(block)
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}

	// Navigate through nested fields
	for _, fieldName := range fd.Path {
		val = val.FieldByName(fieldName)
		if !val.IsValid() {
			return nil
		}
	}

	if val.Kind() == reflect.String {
		str := val.String()
		return &str
	}

	return nil
}
