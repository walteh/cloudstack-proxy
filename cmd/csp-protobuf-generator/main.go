package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"github.com/walteh/cloudstack-proxy/pkg/protogen"
)

func main() {
	fmt.Println("CloudStack Proxy - Protobuf Generator")

	// Define command-line flags
	metadataDir := flag.String("metadata-dir", "", "Directory containing CloudStack API metadata JSON files")
	outputDirFlag := flag.String("output-dir", "", "Output directory for generated proto files (defaults to [root]/proto)")
	formatFlag := flag.Bool("format", true, "Whether to format the generated proto files")
	flag.Parse()

	// Get the root directory
	rootDir, err := os.Getwd()
	if err != nil {
		fmt.Printf("Error getting current directory: %v\n", err)
		os.Exit(1)
	}

	outputDir := *outputDirFlag
	if outputDir == "" {
		outputDir = filepath.Join(rootDir, "proto")
	}

	// Create and run generator
	generator := protogen.NewGenerator(
		protogen.WithMetadataDir(*metadataDir),
		protogen.WithOutputDir(outputDir),
		protogen.WithFormat(*formatFlag),
	)
	err = generator.Run()
	if err != nil {
		fmt.Printf("Error generating protobuf files: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Protobuf generation completed successfully")
}
