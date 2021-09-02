<font size="4">

# 链表

## 单链表
```go
    
type listNode struct {
	value interface{}
	next  *listNode
}

type LinkedList struct {
	head   *listNode
	length uint
}

func NewListNode(val interface{}) *listNode {
	return &listNode{value: val, next: nil}
}

func NewLinkedList() *LinkedList {
	return &LinkedList{length: 0, head: NewListNode(0)}
}

func (node *listNode) GetNext() *listNode {
	return node.next
}

func (node *listNode) GetVal() interface{} {
	return node.value
}

//在节点后插入节点
func (link *LinkedList) InsertAfter(p *listNode, v interface{}) error {
	if p == nil {
		return errors.New("传入节点为空")
	}
	NewNode := NewListNode(v)
	//p当前的next 后续赋值给NewNode
	oldNext := p.next
	p.next = NewNode
	NewNode.next = oldNext
	link.length++
	return nil
}

//在头部添加
func (link *LinkedList) InsertToHead(v interface{}) error {
	return link.InsertAfter(link.head, v)
}

//在尾部添加
func (link *LinkedList) InsertToTail(v interface{}) error {
	cur := link.head.next
	for cur.next != nil {
		cur = cur.next
	}
	return link.InsertAfter(cur, v)
}

//通过index查找
func (link *LinkedList) FindByIndex(index int) *listNode {
	if link.length < uint(index) {
		return nil
	}
	count := 0
	cur := link.head.next
	for cur.next != nil && count < index {
		cur = cur.next
		count++
	}
	return cur
}

func (link *LinkedList) DeleteNode(node *listNode) error {
	if node == nil {
		return errors.New("节点不能为空")
	}
	cur := link.head.next
	pre := link.head
	for cur.next != nil {
		pre = cur
		cur = cur.next
		if cur.next == node.next {
			break
		}
	}
	if cur == nil {
		return errors.New("查无此节点")
	}
	pre.next = cur.next
	link.length--
	return nil

}

func (link *LinkedList) Print() {
	next := link.head.next
	for next != nil {
		fmt.Println(next.GetVal())
		next = next.next
	}
	fmt.Println("done")
}
//链表反转
func (link *LinkedList) Reverse() {
	var pre = &listNode{}
	cur := link.head.next
	for nil != cur {
		tmp := cur.next
		cur.next = pre
		pre = cur
		cur = tmp
	}
	link.head.next = pre
}
```
## 双链表

```go

type DobuleNode struct {
	Data interface{}
	prev *DobuleNode
	next *DobuleNode
}

type DobuleLink struct {
	Head *DobuleNode
	Size uint
	Tail *DobuleNode
}

func (link *DobuleLink) newDobuleNode(v interface{}) *DobuleNode {
	return &DobuleNode{Data: v, prev: nil, next: nil}
}

func NewDobuleLink() *DobuleLink {
	return &DobuleLink{Head: nil, Tail: nil}
}

func (link *DobuleLink) Append(v interface{}) error {
	newNode := link.newDobuleNode(v)
	if link.Size == 0 {
		link.Tail = newNode
		link.Head = newNode
	} else {
		newNode.prev = link.Tail
		link.Tail.next = newNode
		link.Tail = newNode
	}
	link.Size++
	return nil
}

func (link *DobuleLink) Insert(v interface{}) error {
	newNode := link.newDobuleNode(v)
	if link.Size == 0 {
		link.Head = newNode
		link.Tail = newNode
	} else {
		newNode.next = link.Head
		link.Head.prev = newNode
		link.Head = newNode
	}
	link.Size++
	return nil
}

func (link *DobuleLink) Print() {
	cur := link.Head
	for cur != nil {
		fmt.Println(cur.Data)
		cur = cur.next
	}
}

```