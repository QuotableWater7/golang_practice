package main

import "fmt"

func subsets(numbers []int) [][]int {
  collected := make([]int, 0)
  results := make([][]int, 0)

  subsets_runner(numbers, 0, collected, &results)

  return results
}

func subsets_runner(numbers []int, index int, collected []int, all *[][]int) {
  if index == len(numbers) {
    *all = append(*all, collected)
    return
  }

  item := numbers[index]
  collected_with := append(collected, item)

  subsets_runner(numbers, index + 1, collected_with, all)
  subsets_runner(numbers, index + 1, collected, all)
}

func main() {
  numbers := []int {1, 2, 3, 4}
  output := subsets(numbers)

  for i := range output {
    fmt.Printf("%d\n", output[i])
  }
}
