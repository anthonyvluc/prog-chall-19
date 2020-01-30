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
    "os"
    "strings"
)

func main() {

    printResults := func(s1 string, s2 string) {
        var results int = 0

        m1, m2 := make(map[rune]int), make(map[rune]int)

        for _, v := range s1 {
            m1[v] += 1
        }

        for _, v := range s2 {
            m2[v] += 1
        }

        multiples := make(map[rune]int)
        for _, v := range s2 {
            multiples[v] = m1[v] / m2[v]
        }

        var min int = 0
        fmt.Println(multiples)
        fmt.Println(min)
        fmt.Println(1/10000) // TODO
        for _, v := range multiples {
            if min < v {
                min = v
                results = min
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
