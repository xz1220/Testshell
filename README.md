# 操作系统实验
设计和实现进程与资源管理，并完成Test shell的编写，本次版本使用的是go语言。
## 数据结构
进程描述：
结构体 
PCB:
   	pid: 进程名称
	requsetResource: 进程所需求的资源
	status: 进程的状态 ready block running
	parent：父亲进程
	child：子进程
	priority：进程优先级
	
就绪队列 READY_LIST：
数组+单链表：
	数组表示优先级 0（init）、1（user），2（system）
	每个数组元素存储着一个就绪队列里面的head指针
	
阻塞队列 BLOCK_LIST:
list：
	存储着组设的PCB

剩余资源列表 available_list:
map:
	R1: 剩余数目
	R2：剩余数目
	R3：剩余数目
	R4：剩余数目
	
变量：
当前进程 currentPID

## 模块功能设计

### 进程创建和销毁模块：

定义PCB

创建进程

销毁进程

### 资源管理模块：

定义剩余资源列表


### 调度与时间片中断模块：

定义就绪队列以及阻塞队列

### 系统控制模块：

