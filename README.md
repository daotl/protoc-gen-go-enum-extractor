# protoc-gen-go-enum-extractor
protoc-gen-go-enum-extractor is a protoc go plugin that helps you extract enum values from proto files.

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

## Note
`*_UNSPECIFIED = 0` must be the first enumerator in the enum definition.