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
    "strconv"
    "strings"

    "./unit"
)

func challenge02(input []string) {
    var result int = 0
    var err error
    var op1 int
    var op2 int
    var op string
    var unit = unit.New(0, 0, "")

    if len(input) == 1 {
        result, err = strconv.Atoi(input[0])
    } else if len(input) == 2 {
        err = fmt.Errorf("Invalid line format: %s", input)
    } else if len(input) >= 3 {
        if (len(input) - 1) % 2 == 0 {
            err = nil
            for i := 0; i < len(input); i += 2 {
                if i == 0 {
                    op1, err = strconv.Atoi(input[i])
                    op2, err = strconv.Atoi(input[i+1])
                    op = input[i+2]
                    i += 1
                } else {
                    op1 = unit.Compute()
                    op2, err = strconv.Atoi(input[i])
                    op = input[i+1]
                }
                if err == nil {
                    unit.Update(op1, op2, op)
                } else {
                    err = fmt.Errorf("Invalid strconv in: %s", input)
                }
            }
        } else {
            err = fmt.Errorf("Invalid line format: %s", input)
        }
    }

    if err != nil {
        log.Println(err)
    } else {
        if len(input) >= 3 {
            result = unit.Compute()
        }
        fmt.Println(result)
    }
}

func main() {

    scanner := bufio.NewScanner(os.Stdin)
    for scanner.Scan() {
        line := scanner.Text()
        str_arr := strings.Split(line, " ")
        challenge02(str_arr)
    }

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }

}
