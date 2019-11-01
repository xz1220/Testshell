package module

//PCBlist 表示三条PCB单链表的封装
type PCBlist struct {
	inithead   *PCB
	userhead   *PCB
	systemhead *PCB
}

//NewPcbList is
func NewPcbList() *PCBlist {
	initheadpcb := NewPCBnoparam()
	userheadpcb := NewPCBnoparam()
	systemheadpcb := NewPCBnoparam()
	return &PCBlist{
		inithead:   initheadpcb,
		userhead:   userheadpcb,
		systemhead: systemheadpcb,
	}
}

//insertlist is
func insertlist(list *PCB, pcb *PCB) {
	var point *PCB
	for point = list; point.nextpcb != nil; point = point.nextpcb {
		//fmt.Println("经过了", point.pid)
	}
	point.nextpcb = pcb

}

//Insert is
func Insert(pcblist *PCBlist, pcb *PCB) {
	if pcb.prioirty == 0 {
		//fmt.Println("success 0")
		insertlist(pcblist.inithead, pcb)
	} else if pcb.prioirty == 1 {

		//fmt.Println("success 1")
		insertlist(pcblist.userhead, pcb)
	} else if pcb.prioirty == 2 {
		insertlist(pcblist.systemhead, pcb)
	}
}

//注意要保证链表中存在元素
func Pop(pcblist *PCBlist) *PCB {
	if pcblist.systemhead.nextpcb != nil {
		temp := pcblist.systemhead.nextpcb
		pcblist.systemhead.nextpcb = temp.nextpcb
		temp.nextpcb = nil
		return temp
	} else if pcblist.userhead.nextpcb != nil {
		//fmt.Println("tan chu user")
		temp := pcblist.userhead.nextpcb
		pcblist.userhead.nextpcb = temp.nextpcb
		temp.nextpcb = nil
		//fmt.Println("tan chu user", temp.pid)
		return temp
	} else if pcblist.inithead.nextpcb != nil {
		//fmt.Println("tan chu init")
		temp := pcblist.inithead.nextpcb
		pcblist.inithead.nextpcb = temp.nextpcb
		temp.nextpcb = nil
		return temp
	} else {
		return nil
	}
}

func Schedule() {
	Scheduleaftercreate()
}

func Selectfromready() (*PCB, int) {
	if Readylist.systemhead.nextpcb != nil {
		return Readylist.systemhead.nextpcb, 2
	} else if Readylist.userhead.nextpcb != nil {
		return Readylist.userhead.nextpcb, 1
	} else if Readylist.inithead.nextpcb != nil {
		return Readylist.inithead.nextpcb, 0
	} else {
		return nil, -1
	}
}

func Selectfromblock() (*PCB, int) {
	if Blocklist.systemhead.nextpcb != nil {
		return Blocklist.systemhead.nextpcb, 2
	} else if Blocklist.userhead.nextpcb != nil {
		return Blocklist.userhead.nextpcb, 1
	} else if Blocklist.inithead.nextpcb != nil {
		return Blocklist.inithead.nextpcb, 0
	} else {
		return nil, -1
	}
}

func Scheduleaftercreate() {
	selectpcb, selectprioirty := Selectfromready()
	//fmt.Println(selectprioirty)
	if selectpcb != nil {
		//在正在运行的进程优先级小于刚刚创建的进程优先级的时候
		//将正在运行的进程挂起，放置于ready队列的最后端
		//并且将对应的新创建的进程执行
		if CurrentPCB.prioirty < selectprioirty {
			//fmt.Println("jinru")
			//fmt.Println("调度时")
			Insert(Readylist, CurrentPCB)
			CurrentPCB = Pop(Readylist)
			//fmt.Println("输出", CurrentPCB.pid)

		}
	}
}

func Scheduleafterdestory() {
	//fmt.Println("在执行删除后调用")
	Schedule()
}

func ScheduleafterRequest() {
	CurrentPCB = Pop(Readylist)
}

func Scheduleafterrelease() {
	//fmt.Println("开始")
	//简化版 只写了user
	for point := Blocklist.userhead; point.nextpcb != nil; point = point.nextpcb {
		//fmt.Println("循环中")
		//fmt.Println(point.nextpcb.requestResource["R1"], point.nextpcb.requestResource["R2"], point.nextpcb.requestResource["R3"], point.nextpcb.requestResource["R4"])
		if point.nextpcb.requestResource["R1"] <= Availablelist["R1"] && point.nextpcb.requestResource["R2"] <= Availablelist["R2"] && point.nextpcb.requestResource["R3"] <= Availablelist["R3"] && point.nextpcb.requestResource["R4"] <= Availablelist["R4"] {
			//fmt.Println("out before:", point.pid)
			out := Delpcbfromlist(point, point.nextpcb)
			From_block_to_ready(out)
			//fmt.Println("out:", out.pid)
			//fmt.Println("找到了")
			Insert(Readylist, out)
			Schedule()
		}
	}
}

func Timeout() {
	Insert(Readylist, CurrentPCB)
	CurrentPCB = Pop(Readylist)
}
