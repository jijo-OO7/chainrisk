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

	"github.com/jijo-OO7/chainrisk/internal/graph"
	"github.com/jijo-OO7/chainrisk/internal/sbom"
)

func RunBlast(file string, target string) {
	result, err := sbom.ParseSBOM(file)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	g := graph.BuildGraph(result)
	affected := g.BlastRadius(target)

	PrintHeader("BLAST RADIUS", "🚨")

	fmt.Println("🎯 Target: ", target)

	fmt.Println()

	PrintSection("Affected Components:", "📦")
	for _, node := range affected {
		fmt.Println("  •", node)
	}

	fmt.Println()
	fmt.Printf("⚡ Total Impact: %d components\n", len(affected))
}