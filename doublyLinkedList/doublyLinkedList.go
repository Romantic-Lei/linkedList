package main
import (
	"fmt"
)

// 定义一个HeroNode
type HeroNode struct {
	no int
	name string
	nickName string
	pre *HeroNode // 这个表示指向前一个结点
	next *HeroNode // 这个表示指向下一个结点
}

// 给链表插入一个结点
// 编写第一种插入方法， 在双向链表的最后加入
//  head *HeroNode 头结点  newHeroNode *HeroNode 新增结点
func InsertHeroNode(head *HeroNode, newHeroNode *HeroNode) {
	// 1. 先找到这个链表的最后一个结点
	// 2. 创建一个辅助结点
	temp := head
	for {
		if temp.next == nil { 
			// 找到链表的最后一条数据，跳出
			// temp.next 为最后一条数据的 next 属性
			break // 跳出 for 循环
		}

		temp = temp.next // 让 temp 不断的指向下一个结点,循环
	}
	// 3. 将 newHeroNode 加入到链表的最后
	temp.next = newHeroNode
	newHeroNode.pre = temp
} 

// 编写第二种插入方法， 在双向链表中根据 no 的编号从小到大插入
func InsertHeroNode2(head *HeroNode, newHeroNode *HeroNode) {
	// 1. 先找到这个链表的最后一个结点
	// 2. 创建一个辅助结点
	temp := head
	flag := true // 假定编号不存在
	// 让插入的结点的no ，和temp 的下一个结点的 no比较
	for {
		if temp.next == nil {
			break
		} else if temp.next.no > newHeroNode.no {
			// 说明 newHeroNode 就应该插入到temp 的后面
			break
		} else if temp.next.no == newHeroNode.no {
			// 说明链表中已经有了这个 no， 就不让插入
			flag = false
			break
		}
		temp = temp.next
	}
	if !flag {
		fmt.Printf("对不起，编号 %d 已存在\n", temp.next.no)
	} else {
		newHeroNode.next = temp.next
		newHeroNode.pre = temp
		if temp.next != nil {
			// 当前结点是最后一个结点的时候，会出现空指针异常，所以需要加上这个判断
			temp.next.pre = newHeroNode
		}
		temp.next = newHeroNode


	}	
} 

// 双向链表删除一个结点
// currentNode *HeroNode 当前需要删除的结点， id int 结点编号no
func DelHeroNode(currentNode *HeroNode, id int) {
	// 找到，删除
	currentNode.next = currentNode.next.next
	if currentNode.next != nil {
		// 如果当前结点的下一条是最后一个结点，则不需要操作
		currentNode.next.pre = currentNode
	}
}

// 修改一个结点
// currentNode *HeroNode 当前需要修改的结点， id int 结点编号no
func UpdateHeroNode(currentNode *HeroNode, newHeroNode *HeroNode) {
	// 找到，修改
	newHeroNode.next = currentNode.next.next
	newHeroNode.pre = currentNode
	if newHeroNode.next != nil {
		newHeroNode.next.pre = newHeroNode
	}
	currentNode.next = newHeroNode
}

// 显示链表的所有信息
// 可以仍然使用单向链表的显示方法
func ListHeroNode(head *HeroNode) {
	// 1. 创建一个辅助结点
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
	fmt.Println()
}

// 逆序打印双向链表
func ListHeroNode2(head *HeroNode) {
	// 1. 创建一个辅助结点
	temp := head
	// 先判断该链表是否是空链表
	if temp.next == nil {
		fmt.Println("链表为空")
		return 
	}
	// 2. 让 temp 定义到双向链表的最后
	for {
		if temp.next == nil {
			break
		}
		temp = temp.next
	}

	// 2. 遍历这个链表
	for {
		fmt.Printf("[%d, %s, %s] -->", temp.no, temp.name,
			temp.nickName)
		// 判断链表是否到了最后
		temp = temp.pre
		if temp.pre == nil {
			break
		}
	}
	fmt.Println()
}

// 判断当前 id 是否存在，存在即可进行下面的操作
// head *HeroNode 头结点, id int 编号信息
func IsExist(head *HeroNode, id int) (temp *HeroNode, b bool) {
	temp = head
	flag := false // 假定找不到
	// 找到要删除结点的no ，和temp 的下一个结点的 no比较
	for {
		if temp.next == nil { 
			// 说明到了链表的最后
			break
		} else if temp.next.no == id {
			// 说明找到了这个 no
			flag = true
			break
		}
		temp = temp.next
	}
	return temp, flag
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
		no : 2,
		name : "卢俊义",
		nickName : "玉麒麟",
	}
	hero3 := &HeroNode{
		no : 3,
		name : "林冲",
		nickName : "豹子头",
	}
	// 3. 加入
	// InsertHeroNode(head,hero3)
	// InsertHeroNode(head,hero1)
	// InsertHeroNode(head,hero2)
	InsertHeroNode2(head,hero3)
	InsertHeroNode2(head,hero1)
	InsertHeroNode2(head,hero2)

	var key string 
	var no int
	var name string
	var nickName string
	for {
		fmt.Println("1. 输入 insert 插入一条英雄信息")
		fmt.Println("2. 输入 delete 删除一条英雄信息")
		fmt.Println("2. 输入 update 修改一条英雄信息")
		fmt.Println("3. 输入 list 查询所有英雄信息")
		fmt.Println()

		fmt.Println("请输入对应的关键字")
		fmt.Scanln(&key)
		switch key {
		case "insert":
			fmt.Println("请输入英雄编号")
			fmt.Scanln(&no)
			fmt.Println("请输入英雄名字")
			fmt.Scanln(&name)
			fmt.Println("请输入英雄昵称")
			fmt.Scanln(&nickName)
			hero := &HeroNode{
				no : no,
				name : name,
				nickName : nickName,
			}
			InsertHeroNode2(head,hero)
		case "delete":
			fmt.Println("请输入英雄编号")
			fmt.Scanln(&no)
			temp, flag := IsExist(head, no)
			if flag {
				// 删除
				DelHeroNode(temp, no)
			} else {
				fmt.Printf("请您确认您输入的id(%d) 是否正确\n", no)
			}
		case "update":
			fmt.Println("请输入英雄编号")
			fmt.Scanln(&no)
			temp, flag := IsExist(head, no)
			if flag {
				fmt.Println("请输入英雄名字")
				fmt.Scanln(&name)
				fmt.Println("请输入英雄昵称")
				fmt.Scanln(&nickName)
				// 修改
				hero := &HeroNode{
					no : no,
					name : name,
					nickName : nickName,
				}
				UpdateHeroNode(temp, hero)
			} else {
				fmt.Printf("请您确认您输入的id(%d) 是否正确\n", no)
			}

		case "list":
			// 显示
			ListHeroNode(head)
			fmt.Println("逆序打印")
			ListHeroNode2(head)
		}
	}
}