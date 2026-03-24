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

## Features (MVP)

### 1. SBOM Analysis

Parse SPDX/CycloneDX SBOMs to extract dependency information.

```bash
chainrisk sbom-info sbom.json
```

Example output:

```
Total packages: 120
Top-level dependencies: 8
Most used dependency: zlib
```

---

### 2. Dependency Graph

Build and visualize dependency relationships from an SBOM.

```bash
chainrisk graph sbom.json
```

Example:

```
zlib → protobuf → grpc → service
```

This reveals how components are connected and which dependencies are reused.

---

### 3. Blast Radius Simulation

Simulate the impact of a compromised dependency.

```bash
chainrisk simulate sbom.json --target=zlib
```

Example output:

```
🚨 BLAST RADIUS REPORT

Compromised: zlib

Affected:
- protobuf
- grpc
- payment-service

Impact depth: 3
Total affected: 7 packages
```

---

## CLI Usage

```bash
chainrisk sbom-info <file>
chainrisk graph <file>
chainrisk simulate <file> --target=<dependency>
```

---

## Installation

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

## Roadmap

Planned future improvements:

* dependency centrality scoring
* container image analysis
* CI/CD integration
* risk scoring models

---

## License

Licensed under the Apache License 2.0.
