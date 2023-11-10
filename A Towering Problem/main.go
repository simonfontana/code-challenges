package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	line := scanner.Text()
	inputStr := strings.Fields(line)

	var input []int
	for _, str := range inputStr {
		i, _ := strconv.Atoi(str)
		input = append(input, i)
	}

	boxes := input[:6]
	sort.Ints(boxes)
	towers := input[6:]

	for box1 := 0; box1 < len(boxes); box1++ {
		for box2 := box1 + 1; box2 < len(boxes); box2++ {
			for box3 := box2 + 1; box3 < len(boxes); box3++ {
				a := boxes[box1]
				b := boxes[box2]
				c := boxes[box3]
				if a+b+c == towers[0] {
					fmt.Printf("%d %d %d", c, b, a)
					for i := 5; i >= 0; i-- {
						if i != box1 && i != box2 && i != box3 {
							fmt.Printf(" %d", boxes[i])
						}
					}
					return
				}
			}
		}
	}
}
