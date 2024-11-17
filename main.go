package main

import (
    "bufio"
    "fmt"
    "log"
    "math"
    "os"
    "strconv"
)

func main() {
    args := os.Args[1:]
    if len(args) != 1 {
        log.Fatal("Usage: go run main.go <fileName>")
    }

    file, err := os.Open(args[0])
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close() 

    var nums []int
    buf := bufio.NewScanner(file)
    for buf.Scan() {
        strNum := buf.Text()
        num, err := strconv.Atoi(strNum)
        if err != nil {
            log.Fatal(err)
        }
        nums = append(nums, num)
    }

    if len(nums) == 0 {
        log.Fatal("Error: The file is empty")
    }

    var Σx, Σy, Σxy, Σxx, Σyy float64
    for x, y := range nums {
        Σx += float64(x)
        Σy += float64(y)
        Σxy += float64(x) * float64(y)
        Σxx += float64(x) * float64(x)
        Σyy += float64(y) * float64(y)
    }

    n := float64(len(nums))

    m := (n*Σxy - Σx*Σy) / (n*Σxx - (Σx * Σx))

    b := (Σy - m*Σx) / n

	// pearson correlation coefficient
    denominator := math.Sqrt((n*Σxx - Σx*Σx) * (n*Σyy - Σy*Σy))
    if denominator == 0 {
        log.Fatal("Error: Division by zero in Pearson correlation coefficient calculation")
    }
    r := (n*Σxy - Σx*Σy) / denominator

    fmt.Printf("Linear Regression Line: y = %.6fx + %.6f\n", m, b)
    fmt.Printf("Pearson Correlation Coefficient: %.10f\n", r)
}
