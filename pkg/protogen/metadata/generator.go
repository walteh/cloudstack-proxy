package metadata

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"text/template"
	"unicode"
)

// ProtoGenerator generates protobuf definitions from CloudStack API metadata
type ProtoGenerator struct {
	Metadata *Metadata
	RootDir  string
}

// NewProtoGenerator creates a new protobuf generator
func NewProtoGenerator(metadata *Metadata, rootDir string) *ProtoGenerator {
	return &ProtoGenerator{
		Metadata: metadata,
		RootDir:  rootDir,
	}
}

// GenerateProto generates protobuf files for each API category
func (g *ProtoGenerator) GenerateProto() error {
	// Create category-based protobuf files
	for category, commands := range g.Metadata.Categories {
		if err := g.generateCategoryProto(category, commands); err != nil {
			return fmt.Errorf("error generating protobuf for category %s: %w", category, err)
		}
	}

	return nil
}

// generateCategoryProto generates a protobuf file for a specific API category
func (g *ProtoGenerator) generateCategoryProto(category string, commandNames []string) error {
	// Skip categories with no commands
	if len(commandNames) == 0 {
		return nil
	}

	// Create the category directory if it doesn't exist
	categoryDir := filepath.Join(g.RootDir, "proto", "cloudstack", category, "v1")
	if err := os.MkdirAll(categoryDir, 0755); err != nil {
		return fmt.Errorf("error creating directory: %w", err)
	}

	// Service name for this category (e.g., "VPCService" for vpc)
	serviceName := strings.ToUpper(category[:1]) + category[1:] + "Service"

	// Track admin variants to avoid duplicates
	adminVariants := make(map[string]string) // Maps base command to admin command

	// First pass: identify admin variants
	for _, cmdName := range commandNames {
		cmd, ok := g.Metadata.Commands[cmdName]
		if !ok {
			continue
		}

		simpleName := cmd.SimpleName
		if strings.HasSuffix(simpleName, "CmdByAdmin") {
			// Extract base command name (without "CmdByAdmin")
			baseCmd := strings.TrimSuffix(simpleName, "ByAdmin")
			adminVariants[baseCmd] = simpleName
		} else if strings.HasSuffix(simpleName, "AsAdmin") {
			// Extract base command name (without "AsAdmin")
			baseCmd := strings.TrimSuffix(simpleName, "AsAdmin")
			adminVariants[baseCmd] = simpleName
		} else if strings.HasSuffix(simpleName, "Admin") && !strings.HasSuffix(simpleName, "ByAdmin") {
			// Handle other admin variants
			baseCmd := strings.TrimSuffix(simpleName, "Admin")
			adminVariants[baseCmd] = simpleName
		}
	}

	// Collect commands and responses for this category
	var serviceCommands []CommandMetadata
	responseNames := make(map[string]bool)
	processedCommands := make(map[string]bool)

	for _, cmdName := range commandNames {
		cmd, ok := g.Metadata.Commands[cmdName]
		if !ok {
			continue
		}

		// Skip admin variants - they'll be handled with their base commands
		simpleName := cmd.SimpleName
		if isAdminVariant(simpleName) {
			// Check if we already processed the base command
			baseCmd := getBaseCommandName(simpleName)
			if processedCommands[baseCmd] {
				continue
			}
		}

		// Check if this command has an admin variant
		hasAdminVariant := false
		adminVariantName := ""
		if admin, ok := adminVariants[simpleName]; ok {
			hasAdminVariant = true
			adminVariantName = admin
		}

		// Clone and modify command for processing
		cmdCopy := cmd

		// Remove "Cmd" from the name for a cleaner RPC name
		if strings.HasSuffix(cmdCopy.SimpleName, "Cmd") {
			cmdCopy.SimpleName = strings.TrimSuffix(cmdCopy.SimpleName, "Cmd")
		}

		// Track that we've processed this command
		processedCommands[cmd.SimpleName] = true

		// Set admin variant info
		cmdCopy.HasAdminVariant = hasAdminVariant
		cmdCopy.AdminVariantName = adminVariantName

		serviceCommands = append(serviceCommands, cmdCopy)

		// Track needed responses
		if cmdCopy.ResponseName != "" {
			responseNames[cmdCopy.ResponseName] = true
		}

		// Also track admin variant response if it exists
		if hasAdminVariant {
			if adminCmd, ok := g.Metadata.Commands[adminVariantName]; ok && adminCmd.ResponseName != "" {
				responseNames[adminCmd.ResponseName] = true
			}
		}
	}

	// Skip if no commands were found
	if len(serviceCommands) == 0 {
		return nil
	}

	// Extract responses needed for this service
	var serviceResponses []ResponseMetadata
	for respName := range responseNames {
		resp, ok := g.Metadata.Responses[respName]
		if !ok {
			continue
		}
		serviceResponses = append(serviceResponses, resp)
	}

	// Generate the protobuf file
	protoPath := filepath.Join(categoryDir, category+".gen.proto")
	return g.writeProtoFile(protoPath, serviceName, serviceCommands, serviceResponses)
}

