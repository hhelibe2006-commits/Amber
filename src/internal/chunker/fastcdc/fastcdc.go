package fastcdc

func FastCDC(data []byte, maxSize uint64, minSize uint64, bitsS uint64) [][]byte {
	chunks := make([][]byte, 0, maxSize)
	var bitsL uint64
	if bitsS >= 2 {
		bitsL = bitsS - 2
	}
	maskS := uint64((1 << bitsS) - 1)
	maskL := uint64((1 << bitsL) - 1)
	var pos uint64
	for pos < uint64(len(data)) {
		mask := maskS
		gear := NewGearHash()
		if pos+minSize >= uint64(len(data)) {
			chunks = append(chunks, data[pos:])
			break
		}
		for i := pos; ; i++ {
			if i == uint64(len(data)) {
				chunks = append(chunks, data[pos:])
				pos = uint64(len(data))
				break
			}
			if i == pos+maxSize {
				chunks = append(chunks, data[pos:pos+maxSize])
				pos = i + maxSize
				break
			}
			if i-pos >= (maskS+1)/2 {
				mask = maskL
			}
			gear.Next(data[i])
			if i-pos >= minSize {
				if gear.hash&mask == 0 {
					chunks = append(chunks, data[pos:i])
					pos = i
					break
				}
			}
		}
	}
	return chunks
}
