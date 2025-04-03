package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/walteh/cloudstack-proxy/pkg/protogen"
)

func main() {
	fmt.Println("CloudStack Proxy - Protobuf Generator")

	// Get the root directory
	rootDir, err := os.Getwd()
	if err != nil {
		fmt.Printf("Error getting current directory: %v\n", err)
		os.Exit(1)
	}

	// Setup generator options
	opts := protogen.Options{
		SourceDir:      filepath.Join(rootDir, "proto"),
		OutputDir:      filepath.Join(rootDir, "proto"),
		OutputSuffix:   ".gen.proto",
		RootPackage:    "cloudstack",
		GoModule:       "github.com/walteh/cloudstack-proxy",
		GenPackageRoot: "gen/proto",
	}

	// Create and run generator
	generator := protogen.NewGenerator(opts)
	err = generator.Run()
	if err != nil {
		fmt.Printf("Error generating protobuf files: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Protobuf generation completed successfully")
}
