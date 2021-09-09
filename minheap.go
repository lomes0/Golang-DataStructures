package main

import (
	"reflect"
)

type minheap struct {
	/*
	* h[0] = at
	* h[1] = global min
	* h[k] = local min
	 */
	arr []int
}

/*
* @Desc: Return a new empty heap.
*	 heap.arr[0] := heap current index.
 */
func NewMinHeap() *minheap {

	h := &minheap{arr: make([]int, 1)}
	arr := h.arr
	arr[0] = 1

	return h
}

/*
*@Desc: Return heap's number of values.
*	Meaning, return: |{arr[1],...,arr[n]}|.
*
 */
func (h *minheap) size() int {
	arr := h.arr
	return len(arr) - 1
}

/*
*
*
 */
func shift_up(h *minheap, i int) {

	swap := reflect.Swapper(h.arr)

	if i == 1 {
		return
	}

	if h.at(i/2) <= h.at(i) {
		return
	}

	swap(i/2, i)
	shift_up(h, i/2)
}

/*
* @Desc: Insert a new value.
*
 */
func (h *minheap) insert(val int) *minheap {
	h.arr = append(h.arr, val)
	shift_up(h, h.size())
	return h
}

/*
*
*
 */
func shift_down(h *minheap, i int) {

	swap := reflect.Swapper(h.arr)

	if i*2 > h.size() {
		return
	}

	var m int = 2*i + 1
	if m > h.size() || h.at(2*i) <= h.at(m) {
		m = 2 * i
	}

	if h.at(i) <= h.at(m) {
		return
	}

	swap(i, m)
	shift_down(h, i+1)
}

func (h *minheap) empty() bool {
	return h.size() == 0
}

/*
* @Desc: Remove global minimun.
*
 */
func (h *minheap) pop() int {

	swap := reflect.Swapper(h.arr)

	if h.empty() {
		return 0
	}

	ret := h.at(1)
	swap(1, h.size())
	h.arr = h.arr[0:h.size()]

	shift_down(h, 1)

	return ret
}

/*
* @Desc: Return local root index.
*
 */
func (h *minheap) root() int {
	return h.at(0)
}

/*
* @Desc: Set local root position.
*
 */
func (h *minheap) set_root(i int) *minheap {
	h.arr[0] = i
	return h
}

func (h *minheap) get() int {
	return h.at(h.at(0))
}

func (h *minheap) at(i int) int {

	return h.arr[i]
}

/*
*@Desc: Set upper leaf as local root.
*
 */
func (h *minheap) up() *minheap {

	if h.size() == 1 {
		return h
	}

	i := h.root() / 2
	h.set_root(i)
	return h
}

/*
*@Desc: Set left leaf as local root.
*
 */
func (h *minheap) left() *minheap {

	i := h.root() * 2

	if i > h.size() {
		return h
	}

	h.set_root(i)
	return h
}

/*
*@Desc: Set right leaf as local root.
*
 */
func (h *minheap) right() *minheap {

	i := h.root()*2 + 1

	if i > h.size() {
		return h
	}

	h.set_root(i)
	return h
}
