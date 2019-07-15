package Lecture03

import (
	"testing"
	"log"
)


func Test_rbtree(test *testing.T) {
	rbtree := new(RBTree)

	int64arr := [...]int{1, 2, 3, 4, 5, 6, 7, 8}

	for _, num := range int64arr {
		rbtree.Insert(num)
	}

	// r := rand.New(rand.NewSource(time.Now().UnixNano()))
	// for i := 0; i < 50; i++ {
	// 	rbtree.Insert(r.Int63n(100))
	// }
	log.Print("输出红黑树@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@")
	printTreeInLog(rbtree.root, "(root)")

	log.Print("删除节点@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@")
	rbtree.Delete(int(1))
	rbtree.Delete(int(2))
	rbtree.Delete(int(3))
	printTreeInLog(rbtree.root, "(root)")
}