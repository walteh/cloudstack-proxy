package metadata

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// Metadata represents the parsed CloudStack API metadata
type Metadata struct {
	Commands   map[string]CommandMetadata
	Responses  map[string]ResponseMetadata
	Enums      map[string]EnumMetadata
	Categories map[string][]string // Maps category (e.g., "vpc") to command names
}

// CommandMetadata represents the metadata for a CloudStack API command
type CommandMetadata struct {
	ClassName    string
	SimpleName   string
	Description  string
	Parameters   []ParameterMetadata
	ResponseName string
	Category     string // E.g., "vpc", "vm", "network"
}

// ParameterMetadata represents a parameter in a CloudStack API command
type ParameterMetadata struct {
	Name        string
	FieldName   string
	JavaType    string
	Description string
	Required    bool
	Since       string
}

// ResponseMetadata represents the metadata for a CloudStack API response
type ResponseMetadata struct {
	ClassName  string
	SimpleName string
	Fields     []ResponseFieldMetadata
}

// ResponseFieldMetadata represents a field in a CloudStack API response
type ResponseFieldMetadata struct {
	Name           string
	FieldName      string
	JavaType       string
	Description    string
	Since          string
	ResponseObject string
}

// EnumMetadata represents the metadata for a CloudStack API enum
type EnumMetadata struct {
	ClassName   string
	SimpleName  string
	PackageName string
	IsNested    bool
	ParentClass string
	Values      []EnumValueMetadata
}

// EnumValueMetadata represents a value in a CloudStack API enum
type EnumValueMetadata struct {
	Name    string
	Ordinal int
	Extras  map[string]interface{}
}

// ParseMetadata parses the CloudStack API metadata from the given directory
func ParseMetadata(metadataDir string) (*Metadata, error) {
	metadata := &Metadata{
		Commands:   make(map[string]CommandMetadata),
		Responses:  make(map[string]ResponseMetadata),
		Enums:      make(map[string]EnumMetadata),
		Categories: make(map[string][]string),
	}

	// Parse command metadata
	commandPath := filepath.Join(metadataDir, "cloudstack-api-command-metadata.json")
	if err := parseCommandMetadata(commandPath, metadata); err != nil {
		return nil, fmt.Errorf("error parsing command metadata: %w", err)
	}

	// Parse response metadata
	responsePath := filepath.Join(metadataDir, "cloudstack-api-response-metadata.json")
	if err := parseResponseMetadata(responsePath, metadata); err != nil {
		return nil, fmt.Errorf("error parsing response metadata: %w", err)
	}

	// Parse enum metadata
	enumPath := filepath.Join(metadataDir, "cloudstack-api-enum-metadata.json")
	if err := parseEnumMetadata(enumPath, metadata); err != nil {
		return nil, fmt.Errorf("error parsing enum metadata: %w", err)
	}

	return metadata, nil
}

// parseCommandMetadata parses the command metadata JSON file
func parseCommandMetadata(filePath string, metadata *Metadata) error {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return err
	}

	var commandsMap map[string]json.RawMessage
	if err := json.Unmarshal(data, &commandsMap); err != nil {
		return err
	}

	for className, rawData := range commandsMap {
		var cmdData struct {
			ClassName    string `json:"className"`
			SimpleName   string `json:"simpleName"`
			Description  string `json:"description"`
			ResponseName string `json:"responseName"`
			Parameters   []struct {
				Name        string `json:"name"`
				FieldName   string `json:"fieldName"`
				JavaType    string `json:"javaType"`
				Description string `json:"description"`
				Required    bool   `json:"required"`
				Since       string `json:"since"`
			} `json:"parameters"`
		}

		if err := json.Unmarshal(rawData, &cmdData); err != nil {
			return err
		}

		// Extract category from class name (e.g., ListVPCsCmd -> vpc)
		category := extractCategory(cmdData.SimpleName)

		cmd := CommandMetadata{
			ClassName:    cmdData.ClassName,
			SimpleName:   cmdData.SimpleName,
			Description:  cmdData.Description,
			ResponseName: cmdData.ResponseName,
			Category:     category,
			Parameters:   make([]ParameterMetadata, 0, len(cmdData.Parameters)),
		}

		for _, param := range cmdData.Parameters {
			cmd.Parameters = append(cmd.Parameters, ParameterMetadata{
				Name:        param.Name,
				FieldName:   param.FieldName,
				JavaType:    param.JavaType,
				Description: param.Description,
				Required:    param.Required,
				Since:       param.Since,
			})
		}

		metadata.Commands[className] = cmd

		// Add to category map
		if category != "" {
			metadata.Categories[category] = append(metadata.Categories[category], className)
		}
	}

	return nil
}

