package main 

import "fmt"

const (
     engPrefix = "Hello "
     spanPrefix = "Hola "
     frenPrefix = "Bonjour "
     dudePrefix = "Wassssup "

     spanish = "Spanish"
     french = "French"
     dude = "Dude"
)

func Hello(name string, language string) string {
    if name == "" {
        name = "world"
    } 
    return greetingPrefix(language) + name + "!"
}

func greetingPrefix(language string) (prefix string) { 
    switch language {
    case spanish:
        prefix = spanPrefix
    case french:
        prefix = frenPrefix        
    case dude:
        prefix = dudePrefix
    default:
        prefix = engPrefix
    }
    return 
}

func main() {
    fmt.Println(Hello("Chris", ""))
}
