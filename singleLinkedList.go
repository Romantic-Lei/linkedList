package main
import (
	"fmt"
)

// 定义一个HeroNode
type HeroNode struct {
	no int
	name string
	nickName string
	next *HeroNode // 这个表示指向下一个节点
}

// 给链表插入一个节点
// 编写第一种插入方法， 在单链表的最后加入
//  head *HeroNode 头结点  newHeroNode *HeroNode 新增节点
func InsertHeroNode(head *HeroNode, newHeroNode *HeroNode) {
	// 1. 先找到这个链表的最后一个节点
	// 2. 创建一个辅助节点
	temp := head
	for {
		if temp.next == nil { 
			// 找到链表的最后一条数据，跳出
			break
		}

		temp = temp.next // 让 temp 不断的指向下一个节点
	}
	// 3. 将 newHeroNode 加入到链表的最后
	temp.next = newHeroNode
} 

// 显示链表的所有信息
func ListHeroNode(head *HeroNode) {
	// 1. 创建一个辅助节点
	temp := head
	// 先判断该链表是否是空链表
	if temp.next == nil {
		fmt.Println("链表为空")
		return 
	}

	// 2. 遍历这个链表
	for {
		fmt.Printf("[%d, %s, %s] -->", temp.next.no, temp.next.name,
			temp.next.nickName)
		// 判断链表是否到了最后
		temp = temp.next
		if temp.next == nil {
			break
		}
	}
}

func main() {
	// 1. 先创建一个头结点,可以让其都是默认是
	head := &HeroNode{}

	// 2. 创建一个新的HeroNode
	hero1 := &HeroNode{
		no : 1,
		name : "宋江",
		nickName : "及时雨",
	}
	hero2 := &HeroNode{
		no : 1,
		name : "卢俊义",
		nickName : "玉麒麟",
	}
	// 3. 加入
	InsertHeroNode(head,hero1)
	InsertHeroNode(head,hero2)
	// 4. 显示
	ListHeroNode(head)
}