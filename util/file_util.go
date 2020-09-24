package util

import (
	"os"
)

//按int读取文件数据
func ReadInts(file string) []int32{
	f,_ := os.OpenFile(file, os.O_RDONLY, 0744)
	k,_ := f.Stat()
	defer f.Close()
	result := make([]int32, k.Size())
	data := make([]byte, k.Size())
	f.Read(data)
	for i,b := range data{
		result[i] = int32(int8(b))
	}
	return result
}
