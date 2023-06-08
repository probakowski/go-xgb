package main

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"log"
	"sort"
)

// Context represents the protocol we're converting to Go, and a writer
// buffer to write the Go source to.
type Context struct {
	protocol *Protocol
	out      *bytes.Buffer
}

func newContext() *Context {
	return &Context{
		out: bytes.NewBuffer([]byte{}),
	}
}

// Putln calls put and adds a new line to the end of 'format'.
func (c *Context) Putln(format string, v ...interface{}) {
	c.Put(format+"\n", v...)
}

// Put is a short alias to write to 'out'.
func (c *Context) Put(format string, v ...interface{}) {
	_, err := fmt.Fprintf(c.out, format, v...)
	if err != nil {
		log.Fatalf("There was an error writing to context buffer: %s", err)
	}
}

// Morph is the big daddy of them all. It takes in an XML byte slice,
// parse it, transforms the XML types into more usable types,
// and writes Go code to the 'out' buffer.
func (c *Context) Morph(xmlBytes []byte) {
	parsedXML := &XML{}
	err := xml.Unmarshal(xmlBytes, parsedXML)
	if err != nil {
		log.Fatal(err)
	}

	// Parse all imports
	parsedXML.Imports.Eval()

	// Translate XML types to nice types
	c.protocol = parsedXML.Translate(nil)

	c.protocol.AddAlignGaps()

	// Start with Go header.
	c.Putln("// FILE GENERATED AUTOMATICALLY FROM \"%s.xml\"", c.protocol.Name)
	c.Putln("package %s", c.protocol.PkgName())
	c.Putln("")

	// Write imports. We always need to import at least xgb.
	// We also need to import xproto if it's an extension.
	c.Putln("import (")
	if !c.protocol.isExt() {
		c.Putln("	_ \"unsafe\"")
	}
	c.Putln("	\"codeberg.org/gruf/go-xgb\"")
	c.Putln("	\"codeberg.org/gruf/go-xgb/util\"")
	sort.Sort(Protocols(c.protocol.Imports))
	for _, imp := range c.protocol.Imports {
		c.Putln("	\"codeberg.org/gruf/go-xgb/%s\"", imp.Name)
	}
	c.Putln(")")
	c.Putln("")

	if c.protocol.isExt() {
		c.Putln("const (")
		c.Putln("	// ExtName is the user-friendly name string of this X extension.")
		c.Putln("	ExtName = \"%s\"", c.protocol.ExtName)
		c.Putln("")
		c.Putln("	// ExtXName is the name string this extension is known by to the X server.")
		c.Putln("	ExtXName = \"%s\"", c.protocol.ExtXName)
		c.Putln(")")
		c.Putln("")
		c.Putln("var (")
		c.Putln("	// generated index maps of defined event and error numbers -> unmarshalers.")
		c.Putln("	eventFuncs = make(map[uint8]xgb.EventUnmarshaler)")
		c.Putln("	errorFuncs = make(map[uint8]xgb.ErrorUnmarshaler)")
		c.Putln(")")
		c.Putln("")
		c.Putln("func registerEvent(n uint8, fn xgb.EventUnmarshaler) {")
		c.Putln("	if _, ok := eventFuncs[n]; ok {")
		c.Putln("		panic(\"BUG: overlapping event unmarshaler\")")
		c.Putln("	}")
		c.Putln("	eventFuncs[n] = fn")
		c.Putln("}")
		c.Putln("")
		c.Putln("func registerError(n uint8, fn xgb.ErrorUnmarshaler) {")
		c.Putln("	if _, ok := errorFuncs[n]; ok {")
		c.Putln("		panic(\"BUG: overlapping error unmarshaler\")")
		c.Putln("	}")
		c.Putln("	errorFuncs[n] = fn")
		c.Putln("}")
		c.Putln("")
		c.Putln("// Register will query the X server for %s extension support, and register relevant extension unmarshalers with the XConn.", c.protocol.ExtName)
		c.Putln("func Register(xconn *xgb.XConn) error {")
		c.Putln("	// Query the X server for this extension")
		c.Putln("	reply, err := xproto.QueryExtension(xconn, uint16(len(ExtXName)), ExtXName)")
		c.Putln("	if err != nil {")
		c.Putln("		return fmt.Errorf(\"error querying X for \\\"%s\\\": %%w\", err)", c.protocol.ExtName)
		c.Putln("	} else if !reply.Present {")
		c.Putln("		return fmt.Errorf(\"no extension named \\\"%s\\\" is known to the X server: reply=%%+v\", reply)", c.protocol.ExtName)
		c.Putln("	}")
		c.Putln("	")
		c.Putln("	// Clone event funcs map but set our event no. start index")
		c.Putln("	extEventFuncs := make(map[uint8]xgb.EventUnmarshaler, len(eventFuncs))")
		c.Putln("	for n, fn := range eventFuncs {")
		c.Putln("		extEventFuncs[n+reply.FirstEvent] = fn")
		c.Putln("	}")
		c.Putln("	")
		c.Putln("	// Clone error funcs map but set our error no. start index")
		c.Putln("	extErrorFuncs := make(map[uint8]xgb.ErrorUnmarshaler, len(errorFuncs))")
		c.Putln("	for n, fn := range errorFuncs {")
		c.Putln("		extErrorFuncs[n+reply.FirstError] = fn")
		c.Putln("	}")
		c.Putln("	")
		c.Putln("	// Register ourselves with the X server connection")
		c.Putln("	return xconn.Register(xgb.XExtension{")
		c.Putln("		XName:       ExtXName,")
		c.Putln("		MajorOpcode: reply.MajorOpcode,")
		c.Putln("		EventFuncs:  extEventFuncs,")
		c.Putln("		ErrorFuncs:  extErrorFuncs,")
		c.Putln("	})")
		c.Putln("}")
		c.Putln("")
	} else {
		// In the xproto package, we must provide a Setup function that uses
		// SetupBytes in xgb.Conn to return a SetupInfo structure.
		c.Putln("var (")
		c.Putln("	// generated index maps of defined event and error numbers -> unmarshalers.")
		c.Putln("	eventFuncs = make(map[uint8]xgb.EventUnmarshaler)")
		c.Putln("	errorFuncs = make(map[uint8]xgb.ErrorUnmarshaler)")
		c.Putln(")")
		c.Putln("")
		c.Putln("// registerEvent will register an event unmarshaler in global map, panics on overlap")
		c.Putln("func registerEvent(n uint8, fn xgb.EventUnmarshaler) {")
		c.Putln("	if _, ok := eventFuncs[n]; ok {")
		c.Putln("		panic(\"BUG: overlapping event unmarshaler\")")
		c.Putln("	}")
		c.Putln("	eventFuncs[n] = fn")
		c.Putln("}")
		c.Putln("")
		c.Putln("// registerError will register an error unmarshaler in global map, panics on overlap")
		c.Putln("func registerError(n uint8, fn xgb.ErrorUnmarshaler) {")
		c.Putln("	if _, ok := errorFuncs[n]; ok {")
		c.Putln("		panic(\"BUG: overlapping error unmarshaler\")")
		c.Putln("	}")
		c.Putln("	errorFuncs[n] = fn")
		c.Putln("}")
		c.Putln("")
		c.Putln("// sorcery to give us access to package-private functions.")
		c.Putln("//")
		c.Putln("//go:linkname xproto_init codeberg.org/gruf/go-xgb.xproto_init")
		c.Putln("func xproto_init(*xgb.XConn, map[uint8]xgb.EventUnmarshaler, map[uint8]xgb.ErrorUnmarshaler) error")
		c.Putln("")
		c.Putln("// Setup ...")
		c.Putln("func Setup(xconn *xgb.XConn, buf []byte) (*SetupInfo, error) {")
		c.Putln("	// Register ourselves with the X server connection")
		c.Putln("	if err := xproto_init(xconn, eventFuncs, errorFuncs); err != nil {")
		c.Putln("		return nil, err")
		c.Putln("	}")
		c.Putln("	")
		c.Putln("	info := &SetupInfo{}")
		c.Putln("	")
		c.Putln("	// Read setup information from buf")
		c.Putln("	_ = SetupInfoRead(buf, info)")
		c.Putln("	")
		c.Putln("	return info, nil")
		c.Putln("}")
		c.Putln("")
	}

	// Now write Go source code
	sort.Sort(Types(c.protocol.Types))
	sort.Sort(Requests(c.protocol.Requests))
	for _, typ := range c.protocol.Types {
		typ.Define(c)
	}
	for _, req := range c.protocol.Requests {
		req.Define(c)
	}
}
