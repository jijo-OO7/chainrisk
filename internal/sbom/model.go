/*
    Copyright 2026 Suman Mandal (GitHub: jijo-OO7)

   Licensed under the Apache License, Version 2.0 (the "License");
   you may not use this file except in compliance with the License.
   You may obtain a copy of the License at

   http://www.apache.org/licenses/LICENSE-2.0
*/

package sbom

type SBOM struct {
	Packages []Package `Json:"Packages"`
}

type Package struct {
	Name string `Json:"name"`
}
