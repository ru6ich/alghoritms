package medium

import "math/rand"

type RandomizedSet struct {
	rndMap   map[int]int //Для хранения значения и его индекса в rndSlice
	rndSlice []int       //Для рандомного доступа
}

func Constructor() RandomizedSet {
	return RandomizedSet{rndMap: make(map[int]int), rndSlice: []int{}}
}

func (set *RandomizedSet) Insert(val int) bool {
	if _, exists := set.rndMap[val]; exists {
		return false
	}
	set.rndMap[val] = len(set.rndSlice)
	set.rndSlice = append(set.rndSlice, val)
	return true
}

func (set *RandomizedSet) Remove(val int) bool {
	index, exists := set.rndMap[val]
	if !exists {
		return false
	}

	//Перемещаем последний элемент на место удалямого
	lastElement := set.rndSlice[len(set.rndSlice)-1]
	set.rndSlice[index] = lastElement
	set.rndMap[lastElement] = index

	//Удаляем последний элемент из слайса
	set.rndSlice = set.rndSlice[:len(set.rndSlice)-1]

	//Удаляем значение из мапы
	delete(set.rndMap, val)

	return true
}

func (set *RandomizedSet) GetRandom() int {
	return set.rndSlice[rand.Intn(len(set.rndSlice))]
}

/**
 * Your RandomizedSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */
