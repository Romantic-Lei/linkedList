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
		head.next = head // 形成一个环状
		fmt.Println(newCatNode, "加入到环状的链表")
		return 
	}

	// 定义一个临时变量，去找到环形的最后结点
	temp := head
	for {
		if temp.next == head {
			break
		}
		temp = temp.next
	}
	// 加入到链表中
	temp.next = newCatNode
	newCatNode.next = head

}

// 删除一只猫
func DelCatNode(head *CatNode, id int) *CatNode {
	temp := head // 指向环形链表的最前面
	helper := head // 指向环形链表的最后
	// 空链表
	if temp.next == nil {
		fmt.Println("这是一个人空的环形单链表， 不能删除")
		return head
	}

	// 如果只有一个结点
	if temp.next == head { // 只有一个结点
		temp.next = nil
		return head
	}

	// 将 helper 定位到链表的最后
	for {
		if helper.next == head {
			break
		}
		helper = helper.next
	}

	// 如果有两个或者两个以上结点
	flag := true // 默认我们没有删除
	for {
		if temp.next == head {
			// 说明已经到最后一个了，但是还未比较,我们会在下面的 flag 中进行比较删除
			break
		}
		if temp.no == id {
			if temp == head {
				// 说明删除的是头结点
				head = head.next
			}
			// 找到，进行删除
			helper.next = temp.next
			fmt.Printf("猫猫 = %d 被删除\n", id)
			flag = false
			break
		}
		temp = temp.next // 移动 【比较】
		helper = helper.next // 移动 【找到结点后，通过他来删除】
	}
	if flag {
		// 如果 flag 为真，则我们在上面没有进行删除
		if temp.no == id {
			helper.next = temp.next
			fmt.Printf("猫猫 id = %d 被删除\n", id)
		} else {
			fmt.Printf("请确定你输入的猫猫id(%d) 是否正确\n", id)
		}
	}
	return head
}

// 输出环形链表
func ListCircleLink(head *CatNode) {
	fmt.Println("环形单链表显示如下：")
	temp := head
	if temp.next == nil {
		fmt.Println("环形链表为空")
		return
	}
	for {
		fmt.Printf("猫的信息为 = [id = %d name = %s] -> \n", temp.no, temp.name)
		if temp.next == head {
			break
		}
		temp = temp.next
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
	cat2 := &CatNode{
		no : 2,
		name : "Jerry",
	}
	cat3 := &CatNode{
		no : 3,
		name : "Mary",
	}
	InsertCatNode(head, cat1)
	InsertCatNode(head, cat2)
	InsertCatNode(head, cat3)
	ListCircleLink(head)
	head = DelCatNode(head, 5)

	ListCircleLink(head)
}