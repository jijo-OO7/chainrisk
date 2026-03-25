/*
	Copyright 2026 Suman Mandal (GitHub: jijo-OO7)

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

http://www.apache.org/licenses/LICENSE-2.0
*/
package graph

import "github.com/jijo-OO7/chainrisk/internal/sbom"

// BuildGraph constructs dependency graph from SBOM
func BuildGraph(s *sbom.SBOM) *Graph {
	g := &Graph{
		Adjacency: make(map[string][]string),
		Reverse: make(map[string][]string),
	}

	// Build package lookup: SPDXID -> name
	pkgMap := make(map[string]string)

	for _, pkg := range s.Packages {
		pkgMap[pkg.SPDXID] = pkg.Name
	}

	for _, pkg := range s.Packages {
		name := pkg.Name
		if _, exists := g.Adjacency[name]; !exists {
			g.Adjacency[name] = []string{}
		}
	}
 
	// Build edges
	for _, rel := range s.Relationships {
		from := pkgMap[rel.SpdxElementID]
		to := pkgMap[rel.RelatedSpdxElement]

		if from != "" && to != "" {
			g.Adjacency[from] = append(g.Adjacency[from], to)
            //for balast radius 
            g.Reverse[to] = append(g.Reverse[to], from)
		}
	}

	return g
}
