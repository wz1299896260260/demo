package main

import (
	"fmt"

	"web/router"
)

func main() {
	router.Run()
}

type Node struct {
	data int
	next *Node
}

func qq(p *Node){
	for p!=nil{
		fmt.Println(*p)//打印p内容
		p=p.next  //遍历，
	}
}

func main1(){
	var head=new(Node)
	head.data=0//头结点
	var tail *Node
	tail=head//头结点给tail使用
	for i:=1;i<10;i++{
		var node=Node{data:i}
		//头插法
		//node.next=tail  //node的下个结点等于头结点1  0
		//tail=&node//此时头结点是1

		//尾插法
		(*tail).next=&node//*tail是head结点
		tail=&node
	}
	qq(head)//*tail改变的head值
}
func main2(){

}
func main3(){

}





