package productbutthis

import "errors"

// GetProductsOfAllIntsExceptAtIndex returns the product of all the integers in
//		in an array except the current member of the array.
//		Example: in a 4 element array:
//					position 0 would be [1]*[2]*[3]
//					position 1 would be [0]*[2]*[3]
//					position 2 would be [0]*[1]*[3]
//					position 3 would be [0]*[1]*[2]
// O(n) - time
// O(n) - space
func GetProductsOfAllIntsExceptAtIndex(arry []int) ([]int, error) {

	if arry == nil {
		return nil, errors.New("Input array cannot be nil")
	}
	if len(arry) < 2 {
		return nil, errors.New("Need at least two values to multiply")
	}

	result := make([]int, len(arry))
	rolling := 1
	for i := 0; i < len(arry); i++ {
		result[i] = rolling
		rolling *= arry[i]
	}

	rolling = 1
	for i := len(arry) - 1; i >= 0; i-- {
		result[i] *= rolling
		rolling *= arry[i]
	}
	return result, nil
}
