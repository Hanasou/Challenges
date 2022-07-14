package tree

type BinHeap struct {
	list []int
	size int
}

func initBinHeap() *BinHeap {
	return &BinHeap{
		list: []int{0},
		size: 0,
	}
}

func (this *BinHeap) getLeft(p int) int {
	return this.list[2*p]
}

func (this *BinHeap) getRight(p int) int {
	return this.list[2*p+1]
}

func getParentIndex(p int) int {
	return p / 2
}

func (this *BinHeap) getParent(p int) int {
	return this.list[p/2]
}

func (this *BinHeap) swap(i int, j int) {
	temp := this.list[i]
	this.list[i] = this.list[j]
	this.list[j] = temp
}

func (this *BinHeap) insert(n int) {
	index := len(this.list)
	this.list = append(this.list, n)
	parent := this.getParent(index)
	for n > parent {
		p := getParentIndex(index)
		if p == 0 {
			break
		}
		this.swap(index, p)
		index = p
		parent = this.getParent(index)
	}
	this.size++
}
