//本程序主要思路是使用轮巡桶的算法
//时间关系做了一些简化
//命令行input 没有做详细解析，而是直接赋值。
//总的请求数量没有做成参数，而是直接写死为600.
//go 环境下执行 go run main.go  即可看到结果
//在线go环境地址：
package main

import (
	"fmt"
)

type Node struct {
	id     string `json:"id"`
	weight int    `json:"weight"`
	count  int    `json:"count"`
}

func NewNode(id string, weight int, count int) *Node {
	return &Node{
		id:     id,
		weight: weight,
		count:  count,
	}
}
func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}
func GetGcd(n []int) int {
	g := n[0]
	for i := 1; i < len(n)-1; i++ {
		//oldGcd := g
		g = gcd(g, n[i])
	}
	return g
}
func main() {
	var nodearr [3]*Node
	gysarr := []int{}
	table := []string{} // 作为桶使用
	var totalweght int
	var hitrange int
	nodearr[0] = NewNode("a", 150, 0)
	nodearr[1] = NewNode("b", 100, 0)
	nodearr[2] = NewNode("c", 50, 0)

	for _, k := range nodearr {
		cweight := (*k).weight
		gysarr = append(gysarr, cweight)
	}
	gys := GetGcd(gysarr)
    for _, k := range nodearr {
        (*k).weight = (*k).weight / gys
        totalweght += (*k).weight
        fmt.Println("weight is ",(*k).weight)
        for j := 0; j < (*k).weight; j++ {
            table = append(table, (*k).id) //占位
        }
    }
    
	fmt.Println("table size is ",totalweght)
    fmt.Println("table size is ",table)
	for request := 0; request < 600; request++ {
		hitrange = request % totalweght
		for _, k := range nodearr {
			id := (*k).id
            //fmt.Println(table[hitrange])
            //fmt.Println(id)
			if table[hitrange] == id {
				(*k).count += 1
				//fmt.Println((*k).count)
			}
		}
	}
    for _, k := range nodearr {
        fmt.Println((*k).count)
    }
}
