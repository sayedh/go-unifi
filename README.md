# Unifi Go SDK [![GoDoc](https://godoc.org/github.com/sayedh/go-unifi?status.svg)](https://godoc.org/github.com/sayedh/go-unifi)

This project is a fork of [paultyng/go-unifi](https://github.com/paultyng/go-unifi/), updated to support the latest Unifi Network versions.

## Development History

- This project was originally created for use in the [Terraform provider for Unifi](https://github.com/paultyng/terraform-provider-unifi).
- The original project supported Unifi Network versions up to 7.4.162.
- In August 2024, this fork was created to update support for newer Unifi Network versions.
- The code was regenerated using `go generate` in the `unifi` directory, which created many new changes to support Unifi Network version 8.3.32.
- Work is ongoing to update the supplementary Terraform provider code at [sayedh/terraform-provider-unifi](https://github.com/sayedh/terraform-provider-unifi).

## Versioning (from paultyng)

Many of the naming adjustments are breaking changes, but to simplify things, naming errors are treated as minor changes for the 1.0.0 version (probably should have just started at 0.1.0).

## Note on Code Generation (from paultyng)

The data models and basic REST methods are "generated" from JSON files in the JAR that show all fields and the associated regex/validation information.

To regenerate the code, you can bump the Unifi Controller version number in [unifi/gen.go] and run `go generate` inside the `unifi` directory.

This code generation is kind of gross, I wanted to switch to using the java classes in the jar like scala2go but the jar is obfuscated and I couldn't find a way to extract that information from anywhere else. Maybe it exists somewhere in the web UI, but I was unable to find it in there in a way that was extractable in a practical way.

Still planning to dig through the bits some more later on.

## Contributing

Contributions to update and improve this SDK for newer Unifi Network versions are welcome. Please submit issues and pull requests to the [GitHub repository](https://github.com/sayedh/go-unifi).