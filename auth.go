package xgb

/*
auth.go contains functions to facilitate the parsing of .Xauthority files.

It is largely unmodified from the original XGB package that I forked.
*/

import (
	"encoding/binary"
	"errors"
	"io"
	"os"
)

// ReadAuthority reads the X authority file for the DISPLAY.
// Uses os.Hostname() if hostname empty. Returns authority name and data.
func ReadAuthority(hostname, display string) (string, []byte, error) {
	var err error

	// b is a scratch buffer to use and should be at least 256 bytes long
	// (i.e. it should be able to hold a hostname).
	b := make([]byte, 256)

	// As per /usr/include/X11/Xauth.h.
	const familyLocal = 256
	const familyWild = 65535

	// if hostname empty get
	if len(hostname) == 0 || hostname == "localhost" {
		hostname, err = os.Hostname()
		if err != nil {
			return "", nil, err
		}
	}

	// Fetch xauthority from shell env
	fname := os.Getenv("XAUTHORITY")

	// if none, use default $HOME/.Xauthority
	if len(fname) == 0 {
		home := os.Getenv("HOME")
		if len(home) == 0 {
			return "", nil, errors.New("xauthority not found: $XAUTHORITY, $HOME not set")
		}
		fname = home + "/.Xauthority"
	}

	// Open xauthority for reading
	r, err := os.Open(fname)
	if err != nil {
		return "", nil, err
	}
	defer r.Close()

	for {
		var family uint16
		if err := binary.Read(r, binary.BigEndian, &family); err != nil {
			return "", nil, err
		}

		addr, err := getString(r, b)
		if err != nil {
			return "", nil, err
		}

		disp, err := getString(r, b)
		if err != nil {
			return "", nil, err
		}

		name0, err := getString(r, b)
		if err != nil {
			return "", nil, err
		}

		data0, err := getBytes(r, b)
		if err != nil {
			return "", nil, err
		}

		addrmatch := (family == familyWild) ||
			(family == familyLocal && addr == hostname)
		dispmatch := (disp == "") || (disp == display)

		if addrmatch && dispmatch {
			return name0, data0, nil
		}
	}
}

func getBytes(r io.Reader, b []byte) ([]byte, error) {
	var n uint16

	if err := binary.Read(r, binary.BigEndian, &n); err != nil {
		return nil, err
	} else if n > uint16(len(b)) {
		return nil, errors.New("bytes too long for buffer")
	}

	if _, err := io.ReadFull(r, b[0:n]); err != nil {
		return nil, err
	}

	return b[0:n], nil
}

func getString(r io.Reader, b []byte) (string, error) {
	b, err := getBytes(r, b)
	return string(b), err
}
