package arrayqueue

type ArrayQueue struct {
	list []interface{}
}

func New() ArrayQueue {
	var list = make([]interface{}, 0)

	return ArrayQueue{list: list}
}

func (queue *ArrayQueue) Push(value interface{}) {
	queue.list = append(queue.list, value)
}

func (queue *ArrayQueue) Pop() (interface{}, bool) {

	if len(queue.list) == 0 {
		return 0, false
	}

	value := queue.list[0]
	queue.list = queue.list[1:]
	return value, true
}
