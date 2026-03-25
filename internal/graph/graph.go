/*
    Copyright 2026 Suman Mandal (GitHub: jijo-OO7)

   Licensed under the Apache License, Version 2.0 (the "License");
   you may not use this file except in compliance with the License.
   You may obtain a copy of the License at

   http://www.apache.org/licenses/LICENSE-2.0
*/

package graph

import "fmt"

// graph represents adjacency graph
type Graph struct {
	Adjacency map[string][]string
	Reverse   map[string][]string
}

// Print prints the graph
func (g *Graph) Print() {
	for node, neighbors := range g.Adjacency {
		for _, n := range neighbors {
			fmt.Printf("%s → %s\n", node, n)
		}
	}
}
