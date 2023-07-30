package queue

import (
	"errors"
	"fmt"
)

const MAXIMUM_QUEUE_LENGTH = 100

type node struct {
	s    interface{}
	next *node
}

func newNode(s interface{}) *node {
	node := new(node)
	node.s = s
	node.next = nil
	return node
}

type queue struct {
	head *node
	tail *node
	q    []node
}

func Init() *queue {
	q := new(queue)
	q.head = nil
	q.tail = nil
	q.q = make([]node, MAXIMUM_QUEUE_LENGTH)
	return q
}

func (q *queue) EnQueue(s interface{}) {
	n := newNode(s)
	if q.head == nil {
		q.head = n
	}
	if q.tail != nil {
		q.tail.next = n
	}
	q.tail = n
}

func (q *queue) DeQueue() (s interface{}, err error) {
	if q.head == nil {
		err = errors.New("can not dequeue and empty queue")
		return
	}
	s = q.head.s
	err = nil
	q.head = q.head.next
	return
}

func (q *queue) Length() int {
	i := 0
	h := q.head
	for h != nil {
		i++
		h = h.next
	}
	return i
}

func (q *queue) Search(s interface{}) bool {
	h := q.head
	for h != nil {
		if h.s == s {
			return true
		}
		h = h.next
	}
	return false
}

func (q *queue) String() string {
	h := q.head
	var s string
	for h != nil {
		s += fmt.Sprintf("%s\t", h.s)
		h = h.next
	}
	return s
}
