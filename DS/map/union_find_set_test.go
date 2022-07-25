package _map

//测试并查集

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	u := CreateUnionFindSet(10)
	u.Union(1, 9)
	u.Union(0, 9)
	u.Union(2, 1)
	u.Union(3, 2)
	fmt.Println(u.Find(1) == u.Find(3))
	fmt.Println(u.Find(2) == u.Find(0))

}
