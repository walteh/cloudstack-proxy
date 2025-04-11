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
	ClassName        string
	SimpleName       string
	Description      string
	Parameters       []ParameterMetadata
	ResponseName     string
	Category         string // E.g., "vpc", "vm", "network"
	HasAdminVariant  bool   // Whether this command has an admin variant
	AdminVariantName string // The name of the admin variant command, if any
	IsAsync          bool   // Whether this command is asynchronous
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
	ClassName       string
	SimpleName      string
	Description     string
	Fields          []ResponseFieldMetadata
	IsListResponse  bool   // Whether this is a list response
	ListElementTag  string // XML tag for list elements (e.g., "virtualmachine")
	ListElementName string // Name of the list element if specified explicitly
	ItemTypeName    string // Derived item type name for list responses
	ObjectName      string // The underlying object name
}

// ResponseFieldMetadata contains metadata for a field in a CloudStack API response
type ResponseFieldMetadata struct {
	Name           string
	FieldName      string
	JavaType       string
	Description    string
	Since          string
	ResponseObject string // Type of response object for this field
	IsDisplay      bool   // Whether this field is for display purposes
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
	// Read the file
	rawData, err := os.ReadFile(filePath)
	if err != nil {
		return fmt.Errorf("error reading command metadata file: %w", err)
	}

	// Parse the JSON
	var commandData map[string]struct {
		GroupName       string   `json:"groupName"`
		ScopeName       string   `json:"scopeName"`
		ScopedReplicaOf string   `json:"scopedReplicaOf"`
		ClassName       string   `json:"className"`
		SimpleName      string   `json:"simpleName"`
		CommandName     string   `json:"commandName"`
		Description     string   `json:"description"`
		ResponseObject  string   `json:"responseObject"`
		ResponseView    string   `json:"responseView"`
		EntityType      []string `json:"entityType"`
		Parameters      []struct {
			Name                 string                 `json:"name"`
			FieldName            string                 `json:"fieldName"`
			JavaType             string                 `json:"javaType"`
			CommandType          string                 `json:"commandType"`
			Description          string                 `json:"description"`
			Required             bool                   `json:"required"`
			Since                string                 `json:"since"`
			Validations          []interface{}          `json:"validations"`
			AcceptedOnAdminPort  bool                   `json:"acceptedOnAdminPort"`
			Expose               bool                   `json:"expose"`
			IncludeInApiDoc      bool                   `json:"includeInApiDoc"`
			Authorized           []string               `json:"authorized"`
			CollectionType       string                 `json:"collectionType"`
			Length               int                    `json:"length"`
			GetterMethod         interface{}            `json:"getterMethod"`
			IsCommaSeparatedList bool                   `json:"isCommaSeparatedList"`
			Extras               map[string]interface{} `json:"extras"`
		} `json:"parameters"`
		Superclasses             []string               `json:"superclasses"`
		IsAsync                  bool                   `json:"isAsync"`
		RequestHasSensitiveInfo  bool                   `json:"requestHasSensitiveInfo"`
		ResponseHasSensitiveInfo bool                   `json:"responseHasSensitiveInfo"`
		Extras                   map[string]interface{} `json:"extras"`
	}

	if err := json.Unmarshal(rawData, &commandData); err != nil {
		return fmt.Errorf("error parsing command metadata JSON: %w", err)
	}

	// Process each command
	for className, cmd := range commandData {
		// Skip commands that are not meant to be exposed via the API
		if len(cmd.Parameters) == 0 {
			continue
		}

		// Check if this is a scoped replica (admin version) of another command
		isAdminVariant := cmd.ScopedReplicaOf != ""

		// Extract base command name for admin variants
		baseCommand := ""
		if isAdminVariant {
			baseCommand = cmd.ScopedReplicaOf
		}

		// Extract category from command name or group
		category := ""
		if cmd.GroupName != "" {
			category = cmd.GroupName
		} else {
			category = extractCategory(cmd.SimpleName)
		}

		// Create CommandMetadata
		command := CommandMetadata{
			ClassName:        className,
			SimpleName:       cmd.SimpleName,
			Description:      cmd.Description,
			Parameters:       make([]ParameterMetadata, 0, len(cmd.Parameters)),
			ResponseName:     cmd.ResponseObject,
			Category:         category,
			HasAdminVariant:  isAdminVariant,
			AdminVariantName: baseCommand,
			IsAsync:          cmd.IsAsync,
		}

		// Process parameters
		for _, param := range cmd.Parameters {
			// Skip parameters that aren't meant to be exposed
			if !param.Expose || !param.IncludeInApiDoc {
				continue
			}

			parameter := ParameterMetadata{
				Name:        param.Name,
				FieldName:   param.FieldName,
				JavaType:    param.JavaType,
				Description: param.Description,
				Required:    param.Required,
				Since:       param.Since,
			}

			command.Parameters = append(command.Parameters, parameter)
		}

		metadata.Commands[className] = command

		// Add command to category map
		if _, ok := metadata.Categories[category]; !ok {
			metadata.Categories[category] = make([]string, 0)
		}
		metadata.Categories[category] = append(metadata.Categories[category], className)
	}

	return nil
}

