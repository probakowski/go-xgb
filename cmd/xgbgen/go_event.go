package main

import "strings"

// Event types
func (e *Event) Define(c *Context) {
	c.Putln("// %s is the event number for a %s.", e.SrcName(), e.EvType())
	c.Putln("const %s = %d", e.SrcName(), e.Number)
	c.Putln("")
	c.Putln("type %s struct {", e.EvType())
	if !e.NoSequence {
		c.Putln("	Sequence uint16")
	}
	for _, field := range e.Fields {
		field.Define(c)
	}
	c.Putln("}")
	c.Putln("")

	// Read defines a function that transforms a byte slice into this
	// event struct.
	e.Read(c)

	// Write defines a function that transforms this event struct into
	// a byte slice.
	e.Write(c)

	// Makes sure that this event type is an Event interface.
	c.Putln("// SeqID returns the sequence id attached to the %s event.", e.SrcName())
	c.Putln("// Events without a sequence number (KeymapNotify) return 0.")
	c.Putln("// This is mostly used internally.")
	c.Putln("func (v *%s) SeqID() uint16 {", e.EvType())
	if e.NoSequence {
		c.Putln("	return 0")
	} else {
		c.Putln("	return v.Sequence")
	}
	c.Putln("}")
	c.Putln("")

	// Let's the XGB event loop read this event.
	c.Putln("func init() { registerEvent(%d, Unmarshal%s) }", e.Number, e.EvType())
	c.Putln("")
}

func (e *Event) Read(c *Context) {
	c.Putln("// Unmarshal%s constructs a %s value that implements xgb.Event from a byte slice.", e.EvType(), e.EvType())
	c.Putln("func Unmarshal%s(buf []byte) (xgb.XEvent, error) {", e.EvType())
	c.Putln("	if len(buf) != %v {", e.Size())
	c.Putln("		return nil, fmt.Errorf(\"invalid data size %%d for \\\"%s\\\"\", len(buf))", e.EvType())
	c.Putln("	}")
	c.Putln("	")
	c.Putln("	v := &%s{}", e.EvType())
	c.Putln("	b := 1 // don't read event number")
	c.Putln("	")
	for i, field := range e.Fields {
		if i == 1 && !e.NoSequence {
			c.Putln("	v.Sequence = binary.LittleEndian.Uint16(buf[b:])")
			c.Putln("	b += 2")
			c.Putln("	")
		}
		field.Read(c, "v.")
		c.Putln("	")
	}
	c.Putln("	return v, nil")
	c.Putln("}")
	c.Putln("")
}

func (e *Event) Write(c *Context) {
	c.Putln("// Bytes writes a %s value to a byte slice.", e.EvType())
	c.Putln("func (v *%s) Bytes() []byte {", e.EvType())
	c.Putln("	buf := make([]byte, %s)", e.Size())
	c.Putln("	b := 0")
	c.Putln("	")
	c.Putln("	// write event number")
	c.Putln("	buf[b] = %d", e.Number)
	c.Putln("	b += 1")
	c.Putln("	")
	for i, field := range e.Fields {
		if i == 1 && !e.NoSequence {
			c.Putln("	b += 2 // skip sequence number")
			c.Putln("	")
		}
		field.Write(c, "v.")
		c.Putln("	")
	}
	c.Putln("	return buf")
	c.Putln("}")
	c.Putln("")
}

// EventCopy types
func (e *EventCopy) Define(c *Context) {
	c.Putln("// %s is the event number for a %s.", e.SrcName(), e.EvType())
	c.Putln("const %s = %d", e.SrcName(), e.Number)
	c.Putln("")
	c.Putln("type %s %s", e.EvType(), e.Old.(*Event).EvType())
	c.Putln("	")

	// Read defines a function that transforms a byte slice into this
	// event struct.
	e.Read(c)

	// Write defines a function that transoforms this event struct into
	// a byte slice.
	e.Write(c)

	// Makes sure that this event type is an Event interface.
	c.Putln("// SeqID returns the sequence id attached to the %s event.", e.SrcName())
	c.Putln("// Events without a sequence number (KeymapNotify) return 0.")
	c.Putln("// This is mostly used internally.")
	c.Putln("func (v *%s) SeqID() uint16 {", e.EvType())
	if e.Old.(*Event).NoSequence {
		c.Putln("	return uint16(0)")
	} else {
		c.Putln("	return v.Sequence")
	}
	c.Putln("}")
	c.Putln("")

	// Let's the XGB event loop read this event.
	c.Putln("func init() {")
	c.Putln("	registerEvent(%d, Unmarshal%s)", e.Number, e.EvType())
	c.Putln("}")
}

func (e *EventCopy) Read(c *Context) {
	c.Putln("// %sNew constructs a %s value that implements xgb.Event from a byte slice.", e.EvType(), e.EvType())
	c.Putln("func Unmarshal%s(buf []byte) (xgb.XEvent, error) {", e.EvType())

	var pkg string

	oevType := e.Old.(*Event).EvType()

	if strings.Contains(oevType, ".") {
		split := strings.Split(oevType, ".")
		pkg = split[0] + "."
		oevType = split[1]
	}

	c.Putln("	x, err := %sUnmarshal%s(buf)", pkg, oevType)
	c.Putln("	xev, _ := x.(*%s%s)", pkg, oevType)
	c.Putln("	return (*%s)(xev), err", e.EvType())
	c.Putln("}")
	c.Putln("")
}

func (e *EventCopy) Write(c *Context) {
	var pkg string

	oevType := e.Old.(*Event).EvType()

	if strings.Contains(oevType, ".") {
		split := strings.Split(oevType, ".")
		pkg = split[0] + "."
		oevType = split[1]
	}

	c.Putln("// Bytes writes a %s value to a byte slice.", e.EvType())
	c.Putln("func (v *%s) Bytes() []byte {", e.EvType())
	c.Putln("	buf := (*%s%s)(v).Bytes()", pkg, oevType)
	c.Putln("	buf[0] = %d", e.Number)
	c.Putln("	return buf")
	c.Putln("}")
	c.Putln("")
}
