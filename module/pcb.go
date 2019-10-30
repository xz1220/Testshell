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
	occupyResource map[string]int
	requestResource map[string]int
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

	init_occupyResource:=make(map[string]int)
	init_occupyResource["R1"]=0
	init_occupyResource["R2"]=0
	init_occupyResource["R3"]=0
	init_occupyResource["R4"]=0
	return &PCB{
		pid:"None",
		prioirty:-1,
		occupyResource:init_occupyResource,
		requestResource:init_requstResource,
		state:uninit,
		parentID:"None",
		next_pcb:nil,
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
	pcb.pid=pid //进程命名
	pcb.prioirty=prioirty //设定进程优先级
	pcb.parentID=CurrentPCB.pid //设定父亲进程的名称
	CurrentPCB.childID.PushBack(pid) //设定父亲进程的子进程
	pcb.state=Ready
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

func Del_pcb_from_list(pre *PCB,current *PCB){
	temp:=current.next_pcb
	pre.next_pcb=temp
	return temp
}

func find_pcb_list(pid string,pcb_list *PCB_list) *PCB,*PCB{
	for point:=pcb_list.init_head;point.next_pcb!=nil;point=point.next_pcb{
		if point.next_pcb.pid==pid{
			return point,point.next_pcb
		}
	}
	for point:=pcb_list.user_head;point.next_pcb!=nil;point=point.next_pcb{
		if point.next_pcb.pid==pid{
			return point,point.next_pcb
		}
	}
	for point:=pcb_list.system_head;point.next_pcb!=nil;point=point.next_pcb{
		if point.next_pcb.pid==pid{
			return point,point.next_pcb
		}
	}
	return nil,nil
}

func find_pcb(pid string) *PCB,*PCB{
	pre1,point1:=find_pcb_list(pid,Ready_list)
	pre2,point2:=find_pcb_list(pid,Block_list)
	if point1!=nil{
		return pre1,point1
	}else point2!=nil{
		return pre2,point2
	}else if CurrentPCB.pid==pid{
		return nil,CurrentPCB
	}else{
		return nil,nil
	}
}

//这里应该是递归
func kill_child(child list.List){
	for i:=child.Front();i!=nil;i=i.Next(){
		value,ok:=interface{}(i.Value).(string)
		if ok==true{
			pre,point:=find_pcb(value)
			_:=Del_pcb_from_list(pre,point)
		}else{
			fmt.Println("类型转化出错")
		}
	}
}

func kill_tree(pre *PCB,point *PCB){
	if pre==nil && point==nil{
		fmt.Println("进程不存在，无法删除")
	}else if pre==nil && point!=nil{
		fmt.Println("开始删除正在运行的进程")
		kill_child(point.childID)
		CurrentPCB=Pop(Ready_list)
	}else{
		fmt.Println("开始删除队列里面的进程")
		_:=Del_pcb_from_list(pre,point)
		kill_child(point.childID)
	}
}

	
func Destory(pid string){
	pre,point:=find_pcb(pid)
	kill_tree(pre,point)
	Schedule_after_destory()
}
