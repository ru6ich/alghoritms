package medium

import "math/rand"

type RandomizedSet struct {
	rndMap   map[int]int //Для хранения значения и его индекса в rndSlice
	rndSlice []int       //Для рандомного доступа
}

func Constructor() RandomizedSet {
	return RandomizedSet{rndMap: make(map[int]int), rndSlice: []int{}}
}

func (this *RandomizedSet) Insert(val int) bool {
	if _, exists := this.rndMap[val]; exists {
		return false
	}
	this.rndMap[val] = len(this.rndSlice)
	this.rndSlice = append(this.rndSlice, val)
	return true
}

func (this *RandomizedSet) Remove(val int) bool {
	index, exists := this.rndMap[val]
	if !exists {
		return false
	}

	//Перемещаем последний элемент на место удалямого
	lastElement := this.rndSlice[len(this.rndSlice)-1]
	this.rndSlice[index] = lastElement
	this.rndMap[lastElement] = index

	//Удаляем последний элемент из слайса
	this.rndSlice = this.rndSlice[:len(this.rndSlice)-1]

	//Удаляем значение из мапы
	delete(this.rndMap, val)

	return true
}

func (this *RandomizedSet) GetRandom() int {
	return this.rndSlice[rand.Intn(len(this.rndSlice))]
}

/**
 * Your RandomizedSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */
