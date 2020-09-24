package simhash

//二分查找
func BinarySearch(a []int, fromIndex int, toIndex int, key int) int{
	low := fromIndex
	high := toIndex - 1

	for low <= high {
		mid := (low + high) >> 1
		midVal := a[mid]

		if midVal < key{
			low = mid + 1
		} else if midVal > key{
			high = mid - 1
		} else{
			return mid
		}
	}
	return -(low + 1)
}

func Min(a,b int) int{
	if a <= b {
		return a
	}
	return b
}
