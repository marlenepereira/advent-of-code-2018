package main

import (
  "fmt"
  "io"
  "os"
  "strings"
)

// Polymer represents a chain of units.
type Polymer []string

// String implements Stringer to print the polymer chain.
func (p Polymer) String() string {
  polymer := ""
  for _, s := range p {
    polymer += s
  }

  return polymer
}

// React removes consecutive same letters with different capitalization.
func (p Polymer) React() []string {
  var reacted []string
  for _, unit := range p {
    prev := ""
    if len(reacted) != 0 {
      prev = reacted[len(reacted)-1]
    }

    if prev != unit && strings.ToLower(prev) == strings.ToLower(unit) {
      // Remove previous if they are the same but different polarity.
      reacted = reacted[:len(reacted)-1]
      continue
    }

    reacted = append(reacted, unit)
  }

  return reacted
}

func main() {
  chain := readInput()

  // Part one.
  expectsOne := 11252
  printResult(expectsOne, len(Polymer(chain).React()))

  // Part two.
  expectsTwo := 6118
  printResult(expectsTwo, partTwo(chain))
}

// printResult formats the expected and actual result messages.
func printResult[T comparable](e, r T) {
  if e == r {
    fmt.Printf("Correct! %v == %v\n", e, r)
  } else {
    fmt.Printf("Wrong! %v != %v\n", e, r)
  }
}

// partTwo computes the shortest polymer after reacting units.
func partTwo(chain []string) int {
  length := 0
  for i := 'a'; i <= 'z'; i++ {
    var reducedChain []string
    for _, unit := range chain {
      if []rune(strings.ToLower(unit))[0] != i {
        reducedChain = append(reducedChain, unit)
      }
    }

    reactedChainLength := len(Polymer(reducedChain).React())

    if reactedChainLength < length || length == 0 {
      length = reactedChainLength
    }
  }

  return length
}

// readInput reads the input from stdin and build a chain of units.
func readInput() []string {
  var chain []string
  for {
    ch := make([]byte, 1)
    if _, err := os.Stdin.Read(ch); err == io.EOF {
      return chain
    }
    if ch[0] == '\n' {
      break
    }

    chain = append(chain, string(ch[0]))
  }

  return chain
}
