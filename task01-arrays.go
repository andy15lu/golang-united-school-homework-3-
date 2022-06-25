package homework

func average(input [15]float32) (result float32) {
	//Place your code here
	for i := 0; i < 15; i++ {
		result += input[i]
	}
	result /= 15
	return
}
