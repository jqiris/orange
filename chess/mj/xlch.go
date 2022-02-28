package mj

import (
	"errors"
	"math"
	"math/rand"

	"github.com/jqiris/orange/constant"
)

//四川麻将血流成河
type XlchMj struct {
}

func NewXlchMj() *XlchMj {
	return &XlchMj{}
}

func (m *XlchMj) String() string {
	return "血流成河麻将"
}

func (m *XlchMj) GetTotalNums() int {
	return 108
}

func (m *XlchMj) GetMjType(id int) int {
	if id >= 0 && id < 9 {
		return 0 //筒
	} else if id >= 9 && id < 18 {
		return 1 //条
	} else if id >= 18 && id < 27 {
		return 2 //万
	}
	return -1
}

func (m *XlchMj) Shuffle(list []int) error {
	if len(list) != m.GetTotalNums() {
		return errors.New(constant.ErrWrongCardNumber)
	}
	index := 0
	//筒 0-8
	for i := 0; i < 9; i++ {
		for c := 0; c < 4; c++ {
			list[index] = i
			index++
		}
	}
	//条 9-17
	for i := 9; i < 18; i++ {
		for c := 0; c < 4; c++ {
			list[index] = i
			index++
		}
	}
	//万 18-26
	for i := 18; i < 27; i++ {
		for c := 0; c < 4; c++ {
			list[index] = i
			index++
		}
	}

	for i := 0; i < m.GetTotalNums(); i++ {
		lastIndex := m.GetTotalNums() - i - 1
		randIndex := float64(rand.Intn(100) * lastIndex / 100)
		index = int(math.Floor(randIndex))
		list[index], list[lastIndex] = list[lastIndex], list[index]
	}
	return nil
}
