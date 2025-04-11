package metadata

import (
	"fmt"
	"regexp"
	"strings"
	"unicode"
)

// sanitizeEnumValue ensures that an enum value is a valid protobuf identifier
func sanitizeEnumValue(value string) string {
	// Replace special characters with underscore
	re := regexp.MustCompile(`[^a-zA-Z0-9_]`)
	sanitized := re.ReplaceAllString(value, "_")

	// Make sure it starts with a letter
	if len(sanitized) > 0 && !unicode.IsLetter(rune(sanitized[0])) {
		sanitized = "V_" + sanitized
	}

	// Make all uppercase for enum convention
	return strings.ToUpper(sanitized)
}

// identifyEnumsForService determines which enums should be included in the service proto file
func (g *ProtoGenerator) identifyEnumsForService(commands []CommandMetadata, responses []ResponseMetadata, packagePath string) []map[string]interface{} {
	enums := []map[string]interface{}{}

	// Track which enums we've already added
	addedEnums := make(map[string]bool)

	// Helper function to add an enum from Metadata.Enums if it exists
	addEnumIfExists := func(enumName string, description string) bool {
		// First check if we've already added this enum
		if addedEnums[enumName] {
			return true
		}

		// Look through all Java enum types to find one that matches this name
		for _, enumInfo := range g.Metadata.Enums {
			if enumInfo.SimpleName == enumName {
				// Extract enum values from the metadata
				values := []map[string]string{}
				for _, value := range enumInfo.Values {
					desc := value.Name
					// Extract a description from extras if available
					if nameDesc, ok := value.Extras["getName"].(string); ok {
						desc = nameDesc
					}

					// Sanitize the enum value name for protobuf
					// Remove special characters and ensure it's a valid identifier
					sanitizedName := sanitizeEnumValue(value.Name)

					// Convert Name to string to ensure it's valid for snakeAllCaps
					values = append(values, map[string]string{
						"Name":        sanitizedName,
						"Description": desc,
					})
				}

				if len(values) > 0 {
					// Ensure Name is a string
					enumDef := map[string]interface{}{
						"Name":        enumName,
						"Description": description,
						"Values":      values,
					}

					// Debug
					fmt.Printf("Adding enum %s with %d values\n", enumName, len(values))

					enums = append(enums, enumDef)
					addedEnums[enumName] = true
					return true
				}
				break
			}
		}

		return false
	}

	// First pass: Collect all Java enum types used in parameters and response fields

	// Check command parameters
	for _, command := range commands {
		for _, param := range command.Parameters {
			// Check if the parameter uses an enum type
			if strings.HasPrefix(param.JavaType, "com.cloud.") ||
				strings.HasPrefix(param.JavaType, "org.apache.cloudstack.") {
				// Extract the simple name of the enum class
				parts := strings.Split(param.JavaType, ".")
				enumSimpleName := parts[len(parts)-1]

				// If it looks like an enum type, add it
				if strings.Contains(enumSimpleName, "Type") ||
					strings.Contains(enumSimpleName, "State") ||
					strings.Contains(enumSimpleName, "Status") ||
					strings.Contains(enumSimpleName, "Scope") ||
					strings.Contains(enumSimpleName, "Mode") ||
					strings.Contains(enumSimpleName, "Level") {
					addEnumIfExists(enumSimpleName, param.Description)
				}
			}

			// Also check for enum patterns in the description
			// lowerParamName := strings.ToLower(param.FieldName)

			// Special case for fields with known enum values mentioned in description
			if strings.Contains(param.Description, "valid values are") ||
				strings.Contains(param.Description, "Valid values are") ||
				strings.Contains(param.Description, "can be:") ||
				strings.Contains(param.Description, "allowed values") {
				// Extract an enum name from the parameter name
				enumName := toTitleCase(param.FieldName) + "Type"
				if strings.HasSuffix(enumName, "TypeType") {
					enumName = strings.TrimSuffix(enumName, "Type")
				}
				if !addedEnums[enumName] {
					enumDefinition := createEnumDefinitionFromDescription(enumName, param.Description)
					if len(enumDefinition["Values"].([]map[string]string)) > 0 {
						enums = append(enums, enumDefinition)
						addedEnums[enumName] = true
					}
				}
			}

			// // Check for common enum field names
			// if lowerParamName == "state" && !addedEnums["State"] {
			// 	if !addEnumIfExists("State", "the state of the entity") {
			// 		enums = append(enums, createEnumDefinition("State", "the state of the entity"))
			// 		addedEnums["State"] = true
			// 	}
			// } else if strings.Contains(lowerParamName, "permission") && !addedEnums["Permission"] {
			// 	if !addEnumIfExists("Permission", "permission level") {
			// 		enums = append(enums, createEnumDefinition("Permission", "permission level"))
			// 		addedEnums["Permission"] = true
			// 	}
			// } else if lowerParamName == "protocol" && !addedEnums["Protocol"] {
			// 	if !addEnumIfExists("Protocol", "network protocol") {
			// 		enums = append(enums, createEnumDefinition("Protocol", "network protocol"))
			// 		addedEnums["Protocol"] = true
			// 	}
			// } else if strings.Contains(lowerParamName, "interval_type") && !addedEnums["IntervalType"] {
			// 	if !addEnumIfExists("IntervalType", "interval type") {
			// 		enums = append(enums, createEnumDefinition("IntervalType", "interval type"))
			// 		addedEnums["IntervalType"] = true
			// 	}
			// }
		}
	}

	// Check response fields
	for _, response := range responses {
		for _, field := range response.Fields {
			// Check if the field uses an enum type
			if strings.HasPrefix(field.JavaType, "com.cloud.") ||
				strings.HasPrefix(field.JavaType, "org.apache.cloudstack.") {
				// Extract the simple name of the enum class
				parts := strings.Split(field.JavaType, ".")
				enumSimpleName := parts[len(parts)-1]

				// If it looks like an enum type, add it
				if strings.Contains(enumSimpleName, "Type") ||
					strings.Contains(enumSimpleName, "State") ||
					strings.Contains(enumSimpleName, "Status") ||
					strings.Contains(enumSimpleName, "Scope") ||
					strings.Contains(enumSimpleName, "Mode") ||
					strings.Contains(enumSimpleName, "Level") {
					addEnumIfExists(enumSimpleName, field.Description)
				}
			}
		}
	}

	return enums
}

