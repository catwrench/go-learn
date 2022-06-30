package UnionFind

import (
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

// QuickUnion 加权算法
type QuickUnion struct {
	storage []int // 存储结构
	deep    []int // 各个根节点对应的树 的深度
	count   int   // 连通分量数量
}

func NewQuickUnion(max int) *QuickUnion {
	res := &QuickUnion{
		storage: make([]int, max+1, max+1),
		deep:    make([]int, max+1, max+1),
	}
	for i := 0; i < max+1; i++ {
		res.storage[i] = i
		res.deep[i] = 1
	}
	res.count = max + 1
	return res
}

func (q *QuickUnion) Find(data int) int {
	for data != q.storage[data] {
		data = q.storage[data]
	}
	return data
}

func (q *QuickUnion) Connected(first, second int) bool {
	return q.Find(first) == q.Find(second)
}

func (q *QuickUnion) Union(first, second int) {
	dF := q.Find(first)
	dS := q.Find(second)
	if dF == dS {
		return
	}
	// 小树 连接到 大树
	if q.deep[first] < q.deep[second] {
		q.storage[first] = dS
		q.deep[second] += q.deep[first]
	} else {
		q.storage[second] = dF
		q.deep[first] += q.deep[second]
	}
	q.count--
}

func (q *QuickUnion) Count() int {
	return q.count
}

func QuickUnionInput() [][]int {
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

func TestQuickUnion(t *testing.T) {
	input := QuickUnionInput()

	q := NewQuickUnion(9)
	for _, item := range input {
		q.Union(item[0], item[1])
	}

	// 测试数据为0-9，没有6
	// 测试数据为完全连通的，所以连通分量为1，加上不存在的节点6，所以值为2
	assert.Equal(t, 2, q.Count())
}
