package insertdeletegetrandomo1

import "math/rand"

type RandomizedSet struct {
	Map  map[int]int
	List []int
}

func Constructor() RandomizedSet {
	return RandomizedSet{
		Map: make(map[int]int),
	}
}

func (this *RandomizedSet) Insert(val int) bool {
	_, ok := this.Map[val]
	if ok {
		return false
	}

	this.List = append(this.List, val)
	this.Map[val] = len(this.List) - 1
	return true
}

func (this *RandomizedSet) Remove(val int) bool {
	idx, ok := this.Map[val]
	if ok {
		lastVal := this.List[len(this.List)-1]
		this.List[idx] = lastVal
		this.List = this.List[:len(this.List)-1]
		this.Map[lastVal] = idx
		delete(this.Map, val)
		return true
	}
	return false
}

func (this *RandomizedSet) GetRandom() int {
	return this.List[rand.Intn(len(this.List))]
}
