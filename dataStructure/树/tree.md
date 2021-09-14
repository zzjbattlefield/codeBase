<font size="4">

# 二叉树

```go
type Tree struct {
	right *Tree
	value int
	left  *Tree
}

func CreateNode(v int) *Tree {
	return &Tree{value: v}
}

func (tree *Tree) Print() {
	fmt.Println(tree.value)
}

func (tree *Tree) SetValue(v int) {
	tree.value = v
}
//前序遍历
func (tree *Tree) Traverse() {
	if tree == nil {
		return
	}
	tree.Print()
	tree.left.Traverse()
	tree.right.Traverse()
}
```