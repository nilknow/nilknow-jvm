package rt

import "math"

type OperandStack struct {
	size  int32
	slots []Slot
}

func newOperandStack(maxStack uint16) *OperandStack {
	if maxStack > 0 {
		return &OperandStack{
			slots: make([]Slot, maxStack),
		}
	}
	return nil
}

func (os OperandStack) PushSlot(slot Slot) {
	os.slots[os.size] = slot
	os.size++
}

func (os OperandStack) PopSlot() Slot {
	os.size--
	return os.slots[os.size]
}


func (os OperandStack) PushInt(val int32) {
	os.slots[os.size].num = val
	os.size++
}

func (os OperandStack) PopInt() int32 {
	os.size--
	return os.slots[os.size].num
}

func (os OperandStack) PushLong(val int64) {
	os.slots[os.size].num = int32(val)
	os.slots[os.size+1].num = int32(val)
	os.size += 2
}

func (os OperandStack) PopLong() int64 {
	os.size -= 2
	low := uint32(os.slots[os.size].num)
	high := uint32(os.slots[os.size+1].num)
	return int64(high)<<32 | int64(low)
}

func (os OperandStack) PushFloat(val float32) {
	bits := math.Float32bits(val)
	os.slots[os.size].num = int32(bits)
	os.size++
}

func (os OperandStack) PopFloat() float32 {
	os.size--
	bits := uint32(os.slots[os.size].num)
	return math.Float32frombits(bits)
}

func (os OperandStack) PushDouble(val float64) {
	bits := math.Float64bits(val)
	os.PushLong(int64(bits))
}

func (os OperandStack) PopDouble() float64 {
	bits := uint64(os.PopLong())
	return math.Float64frombits(bits)
}

func (os OperandStack) PushRef(ref *Object) {
	os.slots[os.size].ref = ref
	os.size++
}

func (os OperandStack) PopRef() *Object {
	os.size--
	ref := os.slots[os.size].ref
	//for gc
	os.slots[os.size].ref = nil
	return ref
}
