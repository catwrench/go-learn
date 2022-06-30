package UnionFind

import (
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

// QuickFind 算法
type QuickFind struct {
	storage []int // 存储结构
	count   int   // 连通分量数量
}

// max 为最大值，共需要初始化 max+1 的容量
func NewQuickFind(max int) *QuickFind {
	res := &QuickFind{
		storage: make([]int, max+1, max+1),
	}
	for k := range res.storage {
		res.storage[k] = k
	}

	res.count = max + 1
	return res
}

// 获取连通分量数
func (q *QuickFind) Count() int {
	return q.count
}

// 查找节点
func (q *QuickFind) Find(data int) int {
	return q.storage[data]
}

// 判断是否联通
func (q *QuickFind) Connected(first, second int) bool {
	return q.Find(first) == q.Find(second)
}

// QuickFind 算法核心，将值不同的两个节点染色 （值相同的视为在同一个连通分量中）
func (q *QuickFind) Union(first, second int) {
	// 两节点不同时，完成转变
	dF := q.Find(first)
	dS := q.Find(second)
	if dF == dS {
		return
	}

	// 将 first 转变为 second
	for i := 0; i < len(q.storage); i++ {
		if q.storage[i] == dF {
			q.storage[i] = dS
		}
	}
	q.count--
	return
}

func QuickFindInput() [][]int {
	input := "9-0 3-4 5-8 7-2 2-1 5-7 0-3 4-2" // 没有6
	arr := strings.Split(input, " ")
	res := make([][]int, 0)
	for _, item := range arr {
		arr2 := strings.Split(item, "-")
		k, _ := strconv.Atoi(arr2[0])
		v, _ := strconv.Atoi(arr2[1])
		res = append(res, []int{k, v})
	}
	return res
}

func TestQuickFind(t *testing.T) {
	input := QuickFindInput()

	q := NewQuickFind(9)
	for _, item := range input {
		q.Union(item[0], item[1])
	}

	// 测试数据为0-9，没有6
	// 测试数据为完全连通的，所以连通分量为1，加上不存在的节点6，所以值为2
	assert.Equal(t, 2, q.Count())
}
