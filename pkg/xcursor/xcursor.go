package xcursor

import (
	"codeberg.org/gruf/go-xgb"
	"codeberg.org/gruf/go-xgb/xproto"
)

// CreateCursor sets some default colors for nice and easy cursor creation.
// Just supply a cursor constant from 'xcursor/cursordef.go'.
func CreateCursor(xconn *xgb.XConn, cursor uint16) (xproto.Cursor, error) {
	return CreateCursorExtra(xconn, cursor, 0, 0, 0, 0xffff, 0xffff, 0xffff)
}

// CreateCursorExtra features all available parameters to creating a cursor.
// It will return an error if there is a problem with any of the requests
// made to create the cursor.
// (This implies each request is a checked request. The performance loss is
// probably acceptable since cursors should be created once and reused.)
func CreateCursorExtra(xconn *xgb.XConn, cursor, foreRed, foreGreen, foreBlue, backRed, backGreen, backBlue uint16) (xproto.Cursor, error) {
	fontID, err := xproto.NewFontID(xconn)
	if err != nil {
		return 0, err
	}

	cursorID, err := xproto.NewCursorID(xconn)
	if err != nil {
		return 0, err
	}

	err = xproto.OpenFont(xconn, fontID, uint16(len("cursor")), "cursor")
	if err != nil {
		return 0, err
	}

	err = xproto.CreateGlyphCursor(xconn, cursorID, fontID, fontID,
		cursor, cursor+1,
		foreRed, foreGreen, foreBlue,
		backRed, backGreen, backBlue)
	if err != nil {
		return 0, err
	}

	err = xproto.CloseFont(xconn, fontID)
	if err != nil {
		return 0, err
	}

	return cursorID, nil
}
