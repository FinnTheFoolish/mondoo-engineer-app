# Mondoo Engineer — Application

A minimal Go HTTP service used to demonstrate an automated build and release pipeline.

## Application

The application listens on port `8080` by default and responds to HTTP requests with:

```text
Hello from Mondoo Engineer!
```

The listening port can be overridden with the `PORT` environment variable.

Example:

```bash
PORT=9000 go run ./cmd/server
```

## Local development

### Prerequisites

* Go 1.x

Run the application:

```bash
go run ./cmd/server
```

Then:

```bash
curl http://localhost:8080
```

Run tests:

```bash
go test ./...
```

Run the race detector:

```bash
go test -race ./...
```

## CI

Pull requests and pushes to `main` run:

1. Go tests
2. Go race detector
3. Semgrep SAST
4. golangci-lint
5. Binary compilation

The workflows use GitHub Actions dependency caching and concurrency cancellation so obsolete CI runs do not unnecessarily consume runner capacity.

The build is intentionally split into independent quality gates before compilation. This makes failures easier to diagnose and prevents a successful build from being interpreted as evidence that the source passed all validation.

## Releases

Releases are created by pushing a semantic-version tag:

```bash
git tag v1.0.0
git push origin v1.0.0
```

The release workflow builds:

* Linux amd64
* Linux arm64
* macOS amd64
* macOS arm64
* Windows amd64

The release contains the binaries and a `SHA256SUMS` file.

Go's `GOOS` and `GOARCH` variables are used for cross-compilation rather than maintaining separate build environments.

## Design decisions

### Standard library HTTP server

The application deliberately uses Go's standard library. The challenge states that the application is not intended to test advanced Go technique, so introducing a web framework would add dependencies without providing meaningful value.

### Cross-compilation

The release is built from a single Linux GitHub Actions runner using Go's cross-compilation support. This keeps the release workflow straightforward while producing artifacts for multiple operating systems and architectures.

### Matrix builds

The release workflow uses a GitHub Actions matrix rather than duplicating jobs for each operating-system/architecture combination. This makes adding another target a data change rather than a workflow-logic change.

### Race detector

The race detector is included as an inexpensive additional test signal. Although the application is deliberately small, retaining this check establishes a useful CI pattern for future concurrent code.

### Checksums

A SHA-256 checksum file is published with each release. This provides a simple integrity mechanism for downloaded binaries.

### Security tooling

Semgrep provides source-level static analysis. `govulncheck` is used as an additional Go-specific dependency/vulnerability signal.

These checks are deliberately separate from the compilation step so that security and code-quality failures remain visible as independent pipeline stages.

## Trade-offs

This repository does not attempt to implement a complex Go application architecture. Graceful shutdown, structured logging, configuration frameworks and dependency injection would be reasonable production concerns, but they would distract from the platform-engineering objective of this exercise.

Release signing and SBOM publication are suitable extensions to this pipeline and can be added without changing the application's architecture.

Bonne Chance!