package linkedList

type dNode struct {
	data int
	next *dNode
	prev *dNode
}

type DoubleLL struct {
	head   *dNode
	tail   *dNode
	length int
}

func (dl *DoubleLL) Length() int {
	return dl.length
}

func (dl *DoubleLL) InsertDl(data int) {
	newNode := &dNode{data: data}
	if dl.head == nil {
		dl.head = newNode
	} else {
		dl.tail.next = newNode
		newNode.prev = dl.tail
	}
	dl.tail = newNode
	dl.length++
}
