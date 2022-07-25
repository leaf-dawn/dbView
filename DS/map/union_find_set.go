package _map

//并查集,https://www.bilibili.com/video/BV1jv411a7LK?spm_id_from=333.337.search-card.all.click
type UnionFindSet []int

//创建并查集
func CreateUnionFindSet(max int) UnionFindSet {
	answer := make(UnionFindSet, max)
	for i := range answer {
		answer[i] = i
	}
	return answer
}

//查找祖先
func (u UnionFindSet) Find(i int) int {
	if u[i] == i {
		return i
	} else {
		//路径压缩
		u[i] = u.Find(u[i])
		return u[i]
	}
}

//合并
func (u UnionFindSet) Union(i int, j int) {
	fi := u.Find(i)
	fj := u.Find(j)
	u[fi] = fj
}
