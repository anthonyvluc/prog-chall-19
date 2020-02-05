/**
 * # Challenge 00
 *
 * https://www3.nd.edu/~pbui/teaching/cse.30872.fa19/challenge00.html
 * https://code.google.com/codejam/contest/3264486/dashboard#s=p1
 *
 * ## My Idea
 *  1) Convert number n to string s
 *  2) Loop over string to return index q of the last index in s where
 *      it is not tidy.
 *  3a) If the index q is 0 or N-1, then return n
 *  3b) Otherwise return: to_num(s[i:q]+padded_zeros)-1
 *
 * ## Time Complexity
 *      For n numbers with at most a length of s, the largest string in the set,
 *      the time complexity in this program should be O(n*s).
 *
 * ## Space Complexity
 *      There are no data structures used in this program, so there are no
 *      variables which grow based on the input. Therefore, the space complexity
 *      of this program is O(1).
 *
 */

package main

import (
    "bufio"
    "fmt"
    "log"
    "math"
    "os"
    "strconv"
    "sync"
)

func challenge00() {
    var n int          = 0
    var deck int       = 1
    var err error      = nil
    var s string       = ""
    var ch chan string = make(chan string)
    var wg sync.WaitGroup

    generateResult := func(s string, deck int, ch chan<- string) {
        var index int      = 0
        var sub_int int    = 0
        var prev_char rune = 0
        var result string  = ""

        defer wg.Done()

        if s == "" {
            s = "0"
        }

        for pos, char := range(s) {
            if char > prev_char {
                index = pos + 1
            } else if char == prev_char {
                if pos == len(s) - 1 {
                    index = pos + 1
                }
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

        result = fmt.Sprintf("Deck #%d: %s", deck, result)
        fmt.Println(result)
    }

    /* Read from stdin. */
    scanner := bufio.NewScanner(os.Stdin)

    /* Read n_numbers. */
    scanner.Scan()
    n, err = strconv.Atoi(scanner.Text())
    if err != nil {
        log.Println(err)
    }
    _ = n // suppress error

    /* Read and evaluate the rest of the numbers. */
    for scanner.Scan() {
        wg.Add(1)
        s = scanner.Text()
        go generateResult(s, deck, ch)
        deck += 1
    }
    if err := scanner.Err(); err != nil {
        log.Println(err)
    }

    wg.Wait()
}

func main() {
    challenge00()
}
