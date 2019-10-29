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
	parentID string
	childID list.List
	next_pcb *PCB
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
		parentID:"None",
	}
}

func init(init *PCB){
	//初始化资源
	Available_list["R1"]=1
	Available_list["R2"]=2
	Available_list["R3"]=3
	Available_list["R4"]=4

	//
	CurrentPID=init
}
		
//定义创建函数 不确定是否需要在这里READY_list 是否需要导入包
func Create(pid string,prioirty int,ready_list *READY_list) PCB{
	pcb:=New_PCB_no_param()
	pcb.pid=pid
	pcb.prioirty=prioirty
	pcb.parentID=CurrentPCB.pid
	CurrentPCB.childID.PushBack(pid)
	if prioirty==0{
		init(&pcb)
	}else if prioirty==1{
		ready_list.user.current_node=&pcb
	}
	else if prioirty==2{
		ready_list.system.current_node=&pcb
	}
	//不确定函数内部的结构体会不会被回收
	Schedule_after_create()
		
}
		


