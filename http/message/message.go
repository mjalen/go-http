package message

import (
	"regexp"
	"fmt"
	"log"
)

// regex versions of the RFC1945 BNF grammar for HTTP/1.0, 3.2.1
// Likely a better way, but decided to curry the syntax functions.
// Global syntax was kinda painful.
func syntax() (func(string, []byte) bool, func(string, []byte) []byte) {
	var s = map[string]string {
		"hex": `[A-F0-9]`,
		"reserved": `[;/?:@&=+]`,
		"extra": `[!*'(),]`,
		"safe": `[$\-_\.]`,
		"unsafe": `([[:cntrl:]]|[ "])`,
	}
	s["escape"] = fmt.Sprintf("%%(%s{2})", s["hex"]) 
	s["unreserved"] = fmt.Sprintf("([a-z0-9]|%s|%s)", s["safe"], s["extra"]) 

	s["uchar"] = fmt.Sprintf(`(%s|%s)`, s["unreserved"], s["escape"])
	s["pchar"] = fmt.Sprintf(`(%s|[:@&=+])`, s["uchar"])

	s["segment"] = fmt.Sprintf(`%s*`, s["pchar"])
	s["fsegment"] = fmt.Sprintf(`%s+`, s["pchar"])

	s["path"] = fmt.Sprintf(`%s(/%s)*`, s["fsegment"], s["segment"])
	s["rel_path"] = fmt.Sprintf(`(%s)?`, s["path"]) // TODO Add params and queries
	s["abs_path"] = fmt.Sprintf(`/%s`, s["rel_path"])

	s["Request-URI"] = fmt.Sprintf(`%s`, s["abs_path"]) // TODO Must contain absoluteURI
	s["Simple-Request"] = fmt.Sprintf(`GET %s\r\n`, s["Request-URI"]) 
	s["Simple-Response"] = fmt.Sprintf(`.*`) // TODO Verify octet charset

	validate_fn := func(pattern string, in []byte) bool {
		rx, err := regexp.Compile(fmt.Sprintf("^%s$", s[pattern]))	
		if err != nil {
			log.Printf("ERR (Validate(\"%s\")): %s", pattern, err) 
			return false
		}

		return rx.Match(in)
	}

	find_fn := func(pattern string, in []byte) []byte {
		rx, err := regexp.Compile(s[pattern])
		if err != nil {
			log.Printf("ERR (Validate(\"%s\")): %s", pattern, err) 
			return nil 
		}

		return rx.Find(in)
	}

	return validate_fn, find_fn
}

var Validate, Find = syntax() // then extract the curry functions
