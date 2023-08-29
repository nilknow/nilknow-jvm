package rt

type Stack struct {
	size    uint
	maxSize uint
	_top    *Frame
}

func NewStack(size uint) *Stack {
	return &Stack{
		size: size,
	}
}

func (s Stack) push(frame *Frame) {
	//todo
}

func (Stack) pop() *Frame {
	//todo
	return &Frame{}
}

func (Stack) top() *Frame {
	//todo
	return &Frame{}
}
