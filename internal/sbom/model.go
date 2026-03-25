/*
    Copyright 2026 Suman Mandal (GitHub: jijo-OO7)

   Licensed under the Apache License, Version 2.0 (the "License");
   you may not use this file except in compliance with the License.
   You may obtain a copy of the License at

   http://www.apache.org/licenses/LICENSE-2.0
*/

package sbom

type SBOM struct {
	Packages      []Package      `json:"packages"`
	Relationships []Relationship `json:"relationships"`
}

type Package struct {
	SPDXID string `json:"SPDXID"`
	Name   string `json:"name"`
}

type Relationship struct {
	SpdxElementID      string `json:"spdxElementId"`
	RelatedSpdxElement string `json:"relatedSpdxElement"`
	RelationshipType   string `json:"relationshipType"`
}
