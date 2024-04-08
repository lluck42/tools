# 简单的工具类组件

## list（切片处理）

1. 查询list中是否程存在元素
2. list排重

### USAGE InList UniqueList

```
var a uint = 8
var arr = []uint{2, 2, 3, 8}
var rtn = InList(arr, a)
fmt.Println(rtn)
var rtn2 = UniqueList(arr)
fmt.Println(rtn2)

```


## 常用tree结构查询

1. 生成tree 下级节点
2. 生成tree 包含自身节点
3. 查找下级id（包括自身id）

### USAGE TreeSelfGet 生成tree 包含自身节点

```
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

```

### USAGE TreeSelfIds 查找下级id（包括自身id）

```
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

```