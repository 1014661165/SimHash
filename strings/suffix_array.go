package strings

import "sort"

type SuffixArray struct {
	tokens []int32
	mlcc int32
	sa []int32
	height []int32
}

func (sa *SuffixArray) Init(tokens []int32, mlcc int32){
	sa.tokens = tokens
	sa.mlcc = mlcc
	sa.sa = make([]int32, len(sa.tokens))
	sa.height = make([]int32, len(sa.tokens))
}

func (sa *SuffixArray) BuildSuffixArray(){
	for i:= 0; i<len(sa.tokens); i++{
		sa.sa[i] = int32(i)
	}
	sort.Slice(sa.sa, func(i, j int) bool {
		suffix1 := sa.tokens[sa.sa[i]:]
		suffix2 := sa.tokens[sa.sa[j]:]
		sizeI := len(suffix1)
		sizeJ := len(suffix2)
		size := 0
		result := true
		if sizeI <= sizeJ{
			size = len(suffix1)
		}else{
			size = len(suffix2)
			result = false
		}
		for m:=0; m<size; m++ {
			if suffix1[m] < suffix2[m]{
				result = true
				break
			}else if suffix1[m] > suffix2[m] {
				result = false
				break
			}
		}
		return result
	})
}

func (sa *SuffixArray) CalculateHeight(){
	for i:=1; i<len(sa.tokens); i++ {
		pre := sa.tokens[sa.sa[i - 1]:]
		cur := sa.tokens[sa.sa[i]:]
		cnt := 0
		size := 0
		if len(pre) < len(cur){
			size = len(pre)
		}else{
			size = len(cur)
		}
		for j:=0; j<size; j++ {
			if pre[j] != cur[j]{
				break
			}
			cnt++
		}
		sa.height[i] = int32(cnt)
	}
}

func (sa *SuffixArray) Detect() []int32{
	result := make([]int32, 0)
	for i:=1; i<len(sa.height); i++ {
		if sa.height[i] >= sa.mlcc {
			result = append(result, sa.sa[i - 1], sa.sa[i], sa.height[i])
		}
	}
	return result
}
