package collection

func newReverseIterator(items Items) Iterator {
	return &reverseIterator{
		index: len(items) - 1,
		items: items,
	}
}

type reverseIterator struct {
	index int
	items Items
}

func (r *reverseIterator) HasMore() bool {
	return r.index > -1
}

func (r *reverseIterator) Next() Item {
	var item Item
	if r.HasMore() {
		item = r.items[r.index]
		r.index--
	}
	return item
}
