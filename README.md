# ChainRisk

<p align="center">
  <img src="ChainRisk_logo.png" alt="chainrisk logo" width="600">
</p>

<p align="center">
  <strong>ChainRisk is a CLI tool for analyzing software supply chain risk using SBOMs.</strong>
</p>

---

## Overview

ChainRisk helps developers and security engineers understand how software dependencies are connected and how risk propagates through them.

Instead of focusing only on known vulnerabilities, ChainRisk analyzes dependency relationships to answer a more practical question:

> **If one dependency is compromised, what breaks?**

---

## Why ChainRisk

Modern systems rely on deep dependency chains:

```
Application
   ↓
Libraries
   ↓
Transitive Dependencies
```

A failure in a low-level dependency can affect multiple services.

ChainRisk helps you:

* understand dependency relationships
* identify critical shared components
* simulate how failures propagate

---

## Features (v0.1)

### 1. SBOM Analysis

Parse SPDX/CycloneDX SBOMs to extract dependency information.

```bash
chainrisk sbom-info <file>
```

Example:

```bash
chainrisk sbom-info testdata/sample.json
```

Output:

```
Total packages: 3
- zlib
- protobuf
- grpc
```

---

### 2. Dependency Graph *(coming next)*

Build and visualize dependency relationships from an SBOM.

---

### 3. Blast Radius Simulation *(coming next)*

Simulate the impact of a compromised dependency.

---

## CLI Usage

```bash
chainrisk version
chainrisk sbom-info <file>
```

---

## Installation

### Using Go (recommended)

```bash
go install github.com/jijo-OO7/chainrisk/cmd/chainrisk@latest
```

---

### From source

```bash
git clone https://github.com/jijo-OO7/chainrisk.git
cd chainrisk
go build -o chainrisk ./cmd/chainrisk
```

Run:

```bash
./chainrisk
```

---

## Project Structure

```
cmd/chainrisk       → CLI entry point  
internal/sbom       → SBOM parsing logic  
internal/cli        → command handlers  
testdata/           → sample SBOM files  
```

---

## Roadmap

* dependency graph generation
* blast radius simulation
* CI/CD integration
* risk scoring

---

## License

Licensed under the Apache License 2.0.

---

## Author

Suman Mandal
GitHub: https://github.com/jijo-OO7
