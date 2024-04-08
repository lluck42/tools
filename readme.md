# 简单的工具类组件

## list（切片处理）

1. 查询list中是否程存在元素
2. list排重

## 常用tree结构查询

1. 生成tree 下级节点
2. 生成tree 包含自身节点
3. 查找下级id（包括自身id）

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