// createEnumDefinition creates a basic enum definition with common values

// createEnumDefinitionFromDescription extracts enum values from a field description
func createEnumDefinitionFromDescription(name string, description string) map[string]interface{} {
	enumValues := []map[string]string{}

	// Look for common patterns that indicate enum values in descriptions
	valuePart := ""
	if strings.Contains(description, "valid values are") {
		valuePart = strings.Split(description, "valid values are")[1]
	} else if strings.Contains(description, "Valid values are") {
		valuePart = strings.Split(description, "Valid values are")[1]
	} else if strings.Contains(description, "can be:") {
		valuePart = strings.Split(description, "can be:")[1]
	} else if strings.Contains(description, "allowed values") {
		valuePart = strings.Split(description, "allowed values")[1]
	}

	if valuePart != "" {
		// Extract values - look for comma-separated or slash-separated lists
		var values []string
		if strings.Contains(valuePart, ",") {
			// Comma-separated list
			parts := strings.Split(valuePart, ",")
			for _, part := range parts {
				// Clean up the value
				value := cleanEnumValue(part)
				if value != "" {
					values = append(values, sanitizeEnumValue(value))
				}
			}
		} else if strings.Contains(valuePart, "/") {
			// Slash-separated list
			parts := strings.Split(valuePart, "/")
			for _, part := range parts {
				value := cleanEnumValue(part)
				if value != "" {
					values = append(values, sanitizeEnumValue(value))
				}
			}
		} else {
			// Maybe just one or two values
			values = extractPossibleValues(valuePart)
		}

		// Create the enum values
		for _, value := range values {
			// Make sure the value is a valid protobuf enum identifier
			enumValue := strings.ToUpper(value)
			// Ensure no invalid characters remain
			enumValue = sanitizeEnumValue(enumValue)

			enumValues = append(enumValues, map[string]string{
				"Name":        enumValue,
				"Description": value + " value",
			})
		}
	}

	return map[string]interface{}{
		"Name":        name,
		"Description": description,
		"Values":      enumValues,
	}
}

// cleanEnumValue cleans up a string to extract a valid enum value
func cleanEnumValue(value string) string {
	// Remove leading/trailing whitespace
	value = strings.TrimSpace(value)

	// Remove punctuation
	value = strings.Trim(value, ".,;:")

	// Remove quotes
	value = strings.Trim(value, "'\"")

	// Check if it looks like a valid enum value (no spaces, not empty)
	if value != "" && !strings.Contains(value, " ") {
		return value
	}

	// If the value has spaces but looks like a quoted identifier, use it
	if strings.HasPrefix(value, "'") && strings.HasSuffix(value, "'") {
		return strings.Trim(value, "'")
	}

	// Special case for handling common patterns
	if strings.Contains(value, "or") {
		parts := strings.Split(value, "or")
		for _, part := range parts {
			cleaned := cleanEnumValue(part)
			if cleaned != "" {
				return cleaned
			}
		}
	}

	return ""
}

// extractPossibleValues tries to extract individual values from a description string
func extractPossibleValues(text string) []string {
	var values []string

	// Look for words in ALL_CAPS or quoted strings which are likely enum values
	allCapsPattern := regexp.MustCompile(`\b[A-Z_]{2,}\b`)
	allCaps := allCapsPattern.FindAllString(text, -1)
	values = append(values, allCaps...)

	// Look for quoted values - fix regex to avoid invalid character escaping
	quotedPattern := regexp.MustCompile(`['"]([^'"]+)['"]`)
	quotedMatches := quotedPattern.FindAllStringSubmatch(text, -1)
	for _, match := range quotedMatches {
		if len(match) > 1 {
			values = append(values, match[1])
		}
	}

	// Common values we expect to find in certain contexts
	if strings.Contains(strings.ToLower(text), "allow") && strings.Contains(strings.ToLower(text), "deny") {
		values = []string{"ALLOW", "DENY"}
	}

	// Sanitize each value to ensure it's a valid protobuf enum value
	for i, value := range values {
		values[i] = sanitizeEnumValue(value)
	}

	return values
}

// snakeAllCaps converts a string to SNAKE_CASE
func snakeAllCaps(s string) string {
	// First convert to snake_case
	snake := toSnakeCase(s)
	// Then to uppercase
	return strings.ToUpper(snake)
}

// toTitleCase converts snake_case or camelCase to TitleCase
func toTitleCase(s string) string {
	// First convert to snake_case to normalize
	snake := toSnakeCase(s)

	// Then convert to TitleCase
	var title string
	parts := strings.Split(snake, "_")
	for _, part := range parts {
		if part != "" {
			title += strings.ToUpper(part[:1]) + part[1:]
		}
	}

	return title
}
