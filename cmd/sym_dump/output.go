package main

import "fmt"

// dumpTypes outputs the type information recorded by the parser as a C header.
func dumpTypes(p *parser) {
	// Print predeclared identifiers.
	def := p.types["bool"]
	fmt.Printf("%s;\n\n", def.Def())
	// Print enums.
	for _, tag := range p.enumTags {
		t := p.enums[tag]
		fmt.Printf("%s;\n\n", t.Def())
	}
	// Print structs.
	for _, tag := range p.structTags {
		t := p.structs[tag]
		fmt.Printf("%s;\n\n", t.Def())
	}
	// Print unions.
	for _, tag := range p.unionTags {
		t := p.unions[tag]
		fmt.Printf("%s;\n\n", t.Def())
	}
	// Print typedefs.
	for _, def := range p.typedefs {
		fmt.Printf("%s;\n\n", def.Def())
	}
}

// dumpDecls outputs the declarations recorded by the parser as a C header.
func dumpDecls(p *parser) {
	// Print variable declarations.
	for _, v := range p.vars {
		fmt.Printf("%s;\n\n", v)
	}
	// Print function declarations.
	for _, f := range p.funcs {
		fmt.Printf("%s;\n\n", f)
	}
}