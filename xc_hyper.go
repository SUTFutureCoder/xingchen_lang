/*
 * xingchen语hyper版本编译器
 *
 * UTF8全支持
 * 支持命令关键字多字符
 * 仅支持解释运行
 *
 * eg: go run xc_pro.go run 加加加加头右加加加加加加加加加加左减尾右加加加加加加加加加出出减出 => 110
 * 中文输出参考 http://bianma.911cha.com/
 *
 */
package main

import (
	"fmt"
	"os"
	"io/ioutil"
)

//连续字节内存块
type block struct {
	mem []rune
	commandLength []int
	pos int
}

func BlockNewHyper() *block {
	block := new(block)
	block.mem = make([]rune, 10240)
	return block
}

func BuildHyper(commandRune []rune, loopmap map[int]int){
	current := 0
	block := BlockNewHyper()
	var input string
	for current < len(commandRune) {
		command := string(commandRune[current])
		switch command {
		case "加":
			block.mem[block.pos]++
		case "减":
			block.mem[block.pos]--
		case "右":
			block.pos++ //指向的子块前进
			if len(block.mem) < block.pos {
				block.mem = append(block.mem, 0)    //创建一个新的子块
			}
		case "左":
			block.pos-- //指向的子块后退
		case "头":
			if block.mem[block.pos] == 0 {
				current = loopmap[current]  //如果已经是0，结束循环，移动指针到尾部
			}
		case "尾":
			if block.mem[block.pos] != 0 {
				current = loopmap[current]  //移动到头部
			}
		case "出":
			fmt.Printf("%c", block.mem[block.pos])
		case "入":
			fmt.Scanf("%s", &input)
			block.mem[block.pos] = ([]rune(input))[0]
		}
		current++
	}
}


func ParseHyper(command []rune) ([]rune, map[int]int){
	//已解析
	parsed := make([]rune, 0)
	//循环栈，也就是头部集
	loopstack := make([]int, 0)
	//记录循环对应位置
	loopmap := make(map[int]int, 1024)
	current := 0
	for _, runeChar := range command {
		char := string(runeChar)
		switch char {
		case "加", "减", "左", "右", "头", "尾", "出", "入":
			parsed = append(parsed, runeChar)
			if char == "头" {
				loopstack = append(loopstack, current)
			} else if char == "尾" {
				lasthead := len(loopstack) - 1
				headposition := loopstack[lasthead]
				loopstack = loopstack[:lasthead]    //清除已经获取的头部
				loopmap[current] = headposition
				loopmap[headposition] = current
			}
		}
		current++
	}
	return parsed, loopmap
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("请传入指令或.xc文件")
		os.Exit(1)
	}

	command := ""
	//允许直接传入指令或文件，如文件不存在则直接读取
	_, err := os.Stat(os.Args[1])
	if err == nil || os.IsExist(err){
		fileCommand, err := ioutil.ReadFile(os.Args[1])
		if err != nil {
			fmt.Println("读取文件失败")
			os.Exit(1)
		}
		command = string(fileCommand)
	} else {
		command = os.Args[1]
	}
	commandRune := []rune(command)
	BuildHyper(ParseHyper(commandRune))
}
