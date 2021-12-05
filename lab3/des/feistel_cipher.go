package des

func feistelCipherMode(mode byte) func(*uint32, *uint32, [RoundCount]uint64) {
	switch mode {
	case EncMode:
		return feistelCipherEnc
	case DecMode:
		return feistelCipherDec
	}
	return nil
}

func feistelCipherEnc(left, right *uint32, keys48b [RoundCount]uint64) {
	for round := 0; round < RoundCount; round++ {
		roundFeistelCipher(left, right, keys48b[round])
	}
	*left, *right = *right, *left
}

func feistelCipherDec(left, right *uint32, keys48b [RoundCount]uint64) {
	for round := 15; round >= 0; round-- {
		roundFeistelCipher(left, right, keys48b[round])
	}
	*left, *right = *right, *left
}

func roundFeistelCipher(left, right *uint32, key48b uint64) {
	tmp := *right
	*right = funcF(*right, key48b) ^ *left
	*left = tmp
}
