package main 

import "fmt"

const prefix string = "Hello "

func Hello(name string) string {
    if name == "" {
        name = "world"
    }

    return prefix + name + "!"
}

func main() {
    fmt.Println(Hello("Chris"))
}
