package main

import (
	"fmt"
	"math"
	model "soduku/set"
)

var (
	s      *model.Set
	result []int
)

func GetBlock(index int, SoDoKu []int) *model.Set {
	BlockVal := model.New()
	Column := (index + 9) % 9
	Row := int(math.Floor(float64((index+9)/9))) - 1
	var row, col int
	switch Row {
	case 0, 1, 2:
		row = 0
		break
	case 3, 4, 5:
		row = 3
		break
	case 6, 7, 8:
		row = 6
		break

	}
	switch Column {
	case 0, 1, 2:
		col = 0
		break
	case 3, 4, 5:
		col = 3
		break
	case 6, 7, 8:
		col = 6
		break
	}

	for i := row; i < row+3; i++ {
		for ii := col; ii < col+3; ii++ {
			if SoDoKu[i*9+ii] != 0 {
				BlockVal.Add(SoDoKu[i*9+ii])
			}
		}
	}
	return s.Minus(BlockVal)
}

func GetRowVal(index int, SoDoKu []int) *model.Set {
	RowVal := model.New()
	Row := int(math.Floor(float64((index+9)/9))*9 - 9)
	for i := 0; i < 9; i++ {
		if SoDoKu[Row+i] != 0 {
			RowVal.Add(SoDoKu[Row+i])
		}
	}
	return s.Minus(RowVal)
}

func GetColumnVal(index int, SoDoKu []int) *model.Set {
	ColumnVal := model.New()
	Column := (index + 9) % 9
	for i := 0; i < 9; i++ {
		if SoDoKu[9*i+Column] != 0 {
			ColumnVal.Add(SoDoKu[9*i+Column])
		}
	}
	return s.Minus(ColumnVal)
}

func GetSpareVal(index int, SoDoKu []int) []int {
	Block := GetBlock(index, SoDoKu)
	RowVal := GetRowVal(index, SoDoKu)
	ColumnVal := GetColumnVal(index, SoDoKu)
	return Block.Intersect(RowVal, ColumnVal).List()
}

func UpSpare(SoDoKu []int) {
	if result != nil {
		return
	}
	for i := 0; i < len(SoDoKu); i++ {
		newsudu := make([]int, len(SoDoKu))
		copy(newsudu, SoDoKu)
		if SoDoKu[i] == 0 {
			SpareVal := GetSpareVal(i, SoDoKu)
			if len(SpareVal) == 0 {
				return
			}
			if i == 80 {
				newsudu[80] = SpareVal[0]
				result = make([]int, len(SoDoKu))
				copy(result, newsudu)
				return
			}
			for _, Val := range SpareVal {
				newsudu[i] = Val
				UpSpare(newsudu)
			}
			return
		} else {
			if i == 80 {
				result = make([]int, len(SoDoKu))
				copy(result, newsudu)
				return
			}
		}
	}
}

func main() {
	s = model.New(1, 2, 3, 4, 5, 6, 7, 8, 9)
	Source := []int{
		5, 2, 0, 0, 4, 0, 0, 0, 0,
		7, 3, 0, 0, 0, 5, 0, 9, 0,
		1, 0, 0, 6, 0, 0, 5, 2, 3,
		0, 4, 2, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 9, 0, 0, 0, 6, 0,
		0, 7, 0, 0, 0, 0, 8, 0, 0,
		2, 0, 7, 0, 0, 0, 0, 0, 6,
		0, 0, 0, 0, 8, 7, 4, 0, 0,
		0, 0, 0, 0, 5, 0, 7, 0, 0,
	}

	UpSpare(Source)

	for i, v := range result {
		if i >= 9 && i%9 == 0 {
			fmt.Print("\n", v, ",")
		} else {
			fmt.Print(v, ",")
		}
	}

}
