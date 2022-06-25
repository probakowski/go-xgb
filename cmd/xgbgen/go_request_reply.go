package main

import (
	"fmt"
	"strings"
)

func (r *Request) Define(c *Context) {
	if r.Reply != nil {
		c.Putln("// %s sends a checked request.", r.SrcName())
		c.Putln("func %s(c *xgb.XConn, %s) (%s, error) {", r.SrcName(), r.ParamNameTypes(), r.ReplyTypeName())
		c.Putln("	var reply %s", r.ReplyTypeName())
		r.ReadOpcode(c, true)
		c.Putln("	err := c.SendRecv(%s(op, %s), &reply)", r.ReqName(), r.ParamNames())
		c.Putln("	return reply, err")
		c.Putln("}")
		c.Putln("")

		c.Putln("// %sUnchecked sends an unchecked request.", r.SrcName())
		c.Putln("func %sUnchecked(c *xgb.XConn, %s) error {", r.SrcName(), r.ParamNameTypes())
		r.ReadOpcode(c, false)
		c.Putln("	return c.Send(%s(op, %s))", r.ReqName(), r.ParamNames())
		c.Putln("}")
		c.Putln("")

		r.ReadReply(c)
	} else {
		c.Putln("// %s sends a checked request.", r.SrcName())
		c.Putln("func %s(c *xgb.XConn, %s) error {", r.SrcName(), r.ParamNameTypes())
		r.ReadOpcode(c, false)
		c.Putln("	return c.SendRecv(%s(op, %s), nil)", r.ReqName(), r.ParamNames())
		c.Putln("}")
		c.Putln("")

		c.Putln("// %sUnchecked sends an unchecked request.", r.SrcName())
		c.Putln("func %sUnchecked(c *xgb.XConn, %s) error {", r.SrcName(), r.ParamNameTypes())
		r.ReadOpcode(c, false)
		c.Putln("	return c.Send(%s(op, %s))", r.ReqName(), r.ParamNames())
		c.Putln("}")
		c.Putln("")
	}
	r.WriteRequest(c)
}

func (r *Request) ReadReply(c *Context) {
	c.Putln("// %s represents the data returned from a %s request.",
		r.ReplyTypeName(), r.SrcName())
	c.Putln("type %s struct {", r.ReplyTypeName())
	c.Putln("	Sequence uint16 // sequence number of the request for this reply")
	c.Putln("	Length uint32 // number of bytes in this reply")
	for _, field := range r.Reply.Fields {
		field.Define(c)
	}
	c.Putln("}")
	c.Putln("")

	c.Putln("// Unmarshal reads a byte slice into a %s value.", r.ReplyTypeName())
	c.Putln("func (v *%s) Unmarshal(buf []byte) error {", r.ReplyTypeName())
	c.Putln("	if size := %s; len(buf) < size {", r.Reply.Size().Reduce("v."))
	c.Putln("		return fmt.Errorf(\"not enough data to unmarshal \\\"%s\\\": have=%%d need=%%d\", len(buf), size)", r.ReplyTypeName())
	c.Putln("	}")
	c.Putln("	")
	c.Putln("	b := 1 // skip reply determinant")
	c.Putln("	")
	for i, field := range r.Reply.Fields {
		field.Read(c, "	v.")
		c.Putln("	")
		if i == 0 {
			c.Putln("	v.Sequence = binary.LittleEndian.Uint16(buf[b:])")
			c.Putln("	b += 2")
			c.Putln("	")
			c.Putln("	v.Length = binary.LittleEndian.Uint32(buf[b:]) // 4-byte units")
			c.Putln("	b += 4")
			c.Putln("	")
		}
	}
	c.Putln("	return nil")
	c.Putln("}")
	c.Putln("")
}

