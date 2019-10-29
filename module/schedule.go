package module

import(
	"container/list"
)

var CurrentPCB *PCB  // 全局变量 表示当前正在执行的进程


//表示ready队列
type PCB_list struct{
	init_end *PCB
	init_head *PCB
	user_head *PCB
	user_end *PCB
	system_head *PCB
	system_end *PCB
}

func NewPcbList() *PCB_list{
	init_head_pcb=New_PCB_no_param()
	user_head_pcb=New_PCB_no_param()
	system_head_pcb=New_PCB_no_param()
	return &PCB_list{
		init_head:init_head_pcb,
		init_end:init_head_pcb,
		user_head:user_head_pcb,
		user_end:user_head_pcb,
		system_head:system_head_pcb,
		system_end:system_head_pcb,
	}
}

func Insert(pcb_list *PCB_list, pcb *PCB){
	if pcb.prioirty==0{
		pcb_list.init_end.childID=pcb
		pcb_list.init_end=pcb
	}else if pcb.priority==1{
		pcb_list.user_end.childID=pcb
		pcb_list.user_end=pcb
	}else if pcb.prioirty==2{
		pcb_list.system_end.childID=pcb
		pcb_list.system_end=pcb
	}
}

//注意要保证链表中存在元素
func Pop(pcb_list *PCB_list) *PCB{
	if pcb.prioirty==0{
		temp:=pcb_list.init_head.next_pcb
		pcb_list.init_head.next_pcb=temp.next_pcb
		temp.next_pcb=nil
		return temp
	}else if pcb.priority==1{
		temp:=pcb_list.user_head.next_pcb
		pcb_list.user_head.next_pcb=temp.next_pcb
		temp.next_pcb=nil
		return temp
	}else if pcb.priority==2{
		temp:=pcb_list.system_head.next_pcb
		pcb_list.system_head.next_pcb=temp.next_pcb
		temp.next_pcb=nil
		return temp
	}
}


ready_list:=NewReadyList()
block_list:=NewBlockList()


func Select() *pcb{
	if ready_list.system!=nil{
		return ready_list.system.Front(),2
	}else if ready_list.usr!=nil{
		return ready_list.user.Front(),1
	}else if ready_list.init!=nil{
		return ready_list.init.Front(),0
	}else{
		return nil,-1
	}
}

func Schedule_after_create(){
	select_pcb,select_prioirty:=Select()
	if select_pcb!=nil{
		//在正在运行的进程优先级小于刚刚创建的进程优先级的时候
		//将正在运行的进程挂起，放置于ready队列的最后端
		//并且将对应的新创建的进程执行
		if CurrentPID.prioirty<select_prioirty{
			if CurrentPID.prioirty==2{
				ready_list.system.
			}else if CurrentPID.prioirty==1{
				ready_list.user.PushBack(CurrentPID)
			}else if CurrentPID.prioirty==0{
				ready_list.init.PushBack(CurrentPID)
			}

			if select_prioirty==2{
				temp=ready_list.system.Front()
				CurrentPID=temp.Value
				ready_list.system.Remove(temp)
			}else if select_prioirty==1{
				temp=ready_list.user.Front()
				CurrentPID=temp.Value
				ready_list.user.Remove(temp)
			}else if select_prioirty==0{
				temp=ready_list.init.Front()
				CurrentPID=temp.Value
				ready_list.init.Remove(temp)
			}
		}
	}
}

			
			
