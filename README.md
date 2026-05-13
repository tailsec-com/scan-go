# Tailsec Go SDK

[![Go Reference](https://img.shields.io/badge/go-reference-00ADD8?logo=go)](https://pkg.go.dev/github.com/tailsec-com/scan-go)
[![CI](https://img.shields.io/github/actions/workflow/status/tailsec-com/scan-go/ci.yml?branch=main)](https://github.com/tailsec-com/scan-go/actions)
[![License: MIT](https://img.shields.io/badge/license-MIT-blue.svg)](LICENSE)

![logo](logo.svg)

## Features

- Unified Go interface for all Tailsec scanners
- Supports: Kubernetes, Terraform, Docker, CloudFormation, SAST, Dependencies, Licenses
- Idiomatic Go with struct types and error handling
- Great for embedding in Go services and CI/CD tooling

## Installation

```bash
go get github.com/tailsec-com/scan-go
```

## Quick Start

```go
package main

import (
    "fmt"
    "log"

    "github.com/tailsec-com/scan-go"
)

func main() {
    client := scan.NewClient(scan.ScannerK8s)

    result, err := client.Scan("./manifests")
    if err != nil {
        log.Fatal(err)
    }

    fmt.Printf("Found %d issues\n", result.Summary.Total)

    for _, finding := range result.Kubernetes {
        fmt.Printf("[%s] %s\n", finding.Severity, finding.Title)
    }
}
```

## Available Scanners

| Scanner | Constant | Description |
|---------|----------|-------------|
| Kubernetes | `scan.ScannerK8s` | Kubernetes manifest security |
| Terraform | `scan.ScannerTerraform` | Terraform IaC security |
| Docker | `scan.ScannerDocker` | Dockerfile security |
| CloudFormation | `scan.ScannerCfn` | AWS CloudFormation security |
| SAST | `scan.ScannerSast` | Static application security testing |
| Dependencies | `scan.ScannerDeps` | npm dependency security |
| Licenses | `scan.ScannerLicense` | License compliance |

## Result Structure

```go
type Result struct {
    Summary     Summary            `json:"summary"`
    Kubernetes  []K8sFinding       `json:"kubernetes,omitempty"`
    Terraform   []TerraformFinding `json:"terraform,omitempty"`
    Dockerfile  []DockerFinding    `json:"dockerfile,omitempty"`
    // ...
}

type Summary struct {
    Total     int `json:"total"`
    Critical  int `json:"critical"`
    High      int `json:"high"`
    Medium    int `json:"medium"`
    Low       int `json:"low"`
}
```

## Contributing

This SDK is generated from the core Tailsec scanner packages. To contribute:

1. Contribute to the underlying scanner in `packages/scan-*`
2. Update the Go SDK bindings in this repository
3. Run `go generate` to regenerate bindings

## License

MIT