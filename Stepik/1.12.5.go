/*
Дана последовательность, состоящая из целых чисел. Напишите программу, которая подсчитывает количество положительных чисел
среди элементов последовательности.

Входные данные

Сначала задано число N — количество элементов в последовательности (1≤N≤100).
Далее через пробел записаны N чисел — элементы последовательности. Последовательность состоит из целых чисел.

Выходные данные

Необходимо вывести единственное число - количество положительных элементов в последовательности.

Sample Input:
5
1 2 3 -1 -4

Sample Output:
3
*/

package main

import "fmt"

func main() {
	var dataLen int
	fmt.Scan(&dataLen)

	data := make([]int, dataLen)
	answer := 0
	for _, value := range data {
		fmt.Scan(&value)

		if value > 0 {
			answer++
		}
	}

	fmt.Print(answer)
}
