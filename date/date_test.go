package date

import (
	"testing"
	"time"
)

func TestFormats(t *testing.T) {
	cases := []struct {
		in time.Time
		wantRFC1123 string
		wantRFC850 string
		wantASC string
	}{
		{ 
			in: time.Date(2024, time.June, 1, 0, 0, 0, 0, time.UTC), 
			wantRFC1123: "Sat, 01 Jun 2024 00:00:00 GMT", 
			wantRFC850: "Saturday, 01-Jun-24 00:00:00 GMT",
			wantASC: "Sat Jun  1 00:00:00 2024",
		},
		{ 
			in: time.Date(2002, time.September, 25, 0, 0, 0, 0, time.UTC), 
			wantRFC1123: "Wed, 25 Sep 2002 00:00:00 GMT", 
			wantRFC850: "Wednesday, 25-Sep-02 00:00:00 GMT",
			wantASC: "Wed Sep 25 00:00:00 2002",
		},
		{
			in: time.Date(1987, time.April, 30, 5, 43, 22, 33, time.UTC), 
			wantRFC1123: "Thu, 30 Apr 1987 05:43:22 GMT",
			wantRFC850: "Thursday, 30-Apr-87 05:43:22 GMT",
			wantASC: "Thu Apr 30 05:43:22 1987",
		},
		{
			in: time.Date(1987, time.April, 31, 18, 22, 32, 0, time.UTC),
			wantRFC1123: "Fri, 01 May 1987 18:22:32 GMT",
			wantRFC850: "Friday, 01-May-87 18:22:32 GMT",
			wantASC: "Fri May  1 18:22:32 1987",
		},
	}	

	for _, c := range cases {
		gotRFC1123 := c.in.Format(RFC1123) 
		gotRFC850 := c.in.Format(RFC850)
		gotASC := c.in.Format(ASC)
		if gotRFC1123 != c.wantRFC1123 {
			t.Errorf("RFC1123 from (%+v) == %s, want %s.", c.in, gotRFC1123, c.wantRFC1123)
		}
		if gotRFC850 != c.wantRFC850 {
			t.Errorf("RFC850 from (%+v) == %s, want %s.", c.in, gotRFC850, c.wantRFC850)
		}
		if gotASC != c.wantASC {
			t.Errorf("Ansi-C from (%+v) == %s, want %s.", c.in, gotASC, c.wantASC)
		}
	}
}
