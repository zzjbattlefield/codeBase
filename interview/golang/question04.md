<font size="4">

# 常见排序算法

## 冒泡排序

```go
func swap(i, k int) (int, int) {
	return k, i
}

func bubbleSort(list []int) []int {
	for i := 0; i <= len(list)-1; i++ {
		for k := i + 1; k <= len(list)-1; k++ {
			if list[k] > list[i] {
				list[k], list[i] = swap(list[k], list[i])
			}
		}
	}
	return list
}

func TestBubbleSort(t *testing.T) {
	testTable := []struct {
		list []int
		want []int
	}{
		{
			list: []int{3, 6, 2, 5, 8, 9, 1},
			want: []int{1, 2, 3, 5, 6, 8, 9},
		}, {
			list: []int{},
			want: []int{},
		},
	}
	for _, tt := range testTable {
		res := bubbleSort(tt.list)
		if reflect.DeepEqual(res, tt.want) {
			t.Fatalf("want:%v got:%v", tt.want, res)
		}
	}
}
```

## 快速排序
从数列中挑出一个元素,称为 “基准”(pivot)
重新排序数列,所有元素比基准值小的摆放在基准前面,所有元素比基准值大的摆在基准的后面(相同的数可以到任一边).在这个分区退出之后,该基准就处于数列的中间位置.这个称为分区(partition)操作;
递归的把小于基准值元素的子数列和大于基准值元素的子数列排序;
递归的最底部情形,是数列的大小是零或一,也就是永远都已经被排序好了. 虽然一直递归下去,但是这个算法总会退出,因为在每次的迭代中,它至少会把一个元素摆到它最后的位置去.
```go
func partition(list []int, low, high int) int {
	target := list[low]
	for low < high {
		for low < high && list[high] >= target {
			high--
		}
		list[low] = list[high]
		for low < high && list[low] <= target {
			low++
		}
		list[high] = list[low]
	}
	list[low] = target
	return low
}

func fastSearch(list []int, low, high int) {
	if low < high {
		middle := partition(list, low, high)
		fastSearch(list, low, middle-1)
		fastSearch(list, middle+1, low)
	}
}

func TestFastSearch(t *testing.T) {
	list := []int{8, 5, 8, 3, 54, 6, 3, 4, 5, 6}
	fastSearch(list, 0, len(list)-1)
	log.Println(list)
}
```

## 选择排序

```go
func choseSort(list []int) {
	if len(list) == 0 {
		return
	}
	for i := 0; i < len(list); i++ {
		target := i
		for j := i + 1; j < len(list); j++ {
			if list[target] > list[j] {
				target = j
			}
		}
		list[i], list[target] = swap(list[i], list[target])
	}
}

func swap(i, k int) (int, int) {
	return k, i
}
```