func (r *Request) WriteRequest(c *Context) {
	sz := r.Size(c)
	writeSize1 := func() {
		if sz.exact {
			c.Putln("	binary.LittleEndian.PutUint16(buf[b:], uint16(size / 4)) // write request size in 4-byte units")
		} else {
			c.Putln("	blen := b")
		}
		c.Putln("	b += 2")
		c.Putln("	")
	}
	writeSize2 := func() {
		if sz.exact {
			c.Putln("	return buf")
			return
		}
		c.Putln("	b = internal.Pad4(b)")
		c.Putln("	binary.LittleEndian.PutUint16(buf[blen:], uint16(b / 4)) // write request size in 4-byte units")
		c.Putln("	return buf[:b]")
	}
	c.Putln("// Write request to wire for %s", r.SrcName())
	c.Putln("// %s writes a %s request to a byte slice.", r.ReqName(), r.SrcName())
	c.Putln("func %s(opcode uint8, %s) []byte {", r.ReqName(), r.ParamNameTypes())
	c.Putln("	size := %s", sz)
	c.Putln("	b := 0")
	c.Putln("	buf := make([]byte, size)")
	c.Putln("	")
	if c.protocol.isExt() {
		c.Putln("	buf[b] = opcode")
		c.Putln("	b += 1")
		c.Putln("	")
	}
	c.Putln("	buf[b] = %d // request opcode", r.Opcode)
	c.Putln("	b += 1")
	c.Putln("	")
	if len(r.Fields) == 0 {
		if !c.protocol.isExt() {
			c.Putln("	b += 1 // padding")
		}
		writeSize1()
	} else if c.protocol.isExt() {
		writeSize1()
	}
	for i, field := range r.Fields {
		field.Write(c, "")
		c.Putln("")
		if i == 0 && !c.protocol.isExt() {
			writeSize1()
		}
	}
	writeSize2()
	c.Putln("}")
	c.Putln("")
}

func (r *Request) ReadOpcode(c *Context, reply bool) {
	if c.protocol.isExt() {
		c.Putln("	op, ok := c.Ext(\"%s\")", c.protocol.ExtXName)
		c.Putln("	if !ok {")
		if reply {
			c.Putln("		return reply, errors.New(\"cannot issue request \\\"%s\\\" using the uninitialized extension \\\"%s\\\". %s.Register(xconn) must be called first.\")", r.SrcName(), c.protocol.ExtXName, c.protocol.PkgName())
		} else {
			c.Putln("		return errors.New(\"cannot issue request \\\"%s\\\" using the uninitialized extension \\\"%s\\\". %s.Register(xconn) must be called first.\")", r.SrcName(), c.protocol.ExtXName, c.protocol.PkgName())
		}
		c.Putln("	}")
	} else {
		c.Putln("	var op uint8")
	}
}

func (r *Request) ParamNames() string {
	names := make([]string, 0, len(r.Fields))
	for _, field := range r.Fields {
		switch f := field.(type) {
		case *ValueField:
			// mofos...
			if r.SrcName() != "ConfigureWindow" {
				names = append(names, f.MaskName)
			}
			names = append(names, f.ListName)
		case *PadField:
			continue
		case *ExprField:
			continue
		default:
			names = append(names, fmt.Sprintf("%s", field.SrcName()))
		}
	}
	return strings.Join(names, ", ")
}

func (r *Request) ParamNameTypes() string {
	nameTypes := make([]string, 0, len(r.Fields))
	for _, field := range r.Fields {
		switch f := field.(type) {
		case *ValueField:
			// mofos...
			if r.SrcName() != "ConfigureWindow" {
				nameTypes = append(nameTypes,
					fmt.Sprintf("%s %s", f.MaskName, f.MaskType.SrcName()))
			}
			nameTypes = append(nameTypes, fmt.Sprintf("%s []uint32", f.ListName))
		case *PadField:
			continue
		case *ExprField:
			continue
		default:
			nameTypes = append(nameTypes,
				fmt.Sprintf("%s %s", field.SrcName(), field.SrcType()))
		}
	}
	return strings.Join(nameTypes, ", ")
}
