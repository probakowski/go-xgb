package main

import (
	"strings"
)

// Error types
func (e *Error) Define(c *Context) {
	c.Putln("// %s is the error number for a %s.", e.ErrConst(), e.ErrConst())
	c.Putln("const %s = %d", e.ErrConst(), e.Number)
	c.Putln("")
	c.Putln("type %s struct {", e.ErrType())
	c.Putln("	Sequence uint16")
	c.Putln("	NiceName string")
	for _, field := range e.Fields {
		field.Define(c)
	}
	c.Putln("}")
	c.Putln("")

	// Read defines a function that transforms a byte slice into this
	// error struct.
	e.Read(c)

	// Makes sure this error type implements the xgb.Error interface.
	e.ImplementsError(c)

	// Let's the XGB event loop read this error.
	c.Putln("func init() { registerError(%d, Unmarshal%s) }", e.Number, e.ErrType())
	c.Putln("")
}

func (e *Error) Read(c *Context) {
	c.Putln("// Unmarshal%s constructs a %s value that implements xgb.Error from a byte slice.", e.ErrType(), e.ErrType())
	c.Putln("func Unmarshal%s(buf []byte) (xgb.XError, error) {", e.ErrType())
	c.Putln("	if len(buf) != %v {", e.Size())
	c.Putln("		return nil, fmt.Errorf(\"invalid data size %%d for \\\"%s\\\"\", len(buf))", e.ErrType())
	c.Putln("	}")
	c.Putln("	")
	c.Putln("	v := &%s{}", e.ErrType())
	c.Putln("	v.NiceName = \"%s\"", e.SrcName())
	c.Putln("	")
	c.Putln("	b := 1 // skip error determinant")
	c.Putln("	b += 1 // don't read error number")
	c.Putln("	")
	c.Putln("	v.Sequence = binary.LittleEndian.Uint16(buf[b:])")
	c.Putln("	b += 2")
	c.Putln("	")
	for _, field := range e.Fields {
		field.Read(c, "v.")
		c.Putln("")
	}
	c.Putln("	return v, nil")
	c.Putln("}")
	c.Putln("")
}

// ImplementsError writes functions to implement the XGB Error interface.
func (e *Error) ImplementsError(c *Context) {
	c.Putln("// SeqID returns the sequence id attached to the %s error.", e.ErrConst())
	c.Putln("// This is mostly used internally.")
	c.Putln("func (err *%s) SeqID() uint16 {", e.ErrType())
	c.Putln("	return err.Sequence")
	c.Putln("}")
	c.Putln("")
	c.Putln("// BadID returns the 'BadValue' number if one exists for the %s error. If no bad value exists, 0 is returned.", e.ErrConst())
	c.Putln("func (err *%s) BadID() uint32 {", e.ErrType())
	if !c.protocol.isExt() {
		c.Putln("return err.BadValue")
	} else {
		c.Putln("return 0")
	}
	c.Putln("}")
	c.Putln("// Error returns a rudimentary string representation of the %s error.", e.ErrConst())
	c.Putln("func (err *%s) Error() string {", e.ErrType())
	ErrorFieldString(c, e.Fields, e.ErrConst())
	c.Putln("}")
	c.Putln("")
}

// ErrorCopy types
func (e *ErrorCopy) Define(c *Context) {
	c.Putln("// %s is the error number for a %s.", e.ErrConst(), e.ErrConst())
	c.Putln("const %s = %d", e.ErrConst(), e.Number)
	c.Putln("")
	c.Putln("type %s %s", e.ErrType(), e.Old.(*Error).ErrType())
	c.Putln("")

	// Read defines a function that transforms a byte slice into this
	// error struct.
	e.Read(c)

	// Makes sure this error type implements the xgb.Error interface.
	e.ImplementsError(c)

	// Let's the XGB know how to read this error.
	c.Putln("func init() {")
	c.Putln("	registerError(%d, Unmarshal%s)", e.Number, e.ErrType())
	c.Putln("}")
	c.Putln("")
}

func (e *ErrorCopy) Read(c *Context) {
	c.Putln("// %sNew constructs a %s value that implements xgb.Error from a byte slice.", e.ErrType(), e.ErrType())
	c.Putln("func Unmarshal%s(buf []byte) (xgb.XError, error) {", e.ErrType())

	var pkg string

	errType := e.Old.(*Error).ErrType()

	if strings.Contains(errType, ".") {
		split := strings.Split(errType, ".")
		pkg = split[0] + "."
		errType = split[1]
	}

	c.Putln("	x, err := %sUnmarshal%s(buf)", pkg, errType)
	c.Putln("	xerr, _ := x.(*%s%s)", pkg, errType)
	c.Putln("	return (*%s)(xerr), err", e.ErrType())
	c.Putln("}")
	c.Putln("")
}

// ImplementsError writes functions to implement the XGB Error interface.
func (e *ErrorCopy) ImplementsError(c *Context) {
	c.Putln("// SequenceId returns the sequence id attached to the %s error.",
		e.ErrConst())
	c.Putln("// This is mostly used internally.")
	c.Putln("func (err *%s) SeqID() uint16 {", e.ErrType())
	c.Putln("	return err.Sequence")
	c.Putln("}")
	c.Putln("")
	c.Putln("// BadId returns the 'BadValue' number if one exists for the %s error. If no bad value exists, 0 is returned.", e.ErrConst())
	c.Putln("func (err *%s) BadID() uint32 {", e.ErrType())
	if !c.protocol.isExt() {
		c.Putln("	return err.BadValue")
	} else {
		c.Putln("	return 0")
	}
	c.Putln("}")
	c.Putln("")
	c.Putln("// Error returns a rudimentary string representation of the %s error.", e.ErrConst())
	c.Putln("func (err *%s) Error() string {", e.ErrType())
	ErrorFieldString(c, e.Old.(*Error).Fields, e.ErrConst())
	c.Putln("}")
	c.Putln("")
}

// ErrorFieldString works for both Error and ErrorCopy. It assembles all of the
// fields in an error and formats them into a single string.
func ErrorFieldString(c *Context, fields []Field, errName string) {
	c.Putln("	var buf strings.Builder")
	c.Putln("	")
	c.Putln("	buf.WriteString(\"%s{\")", errName)
	c.Putln("	buf.WriteString(\"NiceName: \"+err.NiceName)")
	c.Putln("	buf.WriteByte(' ')")
	c.Putln("	buf.WriteString(\"Sequence: \"+strconv.FormatUint(uint64(err.Sequence), 10))")
	if len(fields) > 0 {
		// Only add a space separator if not final field
		c.Putln("	buf.WriteByte(' ')")
	}
	c.Putln("	")
	for i, field := range fields {
		switch field.(type) {
		case *PadField:
			continue
		default:
			if field.SrcType() == "string" {
				c.Putln("	buf.WriteString(\"%s: \" + err.%s)", field.SrcName(), field.SrcName())
			} else {
				c.Putln("	fmt.Fprintf(&buf, \"%s: %%d\", err.%s)", field.SrcName(), field.SrcName())
			}
		}

		if i != len(fields)-1 {
			// All but last entry need spacing
			c.Putln("	buf.WriteString(\", \")")
		}
	}
	c.Putln("	buf.WriteByte('}')")
	c.Putln("	")
	c.Putln("	return buf.String()")
}
