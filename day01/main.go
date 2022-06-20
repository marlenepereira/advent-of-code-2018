package main

import (
  "bufio"
  "fmt"
  "os"
  "strconv"
)

func main() {
  frequencies, err := readInput()
  if err != nil {
    fmt.Println(err)
  }

  // Part one.
  expectsOne := 529
  resultOne := partOne(frequencies)
  printResult(expectsOne, resultOne)

  // Part two.
  expectsTwo := 464
  resultTwo := partTwo(frequencies)
  printResult(expectsTwo, resultTwo)
}

// printResult prints the result in a specific format.
func printResult(e, r int) {
  if e == r {
    fmt.Printf("Correct! %v == %v\n", e, r)
  } else {
    fmt.Printf("Wrong! %v != %v\n", e, r)
  }
}

// readInput reads input from stdin and returns it.
func readInput() ([]int, error) {
  scanner := bufio.NewScanner(os.Stdin)

  var frequencies []int
  for scanner.Scan() {
    num, err := strconv.Atoi(scanner.Text())
    if err != nil {
      return []int{}, err
    }
    frequencies = append(frequencies, num)
  }

  if scanner.Err() != nil {
    return []int{}, scanner.Err()
  }

  return frequencies, nil
}

// partOne calculates the result of adding all the frequency readings.
func partOne(frequencies []int) (result int) {
  for _, freq := range frequencies {
    result += freq
  }

  return
}

// partTwo finds the first resulting frequency that has been seen (calculated) twice.
func partTwo(frequencies []int) (result int) {
  calculated := map[int]int{}
  found := false
  for !found {
    for _, freq := range frequencies {
      result = result + freq
      _, exists := calculated[result]
      if exists {
        found = true
        break
      }
      calculated[result] = result
    }
  }

  return
}
