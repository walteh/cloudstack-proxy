package metadata

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"text/template"
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

	// Collect commands and responses for this category
	var serviceCommands []CommandMetadata
	responseNames := make(map[string]bool)

	for _, cmdName := range commandNames {
		cmd, ok := g.Metadata.Commands[cmdName]
		if !ok {
			continue
		}

		serviceCommands = append(serviceCommands, cmd)

		// Track needed responses
		if cmd.ResponseName != "" {
			responseNames[cmd.ResponseName] = true
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

// writeProtoFile writes a protobuf file with the given service and message definitions
func (g *ProtoGenerator) writeProtoFile(filePath, serviceName string, commands []CommandMetadata, responses []ResponseMetadata) error {
	// Template for the protobuf file
	const protoTmpl = `edition = "2023";

package cloudstack.{{ .Category }}.v1;

import "cloudstack/annotations/annotations.proto";
import "cloudstack/validate/validate.proto";
import "google/protobuf/descriptor.proto";

// {{ .ServiceName }} provides operations for managing {{ .Category | title }}s
service {{ .ServiceName }} {
{{- range .Commands }}
	// {{ .SimpleName }} {{ .Description }}
	rpc {{ .SimpleName }}({{ .SimpleName }}Request) returns ({{ .SimpleName }}Response) {}
{{- end }}
}

{{ range .Commands }}
// {{ .SimpleName }}Request represents the parameters for {{ .Description | lower }}
message {{ .SimpleName }}Request {
{{- range $index, $param := .Parameters }}
	// {{ $param.Description }}
	{{ javaTypeToProto $param.JavaType }} {{ $param.Name | camelCase }} = {{ add $index 1 }}{{ if $param.Required }} [
		(validate.field).required = true
	]{{ end }};
{{- end }}
}

// {{ .SimpleName }}Response represents the response from {{ .Description | lower }}
message {{ .SimpleName }}Response {
{{- if .ResponseList }}
	// The list of {{ .ResponseItemName }}s
	repeated {{ .ResponseItemName }} {{ .ResponseItemName | camelCase | plural }} = 1;

	// The total count of {{ .ResponseItemName }}s
	int32 total_count = 2;
{{- else }}
	// The {{ .ResponseItemName }}
	{{ .ResponseItemName }} {{ .ResponseItemName | camelCase }} = 1;
{{- end }}
}
{{ end }}

{{ range .Responses }}
// {{ .MessageName }} represents a {{ .Description | title }}
message {{ .MessageName }} {
{{- range $index, $field := .Fields }}
	// {{ $field.Description }}
	{{ javaTypeToProto $field.JavaType }} {{ $field.Name | camelCase }} = {{ add $index 1 }};
{{- end }}
}
{{ end }}
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
		if resp, ok := g.Metadata.Responses[cmd.ResponseName]; ok {
			responseItemName = resp.SimpleName
			// For list responses, try to determine the actual item type
			if isList && strings.HasSuffix(responseItemName, "sResponse") {
				// Extract the item type from the response name (e.g., ListVPCsResponse -> VPC)
				responseItemName = strings.TrimSuffix(responseItemName, "sResponse")
			} else if strings.HasSuffix(responseItemName, "Response") {
				// For non-list responses, just trim "Response"
				responseItemName = strings.TrimSuffix(responseItemName, "Response")
			}
		}

		templateCommands = append(templateCommands, map[string]interface{}{
			"SimpleName":       cmd.SimpleName,
			"Description":      cmd.Description,
			"Parameters":       cmd.Parameters,
			"ResponseList":     isList,
			"ResponseItemName": responseItemName,
		})
	}

	// Transform response data for template
	var templateResponses []map[string]interface{}
	for _, resp := range responses {
		// Skip responses that are just containers for lists
		if strings.HasSuffix(resp.SimpleName, "sResponse") && strings.HasPrefix(resp.SimpleName, "List") {
			continue
		}

		// For item responses in a list, use the base name without "Response"
		messageName := resp.SimpleName
		if strings.HasSuffix(messageName, "Response") {
			messageName = strings.TrimSuffix(messageName, "Response")
		}

		// Build response for template
		templateResponses = append(templateResponses, map[string]interface{}{
			"MessageName": messageName,
			"Description": getResponseDescription(messageName),
			"Fields":      resp.Fields,
		})
	}

	// Prepare template data
	templateData := map[string]interface{}{
		"Category":    category,
		"ServiceName": serviceName,
		"Commands":    templateCommands,
		"Responses":   templateResponses,
	}

	// Create template with functions
	tmpl := template.New("proto")
	tmpl.Funcs(template.FuncMap{
		"add": func(i, j int) int {
			return i + j
		},
		"lower": strings.ToLower,
		"title": strings.Title,
		"camelCase": func(s string) string {
			return toCamelCase(s)
		},
		"plural": func(s string) string {
			return pluralize(s)
		},
		"javaTypeToProto": func(javaType string) string {
			return javaTypeToProto(javaType)
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

// pluralize adds an 's' to the end of a string, or handles special cases
func pluralize(s string) string {
	// Add specific plural cases here if needed
	return s + "s"
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
	case strings.Contains(javaType, "List"):
		// Handle lists - extract the generic type if possible
		return "repeated " + javaTypeToProto("java.lang.String") // Default to string for now
	default:
		return "string" // Default to string for unknown types
	}
}
