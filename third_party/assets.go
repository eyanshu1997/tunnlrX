package third_party

import "embed"

// OpenAPI contains the static OpenAPI UI files.
//
//go:embed OpenAPI/*
var OpenAPI embed.FS
