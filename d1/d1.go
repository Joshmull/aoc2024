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
    var similarities int

    scanner := bufio.NewScanner(file)
    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }

    for scanner.Scan() {
        string := strings.Split(scanner.Text(), "   ")
        left_list = append(left_list, convert_num(string[0]))
        right_list = append(right_list, convert_num(string[1]))
    }
    

    sort.Sort(sort.IntSlice(left_list))
    sort.Sort(sort.IntSlice(right_list))

    for i, left := range left_list {
        right := right_list[i]

        if left > right {
            diff += left - right
        } else {
            diff += right - left
        }

        num_count := 0

        for _, right_num := range right_list {
            if left == right_num {
                num_count += 1
            }
        }
        similarities += left * num_count
    }


    fmt.Println(diff)
    fmt.Println(similarities)
}

func convert_num(str string) int {
    num, err := strconv.Atoi(str)
    if err != nil {
        fmt.Println(err)
    }

    return num
}

