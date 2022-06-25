// Example randr uses the randr protocol to get information about the active
// heads. It also listens for events that are sent when the head configuration
// changes. Since it listens to events, you'll have to manually kill this
// process when you're done (i.e., ctrl+c.)
//
// While this program is running, if you use 'xrandr' to reconfigure your
// heads, you should see event information dumped to standard out.
//
// For more information, please see the RandR protocol spec:
// http://www.x.org/releases/X11R7.6/doc/randrproto/randrproto.txt
package main

import (
	"fmt"

	"codeberg.org/gruf/go-xgb"
	"codeberg.org/gruf/go-xgb/randr"
	"codeberg.org/gruf/go-xgb/xproto"
)

func main() {
	// Open X server connection
	xconn, b, _ := xgb.DefaultDialer.Dial("")

	// Perform initial X proto setup
	setup, err := xproto.Setup(xconn, b)
	if err != nil {
		panic(err)
	}

	// Register randr extension with XConn
	if err := randr.Register(xconn); err != nil {
		panic(err)
	}

	// Get default (root) window
	root := setup.Roots[0].Root

	// Gets the current screen resources. Screen resources contains a list
	// of names, crtcs, outputs and modes, among other things.
	resources, err := randr.GetScreenResources(xconn, root)
	if err != nil {
		panic(err)
	}

	// Iterate through all of the outputs and show some of their info.
	for _, output := range resources.Outputs {
		info, err := randr.GetOutputInfo(xconn, output, 0)
		if err != nil {
			panic(err)
		}

		if info.Connection == randr.ConnectionConnected {
			bestMode := info.Modes[0]
			for _, mode := range resources.Modes {
				if mode.Id == uint32(bestMode) {
					fmt.Printf("Width: %d, Height: %d\n",
						mode.Width, mode.Height)
				}
			}
		}
	}

	fmt.Println()

	// Iterate through all of the crtcs and show some of their info.
	for _, crtc := range resources.Crtcs {
		info, err := randr.GetCrtcInfo(xconn, crtc, 0)
		if err != nil {
			panic(err)
		}

		fmt.Printf("X: %d, Y: %d, Width: %d, Height: %d\n",
			info.X, info.Y, info.Width, info.Height)
	}

	// Tell RandR to send us events. (I think these are all of them, as of 1.3.)
	err = randr.SelectInputUnchecked(xconn, root,
		randr.NotifyMaskScreenChange|
			randr.NotifyMaskCrtcChange|
			randr.NotifyMaskOutputChange|
			randr.NotifyMaskOutputProperty)
	if err != nil {
		panic(err)
	}

	fmt.Println()

	// Listen to events and just dump them to standard out.
	// A more involved approach will have to read the 'U' field of
	// RandrNotifyEvent, which is a union (really a struct) of type
	// RanrNotifyDataUnion.
	for {
		ev, err := xconn.Recv()
		if err != nil {
			panic(err)
		}
		fmt.Println(ev)
	}
}
