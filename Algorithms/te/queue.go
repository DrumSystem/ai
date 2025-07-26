package main

import "errors"

type CircularQueue struct {
	data []interface{}
	front int
	rear int
	size int
}

func NewCircularQueue(cap int) *CircularQueue {
	return &CircularQueue{
		data: make([]interface{}, cap),
		front: 0,
		rear: 0,
		size: 0,
	}
}

// IsEmpty 检查队列是否为空
func (q *CircularQueue) IsEmpty() bool {
	return q.size == 0
}

// IsFull 检查队列是否已满
func (q *CircularQueue) IsFull() bool {
	return q.size == len(q.data)
}


func (q *CircularQueue) Enqueue(item interface{}) error {
	if q.IsFull() {
		return errors.New("queue is full")
	}

	q.data[q.rear] = item
	q.rear = (q.rear + 1 ) % len(q.data)
	q.size++
	return nil
}


func (q *CircularQueue) Dequeue() (interface{}, error) {
	if q.IsEmpty() {
		return nil,  errors.New("queue is empty")
	}

	item := q.data[q.front]
	q.data[q.front] = nil
	q.front = (q.front + 1 ) % len(q.data)
	q.size--
	return item, nil
}
