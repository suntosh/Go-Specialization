package main

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
    "strings"
)

func main() {
    fmt.Println("Peer-graded assignment: Module 3 - slice")
    fmt.Println("To exit enter 'x'")

    state := make([]int, 0)
    scanner := bufio.NewScanner(os.Stdin)

    for {
        fmt.Print("> ")
        scanner.Scan()
        command := scanner.Text()

        if strings.ToLower(strings.Trim(command, " ")) == "x" {
            os.Exit(0)
        }

        input, err := strconv.ParseInt(command, 0, 0)
        if err != nil {
            fmt.Println(fmt.Sprintf("Please enter a valid integer (You entered: %s)", command))
            continue
        }

        state = add(int(input), state)
        fmt.Println("state: ", state)
    }
}

func add(input int, state []int) []int {
    var newState []int
    for i := 0; i < len(state); i++ {
        curr := state[i]
        if input > curr {
            // input is greater than current, so continue
            continue
        }

        // get all elements that preceding the current element
        preceding := state[:i]
        positioned := make([]int, i)
        // create a new copy of the preceding
        copy(positioned, preceding)

        // append the new input value at the end of the preceding elements
        positioned = append(positioned, input)

        // append all the remaining elements, which should be larger that the input value
        newState = append(positioned, state[i:]...)
        return newState // exit the recursion and return the new state
    }

    // the recursion exited without mutating the slice, so we just need to append the input to the state
    return append(state, input)
}
