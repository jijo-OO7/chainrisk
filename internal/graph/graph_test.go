/*
    Copyright 2026 Suman Mandal (GitHub: jijo-OO7)

   Licensed under the Apache License, Version 2.0 (the "License");
   you may not use this file except in compliance with the License.
   You may obtain a copy of the License at

   http://www.apache.org/licenses/LICENSE-2.0
*/

package graph

import "testing"

func TestGraphCreation(t *testing.T) {

	g := NewGraph()

	service := g.AddNode("payment-service", "1.0")
	grpc := g.AddNode("grpc", "1.62")

	g.AddDependency(service, grpc)

	if len(service.Dependencies) != 1 {
		t.Error("expected service to have 1 dependency")
	}

	if service.Dependencies[0].Name != "grpc" {
		t.Error("expected dependency grpc")
	}
}
