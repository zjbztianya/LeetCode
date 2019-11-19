package reverse_bits

func reverseBits(num uint32) uint32 {
	var ans uint32
	for i := 0; i < 32; i++ {
		ans <<= 1
		if (num & (1 << uint32(i))) > 0 {
			ans++
		}
	}
	return ans
}
