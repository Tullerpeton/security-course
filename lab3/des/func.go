package des

func funcF(block32b uint32, key48b uint64) uint32 {
	block48b := expansionPermutation(block32b)
	block48b ^= key48b
	block32b = substitutions(block48b)
	return permutation(block32b)
}

func expansionPermutation(block32b uint32) uint64 {
	var block48b uint64
	for i := 0; i < 48; i++ {
		// 48 = 00000000 00000000 00000000 00000000 00000000 00000000 (правые 16 бит - 0)
		// 32 =                   00000000 00000000 00000000 00000000
		// (63 - i) - стартовая позиция
		// (32 - EP[i]) - позиция нужного бита через дезъюнкцию получаем нужный
		block48b |= ((uint64(block32b) >> (32 - EP[i])) & 1) << (63 - i)
	}
	return block48b
}

func substitutions(block48b uint64) uint32 {
	blocks6b := split48bTo6b(block48b)
	blocks4b := substitution6bTo4b(blocks6b)
	return join4bTo32b(blocks4b)
}

func split48bTo6b(block48b uint64) (blocks6b [8]byte) {
	for i := 0; i < 8; i++ {
		blocks6b[i] = byte((block48b >> (58 - (i * 6))) << 2)
	}
	return
}

func substitution6bTo4b(blocks6b [8]byte) (blocks4b [4]byte) {
	var block2b, block4b byte
	for i, j := 0, 0; i < 8; i, j = i+2, j+1 {
		block2b = extremeBits(blocks6b[i])
		block4b = middleBits(blocks6b[i])
		blocks4b[j] = Sbox[i][block2b][block4b]

		block2b = extremeBits(blocks6b[i+1])
		block4b = middleBits(blocks6b[i+1])
		blocks4b[j] = (blocks4b[j] << 4) | Sbox[i+1][block2b][block4b]
	}
	return
}

func extremeBits(block6b byte) byte {
	// смотрим 8 и 3 бит (первый и последний)
	return ((block6b >> 6) & 2) | ((block6b >> 2) & 1)
}

func middleBits(block6b byte) byte {
	// 0000 1111 (0xF)
	// берём середину после extremeBits
	return (block6b >> 3) & 0xF
}

func join4bTo32b(blocks4b [4]byte) uint32 {
	var block32b uint32
	for _, b := range blocks4b {
		block32b = (block32b << 8) | uint32(b)
	}
	return block32b
}

func permutation(block32b uint32) uint32 {
	var newBlock32b uint32
	for i := 0; i < 32; i++ {
		newBlock32b |= ((block32b >> (32 - P[i])) & 1) << (31 - i)
	}
	return newBlock32b
}
