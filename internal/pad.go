package internal

// Pad4 aligns a length to 4 bytes.
func Pad4(n int) int {
	return (n + 3) &^ 3
}
