package module

import (
	"fmt"
	"container/list"
)

//定义常量
const Ready=1
const Block=2
const Running=0
const Uninit=-1

//定义结构体数组
type PCB struct{
	pid string
	prioirty int
	requstResource map[string]int
	state int
	parentID *PCB
	childID *PCB
}

//定义接收者函数
func New_PCB_no_param() *PCB{
	init_requstResource:=make(map[string]int)
	init_requstResource["R1"]=0
	init_requstResource["R2"]=0
	init_requstResource["R3"]=0
	init_requstResource["R4"]=0    
	return &PCB{
		pid:"None",
		prioirty:-1,
		requstResource:init_requstResource,
		state:uninit,
		parentID:nil,
		childID:nil,
	}
}
		
//定义创建函数
func Create(pid string,prioirty int,currentPID *PCB){
	pcb:=New_PCB_no_param()
	pcb.pid=pid
	pcb.prioirty=prioirty
	pcb.parentID=currentPID
}
		


