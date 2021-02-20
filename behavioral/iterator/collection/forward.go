package collection

func newForwardIterator(items Items) Iterator {
	return &forwardIterator{
		index: 0,
		items: items,
	}
}

type forwardIterator struct {
	index int
	items Items
}

func (f *forwardIterator) HasMore() bool {
	return f.index < len(f.items)
}

func (f *forwardIterator) Next() Item {
	var item Item
	if f.HasMore() {
		item = f.items[f.index]
		f.index++
	}
	return item
}
