package Sequence

import (
	//"errors"
	"fmt"
	//"log"
)


// Выводит массив размерностей строк массива и возвращает true, если размерности равны.
func MatrixId(massiv [][]interface{ }) (razmeristrok []int, onomatritsa bool) {
	razmeristrok = make([]int, 0, len(massiv))
	for _, v := range massiv {
		razmeristrok = append(razmeristrok, len(v))
	}
	a := 0
	for _, v := range razmeristrok {
		if v == razmeristrok[0] {
			a = a + 1
		}
	}
	if a == len(razmeristrok) {
		onomatritsa = true
	} else {
		onomatritsa = false
	}
	return razmeristrok, onomatritsa
}

//Проверка размеров матриц на соответствие для сложения.

/*func CheckMatrixAdd(slice1,slice2 Matrix)(error, *Matrix){
	  a:=&Matrix{
	  	2,
	  	3,
	  	FLT,
	  	nil,
	  }

	  /*
	  processing

	  return errors.New("Some shit happens"), nil

	  return nil, a
}

// Построчная расчепятка массива массивов.
type Matrix struct {
	strlen  int
	rowlen  int
	RecType string
	body    [][]interface{}
}


/*func myShinyFunc (){
	variable, err := CheckMatrixAdd()
	if err!=nil {
		log.Println(err)
		
	}
}

*/


func PrMatrix(matrix [][]interface{})(){
	for _,v:=range matrix{
		fmt.Println(v)
	}
}
