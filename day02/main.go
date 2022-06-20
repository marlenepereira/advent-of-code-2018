package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	ids, err := readInput()
	if err != nil {
		fmt.Println(err)
	}

	// test part one
	expectOne := 6696
	resultOne := partOne(ids)
	fmt.Printf("%v == %v: %v\n", expectOne, resultOne, expectOne == resultOne)

	// test part two
	expectTwo := "bvnfawcnyoeyudzrpgslimtkj"
	resultTwo := partTwo(ids)
	fmt.Printf("%v == %v: %v\n", expectTwo, resultTwo, expectTwo == resultTwo)
}

// readInput reads input from stdin and returns a slice of ids.
func readInput() ([]string, error) {
	var ids []string
	var lines []byte
	for {
		ch := make([]byte, 1)
		if _, err := os.Stdin.Read(ch); err == io.EOF {
			return ids, nil
		}

		if ch[0] == byte('\n') {
			ids = append(ids, string(lines))
			lines = []byte{}
			continue
		}
		lines = append(lines, ch...)
	}
}

// partOne counts the boxes that have an ID containing exactly two or three of any letter.
// It returns a chechsum corresponding to the multiplication of those two counts.
func partOne(ids []string) int {
	twos, threes := 0, 0
	for _, id := range ids {
		tw, th := containsRepeatedLetters(id)
		if tw {
			twos++
		}
		if th {
			threes++
		}
	}

	return twos * threes
}

// ContainsRepeatedLetters returns true if a given id contains two or more of any letter.
func containsRepeatedLetters(id string) (twos, threes bool) {
	m := make(map[string]int)

	for _, v := range id {
		m[string(v)]++
	}

	for _, value := range m {
		if value == 2 {
			twos = true
		}

		if value == 3 {
			threes = true
		}
	}

	return
}

// partTwo computes a string from the common letters between the two ids when they differ by one character at the same position.
func partTwo(ids []string) string {
	var result string
	for i := range ids {
		for j := range ids {
			if i == j {
				continue
			}
			result = compare(ids[i], ids[j])
			if len(result) == len(ids[i])-1 {
				return result
			}
		}
	}

	return result
}

// compare computes a string from the common characters between two ids.
func compare(one, two string) string {
	var chars string
	for i := 0; i < len(one); i++ {
		if one[i] == two[i] {
			chars += string(one[i])
		}
	}
	return chars
}
