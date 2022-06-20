package main

import (
  "bufio"
  "fmt"
  "os"
  "strconv"
  "strings"
)

// Claim represents a claim made on an area of fabric.
type Claim struct {
  ID     int
  X      int
  Y      int
  Width  int
  Height int
}

// Coordinate describes a particular point in the fabric.
type Coordinate struct {
  X int
  Y int
}

// Iterate loops over each x and y points that makes up a claim.
func (c Claim) Iterate(record func(id, i, j int) bool) {
  for i := c.X; i < c.X+c.Width; i++ {
    for j := c.Y; j < c.Y+c.Height; j++ {
      next := record(c.ID, i, j)
      if !next {
        return
      }
    }
  }
}

func main() {
  claims, err := readInput()
  if err != nil {
    fmt.Println(err)
  }

  // Initialise a book to keep track of claimed areas.
  claimBook := map[Coordinate]int{}

  // Register all claims in the claim book.
  for _, claim := range claims {
    claim.Iterate(func(id, i, j int) bool {
      xy := Coordinate{X: i, Y: j}
      claimBook[xy]++
      return true
    })
  }

  // Part one - overlapping claims area.
  overlap := getOverlaps(claimBook)
  expectsOverlap := 118223
  printResult(expectsOverlap, overlap)

  // Part two - get the ID of the unique claim.
  expectsID := 412
  uniqueID := getUniqueClaim(claims, claimBook)
  printResult(expectsID, uniqueID)
}

// printResult prints the result in a specific format.
func printResult(e, r int) {
  if e == r {
    fmt.Printf("Correct! %v == %v\n", e, r)
  } else {
    fmt.Printf("Wrong! %v != %v\n", e, r)
  }
}

// split a string at the specified characters.
func split(r rune) bool {
  return r == '#' || r == '@' || r == ':' || r == 'x' || r == ','
}

// readInput reads the input from stdin and returns a slice of Claim.
func readInput() ([]Claim, error) {
  scanner := bufio.NewScanner(os.Stdin)
  var claims []Claim
  for scanner.Scan() {
    var numbers []int
    line := scanner.Text()

    line = strings.ReplaceAll(line, " ", "")
    for _, field := range strings.FieldsFunc(line, split) {
      num, err := strconv.Atoi(field)
      if err != nil {
        return nil, err
      }
      numbers = append(numbers, num)
    }

    claims = append(claims, Claim{
      ID:     numbers[0],
      X:      numbers[1],
      Y:      numbers[2],
      Width:  numbers[3],
      Height: numbers[4],
    })
  }

  if scanner.Err() != nil {
    return nil, scanner.Err()
  }

  return claims, nil
}

// getOverlaps calculates the number of square inches that overlap between claims.
func getOverlaps(claimBook map[Coordinate]int) int {
  var overlaps int
  for _, entry := range claimBook {
    if entry > 1 {
      overlaps++
    }
  }
  return overlaps
}

// getUniqueClaim finds the unique claim that does not overlap with other claims.
func getUniqueClaim(claims []Claim, claimBook map[Coordinate]int) int {
  var uniqueID int

  for _, claim := range claims {
    if uniqueID != 0 {
      break
    }

    claim.Iterate(func(id, i, j int) bool {
      xy := Coordinate{X: i, Y: j}
      value, ok := claimBook[xy]
      if !ok {
        return ok
      }

      if value == 1 {
        uniqueID = id
      } else {
        uniqueID = 0
      }

      return value == 1
    })
  }

  return uniqueID
}