// isAdminVariant determines if a command is an admin variant
func isAdminVariant(name string) bool {
	return strings.HasSuffix(name, "CmdByAdmin") ||
		strings.HasSuffix(name, "AsAdmin") ||
		(strings.HasSuffix(name, "Admin") && !strings.HasSuffix(name, "ByAdmin"))
}

// getBaseCommandName extracts the base command name from an admin variant
func getBaseCommandName(name string) string {
	if strings.HasSuffix(name, "CmdByAdmin") {
		return strings.TrimSuffix(name, "ByAdmin")
	} else if strings.HasSuffix(name, "AsAdmin") {
		return strings.TrimSuffix(name, "AsAdmin")
	} else if strings.HasSuffix(name, "Admin") && !strings.HasSuffix(name, "ByAdmin") {
		return strings.TrimSuffix(name, "Admin")
	}
	return name
}

// writeProtoFile writes a protobuf file with the given service and message definitions
func (g *ProtoGenerator) writeProtoFile(filePath, serviceName string, commands []CommandMetadata, responses []ResponseMetadata) error {
	// Template for the protobuf file
	const protoTmpl = `edition = "2023";

package cloudstack.{{ .Category }}.v1;

import "cloudstack/validate/validate.proto";
import "google/protobuf/descriptor.proto";
import "cloudstack/annotations/annotations.proto";

// {{ .ServiceName }} provides operations for managing {{ .Category | title }}s
service {{ .ServiceName }} {
{{- if hasAdminService . }}
	option (cloudstack.annotations.Enum.scope) = Scope_ADMIN;
{{- else }}
	option (cloudstack.annotations.Enum.scope) = Scope_USER;
{{- end }}
{{- range .Commands }}
	// {{ .SimpleName }} {{ .Description }}
	rpc {{ .SimpleName }}({{ .SimpleName }}Request) returns ({{ .SimpleName }}Response) {
	{{- if .HasAdminVariant }}
		option (cloudstack.annotations.Enum.scope) = Scope_ADMIN;
	{{- end }}
	}
{{- end }}
}

{{ range .Commands }}
// {{ .SimpleName }}Request represents the parameters for {{ .Description | lower }}
message {{ .SimpleName }}Request {
{{- if hasAdminVariant . }}
	// Whether to run this operation as an administrator
	bool run_as_admin = 1;

{{- range $index, $param := .Parameters }}
	// {{ $param.Description }}
	{{ javaTypeToProto $param.JavaType }} {{ $param.FieldName | toSnakeCase }} = {{ add $index 2 }}{{ getValidationRules $param $.Parameters }};
{{- end }}
{{- else }}
{{- range $index, $param := .Parameters }}
	// {{ $param.Description }}
	{{ javaTypeToProto $param.JavaType }} {{ $param.FieldName | toSnakeCase }} = {{ add $index 1 }}{{ getValidationRules $param $.Parameters }};
{{- end }}
{{- end }}
}

// {{ .SimpleName }}Response represents the response from {{ .Description | lower }}
message {{ .SimpleName }}Response {
{{- if .ResponseList }}
	option (validate.message).disabled = false;
	// The list of {{ .ResponseItemName }}s
	repeated {{ if eq .ResponseItemName "Item" }}Item{{ else }}{{ .ResponseItemName }}{{ end }} items = 1;

	// The total count of {{ .ResponseItemName }}s
	int32 total_count = 2;
{{- else if .ResponseItemName }}
	// The Result
	Result result = 1;
{{- else }}
	// Whether the operation was successful
	bool success = 1;
	
	// Any text associated with the success or failure
	string display_text = 2;
{{- end }}
}
{{ end }}

{{ range .Responses }}
// {{ .MessageName }} represents a {{ .Description | title }}
message {{ .MessageName }} {
{{- range $index, $field := .Fields }}
	// {{ $field.Description }}
	{{ javaTypeToProto $field.JavaType }} {{ $field.FieldName | toSnakeCase }} = {{ add $index 1 }}{{ getFieldValidation $field }};
{{- end }}
}
{{ end }}

// Item represents a generic item in a list response
message Item {
	// The ID of the item
	string id = 1 [(validate.field).string.uuid = true];
	
	// The name of the item
	string name = 2;
	
	// The display name of the item
	string display_name = 3;
	
	// The description of the item
	string description = 4;
	
	// The account ID the item belongs to
	string account_id = 5 [(validate.field).string.uuid = true];
	
	// The domain ID the item belongs to
	string domain_id = 6 [(validate.field).string.uuid = true];
	
	// The domain name the item belongs to
	string domain = 7;
	
	// The project ID the item belongs to
	string project_id = 8 [(validate.field).string.uuid = true];
	
	// The project name the item belongs to
	string project = 9;
	
	// The date the item was created
	string created = 10;
	
	// The state of the item
	string state = 11;
	
	// Additional fields returned by the API
	map<string, string> details = 12;
}

// Result represents a generic operation result
message Result {
	// Whether the operation was successful
	bool success = 1;
	
	// Any text associated with the success or failure
	string display_text = 2;
	
	// The ID of the resource affected by the operation
	string id = 3 [(validate.field).string.uuid = true];
	
	// The job ID for an async operation
	string job_id = 4 [(validate.field).string.uuid = true];
	
	// The status of the job
	string job_status = 5;
}
`

	// Format the category name for the template
	category := filepath.Base(filepath.Dir(filepath.Dir(filePath)))

	// Get all command details
	var templateCommands []map[string]interface{}
	for _, cmd := range commands {
		// Determine if this is a "List" command (returns multiple items)
		isList := strings.HasPrefix(cmd.SimpleName, "List")

		// Get the response item type (for lists, this is the item type)
		responseItemName := ""

		if cmd.ResponseName != "" {
			// For list commands, get the item type from the response object
			if isList {
				// Extract the base type from the response object
				// e.g., org.apache.cloudstack.api.response.AccountResponse -> Account
				parts := strings.Split(cmd.ResponseName, ".")
				responseName := parts[len(parts)-1]

				// Remove "Response" suffix
				if strings.HasSuffix(responseName, "Response") {
					baseTypeName := strings.TrimSuffix(responseName, "Response")

					// Use the base type name (e.g., Account) as the item type
					responseItemName = baseTypeName
				} else {
					// Fallback if the response name doesn't match expected pattern
					responseItemName = "Item"
				}
			} else if resp, ok := g.Metadata.Responses[cmd.ResponseName]; ok {
				// For non-list responses, just use the response name without "Response"
				responseItemName = strings.TrimSuffix(resp.SimpleName, "Response")
				if responseItemName == "" {
					responseItemName = "Result"
				}
			} else {
				responseItemName = "Result"
			}
		} else {
			// Default item name if response metadata not found
			if isList {
				responseItemName = "Item"
			} else {
				responseItemName = "Result"
			}
		}

		templateCommands = append(templateCommands, map[string]interface{}{
			"SimpleName":       cmd.SimpleName,
			"Description":      cmd.Description,
			"Parameters":       cmd.Parameters,
			"ResponseList":     isList,
			"ResponseItemName": responseItemName,
			"ResponseName":     cmd.ResponseName, // Add the full response name for reference
			"HasAdminVariant":  cmd.HasAdminVariant,
			"AdminVariantName": cmd.AdminVariantName,
			"IsAsync":          cmd.IsAsync,
		})
	}

	// Transform response data for template
	var templateResponses []map[string]interface{}

	// First, collect all response types from the Metadata
	for _, resp := range g.Metadata.Responses {
		// Only include responses that have meaningful fields and aren't just collection wrappers
		if len(resp.Fields) > 0 && (!resp.IsListResponse || resp.SimpleName != "ListResponse") {
			// Determine the message name
			messageName := resp.SimpleName
			if strings.HasSuffix(messageName, "Response") {
				messageName = strings.TrimSuffix(messageName, "Response")
			}

			// Add the response type to our template data
			templateResponses = append(templateResponses, map[string]interface{}{
				"MessageName": messageName,
				"Description": getResponseDescription(messageName),
				"Fields":      resp.Fields,
			})
		}
	}

	// Next, collect all item types that are needed for list responses
	// Create a set to track which item types we've already added
	addedItemTypes := make(map[string]bool)

	// Go through all commands and check for list responses
	for _, cmd := range templateCommands {
		if cmd["ResponseList"].(bool) && cmd["ResponseItemName"].(string) != "Item" {
			itemTypeName := cmd["ResponseItemName"].(string)

			// Skip if we've already added this item type
			if _, exists := addedItemTypes[itemTypeName]; exists {
				continue
			}

			// Mark this type as added
			addedItemTypes[itemTypeName] = true

			// Try to find the response metadata for this item type
			foundResponseFields := false

			// First, check if this is a direct match with an existing response type
			for _, resp := range templateResponses {
				if resp["MessageName"] == itemTypeName {
					foundResponseFields = true
					break
				}
			}

			// If we didn't find a direct match, look through all responses for a match
			if !foundResponseFields {
				responseFullName := cmd["ResponseName"].(string)
				if responseFullName != "" {
					// Extract just the simple name part of the fully qualified class name
					parts := strings.Split(responseFullName, ".")
					// responseSimpleName not currently used, but keeping for future reference
					_ = parts[len(parts)-1]

					// Check if we have metadata for this response
					if resp, ok := g.Metadata.Responses[responseFullName]; ok && len(resp.Fields) > 0 {
						// Add this as a message type
						templateResponses = append(templateResponses, map[string]interface{}{
							"MessageName": itemTypeName,
							"Description": getResponseDescription(itemTypeName),
							"Fields":      resp.Fields,
						})
						foundResponseFields = true
					}
				}
			}

			// If we still can't find it, create a generic item type with standard fields
			if !foundResponseFields {
				templateResponses = append(templateResponses, map[string]interface{}{
					"MessageName": itemTypeName,
					"Description": fmt.Sprintf("%s item", itemTypeName),
					"Fields": []ResponseFieldMetadata{
						{
							Name:        "id",
							FieldName:   "id",
							JavaType:    "java.lang.String",
							Description: fmt.Sprintf("The ID of the %s", itemTypeName),
						},
						{
							Name:        "name",
							FieldName:   "name",
							JavaType:    "java.lang.String",
							Description: fmt.Sprintf("The name of the %s", itemTypeName),
						},
						{
							Name:        "displayname",
							FieldName:   "displayName",
							JavaType:    "java.lang.String",
							Description: fmt.Sprintf("The display name of the %s", itemTypeName),
						},
						{
							Name:        "description",
							FieldName:   "description",
							JavaType:    "java.lang.String",
							Description: fmt.Sprintf("The description of the %s", itemTypeName),
						},
						{
							Name:        "created",
							FieldName:   "created",
							JavaType:    "java.lang.String",
							Description: "The date this entity was created",
						},
					},
				})
			}
		}
	}

	// Add Success response for basic operations
	hasSuccessResponse := false
	for _, resp := range templateResponses {
		if resp["MessageName"] == "Success" {
			hasSuccessResponse = true
			break
		}
	}

	if !hasSuccessResponse {
		successResponse := map[string]interface{}{
			"MessageName": "Success",
			"Description": "Success operation response",
			"Fields": []ResponseFieldMetadata{
				{
					Name:        "success",
					FieldName:   "success",
					JavaType:    "java.lang.Boolean",
					Description: "true if operation is executed successfully",
				},
				{
					Name:        "displaytext",
					FieldName:   "displayText",
					JavaType:    "java.lang.String",
					Description: "any text associated with the success or failure",
				},
			},
		}
		templateResponses = append(templateResponses, successResponse)
	}

	// Prepare template data
	templateData := map[string]interface{}{
		"Category":    category,
		"ServiceName": serviceName,
		"Commands":    templateCommands,
		"Responses":   templateResponses,
	}

	// Create template with functions
	tmpl := template.New("proto").Funcs(template.FuncMap{
		"add": func(i, j int) int {
			return i + j
		},
		"lower": strings.ToLower,
		"title": strings.Title,
		"camelCase": func(s string) string {
			return toCamelCase(s)
		},
		"toSnakeCase": func(s string) string {
			return toSnakeCase(s)
		},
		"plural": func(s string) string {
			return pluralize(s)
		},
		"javaTypeToProto": func(javaType string) string {
			return javaTypeToProto(javaType)
		},
		"getValidationRules": func(param ParameterMetadata, allParams []ParameterMetadata) string {
			return getValidationRules(param, allParams)
		},
		"getFieldValidation": func(field ResponseFieldMetadata) string {
			return getFieldValidation(field)
		},
		"hasAdminVariant": func(cmd interface{}) bool {
			if m, ok := cmd.(map[string]interface{}); ok {
				if hasAdmin, ok := m["HasAdminVariant"].(bool); ok {
					return hasAdmin
				}
			}
			return false
		},
		"hasAdminService": func(data interface{}) bool {
			// Check if this is a template data map
			dataMap, ok := data.(map[string]interface{})
			if !ok {
				return false
			}

			// Get the service name and check if it contains "Admin"
			if serviceName, ok := dataMap["ServiceName"].(string); ok {
				if strings.Contains(strings.ToLower(serviceName), "admin") {
					return true
				}
			}

			// Check if the service category is typically admin-only
			if category, ok := dataMap["Category"].(string); ok {
				// List of categories that are typically admin-only
				adminCategories := map[string]bool{
					"host":    true,
					"cluster": true,
					"pod":     true,
					"storage": true,
				}
				if adminCategories[category] {
					return true
				}
			}

			// Check the commands to see if the majority are admin variants
			commands, ok := dataMap["Commands"].([]map[string]interface{})
			if !ok {
				return false
			}

			// Count admin variants
			adminCount := 0
			for _, cmd := range commands {
				if hasAdmin, ok := cmd["HasAdminVariant"].(bool); ok && hasAdmin {
					adminCount++
				}
			}

			// If more than half the commands are admin variants, mark as admin service
			return adminCount > len(commands)/2
		},
	})

	parsed, err := tmpl.Parse(protoTmpl)
	if err != nil {
		return fmt.Errorf("error parsing template: %w", err)
	}

	// Create the file
	file, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("error creating file: %w", err)
	}
	defer file.Close()

	// Execute the template
	if err := parsed.Execute(file, templateData); err != nil {
		return fmt.Errorf("error executing template: %w", err)
	}

	fmt.Printf("Generated protobuf file: %s\n", filePath)
	return nil
}

