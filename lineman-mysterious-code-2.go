package main

import (
	"fmt"
	"regexp"
	"strings"
)

const MARK = '*'

func removeHighFreq(str string) string {
	counter := make(map[string]int)
	max := 0
    for _, c := range str {
        if string(c) != " " {
            counter[string(c)]++
			if max < counter[string(c)] {
				max = counter[string(c)]
			}
        }
    }

	highFreqChars := make([]string, 0)
    for chr := range counter {
		if counter[chr] == max {
			highFreqChars = append(highFreqChars, chr)
		}
    }

	highFreqRegex := regexp.MustCompile(strings.Join(highFreqChars, "|"))
	return highFreqRegex.ReplaceAllString(str,"")
}

func calculateDirection(direction int, current int, limit int) int {
	if current == limit - 1 {
		direction = -1
	} else if current == 0 {
		direction = 1
	}
	return direction
}

func createZigZagArea(row int, col int) [][]int {
	area := make([][]int, row)
	for i := range area {
		area[i] = make([]int, col)
	}
	currentRow := 0
	direction := 1
	count := 0
	for count < col {
		area[currentRow][count] = MARK
		count += 1
		direction = calculateDirection(direction, currentRow, row)
		currentRow += direction
	}
	return area
}

func assignZigZagArea(area [][]int, strs string) {
	col := len(strs)
	row := len(area)
	count := 0
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if area[i][j] == MARK {
				area[i][j] = int(strs[count])
				count += 1
			}
		}
	}
}

func drawArea(area [][]int) {
	n := len(area)
	m := len(area[0])
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if area[i][j] == 0 {
				fmt.Print(" ")
			} else {
				fmt.Printf("%c", area[i][j])
			}
		}
		fmt.Print("\n")
	}
}

func readZigZag(area [][]int) {
	col := len(area[0])
	row := len(area)

	currentRow := 0
	direction := 1
	count := 0
	for count < col {
		fmt.Printf("%c", area[currentRow][count])
		count += 1
		direction = calculateDirection(direction, currentRow, row)
		currentRow += direction
	}
	fmt.Println()
}


func main() {
	whatIsIt := "CYtZBsWZaZliYZocWLZlXuZZYWYeYXZsXeZXtXWpXeRYYYd!ZnYeWXoYXasnX,WXWrWPoAdWesnciGenWr"
	whatIsIt = removeHighFreq(whatIsIt)
	fmt.Println(whatIsIt)
	rails := 4 // from guessing
	length := len(whatIsIt)
	area := createZigZagArea(rails, length)
	assignZigZagArea(area, whatIsIt)
	readZigZag(area)
	// drawArea(area)
}
