/**
 * Challenge 01
 *
 * https://www3.nd.edu/~pbui/teaching/cse.30872.fa19/challenge01.html
 *
 * ## My Idea
 *
 * ## Time Complexity
 *
 * ## Space Complexity
 *
 * ## Potential Improvements
 *
 */

package main

import (
    "bufio"
    "fmt"
    "math"
    "os"
    "strings"
)

func main() {

    printResults := func(s1 string, s2 string) {
        var results int = 0

        m1, m2 := make(map[rune]int), make(map[rune]int)

        /* Count bytes of string 1. */
        for _, v := range s1 {
            m1[v] += 1
        }

        /* Count bytes of string 2. */
        for _, v := range s2 {
            m2[v] += 1
        }

        multiples := make(map[rune]int)
        for _, v := range s2 {
            multiples[v] = m1[v] / m2[v]
        }

        var max int = math.MaxInt32
        for _, v := range multiples {
            if v < max {
                max = v
                results = max
            }
        }

        fmt.Println(results)
    }

    /* BEGIN */

    /* Read from stdin. */
    scanner := bufio.NewScanner(os.Stdin)

    /* Process and print results. */
    for scanner.Scan() {
        line := scanner.Text()
        rune_array := strings.Split(line, " ")
        if len(rune_array) == 1 { continue }
        printResults(rune_array[0], rune_array[1])
    }

    /* END */
}
