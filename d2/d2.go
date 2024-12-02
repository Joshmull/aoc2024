package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
    "strconv"
)

func main(){
    file, err := os.Open("input")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }

    safe_count := 0

    for scanner.Scan() {
        list := strings.Split(scanner.Text(), " ")

        fmt.Println(list)

        if is_safe(list) {
            safe_count += 1
        }
    }

    println(safe_count)
}

func is_safe(list []string) bool {
    increasing := false
    decreasing := false

    for i, num := range list {
        if i != len(list) - 1 {
            num1 := convert_num(num)
            num2 := convert_num(list[i + 1])
            
            diff := 0

            if num1 > num2 {
                decreasing = true
                diff = num1 - num2
            } else {
                increasing = true
                diff = num2 - num1
            }

            if diff > 3 || diff == 0 {
                return false
            }

        }
    }

    if increasing == true && decreasing == true {
                return false
    } else {
        return true
    }
}

func convert_num(str string) int {
    num, err := strconv.Atoi(str)
    if err != nil {
        fmt.Println(err)
    }
    return num
}
