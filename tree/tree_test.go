package tree

import (
	"fmt"
	"testing"
)

func TestA(t *testing.T) {
	var listAll []Node
	var one Node
	one.ID = 1
	one.Name = "1"
	one.ParentId = 0

	var two Node
	two.ID = 2
	two.ParentId = 1
	two.Name = "2"

	var three Node
	three.ID = 3
	three.ParentId = 2
	three.Name = "3"

	listAll = append(listAll, one, two, three)

	// 获取选中节点的树
	var loopCount uint

	var id uint = 1

	var tree, _ = TreeSelfGet(id, &listAll, &loopCount)

	fmt.Printf("%+v", tree)
}

func TestB(t *testing.T) {
	var listAll []Node
	var one Node
	one.ID = 1
	one.Name = "1"
	one.ParentId = 0

	var two Node
	two.ID = 2
	two.ParentId = 1
	two.Name = "2"

	var three Node
	three.ID = 3
	three.ParentId = 2
	three.Name = "3"

	listAll = append(listAll, one, two, three)

	// 获取选中节点的树
	var loopCount uint

	var id uint = 1

	var ids, _ = TreeSelfIds(id, &listAll, &loopCount)

	fmt.Printf("%+v", ids)
}
