package main

import "fmt"

// Функция GroupTemperatures принимает последовательность температурных колебаний и возвращает карту сгруппированных температур с шагом в 10 градусов.
func GroupTemperatures(temperatures []float64) map[int][]float64 {
	groupedTemperatures := make(map[int][]float64)

	for _, value := range temperatures {
		group := (int(value) / 10) * 10
		groupedTemperatures[group] = append(groupedTemperatures[group], value)
	}

	return groupedTemperatures
}

func main() {
	var temperatures = []float64{22.3, 33.7, -12.9, 13.8, -22.1, -23.5}
	fmt.Println(GroupTemperatures(temperatures))
}
