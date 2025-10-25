# protoc-gen-go-enum-extractor
protoc-gen-go-enum-extractor is a protoc go plugin that helps you extract actual enum values with prefix removed.

An example proto package is included in the `proto` directory. The generated files are included in the `gen` directory.

## Motivation
When you have a proto file with [buf style](https://buf.build/docs/best-practices/style-guide) enumerators like this:
```proto
enum FooBar {
    FOO_BAR_UNSPECIFIED = 0;
    FOO_BAR_A = 1;
    FOO_BAR_B = 2;
}
```
You will not be able to get the enum values with the prefix removed (UNSPECIFIED / A / B) directly with the generated go code. This plugin helps you extract the enum values from proto files.

## Usage

1. Install the plugin

```bash
go install github.com/daotl/protoc-gen-go-enum-extractor@latest
```

2. Include this plugin in your buf.gen.yaml

Parameters:
- `include_go_packages`: A list of go packages with enum definitions to include.
- `unspecified_suffix`: Suffix to remove from enum values (e.g. UNSPECIFIED, UNSET, etc.). Default value is `UNSPECIFIED`.

## Note
`*_UNSPECIFIED = 0` must be the first enumerator in the enum definition.

# Author
**Jiayin Zhang**

* <https://github.com/daotl>

# License

Released under [Apache-2.0 License](https://github.com/daotl/protoc-gen-go-enum-extractor/blob/main/LICENSE)