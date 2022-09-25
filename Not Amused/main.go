package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
	"strconv"
)

const (
	close = "CLOSE"
	exit  = "EXIT"
	enter = "ENTER"
	open  = "OPEN"
)

type visit struct {
	enterTime     int
	timeSpent int
}

func (v visit) calculateCost() float32 {
	return 0.1 * float32(v.timeSpent)
}

func main() {
	var visitors map[string]*visit
	day := 1

	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		line := input.Text()
		switch words := strings.Fields(line); words[0] {
		case open:
			if day != 1 {
				fmt.Printf("\n")
			}
			visitors = make(map[string]*visit)
		case close:
			printReport(day, visitors)
			day++
		case enter:
			name := words[1]
			enterTime, _ := strconv.Atoi(words[2])
			if visitor, ok := visitors[name]; ok {
				visitor.enterTime = enterTime
			} else {
				visitors[name] = &visit{enterTime: enterTime}
			}
		case exit:
			name := words[1]
			exitTime, _ := strconv.Atoi(words[2])
			visitor := visitors[name]
			visitor.timeSpent += (exitTime - visitor.enterTime)
		}
	}
}

func sortKeys(visitors map[string]*visit) []string {
	keys := make([]string, 0, len(visitors))
	for k := range visitors {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return keys
}

func printReport(day int, visitors map[string]*visit) {
	if len(visitors) > 0 {
		fmt.Printf("Day %d\n", day)
	}
	keys := sortKeys(visitors)
	for _, key := range keys {
		v := visitors[key]
		fmt.Printf("%s $%.2f\n", key, v.calculateCost())
	}
}
