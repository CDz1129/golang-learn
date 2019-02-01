package queue

type Queue []interface{}

func (q *Queue) Push(v interface{}) {

	*q = append(*q, v)
}

func (q *Queue) Pop() (r interface{}, ok bool) {
	if len(*q) == 0 {
		return 0, false
	}
	head := (*q)[0]
	*q = (*q)[1:]
	return head, true
}

func (q *Queue) IsEmpty() bool {
	return len(*q) == 0
}
