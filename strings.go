/*
    Ken Bailey
    3/7/18

    All about strings in Go

*/

package main

import (
    "strings"
    "fmt"
    "strconv"
    "log"
)

func main() {
    str := "This is a string"
    otherStr := []byte("This can also be a string")

    // Print a string
    fmt.Println(str)

    // Print a casted string
    fmt.Println(string(otherStr))

    // Print a formatted string
    //  Printf() allows for formatting but no new line!
    //  A byte array can be interpreted as a string using Printf
    fmt.Printf("%s -- I see!\n", str)
    fmt.Printf("%s -- Also treated as string\n", otherStr)

    // Converting a number to string
    num := "1234567890"
    newInt, err := strconv.Atoi(num)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println(newInt+1)
    fmt.Println(num)

    // String - Contains
    a := "the giant has awoken"
    b := "ken"
    fmt.Print("(contains) ken v awoken: ")
    fmt.Println(strings.Contains(a, b))

    // String - Find index of substring in s
    fmt.Print("(index) Index of "+b+" in "+a+": ")
    fmt.Println(strings.Index(a, b))


}