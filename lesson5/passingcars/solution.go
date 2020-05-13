package passingcars

const (
	maxPassingCarsCount = 1000000000
	tooManyPassingCars  = -1
)

// Solution returns the number of pairs of passing cars.
// If the number of pairs of passing cars exceeds 1,000,000,000 then the function returns -1.
// Length of the array a is an integer in the range [1..100,000].
// Elements in the array a represent consecutive cars on a road:
// - if a[i] == 0 then the car is moving east
// - if a[i] == 1 then the car is moving west
func Solution(a []int) int {
	passingCars := 0
	carsMovingWest := prefixSums(a)
	totalCarsMovingWest := carsMovingWest[len(a)]
	for i, item := range a {
		if item == 0 {
			passingCars += totalCarsMovingWest - carsMovingWest[i]
		}
	}
	if passingCars > maxPassingCarsCount {
		return tooManyPassingCars
	}
	return passingCars
}

func prefixSums(a []int) []int {
	sum := make([]int, len(a)+1)
	for i, item := range a {
		sum[i+1] = sum[i] + item
	}
	return sum
}
