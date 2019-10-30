package main

import (
	"Testshell/module"
	"bufio"
	"io"
	"os"
	"strconv"
)

func processLine(line []byte) {
	//os.Stdout.Write(line)
	command := string(line[:])
	if len(command) == 7 {
		function := string(command[0:2])
		pcb := string(command[3:4])
		priority, _ := strconv.Atoi(string(command[5:6]))
		//fmt.Println(function, pcb, priority)
		if function == "cr" {
			module.Create(pcb, priority)
			module.Log()
		}
	} else if len(command) == 3 {
		function := string(command[0:2])
		if function == "to" {
			module.Timeout()
			module.Log()
		}
	} else if len(command) == 9 {
		function := string(command[0:3])
		resource := string(command[4:6])
		num, _ := strconv.Atoi(string(command[7:8]))
		if function == "req" {
			module.Request(resource, num)
			module.Log()
		}
	} else if len(command) == 5 {
		function := string(command[0:2])
		pcb := string(command[3:4])
		if function == "de" {
			module.Destory(pcb)
			module.Log()
		}
	}
}

func ReadLine(filePth string, hookfn func([]byte)) error {
	f, err := os.Open(filePth)
	if err != nil {
		return err
	}
	defer f.Close()

	bfRd := bufio.NewReader(f)
	for {
		line, err := bfRd.ReadBytes('\n')
		hookfn(line)    //放在错误处理前面，即使发生错误，也会处理已经读取到的数据。
		if err != nil { //遇到任何错误立即返回，并忽略 EOF 错误信息
			if err == io.EOF {
				return nil
			}
			return err
		}
	}
	return nil
}

func main() {
	module.Create("init", 0)
	module.Log()
	ReadLine("D:/input.txt", processLine)
	module.Show_pcb("r")
}
