package module

import (
	"container/list"
	"fmt"
)

var Availablelist map[string]int

//Readylist is
var Readylist *PCBlist

//Blocklist is
var Blocklist *PCBlist

//CurrentPCB is
var CurrentPCB *PCB

//定义结构体数组
type PCB struct {
	pid             string
	prioirty        int
	occupyResource  map[string]int
	requestResource map[string]int
	parentID        string
	childID         list.List
	nextpcb         *PCB
}

//定义接收者函数
func NewPCBnoparam() *PCB {
	initrequstResource := make(map[string]int)
	initrequstResource["R1"] = 0
	initrequstResource["R2"] = 0
	initrequstResource["R3"] = 0
	initrequstResource["R4"] = 0

	initoccupyResource := make(map[string]int)
	initoccupyResource["R1"] = 0
	initoccupyResource["R2"] = 0
	initoccupyResource["R3"] = 0
	initoccupyResource["R4"] = 0
	return &PCB{
		pid:             "None",
		prioirty:        -1,
		occupyResource:  initoccupyResource,
		requestResource: initrequstResource,
		parentID:        "None",
		nextpcb:         nil,
	}
}

func initpcb(pcb *PCB) {
	Availablelist = make(map[string]int)
	//初始化资源
	Availablelist["R1"] = 1
	Availablelist["R2"] = 2
	Availablelist["R3"] = 3
	Availablelist["R4"] = 4

	//
	Readylist = NewPcbList()
	Blocklist = NewPcbList()

	//
	CurrentPCB = NewPCBnoparam()
	CurrentPCB = pcb

}

//定义创建函数 不确定是否需要在这里READYlist 是否需要导入包
func Create(pid string, prioirty int) {
	pcb := NewPCBnoparam()
	pcb.pid = pid           //进程命名
	pcb.prioirty = prioirty //设定进程优先级
	if prioirty == 0 {
		initpcb(pcb)
	} else {
		pcb.parentID = CurrentPCB.pid    //设定父亲进程的名称
		CurrentPCB.childID.PushBack(pid) //设定父亲进程的子进程
		//fmt.Println("插入时")
		Insert(Readylist, pcb)
		//Log_ready()
	}
	//不确定函数内部的结构体会不会被回收
	Scheduleaftercreate()
}

func Delpcbfromlist(pre *PCB, current *PCB) *PCB {
	//fmt.Println("正在删除", current.pid)
	temp := current.nextpcb
	pre.nextpcb = temp
	current.nextpcb = nil
	return current
}

func findpcblist(pid string, pcblist *PCBlist) (*PCB, *PCB) {
	for point := pcblist.inithead; point.nextpcb != nil; point = point.nextpcb {
		if point.nextpcb.pid == pid {
			return point, point.nextpcb
		}
	}
	for point := pcblist.userhead; point.nextpcb != nil; point = point.nextpcb {
		if point.nextpcb.pid == pid {
			return point, point.nextpcb
		}
	}
	for point := pcblist.systemhead; point.nextpcb != nil; point = point.nextpcb {
		if point.nextpcb.pid == pid {
			return point, point.nextpcb
		}
	}
	return nil, nil
}

func findpcb(pid string) (*PCB, *PCB) {
	pre1, point1 := findpcblist(pid, Readylist)
	pre2, point2 := findpcblist(pid, Blocklist)
	if point1 != nil {
		return pre1, point1
	} else if point2 != nil {
		return pre2, point2
	} else if CurrentPCB.pid == pid {
		return nil, CurrentPCB
	} else {
		return nil, nil
	}
}

//这里应该是递归
func killchild(child list.List) {
	for i := child.Front(); i != nil; i = i.Next() {
		value, ok := interface{}(i.Value).(string)
		fmt.Println("xiong di meng", value)
		if ok == true {
			pre, point := findpcb(value)
			Delpcbfromlist(pre, point)
		} else {
			fmt.Println("类型转化出错")
		}
	}
}

func killtree(pre *PCB, point *PCB) {
	if pre == nil && point == nil {
		fmt.Println("进程不存在，无法删除")
	} else if pre == nil && point != nil {
		fmt.Println("开始删除正在运行的进程")
		killchild(point.childID)
		CurrentPCB = Pop(Readylist)
	} else {
		//fmt.Println("开始删除队列里面的进程")
		Delpcbfromlist(pre, point)
		killchild(point.childID)
	}
}

func Destory(pid string) {
	pre, point := findpcb(pid)
	//fmt.Println(pre.pid, point.pid)

	killtree(pre, point)
	Release(point)
	Scheduleafterdestory()
}

func Log() {
	fmt.Println(CurrentPCB.pid)
}

func Log_ready() {
	fmt.Println("Ready 队列：")
	if Readylist.userhead.nextpcb == nil {
		fmt.Println("为空")
	}
	for point := Readylist.inithead; point.nextpcb != nil; point = point.nextpcb {
		fmt.Println("user里面的", point.nextpcb.pid)
	}
	for point := Readylist.userhead; point.nextpcb != nil; point = point.nextpcb {
		fmt.Println("user里面的", point.nextpcb.pid)
	}
	for point := Readylist.systemhead; point.nextpcb != nil; point = point.nextpcb {
		fmt.Println("user里面的", point.nextpcb.pid)
	}
}

func Log_block() {
	fmt.Println("Block 队列：")
	if Blocklist.userhead.nextpcb == nil {
		fmt.Println("为空")
	}
	for point := Blocklist.inithead; point.nextpcb != nil; point = point.nextpcb {
		fmt.Println("user里面的", point.nextpcb.pid)
	}
	for point := Blocklist.userhead; point.nextpcb != nil; point = point.nextpcb {
		fmt.Println("user里面的", point.nextpcb.pid)
	}
	for point := Blocklist.systemhead; point.nextpcb != nil; point = point.nextpcb {
		fmt.Println("user里面的", point.nextpcb.pid)
	}
}
func Logres() {
	fmt.Println(" R1:", Availablelist["R1"], " R2:", Availablelist["R2"], " R3:", Availablelist["R3"], " R4:", Availablelist["R4"])
}

func List_all_resource() {
	fmt.Println("可用资源")
	Logres()
	fmt.Println("被占用资源")
	Logres()
}

func List_all_process() {
	Log_ready()
	Log_block()
}

func Show_pcb(pid string) {
	_, point := findpcb(pid)
	fmt.Println("进程的PID为：", point.pid,
		"   进程的")
}
