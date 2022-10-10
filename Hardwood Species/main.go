package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

const toPercentage = 100.0

func main() {
	treeInventory := make(map[string]float32)
	var totalTrees float32
	stdin := bufio.NewScanner(os.Stdin)
	for stdin.Scan() {
		tree := stdin.Text()
		treeInventory[tree]++
		totalTrees++
	}

	text := strings.Builder{}
	totalTrees /= toPercentage
	for _, tree := range sortKeys(treeInventory) {
		fractionOfTotal := treeInventory[tree] / totalTrees
		text.WriteString(fmt.Sprintf("%s %.6f\n", tree, fractionOfTotal))
	}
	fmt.Print(text.String())
}

func sortKeys(map_ map[string]float32) []string {
	keys := make([]string, 0, len(map_))
	for k := range map_ {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return keys
}
