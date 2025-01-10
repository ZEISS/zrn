# zrn

![Github Action](https://github.com/zeiss/zrn/workflows/main/badge.svg)
[![Go Reference](https://pkg.go.dev/badge/github.com/zeiss/zrn.svg)](https://pkg.go.dev/github.com/zeiss/zrn)
[![Go Report Card](https://goreportcard.com/badge/github.com/zeiss/zrn)](https://goreportcard.com/report/github.com/zeiss/zrn)
[![Taylor Swift](https://img.shields.io/badge/secured%20by-taylor%20swift-brightgreen.svg)](https://twitter.com/SwiftOnSecurity)
[![Volkswagen](https://auchenberg.github.io/volkswagen/volkswargen_ci.svg?v=1)](https://github.com/auchenberg/volkswagen)
[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://opensource.org/licenses/Apache-2.0)

A simple and fast library to read and write ZRNs (Zeiss Resource Names).

## Install

```bash
go get github.com/zeiss/zrn
```

## Specification

### Format

```plaintext
zrn:partition:product:region:identifier:resource-id
zrn:partition:product:region:identifier:resource-type/resource-id
zrn:partition:product:region:identifier:resource-type:resource-id
```

## License

[Apache 2.0](/LICENSE)
