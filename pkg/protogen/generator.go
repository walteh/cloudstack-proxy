package protogen

import (
	"fmt"
	"path/filepath"

	"github.com/walteh/cloudstack-proxy/pkg/protogen/metadata"
)

//go:opts
type Options struct {
	// MetadataDir is the directory containing CloudStack API metadata
	metadataDir string
	// OutputDir is the directory where generated files will be written
	outputDir string
}

// Generator handles the generation of protobuf files
type Generator struct {
	opts Options
}

// NewGenerator creates a new protobuf generator with the given options
func NewGenerator(opts ...OptOptionsSetter) *Generator {
	return &Generator{
		opts: NewOptions(opts...),
	}
}

// Run executes the protobuf generation process
func (g *Generator) Run() error {
	// If metadata directory is specified, generate from metadata
	if g.opts.metadataDir == "" {
		return fmt.Errorf("metadata directory is required")
	}

	if g.opts.outputDir == "" {
		g.opts.outputDir = filepath.Join("tmp", "proto")
	}

	// Otherwise, process existing proto files
	return g.generateFromMetadata()
}

// generateFromMetadata generates protobuf files from CloudStack API metadata
func (g *Generator) generateFromMetadata() error {
	fmt.Println("Generating protobuf files from CloudStack API metadata...")

	// Parse the metadata
	meta, err := metadata.ParseMetadata(g.opts.metadataDir)
	if err != nil {
		return fmt.Errorf("error parsing metadata: %w", err)
	}

	// Create a proto generator
	protoGen := metadata.NewProtoGenerator(meta, filepath.Dir(g.opts.outputDir))

	// Generate the protobuf files
	if err := protoGen.GenerateProto(); err != nil {
		return fmt.Errorf("error generating protobuf files: %w", err)
	}

	fmt.Println("Successfully generated protobuf files from metadata")
	return nil
}
