/*
 * # Challenge02
 *
 * https://www3.nd.edu/~pbui/teaching/cse.30872.fa19/challenge02.html
 *
 * ## My Idea
 *
 * ## Time Complexity
 *
 * ## Space Complexity
 *
 */

package main

import (
    "bufio"
    "fmt"
    "log"
    "os"
    "strings"
)

func challenge02() {

    scanner := bufio.NewScanner(os.Stdin)
    for scanner.Scan() {
        line := scanner.Text()
        str_arr := strings.Split(line, " ")
        fmt.Println(str_arr)
    }

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }
}

func main() {
    challenge02()
}
