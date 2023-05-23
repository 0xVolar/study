package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

type Node struct {
	nodeID  string
	buckets []*Bucket
	link    *Bucket
}

type Bucket struct {
	ids            [3]*Node
	lchild, rchild int
}

/*
1.先异或生成距离
2.找到对应的桶，在对应的K桶中找到距离最近的n（自定义）个节点
3.返回对应的节点地址
*/
func (s *Node) FindNode(id string) (*Node, *Node) {
	//计算距离获取是第几个桶
	result := findBucket(s.nodeID, id)
	var a *Bucket
	if result >= len(s.buckets) {
		a = s.buckets[len(s.buckets)-1]
	} else {
		a = s.buckets[result]
	}

	//对桶中的节点进行遍历看是否有目标节点，如果没有的话返回随机的两个节点信息
	for _, v := range a.ids {
		if v.nodeID == id {
			return nil, nil
		}
	}
	//获取桶中最近的两个节点并返回
	index1, index2 := GetRandom2()

	return a.ids[index1], a.ids[index2]
}

// 生成两个随机数
func GetRandom2() (int, int) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	// 随机生成两个不重复的整数
	var first, second int
	done := false
	for !done {
		first, second = r.Intn(3), r.Intn(3)
		if first != second {
			done = true
			break
		}
		// 使用随机数生成器进行洗牌，确保随机数不重复
	}
	// 输出随机数
	fmt.Println(first, second)
	return first, second
}

/**
1.插入节点
计算插入的节点距离应该加到那个桶
- 如果是距离最近的桶（数组最后一个元素）
 - 查看桶满没有，没满就加入
 - 满了的话就分裂
 	- 在数组中加入新的桶，新的桶中永远保存距离最近的n个节点，分裂前满的桶装距离远的节点
- 如果是非最近的桶
 - 桶满就放弃加入
 - 桶未满就加入
*/

func (s *Node) InsertNode(nodeId string) bool {
	result := findBucket(s.nodeID, nodeId)
	if result < 0 {
		return false
	}

	var bucket *Bucket
	if result >= len(s.buckets)-1 {
		bucket = s.buckets[len(s.buckets)-1]
		insertIntoClose(bucket)
	} else {
		bucket = s.buckets[result]
		isnertIntoFar(bucket)
	}

	return true
}

func insertIntoClose(bucket *Bucket) {

}

func isnertIntoFar(bucket *Bucket) {

}

func findBucket(selfId, tragetId string) int {
	num, err := strconv.Atoi(selfId)
	num1, err1 := strconv.Atoi(tragetId)
	if err != nil {
		fmt.Println("Error:", err)
		return -1
	}
	if err1 != nil {
		fmt.Println("Error:", err)
		return -1
	}
	result := 160 - (num ^ num1)
	return result
}

/**
1.新节点的加入
	-
*/

func main() {

}
