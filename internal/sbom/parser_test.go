/*
    Copyright 2026 Suman Mandal (GitHub: jijo-OO7)

   Licensed under the Apache License, Version 2.0 (the "License");
   you may not use this file except in compliance with the License.
   You may obtain a copy of the License at

   http://www.apache.org/licenses/LICENSE-2.0
*/
package sbom

import(
	"testing"
)

func TestParseSBOM(t *testing.T){
	sbom,err := ParseSBOM("../../testdata/sample.json")

	if err !=nil {
		t.Fatalf("Expected no errors, got: %v ",err)
	}
	if len(sbom.Packages) != 3 {
	   t.Fatalf("expected 3 packages, got %d", len(sbom.packages))
	}
}