// parseResponseMetadata parses the response metadata JSON file
func parseResponseMetadata(filePath string, metadata *Metadata) error {
	// Read the file
	rawData, err := os.ReadFile(filePath)
	if err != nil {
		return fmt.Errorf("error reading response metadata file: %w", err)
	}

	// Parse the JSON
	var responseData map[string]struct {
		ClassName           string `json:"className"`
		SimpleName          string `json:"simpleName"`
		ResponseName        string `json:"responseName"`
		ObjectName          string `json:"objectName"`
		Description         string `json:"description"`
		DefaultResponseType string `json:"defaultResponseType"`
		IsListResponse      bool   `json:"isListResponse"`
		ListElementTag      string `json:"listElementTag"`
		ListElementName     string `json:"listElementName"`
		IncludeInApiDoc     bool   `json:"includeInApiDoc"`
		Fields              []struct {
			Name           string                 `json:"name"`
			FieldName      string                 `json:"fieldName"`
			Type           string                 `json:"type"`
			Description    string                 `json:"description"`
			ResponseObject string                 `json:"responseObject"`
			Since          string                 `json:"since"`
			IsDisplay      bool                   `json:"isDisplay"`
			UsedForSorting bool                   `json:"usedForSorting"`
			Extras         map[string]interface{} `json:"extras"`
		} `json:"fields"`
	}

	if err := json.Unmarshal(rawData, &responseData); err != nil {
		return fmt.Errorf("error parsing response metadata JSON: %w", err)
	}

	// Process each response
	for className, resp := range responseData {
		// Skip irrelevant responses
		if !resp.IncludeInApiDoc {
			continue
		}

		// Create ResponseMetadata
		response := ResponseMetadata{
			ClassName:       className,
			SimpleName:      resp.SimpleName,
			Description:     resp.Description,
			IsListResponse:  resp.IsListResponse,
			ListElementTag:  resp.ListElementTag,
			ListElementName: resp.ListElementName,
			ObjectName:      resp.ObjectName,
			Fields:          make([]ResponseFieldMetadata, 0, len(resp.Fields)),
		}

		// Process fields
		for _, field := range resp.Fields {
			// Skip internal fields
			if strings.HasPrefix(field.Name, "_") {
				continue
			}

			responseField := ResponseFieldMetadata{
				Name:           field.Name,
				FieldName:      field.FieldName,
				JavaType:       field.Type,
				Description:    field.Description,
				Since:          field.Since,
				ResponseObject: field.ResponseObject,
				IsDisplay:      field.IsDisplay,
			}

			response.Fields = append(response.Fields, responseField)
		}

		// If this is a list response, try to determine the actual item type
		if resp.IsListResponse {
			// First check if we have ListElementName explicitly defined
			if resp.ListElementName != "" {
				response.ItemTypeName = resp.ListElementName
			} else if resp.ListElementTag != "" {
				// Try to infer from the list element tag
				// Convert tag like "virtualmachine" to "VirtualMachine"
				tagParts := strings.Split(resp.ListElementTag, "_")
				for i, part := range tagParts {
					if len(part) > 0 {
						tagParts[i] = strings.ToUpper(part[:1]) + part[1:]
					}
				}
				response.ItemTypeName = strings.Join(tagParts, "")
			} else {
				// As a fallback, try to infer from the response name
				// E.g., "ListVMsResponse" -> "VM"
				baseName := strings.TrimSuffix(resp.SimpleName, "Response")
				if strings.HasPrefix(baseName, "List") {
					itemName := strings.TrimPrefix(baseName, "List")
					if strings.HasSuffix(itemName, "s") && len(itemName) > 2 {
						response.ItemTypeName = itemName[:len(itemName)-1]
					} else {
						response.ItemTypeName = itemName
					}
				}
			}
		}

		metadata.Responses[className] = response
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
