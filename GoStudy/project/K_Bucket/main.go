package main

import (
	"fmt"
	"strconv"
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
	num, err := strconv.Atoi(id)
	num1, err1 := strconv.Atoi(s.nodeID)
	if err != nil {
		fmt.Println("Error:", err)
		return nil, nil
	}
	if err1 != nil {
		fmt.Println("Error:", err)
		return nil, nil
	}

	//计算距离获取是第几个桶
	result := 160 - (num ^ num1)
	var a *Bucket
	if result >= len(s.buckets) {
		a = s.buckets[len(s.buckets)-1]
	} else {
		a = s.buckets[result]
	}

	//对桶中的节点进行遍历计算出距离，
	var distance []int
	for _, v := range a.ids {
		id_int, _ := strconv.Atoi(v.nodeID)
		distance = append(distance, id_int)
	}
	//获取桶中最近的两个节点并返回
	index1, index2 := find2minIndex(distance)

	return a.ids[index1], a.ids[index2]
}

// 返回数组中最小的两个值的下标,两个下标必须不同-
func find2minIndex(arr []int) (int, int) {
	index1, index2 := 0, 1
	min1, min2 := arr[0], arr[1]
	for i := 1; i < len(arr); i++ {
		if arr[i] < min1 {
			min1 = arr[i]
			index1 = i
		}
	}
	for i := 0; i < len(arr); i++ {
		if i == index1 {
			continue
		}
		if arr[i] < min2 {
			min2 = arr[i]
			index2 = i
		}
	}

	return index1, index2
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

	return true
}

/**
1.新节点的加入
	-
*/

func main() {
	arr := []int{0, 5, 20, 8, 15, 3, 25, 12, 11, 3, 3}
	index1, idnex2 := find2minIndex(arr)
	fmt.Print("idnex1 = ", index1, "\n")
	fmt.Print("idnex2 = ", idnex2, "\n")
}
