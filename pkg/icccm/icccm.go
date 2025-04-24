package icccm

import (
	"codeberg.org/gruf/go-byteutil"
	"encoding/binary"
	"fmt"
	"unsafe"

	"github.com/probakowski/go-xgb/xproto"

	"github.com/probakowski/go-xgb/pkg/xprop"
)

const (
	HintInput = (1 << iota)
	HintState
	HintIconPixmap
	HintIconWindow
	HintIconPosition
	HintIconMask
	HintWindowGroup
	HintMessage
	HintUrgency
)

const (
	SizeHintUSPosition = (1 << iota)
	SizeHintUSSize
	SizeHintPPosition
	SizeHintPSize
	SizeHintPMinSize
	SizeHintPMaxSize
	SizeHintPResizeInc
	SizeHintPAspect
	SizeHintPBaseSize
	SizeHintPWinGravity
)

const (
	StateWithdrawn = iota
	StateNormal
	StateZoomed
	StateIconic
	StateInactive
)

// WM_NAME get
func WmNameGet(xconn *xprop.XPropConn, win xproto.Window) (string, error) {
	reply, err := xconn.GetPropName(win, "WM_NAME")
	if err != nil {
		return "", err
	}
	return xprop.PropValStr(reply)
}

// WM_NAME set
func WmNameSet(xconn *xprop.XPropConn, win xproto.Window, name string) error {
	return xconn.ChangePropName(win, 8, "WM_NAME", "STRING", byteutil.S2B(name))
}

// WM_ICON_NAME get
func WmIconNameGet(xconn *xprop.XPropConn, win xproto.Window) (string, error) {
	reply, err := xconn.GetPropName(win, "WM_ICON_NAME")
	if err != nil {
		return "", err
	}
	return xprop.PropValStr(reply)
}

// WM_ICON_NAME set
func WmIconNameSet(xconn *xprop.XPropConn, win xproto.Window, name string) error {
	return xconn.ChangePropName(win, 8, "WM_ICON_NAME", "STRING", byteutil.S2B(name))
}

// NormalHints is a struct that organizes the information related to the
// WM_NORMAL_HINTS property. Please see the ICCCM spec for more details.
type NormalHints struct {
	Flags                                                   uint32
	X, Y                                                    int32
	Width, Height, MinWidth, MinHeight, MaxWidth, MaxHeight uint32
	WidthInc, HeightInc                                     uint32
	MinAspectNum, MinAspectDen, MaxAspectNum, MaxAspectDen  uint32
	BaseWidth, BaseHeight, WinGravity                       uint32
}

// WM_NORMAL_HINTS get
func WmNormalHintsGet(xconn *xprop.XPropConn, win xproto.Window) (*NormalHints, error) {
	reply, err := xconn.GetPropName(win, "WM_NORMAL_HINTS")
	if err != nil {
		return nil, err
	}

	hints, err := xprop.PropValUint32s(reply)
	if err != nil {
		return nil, err
	}

	if len(hints) != 18 {
		return nil, fmt.Errorf("WmNormalHint: There are %d fields in WM_NORMAL_HINTS, but xgbutil expects %d.", len(hints), 18)
	}

	nh := &NormalHints{}
	nh.Flags = hints[0]
	nh.X = int32(hints[1])
	nh.Y = int32(hints[2])
	nh.Width = hints[3]
	nh.Height = hints[4]
	nh.MinWidth = hints[5]
	nh.MinHeight = hints[6]
	nh.MaxWidth = hints[7]
	nh.MaxHeight = hints[8]
	nh.WidthInc = hints[9]
	nh.HeightInc = hints[10]
	nh.MinAspectNum = hints[11]
	nh.MinAspectDen = hints[12]
	nh.MaxAspectNum = hints[13]
	nh.MaxAspectDen = hints[14]
	nh.BaseWidth = hints[15]
	nh.BaseHeight = hints[16]
	nh.WinGravity = hints[17]

	if nh.WinGravity <= 0 {
		nh.WinGravity = xproto.GravityNorthWest
	}

	return nh, nil
}

// WM_NORMAL_HINTS set
// Make sure to set the flags in the NormalHints struct correctly!
func WmNormalHintsSet(xconn *xprop.XPropConn, win xproto.Window, nh *NormalHints) error {
	return xconn.ChangePropName(win, 32, "WM_NORMAL_HINTS", "WM_SIZE_HINTS", xprop.FromData32([]uint32{
		nh.Flags,
		uint32(nh.X), uint32(nh.Y), nh.Width, nh.Height,
		nh.MinWidth, nh.MinHeight,
		nh.MaxWidth, nh.MaxHeight,
		nh.WidthInc, nh.HeightInc,
		nh.MinAspectNum, nh.MinAspectDen,
		nh.MaxAspectNum, nh.MaxAspectDen,
		nh.BaseWidth, nh.BaseHeight,
		nh.WinGravity,
	}))
}

