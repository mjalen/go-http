package date

// The only reason these formats are re-implemented is so I can
// learn the date formats from the HTTP spec (see RFC2616 3.3.1).
// A better reason for the re-implementation is that the this spec makes
// it seem as though HTTP _requires_ GMT. I could be wrong though.
const (
	RFC1123 string = "Mon, 02 Jan 2006 15:04:05 GMT"
	RFC850 string = "Monday, 02-Jan-06 15:04:05 GMT"
	ASC string = "Mon Jan _2 15:04:05 2006"
)
