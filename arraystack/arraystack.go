package arraystack

type ArrayStack struct {
	list []interface{}
}

func New() ArrayStack {
	var list = make([]interface{}, 0)

	return ArrayStack{list: list}
}

func (stack *ArrayStack) Push(value interface{}) {
	stack.list = append(stack.list, value)
}

func (stack *ArrayStack) Pop() (interface{}, bool) {

	if len(stack.list) == 0 {
		return 0, false
	}

	lastIndex := len(stack.list) - 1
	value := stack.list[lastIndex]
	stack.list = stack.list[:lastIndex]
	return value, true
}
