# ChainRisk

**ChainRisk** is a supply-chain risk intelligence tool designed to analyze container images, dependencies, and open-source ecosystems to estimate the **probability of compromise across software supply chains**.

Unlike traditional scanners that focus only on known vulnerabilities, ChainRisk models **structural and ecosystem risk** using dependency graphs, maintainer trust signals, and probabilistic modeling.

---

# Why ChainRisk Exists

Modern cloud systems depend on complex supply chains:

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

A compromise in a low-level component can propagate through multiple layers and impact many services.

ChainRisk is designed to help security teams **identify these risks before they become incidents**.

---

# Core Features

## Container Supply Chain Analysis

Analyze container images and their dependency chains.

Example:

```
chainrisk image ghcr.io/org/service:latest
```

ChainRisk inspects:

* base images
* OS packages
* application dependencies

This allows detection of risky container supply chains.

---

## SBOM-Based Dependency Graphs

ChainRisk converts Software Bill of Materials (SBOMs) into **dependency graphs**.

```
image
  ↓
grpc
  ↓
protobuf
  ↓
zlib
```

This graph structure enables deeper analysis of dependency relationships and risk propagation.

---

## Global Dependency Risk Graph

ChainRisk builds a **global dependency graph across container images**.

```
serviceA        serviceB        serviceC
   |                |               |
  grpc             grpc           grpc
     \              |             /
       -------- protobuf --------
                 |
               zlib
```

This allows detection of **ecosystem chokepoints**—dependencies that impact many workloads.

---

## Ecosystem Centrality Detection

ChainRisk calculates graph metrics such as:

* dependency centrality
* dependency depth
* blast radius

These metrics identify **high-impact dependencies** that represent systemic supply chain risk.

---

## Blast Radius Simulation

ChainRisk can simulate the impact of a compromised dependency.

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

This helps security teams understand **how far compromise could propagate**.

---

## Maintainer Trust Drift Detection

ChainRisk analyzes repository metadata to detect **maintainer trust drift**.

This occurs when:

* maintainers become inactive
* new contributors gain commit access
* release control shifts unexpectedly

These patterns were observed in incidents like the **XZ Utils backdoor attack**.

Trust drift signals increase the predicted probability of compromise.

---

## Probabilistic Risk Modeling

ChainRisk uses a probabilistic model to estimate compromise likelihood.

Dependencies move through states:

```
Healthy
 ↓
Maintainer Risk
 ↓
Vulnerable
 ↓
Compromised
```

This enables **predictive security analysis** rather than reactive vulnerability scanning.

---

## Container Image Verification

ChainRisk integrates with the Chainguard ecosystem to verify container artifacts.

Supported tooling includes:

* **Cosign** for container signature verification
* **Syft** for SBOM generation
* **SLSA provenance metadata**

Example workflow:

```
Container Image
      ↓
Cosign Signature Verification
      ↓
SBOM Extraction (Syft)
      ↓
Dependency Graph Construction
      ↓
Risk Analysis
```

These integrations ensure that analyzed artifacts are **authentic and trustworthy**.

---

## Dependency Surface Analysis

ChainRisk evaluates the **size and complexity of dependency trees**.

Example:

```
Base image: ubuntu
Packages: 210
Risk: HIGH
```

vs

```
Base image: distroless
Packages: 8
Risk: LOW
```

Smaller dependency surfaces generally reduce supply chain attack risk.

---

# Background Indexing Mode

ChainRisk can operate in an **indexing mode** that continuously analyzes container images and builds a global dependency graph.

```
Container Registries
       ↓
SBOM Extraction
       ↓
Dependency Graph
       ↓
Global Risk Graph
```

This allows the CLI to perform fast queries against a continuously updated dataset.

Example query:

```
chainrisk dependency protobuf
```

Output:

```
protobuf@4.25
Images affected: 312
Centrality: HIGH
Risk Score: MEDIUM
```

---

# How ChainRisk Differs from Existing Security Tools

Most security tools focus on **static vulnerability detection**.

Typical tools answer:

```
Does this dependency contain known CVEs?
```

ChainRisk instead focuses on **systemic supply chain risk**.

It answers questions like:

```
Which dependency is most critical in my ecosystem?
Which dependency is most likely to be compromised?
If dependency X is compromised, which services are affected?
Are maintainers showing suspicious trust drift?
```

---

# What Makes ChainRisk Unique

## Ecosystem-Level Risk Analysis

Instead of analyzing one project at a time, ChainRisk analyzes **entire dependency ecosystems**.

---

## Probabilistic Risk Forecasting

ChainRisk predicts **future compromise risk**, not just present vulnerabilities.

---

## Trust Drift Detection

ChainRisk detects risky maintainer patterns that often precede supply chain attacks.

---

## Container-Native Security

ChainRisk is designed specifically for **modern container and Kubernetes environments**.

---

## Supply Chain Intelligence

ChainRisk combines:

* dependency graph analysis
* container supply chain inspection
* maintainer trust signals
* probabilistic modeling

to provide **actionable risk intelligence**.

---

# Benefits for Security Professionals

ChainRisk helps security teams:

* identify high-risk dependencies early
* understand dependency blast radius
* detect ecosystem chokepoints
* monitor maintainer trust changes
* prioritize audits and remediation efforts

This allows organizations to **reduce supply chain risk proactively rather than reacting to vulnerabilities after they appear**.

---

# Vision

ChainRisk aims to become a **supply chain risk intelligence engine** that helps organizations understand and secure the complex dependency ecosystems powering modern cloud infrastructure.
