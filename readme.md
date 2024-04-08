###
简单的工具类组件

###
包含

list（切片处理）

常用tree结构查询

```
// 节点结构（用作数据池查询数据）
type Node struct {
	ID       uint
	ParentId uint
	Name     string
}

// 树结构（用于输出树）
type Tree struct {
	Node
	Children []Tree
}

```