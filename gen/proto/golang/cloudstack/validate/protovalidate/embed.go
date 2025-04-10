package protovalidate

import "embed"

//go:embed  *.go cel/*.go resolve/*.go
var root embed.FS

func AllGoFiles() (embed.FS, error) {
	return root, nil
}
