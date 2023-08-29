package base

import "nilknow-jvm/rt"

func Branch(frame *rt.Frame,offset int) {
	pc:=frame.Thread().PC()
	nextPc:=pc+offset
	frame.SetNextPC(nextPc)
}