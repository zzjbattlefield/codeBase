<font size="4">

# 常见查找算法

## 二分查找
必须是有序的切片
```go
func binarySearch(target int, list []int) int {
	first := 0
	last := len(list) - 1
	for {
		middle := int(math.Ceil((float64(last) + float64(first)) / 2))
		if middle > last {
			return -1
		}
		if list[middle] == target {
			return middle
		}
		if list[middle] > target {
			last = middle - 1
		} else {
			first = middle + 1
		}
	}
}

func TestBinarySearch(t *testing.T) {
	want := 1
	TestTable := []struct {
		list []int
		in   int
		out  int
	}{
		{
			list: []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
			in:   1,
			out:  1,
		}, {
			list: []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
			in:   9,
			out:  9,
		}, {
			list: []int{},
			in:   10,
			out:  -1,
		},
		{
			list: []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
			in:   10,
			out:  -1,
		},
	}

	for _, tt := range TestTable {
		got := binarySearch(tt.list, tt.in)
		if got != tt.out {
			t.Fatalf("want:%d got :%d", want, got)
		}
	}
}
```