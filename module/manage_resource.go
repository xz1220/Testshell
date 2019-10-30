package module

Available_list:=make(map[string]int)


func Request(pcb *PCB,resource string,num int){
	avail:=Available_list[resource]
	if num<=avail{
		Available_list[resource]-=num
		pcb.occupyResource[resource]+=num
	}else{
		pcb.requestResource[resource]+=num
		Insert(Block_list,pcb)
	}
}
		


func Release(pcb *PCB){
	rel1:=pcb.occupyResource["R1"]
	rel2:=pcb.occupyResource["R2"]
	rel3:=pcb.occupyResource["R3"]
	rel4:=pcb.occupyResource["R4"]

	pcb.occupyResource["R1"]=0
	pcb.occupyResource["R2"]=0
	pcb.occupyResource["R3"]=0
	pcb.occupyResource["R4"]=0

	Available_list["R1"]+=rel1
	Available_list["R2"]+=rel2
	Available_list["R3"]+=rel3
	Available_list["R4"]+=rel4

	Schedule_after_release()
}