// Hints is a struct that organizes information related to the WM_HINTS
// property. Once again, I refer you to the ICCCM spec for documentation.
type Hints struct {
	Flags                   uint32
	Input, InitialState     uint32
	IconX, IconY            int32
	IconPixmap, IconMask    xproto.Pixmap
	WindowGroup, IconWindow xproto.Window
}

// WM_HINTS get
func WmHintsGet(xconn *xprop.XPropConn, win xproto.Window) (*Hints, error) {
	reply, err := xconn.GetPropName(win, "WM_HINTS")
	if err != nil {
		return nil, err
	}

	raw, err := xprop.PropValUint32s(reply)
	if err != nil {
		return nil, err
	}

	if len(raw) != 9 {
		return nil, fmt.Errorf("WmHints: There are %d fields in WM_HINTS, but xgbutil expects %d.", len(raw), 9)
	}

	hints := &Hints{}
	hints.Flags = raw[0]
	hints.Input = raw[1]
	hints.InitialState = raw[2]
	hints.IconPixmap = xproto.Pixmap(raw[3])
	hints.IconWindow = xproto.Window(raw[4])
	hints.IconX = int32(raw[5])
	hints.IconY = int32(raw[6])
	hints.IconMask = xproto.Pixmap(raw[7])
	hints.WindowGroup = xproto.Window(raw[8])

	return hints, nil
}

// WM_HINTS set. Make sure to set the flags in the Hints struct correctly!
func WmHintsSet(xconn *xprop.XPropConn, win xproto.Window, hints *Hints) error {
	return xconn.ChangePropName(win, 32, "WM_HINTS", "WM_HINTS", xprop.FromData32([]uint32{
		hints.Flags, hints.Input, hints.InitialState,
		uint32(hints.IconPixmap), uint32(hints.IconWindow),
		uint32(hints.IconX), uint32(hints.IconY),
		uint32(hints.IconMask),
		uint32(hints.WindowGroup),
	}))
}

// WmClass struct contains two data points: the instance and a class of a window.
type WmClass struct {
	Instance, Class string
}

// WM_CLASS get
func WmClassGet(xconn *xprop.XPropConn, win xproto.Window) (*WmClass, error) {
	reply, err := xconn.GetPropName(win, "WM_CLASS")
	if err != nil {
		return nil, err
	}
	raw, err := xprop.PropValStrs(reply)
	if err != nil {
		return nil, err
	}
	if len(raw) != 2 {
		return nil, fmt.Errorf("WmClass: Two string make up WM_CLASS, but xgbutil found %d in '%v'.", len(raw), raw)
	}
	return &WmClass{
		Instance: raw[0],
		Class:    raw[1],
	}, nil
}

// WM_CLASS set
func WmClassSet(xconn *xprop.XPropConn, win xproto.Window, class *WmClass) error {
	raw := make([]byte, len(class.Instance)+len(class.Class)+2)
	copy(raw, class.Instance)
	copy(raw[(len(class.Instance)+1):], class.Class)
	return xconn.ChangePropName(win, 8, "WM_CLASS", "STRING", raw)
}

// WM_TRANSIENT_FOR get
func WmTransientForGet(xconn *xprop.XPropConn, win xproto.Window) (xproto.Window, error) {
	reply, err := xconn.GetPropName(win, "WM_TRANSIENT_FOR")
	if err != nil {
		return 0, err
	}
	return xprop.PropValWindow(reply)
}

// WM_TRANSIENT_FOR set
func WmTransientForSet(xconn *xprop.XPropConn, win xproto.Window, transient xproto.Window) error {
	return xconn.ChangePropName(win, 32, "WM_TRANSIENT_FOR", "WINDOW", xprop.FromData32([]uint32{uint32(transient)}))
}

// WM_PROTOCOLS get
func WmProtocolsGet(xconn *xprop.XPropConn, win xproto.Window) ([]string, error) {
	reply, err := xconn.GetPropName(win, "WM_PROTOCOLS")
	if err != nil {
		return nil, err
	}
	return xprop.PropValAtomNames(xconn, reply)
}

