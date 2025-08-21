package sorting

func merge(arr []int, left int, mid int, right int) {
	leftArraySize := mid - left + 1
	rightArraySize := right - mid

	leftArrayTemp := make([]int, leftArraySize)
	rightArrayTemp := make([]int, rightArraySize)

	copy(leftArrayTemp, arr[left:mid+1])
	copy(rightArrayTemp, arr[mid+1:right+1])

	for i := range leftArrayTemp {
		leftArrayTemp[i] = arr[left+i]
	}

	for i := range rightArraySize {
		rightArrayTemp[i] = arr[mid+1+i]
	}

	currentLeftPosition := 0
	currentRightPosition := 0
	currentArrPosition := left

	for currentLeftPosition < leftArraySize && currentRightPosition < rightArraySize {
		currentLeftValue := leftArrayTemp[currentLeftPosition]
		currentRightValue := rightArrayTemp[currentRightPosition]

		if currentLeftValue < currentRightValue {
			arr[currentArrPosition] = currentLeftValue
			currentLeftPosition++
		} else {
			arr[currentArrPosition] = currentRightValue
			currentRightPosition++
		}

		currentArrPosition++
	}

	for currentLeftPosition < leftArraySize {
		arr[currentArrPosition] = leftArrayTemp[currentLeftPosition]
		currentLeftPosition++
		currentArrPosition++
	}

	for currentRightPosition < rightArraySize {
		arr[currentArrPosition] = rightArrayTemp[currentRightPosition]
		currentRightPosition++
		currentArrPosition++
	}
}

func mergeRange(arr []int, left int, right int) {
	if left >= right {
		return
	}

	mid := left + (right-left)/2

	mergeRange(arr, left, mid)
	mergeRange(arr, mid+1, right)

	merge(arr, left, mid, right)
}

func MergeSort(arr []int) {
	if len(arr) <= 1 {
		return
	}
	mergeRange(arr, 0, len(arr)-1)
}
