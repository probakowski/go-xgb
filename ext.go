package xgb

type XEvent interface {
	// SeqID returns the X sequence ID this event is associated with
	SeqID() uint16
}

type XError interface {
	// SeqID returns the X sequence ID this error is associated with
	SeqID() uint16

	// implement standard error
	error
}

type XReply interface {
	Unmarshal([]byte) error
}

// EventUnmarshaler ...
type EventUnmarshaler func([]byte) (XEvent, error)

// ErrorUnmarshaler ...
type ErrorUnmarshaler func([]byte) (XError, error)

// XExtension ...
type XExtension struct {
	XName       string
	MajorOpcode uint8
	EventFuncs  map[uint8]EventUnmarshaler
	ErrorFuncs  map[uint8]ErrorUnmarshaler
}

// RawXReply is a byte slice type alias that
// fulfills the XReply interface type, allowing
// you to store an X reply for later decoding.
type RawXReply []byte

func (rpl *RawXReply) Unmarshal(data []byte) error {
	(*rpl) = append((*rpl)[:0], data...)
	return nil
}

// ignoreXReply is a noop XReply implementation.
type IgnoreXReply struct{}

func (IgnoreXReply) Unmarshal([]byte) error { return nil }
