// CookieJar - A contestant's algorithm toolbox
// Copyright 2013 Peter Szilagyi. All rights reserved.
//
// CookieJar is dual licensed: you can redistribute it and/or modify it under
// the terms of the GNU General Public License as published by the Free Software
// Foundation, either version 3 of the License, or (at your option) any later
// version.
//
// The toolbox is distributed in the hope that it will be useful, but WITHOUT
// ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or
// FITNESS FOR A PARTICULAR PURPOSE.  See the GNU General Public License for
// more details.
//
// Alternatively, the CookieJar toolbox may be used in accordance with the terms
// and conditions contained in a signed written agreement between you and the
// author(s).
//
// Author: peterke@gmail.com (Peter Szilagyi)

// Command depsmerge implements a command to retrieve and merge all dependencies
// of a package into a single file.
package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"go/ast"
	"go/parser"
	"go/printer"
	"go/token"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"
)

// Package description from the go list command
type Package struct {
	Name     string
	Dir      string
	Standard bool
	Deps     []string
	GoFiles  []string
}

// Loads the details of the Go package.
func details(name string) (*Package, error) {
	// Create the command to retrieve the package infos
	cmd := exec.Command("go", "list", "-e", "-json", name)

	// Retrieve the output, redirect the errors
	out, err := cmd.StdoutPipe()
	if err != nil {
		return nil, err
	}
	cmd.Stderr = os.Stderr

	// Start executing and parse the results
	if err := cmd.Start(); err != nil {
		return nil, err
	}
	defer cmd.Process.Kill()

	info := new(Package)
	if err := json.NewDecoder(out).Decode(&info); err != nil {
		return nil, err
	}
	// Clean up and return
	if err := cmd.Wait(); err != nil {
		return nil, err
	}
	return info, nil
}

// Collects all the imported packages of a file.
func dependencies(path string) (map[string][]string, error) {
	// Retrieve the dependencies of the source file
	info, err := details(path)
	if err != nil {
		return nil, err
	}
	// Iterate over every dependency and gather the sources
	sources := make(map[string][]string)
	for _, dep := range info.Deps {
		// Retrieve the dependency details
		pkgInfo, err := details(dep)
		if err != nil {
			return nil, err
		}
		// Gather external library sources
		if !pkgInfo.Standard {
			for _, src := range pkgInfo.GoFiles {
				sources[pkgInfo.Name] = append(sources[pkgInfo.Name], filepath.Join(pkgInfo.Dir, src))
			}
		}
	}
	return sources, nil
}

// Parses a source file and scopes all global declarations.
func rewrite(src string, pkg string, deps []string) (string, error) {
	fileSet := token.NewFileSet()
	tree, err := parser.ParseFile(fileSet, src, nil, parser.ParseComments)
	if err != nil {
		return "", err
	}
	// Scope all top level declarations if not main file
	if pkg != "" {
		for _, decl := range tree.Decls {
			switch v := decl.(type) {
			case *ast.FuncDecl:
				if v.Recv == nil {
					rename(tree, v.Name.String(), pkg+"•"+v.Name.String())
				}
			case *ast.GenDecl:
				for _, spec := range v.Specs {
					switch spec := spec.(type) {
					case *ast.ValueSpec:
						for _, name := range spec.Names {
							rename(tree, name.String(), pkg+"•"+name.String())
						}
					case *ast.TypeSpec:
						rename(tree, spec.Name.String(), pkg+"•"+spec.Name.String())
					default:
						fmt.Println("Unknown spec:", spec)
					}
				}
			default:
				fmt.Println(v)
			}
		}
	}
	// Generate the new source contents
	out := bytes.NewBuffer(nil)
	for _, decl := range tree.Decls {
		if gen, ok := decl.(*ast.GenDecl); !ok || gen.Tok != token.IMPORT {
			if err := printer.Fprint(out, fileSet, decl); err != nil {
				return "", err
			}
		}
		fmt.Fprintf(out, "\n\n")
	}
	blob := out.Bytes()

	// Scope all externally imported dependencies
	var fail error
	ast.Inspect(tree, func(node ast.Node) bool {
		if imp, ok := node.(*ast.GenDecl); ok && imp.Tok == token.IMPORT {
			for _, spec := range imp.Specs {
				if spec, ok := spec.(*ast.ImportSpec); ok {
					// Figure out the correct name of the import
					path := strings.Trim(spec.Path.Value, "\"")
					if info, err := details(path); err != nil {
						fail = err
						return false
					} else if !info.Standard {
						// Add scope to external import
						scoper := regexp.MustCompile("\\b" + info.Name + "\\.(.+)")
						blob = scoper.ReplaceAll(blob, []byte(info.Name+"•$1"))
					}
				}
			}
		}
		return true
	})
	if fail != nil {
		return "", fail
	}
	fmt.Println(string(blob))
	return "", nil
}

// Renames a top level declaration to something else.
func rename(tree *ast.File, old, new string) {
	// Rename top-level declarations
	for _, decl := range tree.Decls {
		switch decl := decl.(type) {
		case *ast.FuncDecl:
			// If a top level function matches, rename
			if decl.Recv == nil && decl.Name.Name == old {
				decl.Name.Name = new
				decl.Name.Obj.Name = new
			}
		case *ast.GenDecl:
			// Iterate over all the generic declaration
			for _, spec := range decl.Specs {
				switch spec := spec.(type) {
				case *ast.ValueSpec:
					// If a top level variable matches, rename
					for _, name := range spec.Names {
						if name.Name == old {
							name.Name = new
							name.Obj.Name = new
						}
					}
				case *ast.TypeSpec:
					if spec.Name.Name == old {
						spec.Name.Name = new
						spec.Name.Obj.Name = new
					}
				}
			}
		}
	}
	// Walk the AST and rename all internal occurrences
	stack := []ast.Node{}
	ast.Inspect(tree, func(node ast.Node) bool {
		// Keep a traversal stack if need to reference parent
		if node == nil {
			stack = stack[:len(stack)-1]
			return true
		}
		stack = append(stack, node)

		// Look for identifiers to rename
		id, ok := node.(*ast.Ident)
		if ok && id.Obj == nil && id.Name == old {
			// If selected identifier, leave it alone
			if _, ok := stack[len(stack)-2].(*ast.SelectorExpr); ok {
				return true
			}
			id.Name = new
		}
		if ok && id.Obj != nil && id.Name == old && id.Obj.Name == new {
			id.Name = id.Obj.Name
		}
		return true
	})
}

func main() {
	deps, err := dependencies(os.Args[1])
	if err != nil {
		log.Fatalf("Failed to parse dependency chain: %v.", err)
	}
	pkgs := []string{}
	for pkg, _ := range deps {
		pkgs = append(pkgs, pkg)
	}
	// Rewrite all the dependencies
	for pkg, sources := range deps {
		for _, src := range sources {
			fmt.Println(rewrite(src, pkg, pkgs))
		}
	}
	// Rewrite the main file itself
	fmt.Println(rewrite(os.Args[1], "", pkgs))
}
