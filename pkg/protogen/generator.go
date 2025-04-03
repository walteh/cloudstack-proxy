package protogen

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

// Options for the protobuf generator
type Options struct {
	// SourceDir is the root directory containing original proto files
	SourceDir string

	// OutputDir is the directory where generated files will be written
	OutputDir string

	// OutputSuffix is the suffix added to generated files (e.g., ".gen.proto")
	OutputSuffix string

	// RootPackage is the root package name for the protobuf definitions
	RootPackage string

	// GoModule is the name of the Go module
	GoModule string

	// GenPackageRoot is the root directory for generated Go packages
	GenPackageRoot string
}

// Generator handles the generation of protobuf files
type Generator struct {
	opts Options
}

// NewGenerator creates a new protobuf generator with the given options
func NewGenerator(opts Options) *Generator {
	return &Generator{
		opts: opts,
	}
}

// Run executes the protobuf generation process
func (g *Generator) Run() error {
	// Find all .proto files in the source directory
	var protoFiles []string
	err := filepath.WalkDir(g.opts.SourceDir, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		// Skip generated files
		if strings.HasSuffix(path, g.opts.OutputSuffix) {
			return nil
		}

		// Process only .proto files
		if !d.IsDir() && strings.HasSuffix(path, ".proto") {
			relPath, err := filepath.Rel(g.opts.SourceDir, path)
			if err != nil {
				return err
			}
			protoFiles = append(protoFiles, relPath)
		}

		return nil
	})

	if err != nil {
		return fmt.Errorf("error walking source directory: %w", err)
	}

	fmt.Printf("Found %d proto files to process\n", len(protoFiles))

	// Process each proto file
	for _, relPath := range protoFiles {
		err := g.processProtoFile(relPath)
		if err != nil {
			return fmt.Errorf("error processing %s: %w", relPath, err)
		}
	}

	return nil
}

// processProtoFile processes a single proto file
func (g *Generator) processProtoFile(relPath string) error {
	sourcePath := filepath.Join(g.opts.SourceDir, relPath)

	// Read source file
	content, err := os.ReadFile(sourcePath)
	if err != nil {
		return fmt.Errorf("error reading source file: %w", err)
	}

	// Parse and transform content (for now, just pass through)
	genContent := string(content)

	// This is where you would add custom processing:
	// - Add imports
	// - Apply custom hooks
	// - Transform messages/services

	// Determine output path
	dir := filepath.Dir(relPath)
	base := filepath.Base(relPath)
	ext := filepath.Ext(base)
	name := strings.TrimSuffix(base, ext)

	outputFile := name + g.opts.OutputSuffix
	outputPath := filepath.Join(g.opts.OutputDir, dir, outputFile)

	// Ensure output directory exists
	outputDir := filepath.Dir(outputPath)
	err = os.MkdirAll(outputDir, 0755)
	if err != nil {
		return fmt.Errorf("error creating output directory: %w", err)
	}

	// Write output file
	err = os.WriteFile(outputPath, []byte(genContent), 0644)
	if err != nil {
		return fmt.Errorf("error writing output file: %w", err)
	}

	fmt.Printf("Generated: %s\n", outputPath)

	return nil
}
