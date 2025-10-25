package main

import (
	"flag"
	"strings"

	gengo "google.golang.org/protobuf/cmd/protoc-gen-go/internal_gengo"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/types/descriptorpb"
)

type arrayFlags []string

func (i *arrayFlags) String() string {
	return strings.Join(*i, ",")
}

func (i *arrayFlags) Set(value string) error {
	*i = append(*i, value)
	return nil
}

var (
	includeGoPackages arrayFlags
	unspecifiedSuffix string
)

func main() {
	var flags flag.FlagSet

	flags.Var(&includeGoPackages, "include_go_packages",
		"list of go packages with govmomi-related types definitions to include")
	flags.StringVar(&unspecifiedSuffix, "unspecified_suffix", "UNSPECIFIED",
		"suffix to remove from enum values (e.g. UNSPECIFIED, UNSET, etc.)")

	protogen.Options{
		ParamFunc: flags.Set,
	}.Run(func(gen *protogen.Plugin) error {
		gen.SupportedFeatures = gengo.SupportedFeatures
		gen.SupportedEditionsMinimum = descriptorpb.Edition_EDITION_2023
		gen.SupportedEditionsMaximum = descriptorpb.Edition_EDITION_2024

		// if there are no files to generate, return nil
		if len(gen.Files) == 0 {
			return nil
		}

		// check if the first file's GoImportPath is included in the includeTypesGoPackages list
		// if not, return nil
		if !checkGoImportPath(gen, gen.Files[0]) {
			return nil
		}

		for _, f := range gen.Files {
			if !f.Generate {
				continue
			}

			enumNum := len(f.Enums)
			for _, m := range f.Messages {
				enumNum += len(m.Enums)
			}

			if enumNum > 0 {
				genEnumExtractor(gen, f)
			}
		}

		return nil
	})
}

// checkGoImportPath checks if the given file's GoImportPath is included in the includeTypesGoPackages list
func checkGoImportPath(gen *protogen.Plugin, file *protogen.File) bool {
	for _, p := range includeGoPackages {
		if p == string(file.GoImportPath) {
			return true
		}
	}
	return false
}
