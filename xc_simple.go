/*
 * xingchen语简单版本编译器
 *
 * 仅支持ASCII
 * 仅支持解释运行
 *
 *
 * eg:  go run xc_simple.go "++++++++++[>+++++++>++++++++++>+++>+<<<<-]>++.>+.+++++++..+++.>++.<<+++++++++++++++.>.+++.------.--------.>+.>."
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

func BlockNewSimple() *BlockSimple{
	block := new(BlockSimple)
	block.mem = make([]byte, 10240)
	return block
}

func BuildSimple(command []byte, loopmap map[int]int){
	current := 0
	block := BlockNewSimple()
	var input byte
	for current < len(command) {
		switch command[current] {
		case '+':
			block.mem[block.pos]++
		case '-':
			block.mem[block.pos]--
		case '>':
			block.pos++
			if len(block.mem) < block.pos {
				block.mem = append(block.mem, 0)    //如果当前指向为空块，则初始化
			}
		case '<':
			block.pos--
		case '[':
			if block.mem[block.pos] == 0 {
				current = loopmap[current]  //如果循环过后指向已经是0，结束循环，移动指针到尾部
			}
		case ']':
			if block.mem[block.pos] != 0 {
				current = loopmap[current]  //移动到头部
			}
		case '.':
			fmt.Printf("%c", block.mem[block.pos])
		case ',':
			fmt.Scanf("%c", &input)
			block.mem[block.pos] = input
		}
		current++
	}
}

func ParseSimple(command []byte) ([]byte, map[int]int){
	parsed := make([]byte, 0)   //已解析
	loopstack := make([]int, 0) //循环栈，也就是头部集
	loopmap := make(map[int]int, 1024)  //记录循环对应首尾位置
	current := 0
	for _, char := range command {
		switch char {
		case '+', '-', '<', '>', '[', ']', '.', ',':
			parsed = append(parsed, char)
			if char == '[' {
				loopstack = append(loopstack, current)
			} else if char == ']' {
				lasthead := len(loopstack) - 1
				headposition := loopstack[lasthead]
				loopstack = loopstack[:lasthead]
				loopmap[current] = headposition
				loopmap[headposition] = current
			}
		}
		current++
	}
	return parsed, loopmap
}

func main(){
	if len(os.Args) < 2 {
		fmt.Println("请输入指令")
		os.Exit(1)
	}

	command := []byte(os.Args[1])
	BuildSimple(ParseSimple(command))
}


