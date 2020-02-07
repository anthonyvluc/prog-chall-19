package unit

import (
    "log"
)

type unit struct {
    operand1 int
    operand2 int
    operation string
}

func New(op1 int, op2 int, op string) unit {
    return unit{op1, op2, op}
}

func (u *unit) Update(op1 int, op2 int, op string) {
    u.operand1 = op1
    u.operand2 = op2
    u.operation = op
}

func (u *unit) Valid() bool {
    return u.operation != ""
}

func (u *unit) Compute() int {
    result := 0
    if u.Valid() {
        switch u.operation {
        case "+":
            result = u.operand1 + u.operand2
        case "-":
            result = u.operand1 - u.operand2
        case "*":
            result = u.operand1 * u.operand2
        case "/":
            result = u.operand1 / u.operand2
        default:
            log.Println("Invalid operation")
        }
    }
    return result
}
