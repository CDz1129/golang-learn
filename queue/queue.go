package queue

type Queue []int

func (q *Queue) Push(v int) {

	*q = append(*q, v)
}

func (q *Queue) Pop() (r int,ok bool){
	if len(*q)== 0{
		return 0,false
	}
	head := (*q)[0]
	*q = (*q)[1:]
	return head,true
}

func (q *Queue) IsEmpty() bool {
	return len(*q) == 0
}
