package tree

import (
	"errors"
	"fmt"
)

type Node struct {
	ID       uint
	ParentId uint
	Name     string
}

type Tree struct {
	Node
	Children []Tree
}

// 生成tree 下级节点
func TreeChildrenGet(parentId uint, list *[]Node, loopCount *uint) (tree []Tree, err error) {

	*loopCount += 1

	if *loopCount >= 1000 {
		var errStr = fmt.Sprintf("TreeChildrenGet:超出最大循环次数 loopCount:%d 请检查组织架构是否有循环依赖！", *loopCount)
		return nil, errors.New(errStr)
	}

	for _, v := range *list {
		if v.ParentId == parentId {
			var one Tree
			one.ID = v.ID
			one.ParentId = v.ParentId
			one.Name = v.Name
			var childs, err = TreeChildrenGet(v.ID, list, loopCount)
			if err != nil {
				return nil, err
			}

			one.Children = childs

			tree = append(tree, one)
		}
	}

	return tree, nil
}

// 生成tree 包含自身节点
func TreeSelfGet(id uint, list *[]Node, loopCount *uint) (tree []Tree, err error) {

	// 查找自身节点
	for _, v := range *list {
		if v.ID == id {
			var one Tree
			one.ID = v.ID
			one.ParentId = v.ParentId
			one.Name = v.Name
			var childs, err = TreeChildrenGet(v.ID, list, loopCount)
			if err != nil {
				return nil, err
			}
			one.Children = childs

			tree = append(tree, one)
			break
		}
	}

	return tree, nil
}

// 查找下级id（包括自身id）
func TreeSelfIds(id uint, listAll *[]Node, loopCount *uint) (childrenIds []uint, err error) {

	childrenIds = append(childrenIds, id)

	*loopCount += 1

	if *loopCount >= 1000 {
		var errStr = fmt.Sprintf("MyIds:超出最大循环次数 loopCount:%d 请检查组织架构是否有循环依赖！", *loopCount)
		return nil, errors.New(errStr)
	}

	// 查找子节点
	for _, v := range *listAll {
		if v.ParentId == id {
			// 递归
			var childIds, err = TreeSelfIds(v.ID, listAll, loopCount)
			if err != nil {
				return nil, err
			}

			childrenIds = append(childrenIds, childIds...)
			// childrenIds = append(childrenIds, v.ID)
		}
	}

	return childrenIds, nil
}

func Tree2List(tree []Tree, list *[]Node) (err error) {

	if tree == nil {
		return
	}

	for _, v := range tree {

		*list = append(*list, Node{ID: v.ID, ParentId: v.ParentId, Name: v.Name})

		Tree2List(v.Children, list)
	}

	return nil
}

////////////////////////// 泛型类型树 //////////////////////////

// 泛型树生成函数
func BuildAnyTree[L, T any, I comparable](
	parentId I,
	list []L,
	loopCount *uint,
	getId func(L) I,
	getParentId func(L) I,
	getData func(L, []T) T,
) ([]T, error) {

	*loopCount++

	if *loopCount >= 1000 {
		return nil, fmt.Errorf("BuildTree:超出最大循环次数 loopCount:%d", *loopCount)
	}

	var tree []T
	for _, item := range list {
		if getParentId(item) == parentId {
			children, err := BuildAnyTree(getId(item), list, loopCount, getId, getParentId, getData)
			if err != nil {
				return nil, err
			}

			node := getData(item, children)
			tree = append(tree, node)
		}
	}

	return tree, nil
}
