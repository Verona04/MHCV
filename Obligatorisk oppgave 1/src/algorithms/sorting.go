package algorithms

import "fmt"

//Vår bubblesort
func Bubble_sort_modified(list []int) {
	// find the length of list n
	n := len(list)
	swapped := true
	for swapped {
		swapped = false
		for i := 0; i < n; i++ {
			for j := 0; j < n-1; j++ {
				if list[j] > list[j+1] {
					list[j+1], list[j] = list[j], list[j+1]
					swapped = true
				}
			}
		}
		n--
	}
}


func main() {

	list := []int{1, 6, 2, 4, 9, 0, 5, 3, 7, 8}
	fmt.Println("Unsorted list: ", list)
	Bubble_sort_modified(list)
	fmt.Println("Sorted list: ", list)
}

// Implementering av Bubble_sort algoritmen
func BubbleSort(list []int) {
	// find the length of list n
	n := len(list)
	for i := 0; i < n; i++ {
		for j := 0; j < n-1; j++ {
			if list[j] > list[j+1] {
				temp := list[j+1]
				list[j+1] = list[j]
				list[j] = temp
			}
		}
	}
}

// Implementering av Quicksort algoritmen
func QuickSort(values []int) {
	qsort(values, 0, len(values)-1)
}

func qsort(values []int, l int, r int) {
	if l >= r {
		return
	}

	pivot := values[l]
	i := l + 1

	for j := l; j <= r; j++ {
		if pivot > values[j] {
			values[i], values[j] = values[j], values[i]
			i++
		}
	}

	values[l], values[i-1] = values[i-1], pivot

	qsort(values, l, i-2)
	qsort(values, i, r)
}
