package module

func Request(resource string, num int) {
	avail := Availablelist[resource]
	if num <= avail {
		//fmt.Println("足够")
		Availablelist[resource] -= num
		CurrentPCB.occupyResource[resource] += num
	} else {
		//fmt.Println("不够")
		CurrentPCB.requestResource[resource] += num
		Insert(Blocklist, CurrentPCB)
		ScheduleafterRequest()
	}
}

func From_block_to_ready(pcb *PCB) {
	Availablelist["R1"] -= pcb.requestResource["R1"]
	Availablelist["R2"] -= pcb.requestResource["R2"]
	Availablelist["R3"] -= pcb.requestResource["R3"]
	Availablelist["R4"] -= pcb.requestResource["R4"]

	pcb.occupyResource["R1"] += pcb.requestResource["R1"]
	pcb.occupyResource["R2"] += pcb.requestResource["R2"]
	pcb.occupyResource["R3"] += pcb.requestResource["R3"]
	pcb.occupyResource["R4"] += pcb.requestResource["R4"]

	pcb.requestResource["R1"] = 0
	pcb.requestResource["R1"] = 0
	pcb.requestResource["R1"] = 0
	pcb.requestResource["R1"] = 0
}

func Release(pcb *PCB) {
	//fmt.Println("释放的pcb", pcb.pid)
	rel1 := pcb.occupyResource["R1"]
	rel2 := pcb.occupyResource["R2"]
	rel3 := pcb.occupyResource["R3"]
	rel4 := pcb.occupyResource["R4"]

	pcb.occupyResource["R1"] = 0
	pcb.occupyResource["R2"] = 0
	pcb.occupyResource["R3"] = 0
	pcb.occupyResource["R4"] = 0

	Availablelist["R1"] += rel1
	Availablelist["R2"] += rel2
	Availablelist["R3"] += rel3
	Availablelist["R4"] += rel4

	Scheduleafterrelease()
}

func Release_only_one(resource string) {
	Availablelist[resource] += CurrentPCB.occupyResource[resource]
	CurrentPCB.occupyResource[resource] = 0
	Scheduleafterrelease()
}

func Release_not_only_one(resource string, n int) {
	Availablelist[resource] += n
	CurrentPCB.occupyResource[resource] -= n
	Scheduleafterrelease()
}
