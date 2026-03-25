/*
    Copyright 2026 Suman Mandal (GitHub: jijo-OO7)

   Licensed under the Apache License, Version 2.0 (the "License");
   you may not use this file except in compliance with the License.
   You may obtain a copy of the License at

   http://www.apache.org/licenses/LICENSE-2.0
*/
package cli

import "fmt"

func PrintHeader(title string, emoji string) {
	if emoji != "" {
		fmt.Printf("\n=== %s %s ===\n\n", emoji, title)
	} else {
		fmt.Printf("\n=== %s ===\n\n", title)
	}
}

func PrintSection(title string, emoji string) {
	if emoji != "" {
		fmt.Println("-----------------------------")
	    fmt.Printf("%s %s\n", emoji, title)
	    fmt.Println("-----------------------------")
	} else {
		fmt.Println("-----------------------------")
	    fmt.Println(title)
	    fmt.Println("-----------------------------")
	}
}