// Helper functions for the template

// getResponseDescription returns a description for a response type
func getResponseDescription(name string) string {
	return fmt.Sprintf("%s object", name)
}

// toCamelCase converts a snake_case string to camelCase
func toCamelCase(s string) string {
	parts := strings.Split(s, "_")
	for i := 1; i < len(parts); i++ {
		if len(parts[i]) > 0 {
			parts[i] = strings.ToUpper(parts[i][:1]) + parts[i][1:]
		}
	}
	return strings.Join(parts, "")
}

// toSnakeCase converts a camelCase or PascalCase string to snake_case
func toSnakeCase(s string) string {
	// If already contains underscore, assume it's already snake_case
	if strings.Contains(s, "_") {
		return strings.ToLower(s)
	}

	// Handle special case of IDs (e.g., "vpcId" -> "vpc_id", "ID" -> "id")
	s = strings.ReplaceAll(s, "ID", "Id")

	var result []rune
	for i, r := range s {
		if i > 0 && unicode.IsUpper(r) {
			// Add underscore before uppercase letters
			result = append(result, '_')
		}
		result = append(result, unicode.ToLower(r))
	}

	return string(result)
}

// pluralize adds an 's' to the end of a string, or handles special cases
func pluralize(s string) string {
	// Common special cases for proper pluralization
	switch {
	case strings.HasSuffix(s, "y"):
		return strings.TrimSuffix(s, "y") + "ies"
	case strings.HasSuffix(s, "s") || strings.HasSuffix(s, "x") ||
		strings.HasSuffix(s, "z") || strings.HasSuffix(s, "ch") ||
		strings.HasSuffix(s, "sh"):
		return s + "es"
	default:
		return s + "s"
	}
}

