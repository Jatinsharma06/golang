package main

import "fmt"

// Struct definition
type Person struct {
    Name string
    Age  int
}

// Function to greet a person
func (p Person) Greet() {
    fmt.Printf("Hello, my name is %s and I am %d years old.\n", p.Name, p.Age)
}

func main() {
    fmt.Println("Hello, World!")

    // Loop example
    for i := 1; i <= 5; i++ {
        fmt.Println("Count:", i)
    }

    // Struct usage
    person := Person{Name: "Jatin", Age: 22}
    person.Greet()
}
