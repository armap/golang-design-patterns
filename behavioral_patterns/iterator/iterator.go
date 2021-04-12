package main

import "fmt"

/*
COLLECTION
 */
type collection interface {
	createIterator() iterator
}
type itemCollection struct {
	items []*item
}
func (u *itemCollection) createIterator() iterator {
	return &itemIterator{
		items: u.items,
	}
}

/*
ITERATOR
 */
type iterator interface {
	hasNext() bool
	getNext() *item
}
type itemIterator struct {
	index int
	items []*item
}
func (u *itemIterator) hasNext() bool {
	if u.index < len(u.items) {
		return true
	}
	return false
}
func (u *itemIterator) getNext() *item {
	if u.hasNext() {
		item := u.items[u.index]
		u.index++
		return item
	}
	return nil
}


type item struct {
	name string
	age  int
}

/*
CLIENT
 */
func main() {
	item1 := &item{
		name: "A",
		age:  30,
	}
	item2 := &item{
		name: "B",
		age:  20,
	}

	itemCollection := &itemCollection{
		items: []*item{item1, item2},
	}

	iterator := itemCollection.createIterator()

	for iterator.hasNext() {
		item := iterator.getNext()
		fmt.Printf("Item is %+v\n", item)
	}
}
