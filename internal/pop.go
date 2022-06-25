package internal

// PopCount counts the number of bits set in a value list mask.
func PopCount(mask int) (n int) {
	mask0 := uint32(mask)
	for i := uint32(0); i < 32; i++ {
		if mask0&(1<<i) != 0 {
			n++
		}
	}
	return
}
