package strings

type SuffixMatrix struct {
	tokens []int32
	mlcc int32
}

func (sm *SuffixMatrix) Init(tokens []int32, mlcc int32){
	sm.tokens = tokens
	sm.mlcc = mlcc
}



func (sm *SuffixMatrix) scan(curX int, curY int) []int32{
	var x int32 = 0
	var y int32 = 0
	var size int32 = 0
	output := make([]int32, 0)
	for {
		if curX >= len(sm.tokens) || curY >= len(sm.tokens){
			if size > sm.mlcc {
				output = append(output, x, y, size)
			}
			break
		}
		if sm.tokens[curX] == sm.tokens[curY]{
			if size == 0 {
				x = int32(curX)
				y = int32(curY)
			}
			size++
		}else{
			if size > sm.mlcc {
				output = append(output, x, y, size)
			}
			x = 0
			y = 0
			size = 0
		}
		curX++
		curY++
	}
	return output
}

func (sm *SuffixMatrix) Detect() []int32{
	//将矩阵沿对角元素切分，扫描上三角矩阵
	result := make([]int32, 0)
	for i:=1; i<len(sm.tokens); i++ {
		result = append(result, sm.scan(i, 0)...)
	}
	return result
}