// javaTypeToProto converts a Java type to a protobuf type
func javaTypeToProto(javaType string) string {
	switch {
	case javaType == "java.lang.String":
		return "string"
	case javaType == "java.lang.Boolean" || javaType == "boolean":
		return "bool"
	case javaType == "java.lang.Integer" || javaType == "int":
		return "int32"
	case javaType == "java.lang.Long" || javaType == "long":
		return "int64"
	case javaType == "java.lang.Float" || javaType == "float":
		return "float"
	case javaType == "java.lang.Double" || javaType == "double":
		return "double"
	case strings.Contains(javaType, "Map"):
		// Try to extract key and value types
		genericTypes := extractMapTypes(javaType)
		if len(genericTypes) == 2 {
			keyType := javaTypeToProto(genericTypes[0])
			valType := javaTypeToProto(genericTypes[1])
			return fmt.Sprintf("map<%s, %s>", keyType, valType)
		}
		return "map<string, string>" // Default map type
	case strings.Contains(javaType, "List"):
		// If we can extract the generic type, use it
		genericType := extractGenericType(javaType)
		if genericType != "" {
			return "repeated " + javaTypeToProto(genericType)
		}
		return "repeated string" // Default to string for now
	case strings.Contains(javaType, "Date"):
		// Handle Java Date type
		return "string" // Use string for dates
	default:
		return "string" // Default to string for unknown types
	}
}