// WM_PROTOCOLS set
func WmProtocolsSet(xconn *xprop.XPropConn, win xproto.Window, atomNames []string) error {
	data := make([]uint8, len(atomNames)*4)

	for i, name := range atomNames {
		atom, err := xconn.Atom(name, false)
		if err != nil {
			return err
		}
		into := data[i*4:]
		binary.LittleEndian.PutUint32(into, uint32(atom))
	}

	return xconn.ChangePropName(win, 32, "WM_PROTOCOLS", "ATOM", data)
}

// WM_COLORMAP_WINDOWS get
func WmColormapWindowsGet(xconn *xprop.XPropConn, win xproto.Window) ([]xproto.Window, error) {
	reply, err := xconn.GetPropName(win, "WM_COLORMAP_WINDOWS")
	if err != nil {
		return nil, err
	}
	return xprop.PropValWindows(reply)
}

// WM_COLORMAP_WINDOWS set
func WmColormapWindowsSet(xconn *xprop.XPropConn, win xproto.Window, windows []xproto.Window) error {
	values := *(*[]uint32)(unsafe.Pointer(&windows))
	return xconn.ChangePropName(win, 32, "WM_COLORMAP_WINDOWS", "WINDOW", xprop.FromData32(values))
}

// WM_CLIENT_MACHINE get
func WmClientMachineGet(xconn *xprop.XPropConn, win xproto.Window) (string, error) {
	reply, err := xconn.GetPropName(win, "WM_CLIENT_MACHINE")
	if err != nil {
		return "", err
	}
	return xprop.PropValStr(reply)
}

// WM_CLIENT_MACHINE set
func WmClientMachineSet(xconn *xprop.XPropConn, win xproto.Window, client string) error {
	return xconn.ChangePropName(win, 8, "WM_CLIENT_MACHINE", "STRING", byteutil.S2B(client))
}

// WmState is a struct that organizes information related to the WM_STATE
// property. Namely, the state (corresponding to a State* constant in this file)
// and the icon window (probably not used).
type WmState struct {
	State uint32
	Icon  xproto.Window
}

// WM_STATE get
func WmStateGet(xconn *xprop.XPropConn, win xproto.Window) (*WmState, error) {
	reply, err := xconn.GetPropName(win, "WM_STATE")
	if err != nil {
		return nil, err
	}
	raw, err := xprop.PropValUint32s(reply)
	if err != nil {
		return nil, err
	}
	if len(raw) != 2 {
		return nil,
			fmt.Errorf("WmState: Expected two integers in WM_STATE property "+
				"but xgbutil found %d in '%v'.", len(raw), raw)
	}
	return &WmState{
		State: raw[0],
		Icon:  xproto.Window(raw[1]),
	}, nil
}

// WM_STATE set
func WmStateSet(xconn *xprop.XPropConn, win xproto.Window, state *WmState) error {
	return xconn.ChangePropName(win, 32, "WM_STATE", "WM_STATE", xprop.FromData32([]uint32{
		state.State,
		uint32(state.Icon),
	}))
}

// IconSize is a struct the organizes information related to the WM_ICON_SIZE property. Mostly info about its dimensions.
type IconSize struct {
	MinWidth, MinHeight, MaxWidth, MaxHeight, WidthInc, HeightInc uint32
}

// WM_ICON_SIZE get
func WmIconSizeGet(xconn *xprop.XPropConn, win xproto.Window) (*IconSize, error) {
	reply, err := xconn.GetPropName(win, "WM_ICON_SIZE")
	if err != nil {
		return nil, err
	}
	raw, err := xprop.PropValUint32s(reply)
	if err != nil {
		return nil, err
	}
	if len(raw) != 6 {
		return nil,
			fmt.Errorf("WmIconSize: Expected six integers in WM_ICON_SIZE "+
				"property, but xgbutil found "+"%d in '%v'.", len(raw), raw)
	}
	return &IconSize{
		MinWidth: raw[0], MinHeight: raw[1],
		MaxWidth: raw[2], MaxHeight: raw[3],
		WidthInc: raw[4], HeightInc: raw[5],
	}, nil
}

// WM_ICON_SIZE set
func WmIconSizeSet(xconn *xprop.XPropConn, win xproto.Window, icondim *IconSize) error {
	return xconn.ChangePropName(win, 32, "WM_ICON_SIZE", "WM_ICON_SIZE", xprop.FromData32([]uint32{
		icondim.MinWidth, icondim.MinHeight,
		icondim.MaxWidth, icondim.MaxHeight,
		icondim.WidthInc, icondim.HeightInc,
	}))
}
