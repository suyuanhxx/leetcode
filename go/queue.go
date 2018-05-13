package main

import (
	"container/ring"
)

type Queue struct {
	Element []*int
	Size    int
	Front   int
	Rear    int
	MaxSize int
}

func (q *Queue) Poll() *int {
	if q.Size > 0 {
		item := q.Element[q.Front]
		q.Front -= 1
		q.Size -= 1
		q.Element = q.Element[:q.Front]
		return item
	}
	return nil
}

func (q *Queue) Add(item *int) {
	if (q.Rear+1)%q.MaxSize == q.Front {
		return
	}
	q.Element = append(q.Element, item)
	q.Size += 1
	q.Front += 1
}

func Josephus(n, m, k int) int {
	r := ring.New(n) // 环
	for i := 1; i <= n; i++ {
		r.Value,r = i,r.Next()
	}

	for i := 1; i < k; i++ {
		r = r.Next()
	}
	for r.Len() > 1 {
		for i := 1; i < m; i++ {
			r = r.Next()
		}
		p := r.Prev()
		p.Unlink(1) // remove current node r
		//r.Unlink(1)// remove r next node
		r = p.Next()
	}
	return r.Value.(int)
}
