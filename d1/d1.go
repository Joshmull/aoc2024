package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main(){
    file, err := os.Open("./input")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    var left_list []int
    var right_list []int
    var diff int

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        string := strings.Split(scanner.Text(), "   ")
        left_list = append(left_list, convert_num(string[0]))
        right_list = append(right_list, convert_num(string[1]))
    }
    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }

    sort.Sort(sort.IntSlice(left_list))
    sort.Sort(sort.IntSlice(right_list))

    for i, _ := range left_list {
        left := left_list[i]
        right := right_list[i]

        if left > right {
            diff += left - right
        } else {
            diff += right - left
        }
    }

    fmt.Println(diff)
}

func convert_num(str string) int {
    num, err := strconv.Atoi(str)
    if err != nil {
        fmt.Println(err)
    }

    return num
}

