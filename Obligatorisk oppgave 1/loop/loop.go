package main

import (

	"math/rand"
	"reflect"

)

func shuffle(slice interface{}) {
	rv := reflect.ValueOf(slice)
	swap := reflect.Swapper(slice)
	length := rv.Len()
	for i := length - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		swap(i, j)
	}
}

func bubbleSort2(list []int) {
	n := len(list)
	swapped := true
	for swapped {
		//swapped = false
		for j := 0; j < n-1; j++ {
			if list[j] > list[j+1] {
				temp := list[j+1]
				list[j+1] = list[j]
				list[j] = temp
				swapped = true
			}
		}
		n--
	}
}

func main() {

	swapped := true

	for swapped {
		for i := 0; i < 500; i++ {
			tilSortering := rand.Perm(2500)
			shuffle(tilSortering)
			bubbleSort2(tilSortering)
		}
	}

}