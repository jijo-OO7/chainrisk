/*
    Copyright 2026 Suman Mandal (GitHub: jijo-OO7)

   Licensed under the Apache License, Version 2.0 (the "License");
   you may not use this file except in compliance with the License.
   You may obtain a copy of the License at

   http://www.apache.org/licenses/LICENSE-2.0
*/

package graph

type Graph struct {
	Nodes map[string]*Node
}

func NewGraph() *Graph {
	return &Graph{
		Nodes: make(map[string]*Node),
	}
}

func (g *Graph) AddNode(name, version string) *Node {

	id := name + "@" + version

	if node, exists := g.Nodes[id]; exists {
		return node
	}

	node := &Node{
		Name:    name,
		Version: version,
		ID:      id,
	}

	g.Nodes[id] = node

	return node
}

func (g *Graph) AddDependency(from, to *Node) {

	from.Dependencies = append(from.Dependencies, to)
	to.Dependents = append(to.Dependents, from)

}
