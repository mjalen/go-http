package lang

import (
	"go-http/http/syntax"
)

type Language struct {
	G syntax.Grammar
}

// closure for the HTTP syntax
func HTTP() {
	l := Language{
		G: syntax.Grammar{
			Rules: make(map[string]syntax.Rule, 0),		
		},
	}	

	// 
	l.G.Rule("HEX", `[A-F0-9]`)
	l.G.Rule("reserved", `[;/?:@&=+]`)
	l.G.Rule("extra", `[!*'(),]`)
	l.G.Rule("safe", `[$\-_\.]`)
	l.G.Rule("unsafe", `([[:cntrl:]]|[ "])`)

	l.G.Rule("escape", "%%(", ")")
	l.G.Rule("unreserved", "[a-z0-9]", "safe", "extra")
}
