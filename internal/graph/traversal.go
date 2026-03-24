/*
    Copyright 2026 Suman Mandal (GitHub: jijo-OO7)

   Licensed under the Apache License, Version 2.0 (the "License");
   you may not use this file except in compliance with the License.
   You may obtain a copy of the License at

   http://www.apache.org/licenses/LICENSE-2.0
*/

package graph

func (g *Graph) BFS(start *Node, visit func(*Node)) {

	queue := []*Node{start}
	visited := map[string]bool{}

	for len(queue) > 0 {

		node := queue[0]
		queue = queue[1:]

		if visited[node.ID] {
			continue
		}

		visited[node.ID] = true

		visit(node)

		for _, dep := range node.Dependents {
			queue = append(queue, dep)
		}

	}

}
