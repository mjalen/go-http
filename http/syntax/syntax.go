package syntax 

import (
	"regexp"
	"log"
)

type Rule struct{
	raw []string
	compile []*regexp.Regexp
}

type Grammar struct {
	rules map[string]Rule
	
}

func (g *Grammar) Rule(name string, seq ...string) bool {
	r := Rule{
		raw: make([]string, len(seq)),
		compile: make([]*regexp.Regexp, len(seq)),
	}
	r.raw = seq
	for i, val := range seq {
		if _, ok := g.rules[val]; ok {
			continue
		}

		rx, err := regexp.Compile(val)
		if err != nil {
			log.Printf("ERR (Rule): %s", err)	
			return false
		}
		
		r.compile[i] = rx
	}

	g.rules[name] = r	
	return true
}

func (g *Grammar) Parse(start string, target []byte) bool {
	if _, ok := g.rules[start]; !ok {
		return false	
	}

	indices := make([][]int, len(g.rules[start].raw))
	for i := range indices {
		indices[i] = make([]int, 2)
	}

	for i, v := range g.rules[start].raw {
		if _, ok := g.rules[v]; ok {
			continue	
		}

		if curr := g.rules[start].compile[i].FindIndex(target); curr != nil {
			indices[i] = curr
		} else {
			return false
		}
	}

	for i, v := range g.rules[start].raw {
		if _, ok := g.rules[v]; !ok {
			continue
		}

		var b, e int
		if i == 0 {
			b = i	
		} else {
			b = indices[i - 1][1]
		}

		if i == len(g.rules[start].raw) - 1 {
			e = i	
		} else {
			e = indices[i+1][0]
		}

		if !g.Parse(v, target[b:e]) {
			return false
		}
	}

	return true
}

// func ModelPath(s string) {
// 
// }
// 
// // regex versions of the RFC1945 BNF grammar for HTTP/1.0, 3.2.1
// // Likely a better way, but decided to curry the syntax functions.
// // Global syntax was kinda painful.
// func syntax() (func(string, []byte) bool, func(string, []byte) []byte) {
// 	var s = map[string]string {
// 		"LWS": `(\r\n)?([ \t]){1,}`, // untested
// 		"HEX": `[A-F0-9]`,
// 		"tspecials": `[()<>@,;:\"/\[\]?={} \t]`, // untested
// 		"token": `[[:alpha:]^[()<>@,;:\"/\[\]?={} \t]]{1,}`, // untested
// 		"comment": `\([.*^[\(\)]]\)`, // untested
// 		"quoted-string": `"[[:alpha:]]*"` // untested
// 		"reserved": `[;/?:@&=+]`,
// 		"extra": `[!*'(),]`,
// 		"safe": `[$\-_\.]`,
// 		"unsafe": `([[:cntrl:]]|[ "])`,
// 	}
// 	s["escape"] = fmt.Sprintf("%%(%s{2})", s["HEX"]) 
// 	s["unreserved"] = fmt.Sprintf("([a-z0-9]|%s|%s)", s["safe"], s["extra"]) 
// 
// 	s["uchar"] = fmt.Sprintf(`(%s|%s)`, s["unreserved"], s["escape"])
// 	s["pchar"] = fmt.Sprintf(`(%s|[:@&=+])`, s["uchar"])
// 
// 	s["segment"] = fmt.Sprintf(`%s*`, s["pchar"])
// 	s["fsegment"] = fmt.Sprintf(`%s+`, s["pchar"])
// 
// 	s["path"] = fmt.Sprintf(`%s(/%s)*`, s["fsegment"], s["segment"])
// 	s["rel_path"] = fmt.Sprintf(`(%s)?`, s["path"]) // TODO Add params and queries
// 	s["abs_path"] = fmt.Sprintf(`/%s`, s["rel_path"])
// 
// 	s["Request-URI"] = fmt.Sprintf(`%s`, s["abs_path"]) // TODO Must contain absoluteURI
// 	s["Simple-Request"] = fmt.Sprintf(`GET %s\r\n`, s["Request-URI"]) 
// 	s["Simple-Response"] = fmt.Sprintf(`.*`) // TODO Verify octet charset
// 
// 	validate_fn := func(pattern string, in []byte) bool {
// 		rx, err := regexp.Compile(fmt.Sprintf("^%s$", s[pattern]))	
// 		if err != nil {
// 			log.Printf("ERR (Validate(\"%s\")): %s", pattern, err) 
// 			return false
// 		}
// 
// 		return rx.Match(in)
// 	}
// 
// 	find_fn := func(pattern string, in []byte) []byte {
// 		rx, err := regexp.Compile(s[pattern])
// 		if err != nil {
// 			log.Printf("ERR (Validate(\"%s\")): %s", pattern, err) 
// 			return nil 
// 		}
// 
// 		return rx.Find(in)
// 	}
// 
// 	return validate_fn, find_fn
// }
// 
// var Validate, Find = syntax() // then extract the curry functions
