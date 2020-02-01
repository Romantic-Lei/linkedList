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
func ShowBoy(first *Boy) int {
	// 统计一共有多少个小孩
	i := 0
	// 处理一下如果环形链表为空
	if first.next == nil {
		fmt.Println("链表为空，没有小孩")
		return i
	}
	// 创建一个指针， 帮助遍历
	curBoy := first
	for {
		fmt.Printf("小孩编号 = %d ->", curBoy.No)
		i++
		// 退出的条件 当下一条的next指向了头部的时候，遍历完毕
		if curBoy.next == first {
			break
		}
		// curBoy 移动到下一条
		curBoy = curBoy.next
	}
	return i
}

// first *Boy 头结点, startNo int 开始出列小孩编号, countNum int 每次移动个数 boyNum int 小孩总数
// 使用算法，按照要求，在环形链表中留下最后一个人
func PlayGame(first *Boy, startNo int, countNum int, boyNum int) {
	// 1. 空的链表我们单独的处理
	if first.next == nil {
		fmt.Println("链表为空，没有小孩")
		return 
	}
	// 可以在这个地方留一个判断， 判断 startNo <= 小孩的总数
	if startNo > boyNum {
		fmt.Println("效率过低，开始小孩编号小于小孩总数")
		return 
	}
	// 2. 需要定义一个辅助指针， 帮助我们删除小孩
	tail := first
	// 3. 让 tail 指向环形链表的最后一个小孩
	// 因为在删除小孩是，tail 需要用到
	for {
		if tail.next == first {
			// 说明 tail 已经都了最后一个小孩
			break
		}
		tail = tail.next
	}
	// 4. 让 first 移动到 startNo
	for i := 0; i < startNo - 1; i++ {
		first = first.next
		tail = tail.next
	}
	fmt.Println()
	fmt.Println("first = ", first, "tail = ", tail)
	// 5. 开始数 countNum， 然后就删除 first 指向的小孩
	for {
		// 开始数 countNum - 1 次
		for i := 0; i < countNum - 1; i++ {
			first = first.next
			tail = tail.next
		}
		fmt.Printf("编号为 %d 的小孩被移除 \n", first.No)
		// 删除 first 指向的小孩
		first = first.next
		tail.next = first
		// 如果 tail == first ，圈里面就只有一个小孩了
		if tail == first {
			break
		}
	}
	fmt.Printf("最后出圈的小孩编号为 %d 出圈 -> \n", first.No)
}

func main() {
	first := AddBoy(5)
	// 显示小孩
	i := ShowBoy(first)
	PlayGame(first, 2, 3, i)
}