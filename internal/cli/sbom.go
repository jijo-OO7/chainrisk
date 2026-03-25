/*
	Copyright 2026 Suman Mandal (GitHub: jijo-OO7)

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

http://www.apache.org/licenses/LICENSE-2.0
*/
package cli

import (
	"fmt"
    "sort"
	"github.com/jijo-OO7/chainrisk/internal/graph"
	"github.com/jijo-OO7/chainrisk/internal/sbom"
)

func RunSBOMInfo(file string) {
	result, err := sbom.ParseSBOM(file)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	PrintHeader("SBOM INFO", "📊")

	fmt.Println("📦 Total Packages:", len(result.Packages))

	g := graph.BuildGraph(result)

	// Sort nodes
	var nodes []string
	for node := range g.Adjacency {
		nodes = append(nodes, node)
	}
	sort.Strings(nodes)

	fmt.Println("\n🔗 Dependency Graph:\n")

	for _, node := range nodes {
		for _, n := range g.Adjacency[node] {
			fmt.Println("  •", node, "→", n)
		}
	}
}

