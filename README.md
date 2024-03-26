# go-xgb

XGB is the X protocol Go language Binding.

It is the Go equivalent of XCB, the X protocol C-language Binding
(http://xcb.freedesktop.org/).

### gruf's Fork

This is a fork of jezek's `go-xgb` repository. It is not yet complete... 
- ewmh, xcursor, icccm packages are unsupported
- tests are not yet updated

It is a complete rewrite of the dialer, underlying X connection and generated
code itself. Though the existing `xgbgen` code generator was still used as the
basis for parsing `xcb` definition files.

The primary differences intended were:
- more idiomatic Go code
- more idiomatic logging
- support debug output via `debug` build tag
- rely on memory pooling where possible

### jezek's Fork

I've forked the XGB repository from BurntSushi's github to apply some
patches which caused panics and memory leaks upon close and tests were added,
to test multiple server close scenarios.

### BurntSushi's Fork

I've forked the XGB repository from Google Code due to inactivty upstream.

Godoc documentation can be found here:
https://godoc.org/github.com/BurntSushi/xgb

Much of the code has been rewritten in an effort to support thread safety
and multiple extensions. Namely, go_client.py has been thrown away in favor
of an xgbgen package.

The biggest parts that *haven't* been rewritten by me are the connection and
authentication handshakes. They're inherently messy, and there's really no
reason to re-work them. The rest of XGB has been completely rewritten.

I like to release my code under the WTFPL, but since I'm starting with someone
else's work, I'm leaving the original license/contributor/author information
in tact.

I suppose I can legitimately release xgbgen under the WTFPL. To be fair, it is
at least as complex as XGB itself. *sigh*

Unless otherwise noted, the XGB source files are distributed
under the BSD-style license found in the LICENSE file.

Contributions should follow the same procedure as for the Go project:
http://golang.org/doc/contribute.html

