/*
 * xingchen语迷你版本编译器
 *
 * 仅支持ASCII
 * 仅支持三种命令
 * 仅支持解释运行
 *
 * 提示~请用 <?php
 *			echo str_pad("", 72, "+") . '.' . str_pad("", 101 - 72, "+") . '.' . str_pad("", 108 - 101, "+") . '..' . str_pad("", 111 - 108, "+") . '.' . str_pad("", 111 - 32, "-") . '.' . str_pad("", 87 - 32, "+") . '.' . str_pad("", 114 - 87, "+") . '.' . str_pad("", 114 - 111, "-") . '.' . str_pad("", 111 - 108, "-") . '.' . str_pad("", 108 - 100, "-") . '.' . str_pad("", 100 - 33, "-") . '.';
 * eg:  go run xc_mini.go "++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++.+++++++++++++++++++++++++++++.+++++++..+++.-------------------------------------------------------------------------------.+++++++++++++++++++++++++++++++++++++++++++++++++++++++.+++++++++++++++++++++++++++.---.---.--------.-------------------------------------------------------------------."
 *      Hello World!
 *
 */
package main

import (
	"fmt"
	"os"
)

type BlockSimple struct {   //连续字节内存块
	mem []byte
	pos int
}

func BlockNewMini() *BlockSimple{
	block := new(BlockSimple)
	block.mem = make([]byte, 10240)
	return block
}

func BuildMini(command []byte){
	current := 0
	block := BlockNewMini()
	for current < len(command) {
		switch command[current] {
		case '+':
			block.mem[block.pos]++
		case '-':
			block.mem[block.pos]--
		case '.':
			fmt.Printf("%c", block.mem[block.pos])
		}
		current++
	}
}

func ParseMini(command []byte) ([]byte){
	parsed := make([]byte, 0)   //已解析
	current := 0
	for _, char := range command {
		switch char {
		case '+', '-', '.':
			parsed = append(parsed, char)
		}
		current++
	}
	return parsed
}

func main(){
	command := []byte(os.Args[1])
	BuildMini(ParseMini(command))
}