package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
	"unicode"
)

const (
	Awake = iota
	Asleep
)

// Log is the documented state of a guard during a period of time.
type Log struct {
	State  int
	Month  int
	Day    int
	Hour   int
	Minute int
}

// MonthDay returns the log's date in a MM-DD format.
func (l Log) MonthDay() string {
	return fmt.Sprintf("%v-%v", l.Month, l.Day)
}

// Guard represents a guard with corresponding logs.
type Guard struct {
	ID   int
	Logs map[string][]Log
}

// AsleepRecord computes the guard asleep total count, the minute is most often asleep, and the minute is most likely asleep.
func (g Guard) AsleepRecord() (int, int, int) {
	asleepRecord := make(map[int]int)
	asleepCount := 0
	for _, entries := range g.Logs {
		for j := 0; j < len(entries)-1; j++ {
			for i := entries[j].Minute; i < entries[j+1].Minute; i++ {
				if entries[j].State == Asleep {
					asleepRecord[i]++
					asleepCount++
				}
			}
		}
	}

	asleepMostTimes := 0
	asleepMostMinute := 0
	for k, v := range asleepRecord {
		if v > asleepMostTimes {
			asleepMostTimes = v
			asleepMostMinute = k
		}
	}

	return asleepCount, asleepMostTimes, asleepMostMinute
}

func main() {
	guards, err := readInput()
	if err != nil {
		fmt.Println(err)
	}

	// Part one.
	expectsOne := 35184
	resultOne := partOne(guards)
	printResult(expectsOne, resultOne)

	// Part Two.
	expectsTwo := 37886
	resultTwo := partTwo(guards)
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

// partOne computes the ID of the guard that spends most time asleep multiplied by the minute it is most likely to be asleep.
func partOne(guards []Guard) int {
	guardID, minute, total := 0, 0, 0
	for _, g := range guards {
		asleepTotal, _, minuteAsleepMost := g.AsleepRecord()

		if asleepTotal > total {
			total = asleepTotal
			minute = minuteAsleepMost
			guardID = g.ID
		}
	}

	return guardID * minute
}

// partTwo computes the product between the guard ID and the minute when is most frequently asleep.
func partTwo(guards []Guard) int {
	guardID, minute, times := 0, 0, 0
	for _, g := range guards {
		_, asleepMostTimes, asleepMostMinute := g.AsleepRecord()

		if asleepMostTimes > times {
			times = asleepMostTimes
			minute = asleepMostMinute
			guardID = g.ID
		}
	}

	return guardID * minute
}

// split a string at the specified characters.
func splitTimeRecord(s rune) bool {
	return s == '[' || s == ']' || s == ' ' || s == ':' || s == '#' || s == '-'
}

// readInput reads from stdin and parse the readings into a Guard and its Logs.
func readInput() ([]Guard, error) {
	scanner := bufio.NewScanner(os.Stdin)

	var logs [][]int
	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.FieldsFunc(line, splitTimeRecord)

		var digits []int
		state := Awake
		if strings.Contains(line, "asleep") {
			state = Asleep
		}
		digits = append(digits, state)

		for _, f := range fields {
			if unicode.IsDigit(rune([]byte(f)[0])) {
				num, err := strconv.Atoi(f)
				if err != nil {
					return nil, err
				}
				digits = append(digits, num)
			}
		}

		logs = append(logs, digits)
	}

	// Sort log records in chronological order.
	sort.Slice(logs, func(i, j int) bool {
		a := time.Date(logs[i][1], time.Month(logs[i][2]), logs[i][3], logs[i][4], logs[i][5], 0, 0, &time.Location{})
		b := time.Date(logs[j][1], time.Month(logs[j][2]), logs[j][3], logs[j][4], logs[j][5], 0, 0, &time.Location{})
		return a.Before(b)
	})

	// Create logs from log records and adds them to a guard.
	var uniqueID int
	guardMap := make(map[int]Guard)
	for _, log := range logs {
		if len(log) > 6 {
			uniqueID = log[len(log)-1]
		}

		// Initialize a guard if it doesn't exist.
		_, ok := guardMap[uniqueID]
		if !ok {
			guard := Guard{
				ID:   uniqueID,
				Logs: make(map[string][]Log),
			}
			guardMap[uniqueID] = guard
		}

		l := Log{
			Month:  log[2],
			Day:    log[3],
			Hour:   log[4],
			Minute: log[5],
			State:  log[0],
		}
		guardMap[uniqueID].Logs[l.MonthDay()] = append(guardMap[uniqueID].Logs[l.MonthDay()], l)
	}

	var guards []Guard
	for _, guard := range guardMap {
		guards = append(guards, guard)
	}

	return guards, nil
}
