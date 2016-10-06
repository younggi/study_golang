package qsort

import "sync"

// QSort sorts input int array
func QSort(a []int) []int {
	if len(a) <= 1 {
		return a
	}
	lesser, pivotlist, greater := []int{}, []int{}, []int{}
	// choose pivot
	pivotIndex := len(a) / 2
	pivot := a[pivotIndex]

	for i := 0; i < len(a); i++ {
		if i == pivotIndex {
			continue
		}
		if a[i] < pivot {
			// pivot과 비교하여 pivot보다 작을 경우 lesser list
			lesser = append(lesser, a[i])
		} else if a[i] >= pivot {
			// pivot과 비교하여 pivot보다 클 경우 greater list
			greater = append(greater, a[i])
		}
	}
	pivotlist = append(pivotlist, pivot)
	return append(QSort(lesser), append(pivotlist, QSort(greater)...)...)
}

// EnhancedQSort sorts int array with minimal memory
func EnhancedQSort(a []int, left int, right int) {
	if right > left {
		pivotIndex := left + (right-left)/2
		newPivotIndex := partition(a, left, right, pivotIndex)
		var wg sync.WaitGroup
		wg.Add(2)
		go func() {
			defer wg.Done()
			EnhancedQSort(a, left, newPivotIndex-1)
		}()
		go func() {
			defer wg.Done()
			EnhancedQSort(a, newPivotIndex+1, right)
		}()
		wg.Wait()
	}
}

// partition returns sorted pivotIndex
func partition(a []int, left int, right int, pivotIndex int) int {
	pivotValue := a[pivotIndex]
	a[pivotIndex], a[right] = a[right], a[pivotIndex]
	storedIndex := left
	for i := left; i < right; i++ {
		if a[i] <= pivotValue {
			a[i], a[storedIndex] = a[storedIndex], a[i]
			storedIndex++
		}
	}
	a[storedIndex], a[right] = a[right], a[storedIndex]
	return storedIndex
}
