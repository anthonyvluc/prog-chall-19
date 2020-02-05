/**
 * Challenge 01
 *
 * https://www3.nd.edu/~pbui/teaching/cse.30872.fa19/challenge01.html
 *
 * ## My Idea
 *      Given a string s1 and string s2, I want to count up the chars in each
 *      string. Then, I can determine the maximum amount of times a char in s2
 *      can be generated from s1. This can be done with integer division. Last,
 *      I will need to evaluate the chars in s2 and determine which has the
 *      lowest multiplicity, and that will tell me the maximum number of times
 *      I can generate s2 from s1.
 *
 * ## Time Complexity
 *      Given given n input where l1 is the length of the largest s1 string and
 *      l2 is the length of the largest s2 string, then the time complexity of
 *      this program is O(n*max(l1, l2)).
 *
 * ## Space Complexity
 *      The space complexity of this program will be O(1) because although the
 *      maps are generated from s1 and s2, the size of the space of possible
 *      runes is constant. So the space of potential keys k will equal the
 *      number of potential characters in the string (again, not dependent on
 *      the input). Though technically, this can be O(k) if the input strings
 *      change in the number of unique chars that's being used.
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

func challenge01() {

    /* Function to print results. */
    printResults := func(s1 string, s2 string) {
        var results int = 0
        m1, m2 := make(map[rune]int), make(map[rune]int)

        /* Count runes of string 1. */
        for _, v := range s1 {
            m1[v] += 1
        }

        /* Count runes of string 2. */
        for _, v := range s2 {
            m2[v] += 1
        }

        /* Determine multiplicity of characters. */
        multiples := make(map[rune]int)
        for _, v := range s2 {
            multiples[v] = m1[v] / m2[v]
        }

        /* Determine minimum result of multiplicity. */
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
        if len(rune_array) != 2 { continue }
        printResults(rune_array[0], rune_array[1])
    }

    /* END */
}

func main() {
    challenge01()
}
