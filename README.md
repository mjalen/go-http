# Go HTTP (Proper name pending...)

An HTTP implementation written in go. Currently HTTP/0.9 compliant.

# Why?

This project was born from two goals:

- Learn the HTTP specification in a rigorous method. 
- Become comfortable reading specifications and memos.

This repository is a toy project that does not aim to replace current implementations. My main goal is to provide an API that is HTTP/1.0 compliant.

# Features

- HTTP/0.9 compliance. This means that the only possible request is the following a strict `GET <abs_path>\r\n`.
Requests are compliant in the sense that the above pattern is verified by *both* the user agent and origin server using regular expressions. 
Similarly, the origin server can provide responses for the requested `<abs_path>` by a stream of octets.
- The API also supports routing. Origin server routes can be defined and tied with a callback function; they are triggered when the proper route is requested.

# References 

- [RFC 1945](https://www.rfc-editor.org/rfc/rfc1945). The HTTP/1.0 specification.
- [re2 Syntax](https://github.com/google/re2/wiki/Syntax). The regular expression syntax utilized for translating the BNF grammar from RFC 1945.
- [net](https://pkg.go.dev/net). The package utilized for establishing TCP connections between a client and server.
