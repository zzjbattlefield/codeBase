<font size="4">

# 栈
先进后出 只能从栈顶操作
## 栈的实现
```go
    
type Node struct {
	Data interface{}
	next *Node
}

type Stack struct {
	Top  *Node
	Size uint
}

func (s *Stack) NewNode(v interface{}) *Node {
	return &Node{Data: v}
}

func NewStack() *Stack {
	return &Stack{}
}

func (s *Stack) Push(v interface{}) error {
	newNode := s.NewNode(v)
	if s.Size == 0 {
		s.Top = newNode
	} else {
		newNode.next = s.Top
		s.Top = newNode
	}
	s.Size++
	return nil
}

func (s *Stack) Peek() interface{} {
	if s.Size == 0 {
		return nil
	}
	return s.Top.Data
}

func (s *Stack) Pop() interface{} {
	if s.Size == 0 {
		return nil
	}
	node := s.Top
	s.Top = s.Top.next
	s.Size--
	return node.Data
}

```
给定一个包含了右侧三种字符的字符串， ‘(‘, ‘)’, ‘{‘, ‘}’, ‘[‘ and ‘]’，判断字符串是否合法。合法的判断条件如下：
必须使用相同类型的括号关闭左括号。
必须以正确的顺序关闭打开括号。
```go
func isValid(s string) bool {
	stack := NewStack()
	bracketMap := map[uint8]uint8{'{': '}', '[': ']', '(': ')'}
	for i := 0; i < len(s); i++ {
		if stack.Size > 0 {
			if val, ok := bracketMap[stack.Peek().(uint8)]; ok && val == s[i] {
				stack.Pop()
				continue
			}
		}
		stack.Push(s[i])
	}
	return stack.Size == 0
}
```

## 用slice实现

```go

type Item interface{}

type ItemStack struct {
	items []Item
}

func NewStack() *ItemStack {
	return &ItemStack{[]Item{}}
}

func (s *ItemStack) Pop() Item {
	if len(s.items) == 0 {
		return nil
	}
	item := s.items[len(s.items)-1]
	s.items = s.items[0 : len(s.items)-1]
	return item
}

func (s *ItemStack) Push(v interface{}) {
	s.items = append(s.items, v)
}
```