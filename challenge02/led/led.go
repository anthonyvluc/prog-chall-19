package led

import (
    "fmt"
    "strconv"
    "strings"
)

var LED0 = [][]rune{
    {' ', '_', ' '},
    {'|', ' ', '|'},
    {'|', '_', '|'},
}

var LED1 = [][]rune{
    {' ', ' ', ' '},
    {' ', ' ', '|'},
    {' ', ' ', '|'},
}

var LED2 = [][]rune{
    {' ', '_', ' '},
    {' ', '_', '|'},
    {'|', '_', ' '},
}

var LED3 = [][]rune{
    {' ', '_', ' '},
    {' ', '_', '|'},
    {' ', '_', '|'},
}

var LED4 = [][]rune{
    {' ', ' ', ' '},
    {'|', '_', '|'},
    {' ', ' ', '|'},
}

var LED5 = [][]rune{
    {' ', '_', ' '},
    {'|', '_', ' '},
    {' ', '_', '|'},
}

var LED6 = [][]rune{
    {' ', '_', ' '},
    {'|', '_', ' '},
    {'|', '_', '|'},
}

var LED7 = [][]rune{
    {' ', '_', ' '},
    {' ', ' ', '|'},
    {' ', ' ', '|'},
}

var LED8 = [][]rune{
    {' ', '_', ' '},
    {'|', '_', '|'},
    {'|', '_', '|'},
}

var LED9 = [][]rune{
    {' ', '_', ' '},
    {'|', '_', '|'},
    {' ', ' ', '|'},
}

var LEDNEG = [][]rune{
    {' ', ' ', ' '},
    {' ', '_', ' '},
    {' ', ' ', ' '},
}

func LEDFormat(n int) string {
    result := ""
    strn := strconv.Itoa(n)
    var cur_led_digit *[][]rune

    var digit_to_led = func(d rune) *[][]rune {
        switch d {
        case '0':
            cur_led_digit = &LED0
        case '1':
            cur_led_digit = &LED1
        case '2':
            cur_led_digit = &LED2
        case '3':
            cur_led_digit = &LED3
        case '4':
            cur_led_digit = &LED4
        case '5':
            cur_led_digit = &LED5
        case '6':
            cur_led_digit = &LED6
        case '7':
            cur_led_digit = &LED7
        case '8':
            cur_led_digit = &LED8
        case '9':
            cur_led_digit = &LED9
        case '-':
            cur_led_digit = &LEDNEG
        }
        return cur_led_digit
    }

    for r := 0; r < 3; r++ {
        for _, d := range strn {
            cur_led_digit = digit_to_led(d)
            for c := 0; c < 3; c++ {
                result = fmt.Sprintf("%s%c", result, (*cur_led_digit)[r][c])
            }
        }
        result = fmt.Sprintf("%s\n", result)
    }

    // Chomp newline
    result = strings.TrimSuffix(result, "\n")

    return result
}
