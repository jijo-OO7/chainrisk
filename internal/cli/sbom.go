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

	"github.com/jijo-OO7/chainrisk/internal/sbom"
)

func RunSBOMInfo(file string) {
	result, err := sbom.ParseSBOM(file)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Total packages:", len(result.Packages))

	for _, pkg := range result.Packages {
		fmt.Println("-", pkg.Name)
	}
}