# ChainRisk

<p align="center">
  <img src="docs/chainrisk-logo.png" alt="chainrisk logo" width="300">
</p>

<p align="center">
  <strong>ChainRisk is a supply-chain risk intelligence CLI for container ecosystems.</strong>
</p>

<p align="center">
  Analyze container images, dependencies, and open-source ecosystems to predict and understand software supply chain compromise risk.
</p>

## Overview

ChainRisk is a modern CLI tool designed to help security teams analyze the risk landscape of software supply chains.

Instead of only detecting known vulnerabilities, ChainRisk models dependency ecosystems, maintainer trust signals, and probabilistic risk propagation to identify dependencies most likely to cause large-scale compromise.

ChainRisk is designed for modern cloud environments built on containers and Kubernetes.

## Why ChainRisk

Modern infrastructure relies on deep dependency chains:

```
Kubernetes Workload
        ↓
Container Image
        ↓
Base Image
        ↓
OS Libraries
        ↓
Application Dependencies
```

A compromise in a low-level dependency can propagate across many systems.

ChainRisk helps security teams understand:

- which dependencies are most critical
- which components affect the most workloads
- which projects show suspicious maintainer activity
- how compromise could propagate through infrastructure

## Features

### Container Supply Chain Analysis

Analyze container images and their dependency trees.

```bash
chainrisk image ghcr.io/org/service:latest
```

ChainRisk inspects:

- container base images
- OS packages
- application dependencies

### Global Dependency Risk Graph

ChainRisk builds a global dependency graph across container images.

Example ecosystem graph:

```
serviceA        serviceB        serviceC
   |                |               |
  grpc             grpc           grpc
     \              |             /
       -------- protobuf --------
                 |
                zlib
```

This reveals ecosystem chokepoints where compromise would have the largest impact.

### Dependency Centrality Detection

ChainRisk calculates graph metrics such as:

- dependency centrality
- dependency depth
- blast radius

These identify critical dependencies that represent systemic risk.

### Blast Radius Simulation

Simulate the impact of a compromised dependency.

Example:

```
zlib compromised
 ↓
protobuf affected
 ↓
grpc affected
 ↓
payment-service affected
```

This helps security teams understand how compromise propagates.

### Maintainer Trust Drift Detection

ChainRisk analyzes repository activity to detect maintainer trust drift.

Indicators include:

- maintainer inactivity
- sudden contributor dominance
- release ownership changes

These signals often precede supply-chain compromise.

### Probabilistic Risk Modeling

ChainRisk models dependency risk using probabilistic state transitions.

```
Healthy
 ↓
Maintainer Risk
 ↓
Vulnerable
 ↓
Compromised
```

This allows predictive security analysis, not just reactive scanning.

### Container Image Verification

ChainRisk integrates with supply-chain security tooling such as:

- Cosign for container signature verification
- Syft for SBOM generation
- SLSA provenance metadata

Example pipeline:

```
Container Image
      ↓
Cosign verification
      ↓
SBOM generation
      ↓
Dependency graph analysis
      ↓
Risk intelligence report
```

## CLI Usage

Example commands:

```bash
chainrisk image <container-image>
chainrisk dependency <package>
chainrisk ecosystem
chainrisk simulate
```

Example:

```bash
chainrisk image ghcr.io/org/payment-service:latest
```

## Installation

### Install from source

```bash
git clone https://github.com/jijo-OO7/chainrisk.git
cd chainrisk
go build -o chainrisk ./cmd/chainrisk
```

Run:

```bash
./chainrisk
```

## Architecture

```
Container Image / Repository
            ↓
SBOM Extraction
            ↓
Dependency Graph
            ↓
Global Dependency Graph
            ↓
Graph Metrics
            ↓
Risk Signals
            ↓
Markov Risk Engine
            ↓
Risk Intelligence Report
```

## What Makes ChainRisk Unique

ChainRisk differs from traditional security scanners.

Most tools answer:

> Does this dependency contain known vulnerabilities?

ChainRisk answers:

- Which dependency is most critical in the ecosystem?
- Which dependency is most likely to be compromised?
- If dependency X fails, which services are affected?
- Are maintainers showing suspicious trust drift?

### Key innovations:

- ecosystem-level dependency graph analysis
- probabilistic compromise prediction
- maintainer trust drift detection
- container-native supply-chain analysis

## Roadmap

Upcoming development milestones:

- container image analysis pipeline
- dependency graph engine
- ecosystem indexing mode
- probabilistic risk engine
- Kubernetes workload scanning

## Contributing

Contributions are welcome.

Please read CONTRIBUTING.md for guidelines on reporting issues and submitting pull requests.

## Security

If you discover a security vulnerability in ChainRisk, please follow the responsible disclosure process described in SECURITY.md.

## License

ChainRisk is licensed under the Apache License 2.0.

See LICENSE for details.

## Author

Suman Mandal  
GitHub: https://github.com/jijo-OO7

## Inspiration

ChainRisk is inspired by the need for better supply-chain intelligence in modern container ecosystems.

It builds on ideas from:

- container security
- dependency graph analysis
- probabilistic risk modeling