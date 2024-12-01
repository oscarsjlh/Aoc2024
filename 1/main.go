package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	dat, err := os.Open("./input")
	check(err)
	defer dat.Close()

	r := bufio.NewReader(dat)
	c1 := []int{}
	c2 := []int{}

	for {
		line, _, err := r.ReadLine()
		if len(line) > 0 {
			s := strings.Split(string(line), " ")
			s1, err := strconv.Atoi(s[0])
			check(err)
			s2, err := strconv.Atoi(s[3])
			check(err)

			c1 = append(c1, s1)
			c2 = append(c2, s2)
		}
		if err != nil {
			break
		}

	}

	sol1 := prob1(c1, c2)
	fmt.Println("Problem 1 solution")
	fmt.Println(sol1)
	sol2 := prob2(c1, c2)
	fmt.Println("Problem 2 solution")
	fmt.Println(sol2)
}

func prob1(a, b []int) int {
	sort.Sort(sort.IntSlice(a))
	sort.Sort(sort.IntSlice(b))

	total := 0
	for i, k := range a {
		d := diff(k, b[i])
		total = total + d
	}
	return total
}

func diff(a, b int) int {
	if a < b {
		return b - a
	}
	return a - b
}

func prob2(a, b []int) int {
	total := 0
	for _, k := range a {
		for _, secondColumn := range b {
			multiplier := 0
			if k == secondColumn {
				mult := 0
				multiplier = mult + 1
				total = total + (k * multiplier)
			}
		}
	}
	return total
}
