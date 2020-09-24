package simhash


var (
	generators map[uint64]*FPGenerator
	zero uint64
	one uint64
	polynomials [][]uint64
	std64 *FPGenerator
)

func init(){
	generators = make(map[uint64]*FPGenerator)
	zero = 0
	one = 0x8000000000000000
	polynomials = [][]uint64{ nil,
		{ 0xC000000000000000, 0xC000000000000000},
		{ 0xE000000000000000, 0xE000000000000000 },
		{ 0xD000000000000000, 0xB000000000000000 },
		{ 0xF800000000000000, 0xF800000000000000 },
		{ 0xEC00000000000000, 0xBC00000000000000 },
		{ 0xDA00000000000000, 0xB600000000000000 },
		{ 0xE500000000000000, 0xE500000000000000 },
		{ 0x9680000000000000, 0xD480000000000000 },
		{ 0x80C0000000000000, 0x8840000000000000 },
		{ 0xB0A0000000000000, 0xE9A0000000000000 },
		{ 0xD9F0000000000000, 0xC9B0000000000000 },
		{ 0xE758000000000000, 0xDE98000000000000 },
		{ 0xE42C000000000000, 0x94E4000000000000 },
		{ 0xD4CE000000000000, 0xB892000000000000 },
		{ 0xE2AB000000000000, 0x9E39000000000000 },
		{ 0xCCE4800000000000, 0x9783800000000000 },
		{ 0xD8F8C00000000000, 0xA9CDC00000000000 },
		{ 0x9A28200000000000, 0xFD79E00000000000 },
		{ 0xC782500000000000, 0x96CD300000000000 },
		{ 0xBEE6880000000000, 0xE902C80000000000 },
		{ 0x86D7E40000000000, 0xF066340000000000 },
		{ 0x9888060000000000, 0x910ABE0000000000 },
		{ 0xFF30E30000000000, 0xB27EFB0000000000 },
		{ 0x8E375B8000000000, 0xA03D948000000000 },
		{ 0xD1415C4000000000, 0xF5357CC000000000 },
		{ 0x91A916E000000000, 0xB6CE66E000000000 },
		{ 0xE6D2FC5000000000, 0xD55882B000000000 },
		{ 0x9A3BA0B800000000, 0xFBD654E800000000 },
		{ 0xAEA5D2E400000000, 0xF0E533AC00000000 },
		{ 0xDA88B7BE00000000, 0xAA3AAEDE00000000 },
		{ 0xBA75BB4300000000, 0xF5A811C500000000 },
		{ 0x9B6C9A2F80000000, 0x9603CCED80000000 },
		{ 0xFABB538840000000, 0xE2747090C0000000 },
		{ 0x8358898EA0000000, 0x8C698D3D20000000 },
		{ 0xDA8ABD5BF0000000, 0xF6DF3A0AF0000000 },
		{ 0xB090C3F758000000, 0xD3B4D3D298000000 },
		{ 0xAD9882F5BC000000, 0x88DA4FB544000000 },
		{ 0xC3C366272A000000, 0xDCCF2A2262000000 },
		{ 0x9BC0224A97000000, 0xAF5D96F273000000 },
		{ 0x8643FFF621800000, 0x8E390C6EDC800000 },
		{ 0xE45C01919BC00000, 0xCBB34C8945C00000 },
		{ 0x80D8141BC2E00000, 0x886AFC3912200000 },
		{ 0xF605807C26500000, 0xA3B92D28F6300000 },
		{ 0xCE9A2CFC76280000, 0x98400C1921280000 },
		{ 0xF61894904C040000, 0xC8BE6DBCEC8C0000 },
		{ 0xE3A44C104D160000, 0xCA84A59443760000 },
		{ 0xC7E84953A11B0000, 0xD9D4F6AA02CB0000 },
		{ 0xC26CDD1C9A358000, 0x8BE8478434328000 },
		{ 0xAE125DBEB198C000, 0xFCC5DBFD5AAAC000 },
		{ 0x86DE52A079A6A000, 0xC5F16BD883816000 },
		{ 0xDF82486AAFE37000, 0xA293EC735692D000 },
		{ 0xE91ABA275C272800, 0xD686192874E3C800 },
		{ 0x963D0960DAB3FC00, 0xBA9DE62072621400 },
		{ 0xA2188C4E8A46CE00, 0xD31F75BCB4977E00 },
		{ 0xC43A416020A6CB00, 0x99F57FECA12B3900 },
		{ 0xA4F72EF82A58AE80, 0xCECE4391B81DA380 },
		{ 0xB39F119264BC0940, 0x80A277D20DABB9C0 },
		{ 0xFD6616C0CBFA0B20, 0xED16E64117DC11A0 },
		{ 0xFFA8BC44327B5390, 0xEDFB13DB3B66C210 },
		{ 0xCAE8EB99E73AB548, 0xC86135B6EA2F0B98 },
		{ 0xBA49BADCDD19B16C, 0x8F1944AFB18564C4 },
		{ 0xECFC86D765EABBEE, 0x9190E1C46CC13702 },
		{ 0xE1F8D6B3195D6D97, 0xDF70267FF5E4C979 },
		{ 0xD74307D3FD3382DB, 0x9999B3FFDC769B48 }}
	std64 = Make(polynomials[64][0], 64)
}


type FPGenerator struct {
	empty uint64
	degree int
	polynomial uint64
	byteModTable [16][256]uint64
}

func Make(polynomial uint64, degree int) *FPGenerator{
	fpgen, ok := generators[polynomial]
	if !ok{
		ins := FPGenerator{}
		ins.init(polynomial, degree)
		fpgen = &ins
		generators[polynomial] = fpgen
	}
	return fpgen
}

func (ins *FPGenerator) reduce(fp uint64) uint64{
	N := 8 - ins.degree / 8
	var local uint64
	if N == 8 {
		local = 0
	}else{
		local = fp & uint64(0xffffffffffffffff << (8 * N))
	}
	temp := zero
	for i:=0; i<N; i++ {
		temp ^= ins.byteModTable[8 + i][fp & 0xff]
		fp = fp >> 8
	}
	return local ^ temp
}

func (ins *FPGenerator) extendByte(f uint64, v int) uint64{
	f ^= uint64(0xff & v)
	i := int(f)
	result := f >> 8
	result ^= ins.byteModTable[7][i & 0xff]
	return result
}

func (ins *FPGenerator) fp(buf []int8, start int, n int) uint64{
	return ins.extend(ins.empty, buf, start, n)
}

func (ins *FPGenerator) extend(f  uint64, buf []int8, start int, n int) uint64{
	for i:=0; i<n; i++ {
		f = ins.extendByte(f, int(buf[start + i]))
	}
	return ins.reduce(f)
}

func (ins *FPGenerator) init(polynomial uint64, degree int){
	ins.degree = degree
	ins.polynomial = polynomial
	var powerTable [128]uint64

	x_to_the_i := one
	x_to_the_degree_minus_one := one >> (degree - 1)
	for i:=0; i<128; i++ {
		powerTable[i] = x_to_the_i
		overflow := (x_to_the_i & x_to_the_degree_minus_one) != 0
		x_to_the_i = x_to_the_i >> 1
		if overflow {
			x_to_the_i ^= polynomial
		}
	}
	ins.empty = powerTable[64]

	for i:=0; i<16; i++ {
		for j:=0; j<256; j++ {
			v := zero
			for k:=0; k<8; k++{
				if (j & (1 << k)) != 0 {
					v ^= powerTable[127 - i * 8 - k]
				}
			}
			ins.byteModTable[i][j] = v
		}
	}
}






