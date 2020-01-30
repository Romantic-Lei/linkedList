package main
import (
	"fmt"
)

// 定义猫的结构体
type CatNode struct {
	no int
	name string
	next *CatNode 
}

func InsertCatNode(head *CatNode, newCatNode *CatNode) {

	// 判断是不是添加第一只猫
	if head.next == nil {
		head.no = newCatNode.no
		head.name = newCatNode.name
		head.next = newCatNode // 形成一个环状
		fmt.Println(newCatNode, "加入到环状的链表")
		return 
	}

	// 定义一个临时变量，去找到环形的最后结点
	temp := head
	for {
		if temp.next == head {
			break
		}
		// 加入到链表中
		temp.next = newCatNode
	}

}

func main() {
	// 这里我们初始化一个环形链表的头结点
	head := &CatNode{}

	// 创建一只猫
	cat1 := &CatNode{
		no : 1,
		name : "Tom",
	}
	InsertCatNode(head, cat1)
}