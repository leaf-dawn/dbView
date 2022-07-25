package tree

//字典树
type Trie struct {
	next  [26]*Trie
	isEnd bool
}

func Constructor() Trie {
	return Trie{
		isEnd: false,
		next:  [26]*Trie{},
	}
}

//添加记录到前缀树
func (this *Trie) Insert(word string) {
	//记录字典当前位置
	node := this
	for _, v := range word {
		vi := v - 'a'
		//检验下一个是否含有
		if node.next[vi] != nil {
			node = node.next[vi]
		} else {
			t := Constructor()
			node.next[vi] = &t
			node = node.next[vi]
		}
	}
	//最后一个
	node.isEnd = true
}

//搜索前缀树中是否含有某个word
func (this *Trie) Search(word string) bool {
	node := this
	for _, v := range word {
		vi := v - 'a'
		if node.next[vi] != nil {
			node = node.next[vi]
		} else {
			return false
		}
	}
	return node.isEnd
}

//搜索前缀树中的多有记录是否含有该prefix
func (this *Trie) StartsWith(prefix string) bool {
	node := this
	for _, v := range prefix {
		vi := v - 'a'
		if node.next[vi] != nil {
			node = node.next[vi]
		} else {
			return false
		}
	}
	return true
}
