package simhash

const (
	DefaultLoadFactor float64 = 0.75
	DefaultGrowthFactor int = 16
	Free int8 = 0
	Occupied int8 = -1
)

var (
	Primes []int
)

func init(){
	Primes = []int{ 3, 3, 3, 3, 3, 3, 3,
		3, 3, 3, 5, 5, 5, 5, 5, 5, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 13, 13, 13,
		13, 13, 13, 13, 13, 19, 19, 19, 19, 19, 19, 19, 19, 19, 19, 19, 19, 31,
		31, 31, 31, 31, 31, 31, 43, 43, 43, 43, 43, 43, 43, 43, 61, 61, 61, 61,
		61, 73, 73, 73, 73, 73, 73, 73, 103, 103, 109, 109, 109, 109, 109, 139,
		139, 151, 151, 151, 151, 181, 181, 193, 199, 199, 199, 229, 241, 241,
		241, 271, 283, 283, 313, 313, 313, 349, 349, 349, 349, 421, 433, 463,
		463, 463, 523, 523, 571, 601, 619, 661, 661, 661, 661, 661, 823, 859,
		883, 883, 883, 1021, 1063, 1093, 1153, 1153, 1231, 1321, 1321, 1429,
		1489, 1489, 1621, 1699, 1789, 1873, 1951, 2029, 2131, 2143, 2311, 2383,
		2383, 2593, 2731, 2803, 3001, 3121, 3259, 3391, 3583, 3673, 3919, 4093,
		4273, 4423, 4651, 4801, 5023, 5281, 5521, 5743, 5881, 6301, 6571, 6871,
		7129, 7489, 7759, 8089, 8539, 8863, 9283, 9721, 10141, 10531, 11071,
		11551, 12073, 12613, 13009, 13759, 14323, 14869, 15649, 16363, 17029,
		17839, 18541, 19471, 20233, 21193, 22159, 23059, 24181, 25171, 26263,
		27541, 28753, 30013, 31321, 32719, 34213, 35731, 37309, 38923, 40639,
		42463, 44281, 46309, 48313, 50461, 52711, 55051, 57529, 60091, 62299,
		65521, 68281, 71413, 74611, 77713, 81373, 84979, 88663, 92671, 96739,
		100801, 105529, 109849, 115021, 120079, 125509, 131011, 136861, 142873,
		149251, 155863, 162751, 169891, 177433, 185071, 193381, 202129, 211063,
		220021, 229981, 240349, 250969, 262111, 273643, 285841, 298411, 311713,
		325543, 339841, 355009, 370663, 386989, 404269, 422113, 440809, 460081,
		480463, 501829, 524221, 547399, 571603, 596929, 623353, 651019, 679909,
		709741, 741343, 774133, 808441, 844201, 881539, 920743, 961531, 1004119,
		1048573, 1094923, 1143283, 1193911, 1246963, 1302181, 1359733, 1420039,
		1482853, 1548541, 1616899, 1688413, 1763431, 1841293, 1922773, 2008081,
		2097133, 2189989, 2286883, 2388163, 2493853, 2604013, 2719669, 2840041,
		2965603, 3097123, 3234241, 3377191, 3526933, 3682363, 3845983, 4016041,
		4193803, 4379719, 4573873, 4776223, 4987891, 5208523, 5439223, 5680153,
		5931313, 6194191, 6468463, 6754879, 7053331, 7366069, 7692343, 8032639,
		8388451, 8759953, 9147661, 9552733, 9975193, 10417291, 10878619,
		11360203, 11863153, 12387841, 12936529, 13509343, 14107801, 14732413,
		15384673, 16065559, 16777141, 17519893, 18295633, 19105483, 19951231,
		20834689, 21757291, 22720591, 23726449, 24776953, 25873963, 27018853,
		28215619, 29464579, 30769093, 32131711, 33554011, 35039911, 36591211,
		38211163, 39903121, 41669479, 43514521, 45441199, 47452879, 49553941,
		51747991, 54039079, 56431513, 58930021, 61539091, 64263571, 67108669,
		70079959, 73182409, 76422793, 79806229, 83339383, 87029053, 90881083,
		94906249, 99108043, 103495879, 108077731, 112863013, 117860053,
		123078019, 128526943, 134217439, 140159911, 146365159, 152845393,
		159612601, 166679173, 174058849, 181765093, 189812341, 198216103,
		206991601, 216156043, 225726379, 235720159, 246156271, 257054491,
		268435009, 280319203, 292730833, 305691181, 319225021, 333358513,
		348117151, 363529759, 379624279, 396432481, 413983771, 432312511,
		451452613, 471440161, 492312523, 514109251, 536870839, 560640001,
		585461743, 611382451, 638450569, 666717199, 696235363, 727060069,
		759249643, 792864871, 827967631, 864625033, 902905501, 942880663,
		984625531, 1028218189, 1073741719, 1121280091, 1170923713, 1222764841,
		1276901371, 1333434301, 1392470281, 1454120779, 1518500173, 1585729993,
		1655935399, 1729249999, 1805811253, 1885761133, 1969251079, 2056437379,
		2147482951}
}

