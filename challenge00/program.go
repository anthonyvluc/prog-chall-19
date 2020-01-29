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
    "os"
)

func main() {
    scanner := bufio.NewScanner(os.Stdin)
    for scanner.Scan() {
        fmt.Println(scanner.Text())
    }
    if err := scanner.Err(); err != nil {
        log.Println(err)
    }
}

