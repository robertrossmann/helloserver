//go:build ignore

// This program generates GraphQL components from a schema file. It works the same way as using
// `github.com/99designs/gqlgen generate` directly, but it allows us to configure the generator in Go instead of YAML,
// which means we have strongly typed configuration available in our IDEs. A secondary benefit is that we can write
// gqlgen plugins without refactoring the generation pipeline.
package main

import (
	"fmt"
	"os"

	"github.com/99designs/gqlgen/api"
	"github.com/99designs/gqlgen/codegen/config"
)

func main() {
	cfg := config.DefaultConfig()
	cfg.SchemaFilename = []string{"graph/*.graphqls"}
	cfg.Exec.Layout = config.ExecLayoutSingleFile
	cfg.Exec.Filename = "graph/exec_gen.go"
	cfg.Exec.DirName = "graph"
	cfg.Exec.Package = "graph"
	cfg.Model.Filename = "graph/gqlm/models_gen.go"
	cfg.Model.Package = "gqlm"
	cfg.Resolver.Layout = config.LayoutFollowSchema
	cfg.Resolver.Package = "graph"
	// cfg.Resolver.Filename = "graph/resolvers.go"
	cfg.Resolver.FilenameTemplate = "{name}.resolver.go"

	cfg.SkipModTidy = true
	cfg.OmitComplexity = true
	cfg.OmitSliceElementPointers = true
	cfg.OmitGQLGenVersionInFileNotice = true
	cfg.NullableInputOmittable = true

	if err := config.CompleteConfig(cfg); err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(2)
	}

	if err := api.Generate(cfg); err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(3)
	}
}
