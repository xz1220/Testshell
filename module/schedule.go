package module

import(
	"container/list"
)

 
type PCB_list struct{
	init_head *PCB
	user_head *PCB
	system_head *PCB
}

func NewPcbList() *PCB_list{
	init_head_pcb=New_PCB_no_param()
	user_head_pcb=New_PCB_no_param()
	system_head_pcb=New_PCB_no_param()
	return &PCB_list{
		init_head:init_head_pcb,
		user_head:user_head_pcb,
		system_head:system_head_pcb,
	}
}

func insert_list(list *PCB,pcb *PCB){
	for point:=list;;point=point.next_pcb{
		if point.next_pcb==nil{
			point.next_pcb=pcb
			break
		}
	}
}

func Insert(pcb_list *PCB_list, pcb *PCB){
	if pcb.prioirty==0{
		insert_list(pcb_list.init_head,pcb)
	}else if pcb.priority==1{
		insert_list(pcb_list.user_head,pcb)
	}else if pcb.prioirty==2{
		insert_list(pcb_list.system_head,pcb)
	}
}

//注意要保证链表中存在元素
func Pop(pcb_list *PCB_list) *PCB{
	if pcb_list.system_head.next_pcb!=nil{
		temp:=pcb_list.system_head.next_pcb
		pcb_list.system_head.next_pcb=temp.next_pcb
		temp.next_pcb=nil
		return temp
	}else if pcb_list.user_head.next_pcb!=nil{
		temp:=pcb_list.user_head.next_pcb
		pcb_list.user_head.next_pcb=temp.next_pcb
		temp.next_pcb=nil
		return temp
	}else if pcb_list.init_head.next_pcb!=nil{
		temp:=pcb_list.init_head.next_pcb
		pcb_list.init_head.next_pcb=temp.next_pcb
		temp.next_pcb=nil
		return temp
	}else{
		return nil
	}
}



Ready_list:=NewPcbList()
Block_list:=NewPcbList()
var CurrentPCB *PCB

func Schedule(){
	Schedule_after_create()
}

func Select_from_ready() *pcb{
	if ready_list.system!=nil{
		return Pop(Ready_list,2),2
	}else if ready_list.usr!=nil{
		return Pop(Ready_list,1),1
	}else if ready_list.init!=nil{
		return Pop(Ready_list,0),0
	}else{
		return nil,-1
	}
}

func Schedule_after_create(){
	select_pcb,select_prioirty:=Select_from_ready()
	if select_pcb!=nil{
		//在正在运行的进程优先级小于刚刚创建的进程优先级的时候
		//将正在运行的进程挂起，放置于ready队列的最后端
		//并且将对应的新创建的进程执行
		if CurrentPID.prioirty<select_prioirty{
			Insert(Ready_list,CurrentPCB)
			CurrentPCB=Pop(Ready_list)
			
		}
	}
}

func Schedule_after_destory(){
	fmt.Println("在执行删除后调用")
	Schedule()
}

func Schedule_after_release(){
	//简化版 只写了user
	for point:=Block_list.user_head;point.next_pcb!=nil;point=point.next_pcb{
		if point.next_pcb.requestResource["R1"]>Available_list["R1"] && point.next_pcb.requestResource["R2"]>Available_list["R2"] && point.next_pcb.requestResource["R3"]>Available_list["R3"] && point.next_pcb.requestResource["R4"]>Available_list["R4"]{
			out=Del_pcb_from_list(point,point.next_pcb)
			Insert(Ready_list,out)
			Schedule()
		}
	}
}

			
func Timeout(){
	Insert(Ready_list,CurrentPCB)
	CurrentPCB=Pop(Ready_list)
}
