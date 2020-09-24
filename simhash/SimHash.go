package simhash

const (
	HASH_SIZE int = 64
	FIXED_BGRAM_LENGTH int = 8
	FIXED_CGRAM_LENGTH int = 4
)

//计算哈希值
func ComputeOptimizedSimHashForString(s string) uint64{
	shingles := LongOpenHashSet{}
	shingles.Init2(Min(len(s), 100000))
	length := len(s)
	for i := 0; i < length - FIXED_CGRAM_LENGTH + 1; i++{
		var shingle = uint64(s[i])
		shingle <<= 16
		shingle |= uint64(s[i+1])
		shingle <<= 16
		shingle |= uint64(s[i+2])
		shingle <<= 16
		shingle |= uint64(s[i+3])
		shingles.Add(shingle)
	}

	var v [HASH_SIZE]int
	longAsBytes := make([]int8, 8)

	shingles.setInit()
	for  {
		shingle := shingles.Next()
		if shingle == 0{
			break
		}

		longAsBytes[0] = int8(shingle >> 56)
		longAsBytes[1] = int8(shingle >> 48)
		longAsBytes[2] = int8(shingle >> 40)
		longAsBytes[3] = int8(shingle >> 32)
		longAsBytes[4] = int8(shingle >> 24)
		longAsBytes[5] = int8(shingle >> 16)
		longAsBytes[6] = int8(shingle >> 8)
		longAsBytes[7] = int8(shingle)

		longHash := Std64.Fp(longAsBytes, 0, 8)
		for i:=0; i<HASH_SIZE; i++ {
			bitSet := ((longHash >> i) & 1) == 1
			if bitSet {
				v[i] += 1
			}else{
				v[i] += -1
			}
		}
	}
	var sim uint64 = 0
	for i:=0; i<HASH_SIZE; i++ {
		if v[i] > 0 {
			sim |= 1 << i
		}
	}
	return sim
}

//计算海明距离
func HammingDistance(hash1 uint64, hash2 uint64) int{
	bits := hash1 ^ hash2
	count := 0
	for bits != 0 {
		bits &= bits - 1
		count++
	}
	return count
}

func Divide4(hash uint64) []uint64{
	hashes := make([]uint64, 4)
	hashes[0] = hash & 0xffff000000000000
	hashes[1] = hash & 0x0000ffff00000000
	hashes[2] = hash & 0x00000000ffff0000
	hashes[3] = hash & 0x000000000000ffff
	return hashes
}



