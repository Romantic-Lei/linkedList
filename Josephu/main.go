package main
import (
	"fmt"
)

type Boy struct{
	No int // 编号
	next *Boy // 指向下一个小孩的指针
}

// 编写一个函数，构成单向的环形链表
// num : 表示小孩的个数
// *Boy ： 返回该环形链表的第一个小孩的指针
func AddBoy(num int) *Boy {
	first := &Boy{} // 当成头指针
	curBoy := &Boy{} // 当成临时指针
	// 判断 
	if num < 1 {
		fmt.Print("num 值不对")
		return first
	}
	// 循环构建这个环形链表
	for i := 1; i <= num; i++ {
		boy := &Boy{
			No : i,
		}
		// 1. 因为是第一个小孩，比较特殊
		if i == 1 { // 第一个小孩
			first = boy
			curBoy = boy
			curBoy.next = first // 连成环形
		} else {
			curBoy.next = boy
			curBoy = boy
			curBoy.next = first // 构成环形链表
		}
	}
	return first
}

// 显示单向环形链表[遍历]
func ShowBoy(first *Boy) {
	// 处理一下如果环形链表为空
	if first.next == nil {
		fmt.Println("链表为空，没有小孩")
		return 
	}
	// 创建一个指针， 帮助遍历
	curBoy := first
	for {
		fmt.Printf("小孩编号 = %d ->", curBoy.No)
		// 退出的条件
		if curBoy.next == nil {
			
		}
	}
}

func main() {

}