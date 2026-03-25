/*
    Copyright 2026 Suman Mandal (GitHub: jijo-OO7)

   Licensed under the Apache License, Version 2.0 (the "License");
   you may not use this file except in compliance with the License.
   You may obtain a copy of the License at

   http://www.apache.org/licenses/LICENSE-2.0
*/

package main

import (
	"fmt"
	"os"

	"github.com/jijo-OO7/chainrisk/internal/cli"
)

var version = "0.1.0"

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: chainrisk <command>")
		return
	}

	command := os.Args[1]

	switch command {
	case "version":
		fmt.Println("chainrisk version", version)

	case "sbom-info":
		if len(os.Args) < 3 {
			fmt.Println("Usage: chainrisk sbom-info <file>")
			return
		}
		file := os.Args[2]
		cli.RunSBOMInfo(file)
	
	case "blast":
	if len(os.Args) < 4 {
		fmt.Println("Usage: chainrisk blast <file> --target=<dependency>")
		return
	}

	file := os.Args[2]

	var target string
	for _, arg := range os.Args[3:] {
		if len(arg) > 9 && arg[:9] == "--target=" {
			target = arg[9:]
		}
	}

	if target == "" {
		fmt.Println("Error: --target is required")
		return
	}

	cli.RunBlast(file, target)

	default:
		fmt.Println("Unknown command:", command)
	}
}
