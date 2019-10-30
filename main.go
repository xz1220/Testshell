package main

import (
	"bufio"
	"io"
	"os"
	"fmt"
	"strconv"
	"Testshell/module"
)

func processLine(line []byte) {
	//os.Stdout.Write(line)
	command:=string(line[:])
	if len(command)==7{
		function:=string(command[0:2])
		pcb:=string(command[3:4])
		prioirty,_:=strconv.Atoi(string(command[5:6]))
		fmt.Println(function,pcb,priority)
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
		hookfn(line) //放在错误处理前面，即使发生错误，也会处理已经读取到的数据。
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
	_:=Create("init",0,Ready_list)
	ReadLine("/root/input.txt", processLine)
}



