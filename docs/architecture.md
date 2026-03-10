ChainRisk Design Document

# Overview

chainrisk is a CLI tool designed to analyze and forecast risk in open-source software supply chains.

Traditional security tools focus on known vulnerabilities. However, recent incidents such as the XZ Utils backdoor incident demonstrate that compromise often emerges from ecosystem dynamics, including maintainer trust changes, build pipeline manipulation, and dependency propagation.

chainrisk aims to estimate the probability of supply-chain compromise by combining:

dependency graph analysis

maintainer ecosystem signals

vulnerability evolution

probabilistic state modeling

The system produces a risk forecast instead of a static vulnerability list.

# Problem Statement

Modern software systems depend heavily on open-source components.

A typical application may include:

Application
 ├── Dependency A
 │     ├── Dependency C
 │     └── Dependency D
 └── Dependency B

A compromise in a single upstream component can propagate through the dependency graph and impact downstream systems.

Current security tools primarily answer:

Do vulnerabilities exist in dependencies?

chainrisk attempts to answer:

What is the probability that a dependency becomes compromised?
System Architecture

The tool processes supply-chain data through several stages.

+----------------------+
| Input                |
| SBOM / Repository    |
+----------+-----------+
           |
           v
+----------------------+
| SBOM Parser          |
| Extract components   |
+----------+-----------+
           |
           v
+----------------------+
| Dependency Graph     |
| Construct graph of   |
| components           |
+----------+-----------+
           |
           v
+----------------------+
| Risk Signals         |
| Maintainer signals   |
| Vulnerability data   |
| Release activity     |
+----------+-----------+
           |
           v
+----------------------+
| Markov Risk Engine   |
| State transitions    |
| Risk simulation      |
+----------+-----------+
           |
           v
+----------------------+
| Risk Report          |
| Compromise score     |
| Attack path          |
+----------------------+
Core Concepts
Dependency Graph

Dependencies are represented as a directed graph.

A → B → C
A → D

Where:

nodes represent software components

edges represent dependency relationships

The graph allows the system to simulate how compromise propagates across the ecosystem.

Component States

Each component in the graph can exist in a risk state.

Healthy
↓
Maintainer Risk
↓
Vulnerable
↓
Compromised

Transitions between states are probabilistic.

# Markov Risk Model

The risk engine models component state transitions using a Markov process.

Example transition matrix:

              H     MR    V     C
Healthy      0.85  0.10  0.05  0.00
MaintRisk    0.00  0.70  0.20  0.10
Vulnerable   0.00  0.00  0.75  0.25
Compromised  0.00  0.00  0.00  1.00

Where:

H = Healthy

MR = Maintainer Risk

V = Vulnerable

C = Compromised

This allows simulation of future compromise probability.

# CLI Interface

The tool exposes several commands.

Analyze SBOM
chainrisk analyze sbom.json

Extracts dependencies and builds a dependency graph.

Repository Analysis
chainrisk repo github.com/org/project

Analyzes repository activity and maintainer ecosystem signals.

Risk Simulation
chainrisk simulate sbom.json

Runs probabilistic simulation to estimate compromise risk.

Dependency Blast Radius
chainrisk blast sbom.json

Estimates downstream impact of a compromised dependency.

Internal Module Layout

The repository is organized using standard Go project structure.

cmd/chainrisk
internal/sbom
internal/graph
internal/risk
internal/markov
docs/

Responsibilities:

sbom    → SBOM parsing
graph   → dependency graph construction
risk    → signal extraction
markov  → probabilistic risk model
Development Roadmap

The system will be developed in stages.

# Phase 1 — SBOM Analysis

Features:

SBOM parsing

dependency extraction

graph construction

# Phase 2 — Risk Signals

Features:

maintainer activity analysis

vulnerability history signals

# Phase 3 — Probabilistic Risk Engine

Features:

Markov state model

Monte Carlo simulation

compromise probability calculation

# Phase 4 — Attack Path Analysis

Features:

dependency blast radius

compromise propagation simulation

Example Output
Supply Chain Risk Report
------------------------

Root Artifact: payment-service

Compromise Probability (90 days): 0.18

Top Risk Dependencies:
1. liblzma
2. openssl
3. grpc

Most Likely Attack Path:
liblzma → systemd → ssh
Future Work

Potential extensions include:

container image analysis

CI/CD pipeline risk modeling

Kubernetes supply chain analysis

GitHub Actions integration

Design Principles

The architecture follows a strict pipeline:

Input
↓
Graph
↓
Signals
↓
Risk Model
↓
Output

This separation ensures that each subsystem remains modular and extensible.

# Conclusion

chainrisk aims to move supply-chain security from reactive vulnerability scanning to predictive risk modeling.

By combining dependency graphs with probabilistic modeling, the system attempts to forecast compromise risk across open-source ecosystems.