// parseResponseMetadata parses the response metadata JSON file
func parseResponseMetadata(filePath string, metadata *Metadata) error {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return err
	}

	var responsesMap map[string]json.RawMessage
	if err := json.Unmarshal(data, &responsesMap); err != nil {
		return err
	}

	for className, rawData := range responsesMap {
		var respData struct {
			ClassName  string `json:"className"`
			SimpleName string `json:"simpleName"`
			Fields     []struct {
				Name           string `json:"name"`
				FieldName      string `json:"fieldName"`
				JavaType       string `json:"javaType"`
				Description    string `json:"description"`
				Since          string `json:"since"`
				ResponseObject string `json:"responseObject"`
			} `json:"fields"`
		}

		if err := json.Unmarshal(rawData, &respData); err != nil {
			return err
		}

		resp := ResponseMetadata{
			ClassName:  respData.ClassName,
			SimpleName: respData.SimpleName,
			Fields:     make([]ResponseFieldMetadata, 0, len(respData.Fields)),
		}

		for _, field := range respData.Fields {
			resp.Fields = append(resp.Fields, ResponseFieldMetadata{
				Name:           field.Name,
				FieldName:      field.FieldName,
				JavaType:       field.JavaType,
				Description:    field.Description,
				Since:          field.Since,
				ResponseObject: field.ResponseObject,
			})
		}

		metadata.Responses[className] = resp
	}

	return nil
}

// parseEnumMetadata parses the enum metadata JSON file
func parseEnumMetadata(filePath string, metadata *Metadata) error {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return err
	}

	var enumsMap map[string]json.RawMessage
	if err := json.Unmarshal(data, &enumsMap); err != nil {
		return err
	}

	for className, rawData := range enumsMap {
		var enumData struct {
			ClassName   string `json:"className"`
			SimpleName  string `json:"simpleName"`
			PackageName string `json:"packageName"`
			IsNested    bool   `json:"isNested"`
			ParentClass string `json:"parentClass"`
			Values      []struct {
				Name    string                 `json:"name"`
				Ordinal int                    `json:"ordinal"`
				Extras  map[string]interface{} `json:"extras"`
			} `json:"values"`
		}

		if err := json.Unmarshal(rawData, &enumData); err != nil {
			return err
		}

		enum := EnumMetadata{
			ClassName:   enumData.ClassName,
			SimpleName:  enumData.SimpleName,
			PackageName: enumData.PackageName,
			IsNested:    enumData.IsNested,
			ParentClass: enumData.ParentClass,
			Values:      make([]EnumValueMetadata, 0, len(enumData.Values)),
		}

		for _, value := range enumData.Values {
			enum.Values = append(enum.Values, EnumValueMetadata{
				Name:    value.Name,
				Ordinal: value.Ordinal,
				Extras:  value.Extras,
			})
		}

		metadata.Enums[className] = enum
	}

	return nil
}

// extractCategory extracts the API category from a command name
func extractCategory(cmdName string) string {
	// Simple heuristic: look for common patterns like ListVPCsCmd or CreateNetworkCmd
	lowerName := strings.ToLower(cmdName)

	// Common CloudStack API categories
	categories := []string{
		"vpc", "vm", "network", "account", "cluster", "host", "pod", "storage",
		"template", "iso", "volume", "snapshot", "firewall", "loadbalancer",
		"vpn", "user", "project", "zone", "domain", "tag", "configuration",
		"event", "offering", "disk", "service", "resource", "systemvm", "router",
	}

	for _, category := range categories {
		if strings.Contains(lowerName, category) {
			return category
		}
	}

	// Default empty string if no match
	return ""
}
