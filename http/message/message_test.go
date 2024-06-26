package message

import (
	"testing"
)

func TestSyntax(t *testing.T) {
	cases := []struct {
		r string
		in string
		want bool
	}{
		{ r: "escape", in: `%F3`, want: true, },
		{ r: "escape", in: `%1G`, want: false, },

		{ r: "reserved", in: "?", want: true, },
		{ r: "reserved", in: "!", want: false, },

		{ r: "extra", in: "*", want: true, },
		{ r: "extra", in: "", want: false, }, 

		{ r: "safe", in: "_", want: true, },
		{ r: "safe", in: "+", want: false, },

		{ r: "unsafe", in: `"`, want: true, },
		{ r: "unsafe", in: "a", want: false, },

		{ r: "unreserved", in: "a", want: true },
		{ r: "unreserved", in: ";", want: false },
		
		{ r: "uchar", in: "0", want: true },
		{ r: "uchar", in: `\r`, want: false },

		{ r: "pchar", in: "0", want: true },
		{ r: "pchar", in: "+", want: true },

		{ r: "segment", in: "as7f8rh4", want: true },
		{ r: "segment", in: "", want: true },
		{ r: "segment", in: "a;/", want: false },

		{ r: "fsegment", in: "as8s3hf", want: true },
		{ r: "fsegment", in: "", want: false },
		{ r: "fsegment", in: "as/faser5", want: false },

		{ r: "path", in: "asdf/ew93i@", want: true },
		{ r: "path", in: "", want: false },
		{ r: "path", in: "a/", want: true },

		{ r: "rel_path", in: "", want: true },
		{ r: "rel_path", in: "asdf/hji91/as@.asd/", want: true },
		
		{ r: "abs_path", in: "/", want: true },
		{ r: "abs_path", in: "asd/", want: false }, // for now, implies a functional Request-URI 

		{ r: "Simple-Request", in: "GET /\r\n", want: true },
		{ r: "Simple-Request", in: "POST /\r\n", want: false },
		{ r: "Simple-Request", in: `GET /`, want: false},
		{ r: "Simple-Request", in: "GET /test/init\r\n", want: true },

		{ r: "Simple-Response", in: `a81infa[df91`, want: true },
	}

	for _, c := range cases {
		got := Validate(c.r, []byte(c.in))
		if got != c.want {
			t.Errorf("\"%s\".MatchString(\"%s\") == %t, want %t", c.r, c.in, got, c.want)	
		}
	}
}

