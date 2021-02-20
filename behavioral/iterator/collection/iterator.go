package collection

type Iterable interface {
	ForwardIterator() Iterator
	ReverseIterator() Iterator
}

type Iterator interface {
	HasMore() bool
	Next() Item
}

func NewItemsCollection(items ...Item) Iterable {
	return (Items)(items)
}

type Item int

type Items []Item

func (i Items) ForwardIterator() Iterator { return newForwardIterator(i) }

func (i Items) ReverseIterator() Iterator { return newReverseIterator(i) }
