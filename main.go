package main

import "fmt"

type sudo struct {
	Value int
	SpareVal []int
}

var (
	SoDoKu [9][9]sudo
	Source [9][9]int
)

//初始化数独
func InitSudo()  {
	for i1, i2 := range Source {
		for i, i3 := range i2 {
			SoDoKu[i1][i].Value = i3
		}
	}
	
}

//行 是否存在指定值
func RowIseExist(Row int,val int) bool {
	var exist bool
	for _, i := range SoDoKu[Row] {
		if i.Value == val {
			return false
		}else {
			exist = true
		}
	}
	return exist
}

//列 是否存在指定值
func ColumnIseExist(Column int,val int) bool {
	var exist bool
	for _, i := range SoDoKu {
		if i[Column].Value == val {
			return false
		}else {
			exist = true
		}
	}
	return exist
}

//当前九宫 是否存在指定值 及 可能值数组中是否存在指定值
func BlockIseExist(Row,Column,val int,IsSpare bool) bool {
	exist := true
	var row,col int
	switch Row {
	case 0,1,2: row =0; break
	case 3,4,5: row =3; break
	case 6,7,8: row =6; break

	}
	switch Column {
	case 0,1,2: col = 0; break
	case 3,4,5: col = 3; break
	case 6,7,8: col = 6; break
	}
	for i := row ; i < row + 3 ; i++ {
		for ii := col ; ii < col + 3 ; ii++ {
			if IsSpare {
				for _, i3 := range SoDoKu[i][ii].SpareVal {
					if i3 == val {
						if i == Row && ii == Column {
							break
						}else {
							return false
						}
					}else {
						exist = true
					}
				}
			}else {
				if SoDoKu[i][ii].Value == val {
					return false
				}else {
					exist = true
				}
			}

		}
	}
	return exist
}

//更新 可能值数组
func UpSpare()  {
	for row, sudos := range SoDoKu {
		for Column, sudo := range sudos {
			if sudo.Value == 0 {
				SoDoKu[row][Column].SpareVal = []int{}
				sudo.SpareVal = []int{}
				for i:=1; i<=9; i++ {
					if RowIseExist(row,i) && ColumnIseExist(Column,i) && BlockIseExist(row,Column,i,false){
						sudo.SpareVal = append(sudo.SpareVal, i)
					}
				}
				SoDoKu[row][Column].SpareVal = sudo.SpareVal
			}
		}
	}
}

//解数独
func Solved()  {
	for row, sudos := range SoDoKu {
		for Column, sudo := range sudos {
			if sudo.Value == 0 {
				for _, i := range sudo.SpareVal {
					if BlockIseExist(row,Column,i,true){
							SoDoKu[row][Column].Value = i
							SoDoKu[row][Column].SpareVal = []int{}
							UpSpare()
							return
						}
					}
			}
		}
	}
}

//判断是否完成
func IsSolvedOk() bool {
	for _, sudos := range SoDoKu {
		for _, sudo := range sudos {
			if sudo.Value == 0 {
				return false
			}
		}
	}
	return true
}


func main() {
	Source = [9][9]int{
		{0,8,0,0,0,5,0,1,7},
		{0,0,0,0,7,9,2,0,0},
		{7,9,0,4,0,0,0,0,6},
		{2,0,0,9,4,0,0,3,0},
		{0,3,4,2,0,1,6,7,0},
		{0,5,0,0,3,7,0,0,1},
		{6,0,0,0,0,3,0,4,5},
		{0,0,9,5,1,0,0,0,0},
		{1,7,0,6,0,0,0,9,0},
	}
	InitSudo()
	UpSpare()

	for {
		Solved()
		if IsSolvedOk(){
			break
		}
	}

	for i, i2 := range SoDoKu {
		for i3, i4 := range i2 {
			Source[i][i3] = i4.Value
		}
		fmt.Println(Source[i])
	}

}
