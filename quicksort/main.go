package main

import "fmt"

func partition(low int, high int, numbers []int) {
  current_low := low

  for i := low; i < high; i++ {
    if numbers[i] < numbers[high] {
      if i != current_low {
        numbers[i], numbers[current_low] = numbers[current_low], numbers[i]
      }

      current_low++
    }
  }

  numbers[current_low], numbers[high] = numbers[high], numbers[current_low]

  if current_low - low > 1 { partition(low, current_low - 1, numbers) }
  if high - current_low > 1 { partition(current_low, high, numbers) }
}

func quicksort(numbers []int) []int {
  partition(0, len(numbers) - 1, numbers)

  return numbers
}

func main() {
  var inputs = []int {2, 5, 3, 4, 10, 1, 7, 9, 8, 6, 15, 13, 14, 11, 12}

  output := quicksort(inputs)
  for i := range output {
    fmt.Printf("%d ", output[i])
  }
}
