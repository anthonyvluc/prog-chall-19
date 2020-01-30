/**
 * https://www3.nd.edu/~pbui/teaching/cse.30872.fa19/challenge00.html
 * My Idea:
 *  1) Convert number n to string s
 *  2) Loop over string to return index q of the last index in s where
 *     it is not tidy.
 *  3a) If the index q is 0 or N-1, then return n
 *  3b) Otherwise return: to_num(s[i:q]+padded_zeros)-1
 * Time Complexity:
 * Space Complexity:
 */

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
    var n int          = 0
    var deck int       = 1
    var err error      = nil
    var s string       = ""

    printResults := func(s string) {
        var index int      = 0
        var sub_int int    = 0
        var prev_char rune = 0
        var result string  = ""

        for pos, char := range(s) {
            if char > prev_char {
                index = pos + 1
            } else if char == prev_char {
                continue
            } else {
                break
            }
            prev_char = char
        }

        if (index == 1 && len(s) == 1) || index == len(s) {
            result = s
        } else {
            sub_int, err = strconv.Atoi(s[:index])
            if err != nil {
                log.Println(err)
            }
            result = strconv.FormatInt(int64(sub_int*int(math.Pow10(len(s)-index)) - 1), 10)
        }

        fmt.Println(fmt.Sprintf("Deck #%d: %s", deck, result))
        deck += 1
    }

    /* Read from stdin. */
    scanner := bufio.NewScanner(os.Stdin)

    /* Read n_numbers. */
    scanner.Scan()
    n, err = strconv.Atoi(scanner.Text())
    if err != nil {
        log.Println(err)
    }
    n += 1

    /* Read and evaluate the rest of the numbers. */
    for scanner.Scan() {
        s = scanner.Text()
        printResults(s)
    }
    if err := scanner.Err(); err != nil {
        log.Println(err)
    }
}
