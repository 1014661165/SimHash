package strings

type SuffixMatrixParallel struct {
	tokens []int32
	mlcc int
	output chan []int32
	exit chan bool
}

func (smp *SuffixMatrixParallel) Init(tokens []int32, mlcc int){
	smp.tokens = tokens
	smp.mlcc = mlcc
	smp.output = make(chan []int32, 1)
	smp.output <- []int32{}
	smp.exit = make(chan bool)
}

func (smp *SuffixMatrixParallel) scan(curX int, curY int){
	var x int32 = 0
	var y int32 = 0
	var size int32 = 0
	output := make([]int32, 0)
	for {
		if curX >= len(smp.tokens) || curY >= len(smp.tokens){
			if size > int32(smp.mlcc) {
				output = append(output, x, y, size)
			}
			break
		}
		if smp.tokens[curX] == smp.tokens[curY]{
			if size == 0 {
				x = int32(curX)
				y = int32(curY)
			}
			size++
		}else{
			if size > int32(smp.mlcc){
				output = append(output, x, y, size)
			}
			x = 0
			y = 0
			size = 0
		}
		curX++
		curY++
	}
	res := <-smp.output
	res = append(res, output...)
	smp.output <- res
	smp.exit<-true
}

func (smp *SuffixMatrixParallel) Detect() []int32{
	//将矩阵沿对角元素切分，扫描上三角矩阵
	for i:=1; i<len(smp.tokens)-smp.mlcc; i++ {
		go smp.scan(i, 0)
	}

	for i:=1; i<len(smp.tokens)-smp.mlcc; i++{
		<-smp.exit
	}
	close(smp.output)
	close(smp.exit)
	return <-smp.output
}