// extractMapTypes tries to extract the key and value types from Java Map generic
func extractMapTypes(javaType string) []string {
	start := strings.Index(javaType, "<")
	end := strings.LastIndex(javaType, ">")

	if start > 0 && end > start {
		content := javaType[start+1 : end]
		types := strings.Split(content, ",")

		if len(types) == 2 {
			return []string{
				strings.TrimSpace(types[0]),
				strings.TrimSpace(types[1]),
			}
		}
	}

	return nil
}

// extractGenericType tries to extract the generic type from Java generics
func extractGenericType(javaType string) string {
	start := strings.Index(javaType, "<")
	end := strings.LastIndex(javaType, ">")

	if start > 0 && end > start {
		return javaType[start+1 : end]
	}

	return ""
}

// getValidationRules generates validation rules for a parameter based on its type and name
func getValidationRules(param ParameterMetadata, allParams []ParameterMetadata) string {
	var rules []string

	// Add required validation if parameter is required
	if param.Required {
		rules = append(rules, "(validate.field).required = true")
	}

	// Add specific validations based on parameter name and type
	fieldName := param.FieldName
	lowerName := strings.ToLower(param.Name)
	javaType := param.JavaType

	// UUID validation for IDs
	if strings.Contains(lowerName, "id") && !strings.Contains(lowerName, "cidr") &&
		!strings.Contains(lowerName, "valid") && javaType == "java.lang.String" {
		rules = append(rules, "(validate.field).string.uuid = true")
	}

	// IP address validations
	if (strings.Contains(lowerName, "ip") || strings.Contains(lowerName, "gateway")) &&
		!strings.Contains(lowerName, "script") && javaType == "java.lang.String" {
		if strings.Contains(lowerName, "cidr") {
			rules = append(rules, "(validate.field).string.ipv4_prefix = true")
		} else if strings.Contains(lowerName, "ipv6") {
			rules = append(rules, "(validate.field).string.ipv6 = true")
		} else {
			rules = append(rules, "(validate.field).string.ipv4 = true")
		}
	}

	// String length validations for names and descriptions
	if (strings.Contains(lowerName, "name") || strings.Contains(lowerName, "description")) &&
		javaType == "java.lang.String" {
		rules = append(rules, "(validate.field).string.min_len = 1")

		if strings.Contains(lowerName, "description") {
			rules = append(rules, "(validate.field).string.max_len = 1024")
		} else {
			rules = append(rules, "(validate.field).string.max_len = 255")
		}
	}

	// Numeric range validations
	if strings.Contains(javaType, "Integer") || strings.Contains(javaType, "int") {
		if strings.Contains(lowerName, "port") {
			rules = append(rules, "(validate.field).int32.gte = 1")
			rules = append(rules, "(validate.field).int32.lte = 65535")
		}
	}

	// Boolean field formatting
	if javaType == "java.lang.Boolean" || javaType == "boolean" {
		// For booleans, use the features.field_presence to make it explicit
		rules = append(rules, "features.field_presence = EXPLICIT")
	}

	// Check for "Must be used with the domainId parameter" in description
	if strings.Contains(param.Description, "Must be used with the domainId parameter") ||
		strings.Contains(param.Description, "Must be used with the domainId") {
		// Add CEL validation to verify domainId is set when this parameter is set
		snakeName := toSnakeCase(fieldName)
		rules = append(rules, fmt.Sprintf("(validate.field).cel = {\n\t\t\tid: \"%s_with_domain_id\",\n\t\t\texpression: \"!has(%s) || has(domain_id)\",\n\t\t\tmessage: \"%s must be used with domain_id parameter\"\n\t\t}", snakeName, snakeName, snakeName))
	}

	// Format the rules if any exist
	if len(rules) > 0 {
		return " [\n\t\t" + strings.Join(rules, ",\n\t\t") + "\n\t]"
	}

	return ""
}

// getFieldValidation generates validation rules for a response field
func getFieldValidation(field ResponseFieldMetadata) string {
	var rules []string

	// Add validations based on field name and type
	lowerName := strings.ToLower(field.Name)
	javaType := field.JavaType

	// UUID validation for IDs
	if strings.Contains(lowerName, "id") && !strings.Contains(lowerName, "cidr") &&
		!strings.Contains(lowerName, "valid") && javaType == "java.lang.String" {
		rules = append(rules, "(validate.field).string.uuid = true")
	}

	// IP address validations
	if (strings.Contains(lowerName, "ip") || strings.Contains(lowerName, "gateway")) &&
		!strings.Contains(lowerName, "script") && javaType == "java.lang.String" {
		if strings.Contains(lowerName, "cidr") {
			rules = append(rules, "(validate.field).string.ipv4_prefix = true")
		} else if strings.Contains(lowerName, "ipv6") {
			rules = append(rules, "(validate.field).string.ipv6 = true")
		} else {
			rules = append(rules, "(validate.field).string.ipv4 = true")
		}
	}

	// Format the rules if any exist
	if len(rules) > 0 {
		return " [\n\t\t" + strings.Join(rules, ",\n\t\t") + "\n\t]"
	}

	return ""
}
