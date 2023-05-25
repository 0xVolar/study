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
}

type Bucket struct {
	ids []*Node
}

func initNode(nodeId string) {

}

func (s *Node) FindNode(nodeID string, array []*Node) []*Node {
	if s.nodeID == nodeID {
		return nil
	}

	//执行insert操作
	s.InsertNode(nodeID)

	//寻找到对应的桶
	result := findBucket(s.nodeID, nodeID)
	var bucket *Bucket
	if result >= (len(s.buckets) - 1) {
		result = len(s.buckets) - 1
		bucket = s.buckets[len(s.buckets)-1]
	} else {
		bucket = s.buckets[result]
	}

	//判断桶中是否存在该节点
	for _, v := range bucket.ids {
		if v.nodeID == nodeID {
			return nil
		}
	}

	//不存在就进行递归,桶中选取
	for {
		//对桶中进随机挑选

		return nil
	}

	return nil
}

/*
1.先异或生成距离
2.找到对应的桶，在对应的K桶中找到距离最近的n（自定义）个节点
3.返回对应的节点地址
*/
func (s *Node) getNodeAdd(id string) (*Node, *Node) {
	//计算距离获取是第几个桶
	result := findBucket(s.nodeID, id)
	var a *Bucket
	if result >= len(s.buckets) {
		a = s.buckets[len(s.buckets)-1]
	} else {
		a = s.buckets[result]
	}

	//获取桶中的两条节点信息
	//信息足够直接返回，信息不够从附近的桶中进行随机选取2个节点信息进行返回
	//获取桶中最近的两个节点并返回
	index1, index2 := GetRandom2()

	return a.ids[index1], a.ids[index2]
}

// 生成两个随机数，0~2之间
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
	if s.nodeID == nodeId {
		return false
	}
	new_node := Node{nodeID: nodeId}
	//判断是否为第一次加入节点,是的话就进行一个初始化功能
	if len(s.buckets) == 0 {
		bucket := Bucket{}
		bucket.ids = append(bucket.ids, &new_node)
		s.buckets = append(s.buckets, &bucket)
		return true
	}
	result := findBucket(s.nodeID, nodeId)
	if result < 0 {
		return false
	}

	var index int
	if result >= (len(s.buckets) - 1) {
		index = len(s.buckets) - 1
		insertIntoClose(index, &new_node, s)
	} else {
		index = result
		isnertIntoFar(index, &new_node, s)
	}
	return true
}

func insertIntoClose(index int, new_node *Node, target_node *Node) {
	bucket := target_node.buckets[index]
	//判断桶中是否已经存在要加入的节点
	for _, v := range bucket.ids {
		if v.nodeID == new_node.nodeID {
			return
		}
	}
	//判断桶是否已满，没满的话加入桶中，满的话进行扩充
	if len(bucket.ids) < 3 {
		bucket.ids = append(bucket.ids, new_node)
	} else {
		//如果桶的数量大于160个的话就不会进行分裂
		if len(target_node.buckets) >= 160 {
			return
		}
		bucket_far := Bucket{}
		bucket_near := Bucket{}

		//将所有节点放到一起
		bucket.ids = append(bucket.ids, new_node)
		//将每个节点的距离计算出来并加入数组之中
		var distance []int
		for _, v := range bucket.ids {
			num1, _ := strconv.ParseInt(v.nodeID, 2, 0)
			num2, _ := strconv.ParseInt(target_node.nodeID, 2, 0)
			distance = append(distance, int(num1^num2))
		}
		//对距离数组进行筛选，选出最近的1个节点加入最远桶
		temp := distance[0]
		index_max := 0
		for i := 1; i < len(distance); i++ {
			if distance[i] > temp {
				temp = distance[i]
				index_max = i
			}
		}

		//将最远的节点加入最远桶，最近的加入最近桶
		bucket_far.ids = append(bucket_far.ids, bucket.ids[index_max])
		for i, _ := range distance {
			if i == index_max {
				continue
			}
			bucket_near.ids = append(bucket_near.ids, bucket.ids[i])
		}

		//将bucket进行更新
		target_node.buckets[index] = &bucket_far
		target_node.buckets = append(target_node.buckets, &bucket_near)
	}
}

func isnertIntoFar(index int, new_node *Node, target_node *Node) {
	bucket := target_node.buckets[index]
	//查看桶中是否已经存在新的节点
	for _, v := range bucket.ids {
		if v.nodeID == new_node.nodeID {
			return
		}
	}
	//小于三个就更新，满了就不管（简化），事实上要进行心跳监测
	if len(bucket.ids) < 3 {
		bucket.ids = append(bucket.ids, new_node)
	}
}

func findBucket(selfId, targetId string) int {
	num, err := strconv.ParseInt(selfId, 2, 0)
	num1, err1 := strconv.ParseInt(targetId, 2, 0)
	if err != nil {
		fmt.Println("Error:", err)
		return -1
	}
	if err1 != nil {
		fmt.Println("Error:", err)
		return -1
	}
	result := strconv.FormatUint(uint64(num)^uint64(num1), 2)
	return 160 - len([]byte(result))
}

// 打印桶中的id
func (s *Bucket) printBucketContents() {
	for _, v := range s.ids {
		fmt.Printf("nodeID = %s \n", v.nodeID)
	}
}

func main() {
	//测试insert方法
	node := Node{nodeID: "000000000001"}

	node.InsertNode("000000000111")
	node.InsertNode("100000000110")
	node.InsertNode("000000000100")
	node.InsertNode("000000001111")
	node.InsertNode("111111111110")
	node.InsertNode("000000001110")
	// a := node.FindNode("000000000111", nil)

	for i, v := range node.buckets {
		fmt.Printf("buckets num is = %d \n", i)
		v.printBucketContents()
		fmt.Println("--------------------------")
	}

	// num1, _ := strconv.ParseInt("000000000001", 2, 0)
	// num2, _ := strconv.ParseInt("000000001111", 2, 0)
	// result := strconv.FormatUint(uint64(num1)^uint64(num2), 2)
	// fmt.Println("result = ", result)
	// // fmt.Print("count = ", count)
}
