package rt

type Frame struct {
	lower        *Frame
	localVars    LocalVars
	operandStack *OperandStack
	thread       *Thread
	nextPC       int
}

func (f Frame) Thread() *Thread {
	return f.thread
}

func (f Frame) NextPC() int {
	return f.nextPC
}

func (f Frame) SetNextPC(nextPC int) {
	f.nextPC = nextPC
}

func NewFrame(maxLocals uint16, maxStack uint16) *Frame {
	return &Frame{
		localVars:    newLocalVars(maxLocals),
		operandStack: newOperandStack(maxStack),
	}
}

func newFrame(thread *Thread, maxLocals uint16, maxStack uint16) *Frame {
	return &Frame{
		thread:       thread,
		localVars:    newLocalVars(maxLocals),
		operandStack: newOperandStack(maxStack),
	}
}

func (f Frame) LocalVars() LocalVars {
	return f.localVars
}

func (f Frame) OperandStack() *OperandStack {
	return f.operandStack
}
