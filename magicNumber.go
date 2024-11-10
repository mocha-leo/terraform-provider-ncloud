package main

import "fmt"

func main() {
    userAge := 21

    if userAge > 18 {
        fmt.Println("You are an adult. Access granted.")
    } else {
        fmt.Println("You are not an adult. Access denied.")
    }

    if userAge > 65 {
        fmt.Println("Eligible for senior citizen benefits.")
    } else if userAge < 13 {
        fmt.Println("You are a child. Special permissions may apply.")
    } else {
        fmt.Println("Standard user permissions apply.")
    }
}
