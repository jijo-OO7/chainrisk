ChainRisk Design Document

# Introduction

chainrisk is a CLI tool for probabilistic risk analysis of open-source software supply chains.

Modern software systems rely heavily on open-source dependencies. A single compromised component can propagate across the dependency graph and impact thousands of downstream systems.

Recent incidents such as the XZ Utils backdoor incident demonstrate how attackers can compromise widely used dependencies to maximize impact.

Traditional security tools focus primarily on known vulnerabilities, but supply-chain compromise often emerges through ecosystem dynamics such as:

maintainer trust escalation

malicious releases

dependency propagation

ecosystem chokepoints

chainrisk aims to model these risks using dependency graphs, ecosystem signals, and probabilistic modeling.

The goal is to shift supply-chain security from reactive vulnerability detection to predictive risk forecasting.

# Design Goals

Predictive Security

Most tools answer:

Are there vulnerabilities in my dependencies?

chainrisk attempts to answer:

What is the probability that a dependency becomes compromised?

This requires analyzing ecosystem behavior and structural risk, not just vulnerability databases.

Graph-Based Modeling

Software supply chains behave as dependency graphs, not flat lists.

Example:

Application
 ├── grpc
 │     └── protobuf
 └── openssl

A compromise in a low-level dependency can propagate upstream.

Graph modeling allows the system to:

simulate compromise propagation

detect ecosystem chokepoints

compute dependency blast radius

# Modular Architecture

The system follows a strict pipeline architecture.

Input
 ↓
Dependency Graph
 ↓
Graph Metrics
 ↓
Risk Signals
 ↓
Risk Model
 ↓
Risk Report

Each subsystem is independent and extensible.

Core Data Model
Graph Node

The fundamental unit of analysis is a package version.

Example nodes:

openssl@3.0.12
grpc@1.62.0
liblzma@5.6.0

This allows the system to detect version-specific compromises.

For example:

liblzma@5.6.0  → compromised
liblzma@5.5.1  → safe
Graph Edge

Dependencies form directed edges.

A → B

Meaning:

A depends on B

Example dependency graph:

payment-service
     |
   grpc@1.62
     |
protobuf@4.25
     |
   zlib@1.2
Dependency Graph Structure

The dependency graph is stored as a directed graph.

Responsibilities of the graph module:

node creation

dependency edge creation

dependency traversal

reverse dependency lookup

graph analysis

The graph must support efficient traversal for:

risk propagation

centrality metrics

blast radius analysis

Component Risk States

Each component exists in a risk state.

Healthy
 ↓
Maintainer Risk
 ↓
Vulnerable
 ↓
Compromised

These states represent the lifecycle of supply-chain compromise.

Risk Signals

The system aggregates multiple signals to evaluate risk.

# Maintainer Signals

Repository activity indicators:

maintainer inactivity

contributor privilege escalation

release ownership concentration

sudden commit bursts

These signals may indicate potential maintainer compromise.

Vulnerability Signals

Vulnerability history is used as a risk indicator.

Examples include:

CVE frequency

patch cadence

exploit availability

Graph Signals

Graph structure provides important contextual signals:

dependency centrality

dependency depth

fan-out impact

These metrics identify high-impact dependencies.

Centralized Dependency Risk Model

Modern software ecosystems often contain highly centralized dependencies that many projects rely on.

Example ecosystem:

serviceA        serviceB        serviceC
   |                |               |
  grpc             grpc           grpc
     \              |             /
       -------- protobuf --------
                 |
               zlib

If zlib is compromised:

zlib
 ↓
protobuf
 ↓
grpc
 ↓
many projects

Such dependencies represent ecosystem chokepoints.

This concept is critical for supply-chain risk analysis.

Libraries such as OpenSSL and XZ Utils illustrate this phenomenon.

Graph Metrics

To detect ecosystem chokepoints, the system computes several graph metrics.

Degree Centrality

Measures how many components depend on a node.

Example:

protobuf
 ↑  ↑  ↑
grpc libA libB

Higher degree implies greater ecosystem importance.

Betweenness Centrality

Measures how frequently a node lies on dependency paths.

Example:

service
 ↓
grpc
 ↓
protobuf
 ↓
zlib

If protobuf fails, many dependency paths break.

# Dependency Depth

Measures how deeply nested a dependency is.

Example:

service
 ↓
grpc
 ↓
protobuf
 ↓
zlib

Deep dependencies are often harder to audit.

# Risk Scoring

Risk scores combine multiple signals.

Example scoring formula:

RiskScore =
  CentralityScore
+ VulnerabilityScore
+ MaintainerRiskScore

This allows identifying dependencies that pose systemic supply-chain risk.

Markov Risk Model

chainrisk models dependency compromise using a Markov process.

Example transition matrix:

              H     MR    V     C
Healthy      0.85  0.10  0.05  0.00
MaintRisk    0.00  0.70  0.20  0.10
Vulnerable   0.00  0.00  0.75  0.25
Compromised  0.00  0.00  0.00  1.00

Where:

H  = Healthy
MR = Maintainer Risk
V  = Vulnerable
C  = Compromised

The Markov engine simulates future compromise probabilities.

Blast Radius Analysis

Blast radius estimates downstream impact of compromise.

Example:

zlib compromised
 ↓
protobuf affected
 ↓
grpc affected
 ↓
payment-service affected

This allows identifying the most dangerous dependencies.

# CLI Interface

chainrisk exposes several commands.

Analyze SBOM
chainrisk analyze sbom.json

Parses the SBOM and constructs the dependency graph.

Repository Analysis
chainrisk repo github.com/org/project

Analyzes maintainer ecosystem signals.

Risk Simulation
chainrisk simulate sbom.json

Runs probabilistic risk simulation.

Blast Radius
chainrisk blast sbom.json

Estimates downstream compromise impact.

Implementation Layout

Repository structure:

cmd/chainrisk
internal/sbom
internal/graph
internal/metrics
internal/risk
internal/markov
docs/

Responsibilities:

sbom     → SBOM parsing
graph    → dependency graph
metrics  → graph centrality
risk     → risk signals
markov   → probabilistic engine
Development Phases
# Phase 1 — Graph Foundation

SBOM parsing

dependency graph construction

CLI analyze command

# Phase 2 — Risk Signals

maintainer ecosystem analysis

vulnerability signals

# Phase 3 — Risk Modeling

Markov engine

compromise probability simulation

# Phase 4 — Advanced Analysis

attack path detection

blast radius simulation

ecosystem chokepoint detection

### Future Work

Potential extensions include:

container image analysis

ecosystem-wide dependency graphs

CI/CD integration

Kubernetes supply-chain analysis

### Conclusion

chainrisk aims to advance supply-chain security from reactive vulnerability scanning to predictive ecosystem risk modeling.

By combining dependency graph analysis, ecosystem signals, and probabilistic modeling, the system provides a deeper understanding of supply-chain compromise risk.