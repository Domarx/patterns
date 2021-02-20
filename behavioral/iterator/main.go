package main

import (
	"fmt"

	"github.com/domarx/patterns/behavioral/iterator/collection"
)

type filter func(item collection.Item)

func Process(iterator collection.Iterator, filter filter) {
	for iterator.HasMore() {
		filter(iterator.Next())
	}
}

func main() {
	iterable := collection.NewItemsCollection(1, 2, 3, 4, 5, 6)

	Process(iterable.ForwardIterator(), func(item collection.Item) {
		if item%2 == 0 {
			fmt.Println("[forward] even item-> ", item)
		}
	})
	Process(iterable.ReverseIterator(), func(item collection.Item) {
		fmt.Println("[reverse] item", item)
	})
	Process(iterable.ReverseIterator(), func(item collection.Item) {
		if item%2 != 0 {
			fmt.Println("[reverse] uneven item", item)
		}
	})
}
