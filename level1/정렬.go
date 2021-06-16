package main

import (
	"fmt"
	"sort"
)

func main() {
	arr := [] int{1, 5, 2, 6, 3, 7, 4}
	commands := [][] int {{2, 5, 3}, {4, 4, 1}, {1, 7, 3}}

	solution(arr,commands)
}

func solution(array []int, commands [][]int) []int {
	var answer [] int
	for _, command := range commands{
		start := command[0]-1
		end := command[1]
		index := command[2]
		new_arr := append([]int{}, array[start:end] ...)
		sort.Ints(new_arr)
		answer = append(answer, new_arr[index-1])

		}
	fmt.Println(answer)
	return answer
}