package stats

func test() {
	var data = []float64{1, 2, 3, 4, 4, 5}

	mean, _ := Mean(data)
	fmt.Println(mean) // 3.5
}
