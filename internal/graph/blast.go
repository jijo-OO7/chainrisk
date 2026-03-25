/*
	Copyright 2026 Suman Mandal (GitHub: jijo-OO7)

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

http://www.apache.org/licenses/LICENSE-2.0
*/

package graph

func (g *Graph) BlastRadius(start string) []string {
	visited := make(map[string]bool)
	queue := []string{start}
	result := []string{}

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		if visited[node] {
			continue
		}

		visited[node] = true
		result = append(result, node)

		for _, next := range g.Reverse[node] {
			if !visited[next] {
				queue = append(queue, next)
			}
		}
	}

	return result
}