type LongOpenHashSet struct {
	key []uint64
	state []int8
	f float64
	p int
	maxFill int
	free int
	count int
	growthFactor int
	pos int
	last int
	c int
}

func (ins *LongOpenHashSet) init(n int, f float64){
	ins.growthFactor = DefaultGrowthFactor
	if f <= 0 || f > 1 {
		panic("Load factor must be greater than 0 and smaller than or equal to 1")
	}
	if n < 0 {
		panic("Hash table size must be nonnegative")
	}
	l := BinarySearch(Primes, 0, len(Primes),int(float64(n) / f) + 1)
	if l < 0 {
		l = -l - 1
	}
	ins.p = l
	ins.free = Primes[l]
	ins.f = f
	ins.maxFill = int(float64(ins.free) * f)
	ins.key = make([]uint64, ins.free)
	ins.state = make([]int8, ins.free)
}

func (ins *LongOpenHashSet) init2(n int){
	ins.init(n, DefaultLoadFactor)
}

func (ins *LongOpenHashSet) findInsertionPoint(k uint64) int{
	key := ins.key
	state := ins.state
	n := len(key)
	k2i := longHash2IntHash(k) & 0x7FFFFFFF
	h1 := k2i % n
	if state[h1] == Occupied && !(k == key[h1]) {
		h2 := (k2i % (n - 2)) + 1
		for {
			h1 = (h1 + h2) % n
			if !(state[h1] == Occupied && !((k) == (key[h1]))){
				break
			}
		}
	}
	if state[h1] == Free {
		return h1
	}
	if state[h1] == Occupied {
		return -h1 - 1
	}
	i := h1
	if !(k == key[h1]) {
		// The secondary hash.
		h2 := (k2i % (n - 2)) + 1
		for {
			h1 = (h1 + h2) % n
			if !(state[h1] != Free && !((k) == (key[h1]))){
				break
			}
		}
	}
	if state[h1] == Occupied {
		return -h1 - 1
	}
	return i
}

func (ins *LongOpenHashSet) findKey(k uint64) int{
	key := ins.key
	state := ins.state
	n := len(key)
	k2i := longHash2IntHash(k) & 0x7FFFFFFF
	h1 := k2i % n
	if state[h1] != Free && !(k == key[h1]){
		h2 := (k2i % (n - 2)) + 1
		for {
			h1 = (h1 + h2) % n
			if !(state[h1] != Free && !(k == key[h1])){
				break
			}
		}
	}
	if state[h1] == Occupied {
		return h1
	}
	return -1
}

func (ins *LongOpenHashSet) add(k uint64) bool{
	i := ins.findInsertionPoint(k)
	if i < 0 {
		return false
	}
	if ins.state[i] == Free{
		ins.free--
	}
	ins.state[i] = Occupied
	ins.key[i] = k
	ins.count++
	if ins.count >= ins.maxFill {
		newP := Min(ins.p + ins.growthFactor, len(Primes) - 1)
		for Primes[newP] == Primes[ins.p]{
			newP++
		}
		ins.rehash(newP)
	}
	if ins.free == 0 {
		ins.rehash(ins.p)
	}
	return true
}

func (ins *LongOpenHashSet) setInit(){
	state := ins.state
	n := len(state)
	ins.c = ins.count
	if ins.c != 0 {
		for ins.pos < n && state[ins.pos] != Occupied{
			ins.pos++
		}
	}
}

func (ins *LongOpenHashSet) next() uint64{
	return ins.nextLong()
}

func (ins *LongOpenHashSet) hasNext() bool{
	return ins.c != 0 && ins.pos < len(ins.state)
}

func (ins *LongOpenHashSet) nextLong() uint64{
	var retVal uint64
	state := ins.state
	n := len(state)
	if !ins.hasNext() {
		return 0
	}
	ins.last = ins.pos
	retVal = ins.key[ins.last]
	ins.c--
	if ins.c != 0 {
		for {
			ins.pos++
			if !(ins.pos < n && state[ins.pos] != Occupied){
				break
			}
		}
	}
	return retVal
}

func (ins *LongOpenHashSet) rehash(newP int){
	i := 0
	j := ins.count
	var k2i, h1, h2 int
	state := ins.state
	var k uint64
	newN := Primes[newP]
	key := ins.key
	newKey := make([]uint64, newN)
	newState := make([]int8, newN)
	for {
		if j == 0 {
			break
		}
		j--
		for state[i] != Occupied{
			i++
		}
		k = key[i]
		k2i = longHash2IntHash(k) & 0x7FFFFFFF
		h1 = k2i % newN
		if newState[h1] != Free{
			h2 = (k2i % (newN - 2)) + 1
			for {
				h1 = (h1 + h2) % newN
				if newState[h1] == Free{
					break
				}
			}
		}
		newState[h1] = Occupied
		newKey[h1] = k
		i++
	}
	ins.p = newP;
	ins.free = newN - ins.count
	ins.maxFill = int(float64(newN) * ins.f)
	ins.key = newKey
	ins.state = newState
}

func longHash2IntHash(l uint64) int{
	return int(l^(l >> 32))
}









