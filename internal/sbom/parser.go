/*
	Copyright 2026 Suman Mandal (GitHub: jijo-OO7)

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

http://www.apache.org/licenses/LICENSE-2.0
*/
package sbom

import (
	"encoding/json"
	"os"
)

// ParseSBOM reads a file and parses it into SBOM struct
func ParseSBOM(file string) (*SBOM, error) {
	data, err := os.ReadFile(file)
	if err != nil {
		return nil, err
	}

	var sbom SBOM
	err = json.Unmarshal(data, &sbom)
	if err != nil {
		return nil, err
	}

	return &sbom, nil
}
